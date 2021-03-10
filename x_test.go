package x
import "testing"

func TestInfo(t *testing.T) {
   LogInfo("Run", "ffmpeg", []string{"-i", "in file.flac", "out file.opus"})
}
