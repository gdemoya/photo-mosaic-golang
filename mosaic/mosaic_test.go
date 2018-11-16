package mosaic

import (
	"log"
	"testing"
)

func BenchmarkMosaic(b *testing.B) {
	_, err := NewMosaic(
		"../img/cage.jpg",
		"../img/tiles/",
		80,
		50,
		4,
		180,
	)

	if err != nil {
		log.Fatal(err)
	}
}
