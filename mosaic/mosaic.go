package mosaic

import (
	"log"
	"math"
	"os"
	"sync"

	"golang-challenge.org/challenge3/util"

	"image"
	"image/color"
	"image/draw"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

// ALGORITHM
// Take a scanned photo (source image) to process.
// Make a grid over this photo.
// Look at every cell on the grid.
// Calculate the average color of each cell.
// Find the nearest image in the library with the same average color of the cell and substitute it in that cell.
// Repeat the process for each cell in the grid.

type Mosaic struct {
	master                         MasterImage
	tilerCollection                TilerCollection
	cv, ch                         int
	tiler_h_pixels, tiler_v_pixels int
	mask                           image.Image
	Threads                        int
}

func NewMosaic(masterPath, tilerPath string, ch, cv, threads int, mask uint8) (Mosaic, error) {
	log.Printf("Creating new Mosaic...")

	m, err := NewMasterImage(masterPath, ch, cv)
	if err != nil {
		log.Fatal(err)
	}

	bounds := m.Bounds()

	tiler_h_pixels := (bounds.Max.X - bounds.Min.X) / ch
	tiler_v_pixels := (bounds.Max.Y - bounds.Min.Y) / cv

	t, err := NewTilerCollection(tilerPath, tiler_h_pixels, tiler_v_pixels, threads)
	if err != nil {
		log.Fatal(err)
	}

	return Mosaic{
		master:          m,
		tilerCollection: t,
		cv:              cv,
		ch:              ch,
		tiler_h_pixels:  tiler_h_pixels,
		tiler_v_pixels:  tiler_v_pixels,
		mask:            &image.Uniform{color.RGBA{mask, mask, mask, mask}},
		Threads:         threads,
	}, nil
}

func (img Mosaic) Generate() {
	cv := img.cv / img.Threads
	for i := 0; i < img.Threads; i++ {
		go func(i int) {
			cv_init := i * cv
			cv_final := cv_init + cv
			if i == (img.Threads - 1) {
				cv_final += (img.cv % img.Threads)
			}
			log.Printf("%v hasta %v", cv_init, cv_final)
			for y := cv_init; y < cv_final; y++ {
				for x := 0; x < img.ch; x++ {
					x0_tmp := x * img.tiler_h_pixels
					x1_tmp := (x + 1) * img.tiler_h_pixels
					y0_tmp := y * img.tiler_v_pixels
					y1_tmp := (y + 1) * img.tiler_v_pixels

					tmp_rect := image.Rect(x0_tmp, y0_tmp, x1_tmp, y1_tmp)
					slice := img.master.SubImage(tmp_rect)
					a_color := util.AverageColor(slice)

					tiler := img.tilerCollection.SearchClosestColorTiler(a_color)

					draw.DrawMask(img.master, tmp_rect, tiler, image.ZP, img.mask, image.ZP, draw.Over)
				}
			}
		}(i)

	}
}

func (img Mosaic) Get() image.Image {
	return &img.master
}

type MasterImage struct {
	*image.RGBA
}

func NewMasterImage(masterPath string, ch, cv int) (MasterImage, error) {
	reader, err := os.Open(masterPath)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	img, _, err := image.Decode(reader)
	if err != nil {
		log.Fatalf("Error with file %s: %v", masterPath, err)
	}

	bounds := img.Bounds()

	width := (bounds.Max.X - bounds.Min.X)
	fix_width := (width / ch) * ch

	heigth := (bounds.Max.Y - bounds.Min.Y)
	fix_heigth := (heigth / cv) * cv

	dst := image.NewRGBA(image.Rect(0, 0, fix_width, fix_heigth))
	util.ResizeImage(dst, img)

	return MasterImage{dst}, nil
}

type TilerCollection struct {
	tilerImages []TilerImage
}

func NewTilerCollection(tilerPath string, tiler_h_pixels, tiler_v_pixels, threads int) (TilerCollection, error) {
	var tilerCollection []TilerImage

	f, err := os.Open(tilerPath)
	if err != nil {
		return TilerCollection{nil}, err
	}
	fileInfo, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return TilerCollection{nil}, err
	}

	var wg sync.WaitGroup
	wg.Add(threads)
	nfiles := len(fileInfo) / threads
	for i := 0; i < threads; i++ {
		go func(i int) {
			defer wg.Done()
			nfile_init := i * nfiles
			nfile_final := nfile_init + nfiles
			if i == (threads - 1) {
				nfile_final += (nfiles % threads)
			}
			log.Printf("%v hasta %v", nfile_init, nfile_final)

			for y := nfile_init; y < nfile_final; y++ {
				file := fileInfo[y]
				t, err := NewTilerImage(tilerPath+file.Name(), tiler_h_pixels, tiler_v_pixels)
				if err != nil {
					log.Fatal(err)
				}
				tilerCollection = append(tilerCollection, t)
			}
		}(i)
	}

	log.Printf("Waiting...")
	wg.Wait()
	log.Printf("Finished")
	return TilerCollection{tilerCollection}, nil
}

func (tc TilerCollection) SearchClosestColorTiler(color color.Color) TilerImage {
	var out_index int
	closest_dist := math.MaxFloat64

	for index, tiler := range tc.tilerImages {
		tmp_dist := util.ColorDistance(tiler.AverageColor, color)
		if tmp_dist < closest_dist {
			closest_dist = tmp_dist
			out_index = index
		}
	}

	return tc.tilerImages[out_index]
}

type TilerImage struct {
	*image.RGBA
	AverageColor color.Color
}

func NewTilerImage(tilerPath string, tiler_h_pixels, tiler_v_pixels int) (TilerImage, error) {
	// log.Printf("Creating new Tiller...")

	reader, err := os.Open(tilerPath)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatalf("Error with file %s: %v", tilerPath, err)
	}

	dst := image.NewRGBA(image.Rect(0, 0, tiler_h_pixels, tiler_v_pixels))
	util.ResizeImage(dst, m)

	a_color := util.AverageColor(dst)

	return TilerImage{dst, a_color}, nil
}
