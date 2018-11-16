package mosaic

import (
	"testing"
	"log"
)


func BenchmarkMosaic(b *testing.B) {
	_, err := NewMosaic(
			"../img/cage.jpg", 
			"../img/tiles/", 
			80 ,
			50,
			180,
		)
	
	if err != nil {
		log.Fatal(err)
	}
}