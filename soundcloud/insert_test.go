package soundcloud
import "testing"

var tests = []struct{
   url string
   artwork string
}{
   {"bluewednesday/murmuration-feat-shopan", "artworks-000507498393-dd22sy"},
   {"four-tet/burial-four-tet-nova", "avatars-000014893963-91gp52"},
}

func TestInsert(t *testing.T) {
   for _, in := range tests {
      out, e := Insert("https://soundcloud.com/" + in.url)
      if e != nil {
         t.Error(e)
      }
      if out.Artwork != in.artwork {
         t.Error(out)
      }
   }
}
