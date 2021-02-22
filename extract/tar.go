// Extract
package extract

import (
   "archive/tar"
   "compress/bzip2"
   "compress/gzip"
   "fmt"
   "github.com/klauspost/compress/zstd"
   "github.com/xi2/xz"
   "io"
   "os"
   "path"
   "strings"
)

func replace(s string, strip int, add string) string {
   split := strings.SplitN(s, "/", strip + 1)
   if len(split) <= strip {
      return ""
   }
   return path.Join(add, split[strip])
}

type Tar struct {
   Strip int
}

func (t Tar) Bz2(source, dest string) error {
   open, err := os.Open(source)
   if err != nil {
      return err
   }
   return t.readFrom(bzip2.NewReader(open), dest)
}

func (t Tar) Gz(source, dest string) error {
   open, err := os.Open(source)
   if err != nil {
      return err
   }
   read, err := gzip.NewReader(open)
   if err != nil {
      return err
   }
   return t.readFrom(read, dest)
}

func (t Tar) Xz(source, dest string) error {
   open, err := os.Open(source)
   if err != nil {
      return err
   }
   read, err := xz.NewReader(open, 0)
   if err != nil {
      return err
   }
   return t.readFrom(read, dest)
}

func (t Tar) Zst(source, dest string) error {
   open, err := os.Open(source)
   if err != nil {
      return err
   }
   read, err := zstd.NewReader(open)
   if err != nil {
      return err
   }
   return t.readFrom(read, dest)
}

func (t Tar) readFrom(source io.Reader, dest string) error {
   read := tar.NewReader(source)
   for {
      file, err := read.Next()
      if err == io.EOF {
         break
      } else if err != nil {
         return err
      }
      name := replace(file.Name, t.Strip, dest)
      if name == "" {
         continue
      }
      switch file.Typeflag {
      case tar.TypeLink:
         _, err = os.Stat(name)
         if err == nil {
            err = os.Remove(name)
            if err != nil {
               return fmt.Errorf("readFrom %v", err)
            }
         }
         err = os.Link(replace(file.Linkname, t.Strip, dest), name)
         if err != nil {
            return err
         }
      case tar.TypeReg:
         err = os.MkdirAll(path.Dir(name), os.ModeDir)
         if err != nil {
            return err
         }
         create, err := os.Create(name)
         if err != nil {
            return err
         }
         _, err = create.ReadFrom(read)
         if err != nil {
            return err
         }
      }
   }
   return nil
}
