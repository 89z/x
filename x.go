// Small Go packages
package x

import (
   "bufio"
   "bytes"
   "encoding/json"
   "fmt"
   "net/http"
   "os"
   "os/exec"
   "path/filepath"
   "regexp"
)

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
   LogInfo("Get", source)
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

func LogFail(verb, msg string) {
   fmt.Println("\x1b[30;101m " + verb + " \x1b[m", msg)
}

func LogInfo(verb, msg string) {
   fmt.Println("\x1b[30;103m " + verb + " \x1b[m", msg)
}

func LogPass(verb, msg string) {
   fmt.Println("\x1b[30;102m " + verb + " \x1b[m", msg)
}

func Popen(name string, arg ...string) (*bufio.Scanner, error) {
   cmd := exec.Command(name, arg...)
   pipe, err := cmd.StdoutPipe()
   if err != nil {
      return nil, fmt.Errorf("StdoutPipe %v", err)
   }
   return bufio.NewScanner(pipe), cmd.Start()
}
