// YouTube
package youtube

import (
   "encoding/json"
   "fmt"
   "github.com/89z/x"
   "io/ioutil"
   "math"
   "net/http"
   "net/url"
   "strconv"
   "time"
)

func numberFormat(x float64) string {
   n := int(math.Log10(x)) / 3
   a := x / math.Pow10(n * 3)
   return fmt.Sprintf("%.3f", a) + []string{"", " k", " M", " B"}[n]
}

func sinceHours(left string) (float64, error) {
   right := "1970-01-01T00:00:00Z"[len(left):]
   t, err := time.Parse(time.RFC3339, left + right)
   if err != nil {
      return 0, err
   }
   return time.Since(t).Hours(), nil
}

func Info(id string) (Player, error) {
   get, err := http.Get("https://www.youtube.com/get_video_info?video_id=" + id)
   if err != nil {
      return Player{}, err
   }
   body, err := ioutil.ReadAll(get.Body)
   if err != nil {
      return Player{}, err
   }
   query, err := url.ParseQuery(string(body))
   if err != nil {
      return Player{}, err
   }
   var resp response
   err = json.Unmarshal(
      []byte(query.Get("player_response")), &resp,
   )
   if err != nil {
      return Player{}, err
   }
   return resp.Microformat.PlayerMicroformatRenderer, nil
}

func (p Player) Views() (string, error) {
   view, err := strconv.ParseFloat(p.ViewCount, 64)
   if err != nil {
      return "", err
   }
   hour, err := sinceHours(p.PublishDate)
   if err != nil {
      return "", err
   }
   view /= hour / 24 / 365
   format := numberFormat(view)
   if view > 8_000_000 {
      return x.ColorRed(format), nil
   }
   return x.ColorGreen(format), nil
}
