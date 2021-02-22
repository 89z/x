package extract

import (
   "os"
   "path"
   "testing"
)

func TestZip(t *testing.T) {
   cache, err := os.UserCacheDir()
   if err != nil {
      t.Error(err)
   }
   source := path.Join(cache, "x", "readme.zip")
   dest := path.Join(cache, "x", "zip")
   err = Zip(source, dest)
   if err != nil {
      t.Error(err)
   }
}
