// MusicBrainz
package musicbrainz

import (
   "encoding/json"
   "net/http"
   "net/url"
   "path"
   "sort"
   "strings"
)

func NewRelease(in string) (Release, error) {
   id := path.Base(in)
   value := url.Values{}
   value.Set("fmt", "json")
   value.Set("inc", "artist-credits recordings")
   if strings.Contains(in, "/release/") {
      get, err := http.Get(
         "http://musicbrainz.org/ws/2/release/" + id + "?" + value.Encode(),
      )
      if err != nil {
         return Release{}, err
      }
      var album Release
      return album, json.NewDecoder(get.Body).Decode(&album)
   }
   value.Set("release-group", id)
   get, err := http.Get(
      "http://musicbrainz.org/ws/2/release?" + value.Encode(),
   )
   if err != nil {
      return Release{}, err
   }
   var group ReleaseGroup
   err = json.NewDecoder(get.Body).Decode(&group)
   if err != nil {
      return Release{}, err
   }
   Sort(group.Releases)
   return group.Releases[0], nil
}

func Sort(releases []Release) {
   sort.Slice(releases, func (i, j int) bool {
      first, second := releases[i], releases[j]
      // 1. STATUS
      if status(first) > status(second) {
         return true
      }
      if status(first) < status(second) {
         return false
      }
      // 2. YEAR
      if date(first, 4) < date(second, 4) {
         return true
      }
      if date(first, 4) > date(second, 4) {
         return false
      }
      // 3. TRACKS
      if trackLen(first) < trackLen(second) {
         return true
      }
      if trackLen(first) > trackLen(second) {
         return false
      }
      // 4. DATE
      return date(first, 10) < date(second, 10)
   })
}

func date(album Release, width int) string {
   start := len(album.Date)
   right := "9999-12-31"[start:]
   return (album.Date + right)[:width]
}

func status(album Release) int {
   if album.Status != "Official" {
      return 0
   }
   return 1
}

func trackLen(album Release) (track int) {
   for _, each := range album.Media {
      track += each.TrackCount
   }
   return
}
