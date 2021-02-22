package x

import (
   "bytes"
   "encoding/json"
   "fmt"
   "net/http"
   "os"
   "path/filepath"
   "regexp"
)

func ColorCyan(s string) string {
   return "\x1b[96m" + s + "\x1b[m"
}

// Foreground bright green
func ColorGreen(s string) string {
   return "\x1b[92m" + s + "\x1b[m"
}

// Foreground bright red
func ColorRed(s string) string {
   return "\x1b[91m" + s + "\x1b[m"
}

func Copy(source, dest string) (int64, error) {
   info, err := os.Stat(dest)
   if err == nil {
      if info.Mode().IsRegular() {
         return 0, os.ErrExist
      }
      err = os.Remove(dest)
      if err != nil {
         return 0, fmt.Errorf("%q %q %v", source, dest, err)
      }
   }
   err = os.MkdirAll(filepath.Dir(dest), os.ModeDir)
   if err != nil {
      return 0, err
   }
   create, err := os.Create(dest)
   if err != nil {
      return 0, err
   }
   defer create.Close()
   fmt.Println(ColorCyan("Get"), source)
   get, err := http.Get(source)
   if err != nil {
      return 0, err
   }
   return create.ReadFrom(get.Body)
}

func FindSubmatch(re string, input []byte) []byte {
   a := regexp.MustCompile(re).FindSubmatch(input)
   if len(a) < 2 {
      return nil
   }
   return a[1]
}

func JsonMarshal(v interface{}) ([]byte, error) {
   var dst bytes.Buffer
   enc := json.NewEncoder(&dst)
   enc.SetEscapeHTML(false)
   enc.SetIndent("", " ")
   err := enc.Encode(v)
   if err != nil {
      return nil, err
   }
   return dst.Bytes()[:dst.Len() - 1], nil
}
