package mp3

import "io"

type SizedSeeker interface {
	io.Seeker
	Length() int64
}
