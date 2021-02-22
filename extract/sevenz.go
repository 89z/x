package extract

import (
   "bytes"
   "github.com/saracen/go7z"
   "io"
   "io/ioutil"
   "os"
   "path"
)

var Magic7z = [6]byte{'7', 'z', 0xBC, 0xAF, 0x27, 0x1C}

func Exe(source, dest string) error {
   data, err := ioutil.ReadFile(source)
   if err != nil {
      return err
   }
   index := bytes.Index(data, Magic7z[:])
   bt := bytes.NewReader(data[index:])
   sz, err := go7z.NewReader(bt, int64(bt.Len()))
   if err != nil {
      return err
   }
   return readFrom(sz, dest)
}

func SevenZ(source, dest string) error {
   sz, err := go7z.OpenReader(source)
   if err != nil {
      return err
   }
   return readFrom(&sz.Reader, dest)
}

func readFrom(source *go7z.Reader, dest string) error {
   for {
      file, err := source.Next()
      if err == io.EOF {
         break
      } else if err != nil {
         return err
      } else if file.IsEmptyStream {
         continue
      }
      join := path.Join(dest, file.Name)
      err = os.MkdirAll(path.Dir(join), os.ModeDir)
      if err != nil {
         return err
      }
      create, err := os.Create(join)
      if err != nil {
         return err
      }
      _, err = create.ReadFrom(source)
      if err != nil {
         return err
      }
   }
   return nil
}
