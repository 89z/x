package musicbrainz

type ReleaseGroup struct {
   Releases []Release
}

type Release struct {
   ArtistCredit []struct {
      Artist struct {
         Id string
      }
      Name string
   } `json:"artist-credit"`
   Date string
   Media []struct {
      TrackCount int `json:"track-count"`
      Tracks []struct {
         Length int
         Title string
      }
   }
   Status string
   Title string
}
