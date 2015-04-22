package compositor

import (
	"os"
	"fmt"
	"github.com/Lostmyname/magick"
)

const NCPU = 8

func GeneratePage(assets []string, page int) {
	image_001, _ := magick.NewFromFile(assets[0])
	defer image_001.Destroy()

	for _, layerName := range assets[1:] {
		layer, _ := magick.NewFromFile(layerName)
		defer layer.Destroy()

		image_001.Compose(magick.SrcAtopCompositeOp, layer, 0, 0)
	}

	outFile := fmt.Sprintf("test/page_%02d.png", page)
	os.Remove(outFile)
	image_001.ToFile(outFile)	
}

func GeneratePagesAsync(assets []string, count int) {
	c := make(chan int, NCPU)

	for i := 1; i < count; i++ {
		go func() {
			GeneratePage(assets, i)
			c <- 1
		}()
	}

	for i := 1; i < count; i++ {
		<-c
	}
}

func GeneratePagesSync(assets []string, count int) {
	for i := 0; i < count; i++ {
		GeneratePage(assets, i+1)
	}
}
