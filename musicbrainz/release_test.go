package musicbrainz

import (
   "fmt"
   "testing"
)

func TestRelease(t *testing.T) {
   album, err := NewRelease("/release/5e2e02a7-2d83-4614-a7fd-8af265d85c34")
   if err != nil {
      t.Error(err)
   }
   fmt.Printf("%+v\n", album)
}
