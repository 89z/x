// SoundCloud
package soundcloud

import (
   "github.com/89z/x"
   "io/ioutil"
   "net/http"
)

type Player struct{
   Artwork string
   Id string
   Pubdate string
   Title string
}

func Insert(url string) (Player, error) {
   get, err := http.Get(url)
   if err != nil {
      return Player{}, err
   }
   body, err := ioutil.ReadAll(get.Body)
   if err != nil {
      return Player{}, err
   }
   artwork := x.FindSubmatch("/([^/]+)-t500x500.jpg", body)
   id := x.FindSubmatch(`/tracks/([^"]+)"`, body)
   pubdate := x.FindSubmatch(` pubdate>(\d{4})-`, body)
   title := x.FindSubmatch("<title>([^|]+) by ", body)
   return Player{
      string(artwork),
      string(id),
      string(pubdate),
      string(title),
   }, nil
}
