package extract

import (
   "archive/zip"
   "os"
   "path"
)

func Zip(source, dest string) error {
   read, err := zip.OpenReader(source)
   if err != nil {
      return err
   }
   for _, file := range read.File {
      if file.Mode().IsDir() {
         continue
      }
      join := path.Join(dest, file.Name)
      err = os.MkdirAll(path.Dir(join), os.ModeDir)
      if err != nil {
         return err
      }
      open, err := file.Open()
      if err != nil {
         return err
      }
      create, err := os.Create(join)
      if err != nil {
         return err
      }
      _, err = create.ReadFrom(open)
      if err != nil {
         return err
      }
   }
   return nil
}
