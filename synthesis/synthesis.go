package synthesis

import (
	"github.com/nfnt/resize"
	"image"
	"log"
	"os"

	"image/draw"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

type ImageInfo struct {
	Img  image.Image
	Size image.Point
}

func GetImageInfo(filePath string) ImageInfo {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	var imageInfo ImageInfo
	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	imageInfo.Img = img
	imageInfo.Size = img.Bounds().Size()

	return imageInfo
}

func ResizeImage(imgInfo *ImageInfo, x uint, y uint) {
	(*imgInfo).Img = resize.Resize(x, y, (*imgInfo).Img, resize.Lanczos3)
	(*imgInfo).Size = (*imgInfo).Img.Bounds().Size()
}

func SynthesizeImages(baseInfo ImageInfo, overInfo ImageInfo) *image.RGBA {
	startPoint := image.Point{0, int(baseInfo.Size.Y / 10)}
	overRectangle := image.Rectangle{startPoint, startPoint.Add(overInfo.Size)}
	baseRectangle := image.Rectangle{image.Point{0, 0}, baseInfo.Size}

	rgba := image.NewRGBA(baseRectangle)
	draw.Draw(rgba, baseRectangle, baseInfo.Img, image.Point{0, 0}, draw.Src)
	draw.Draw(rgba, overRectangle, overInfo.Img, image.Point{0, 0}, draw.Over)

	return rgba
}
