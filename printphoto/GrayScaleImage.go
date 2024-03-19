package printphoto

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/nfnt/resize"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"image"
	"log"
	"os"
)

func outputImage(img image.Image) {

}

func Printgrayscaleimage(Width uint, Height uint) {
	buf := bytes.NewBuffer(nil)
	ffmpeg.Input("dragon.mp4").
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", 1)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	img, err := imaging.Decode(buf)
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
	}
	img = resize.Resize(Width, Height, img, resize.Lanczos3)
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()
	// 每个像素对应的字符
	pixelChar := []string{" ", ".", "-", ":", "*", "=", "+", "#", "%", "@"}
	var buffer bytes.Buffer
	// 输出每个像素的字符到终端
	for y := range height {
		for x := range width {
			r, g, b, _ := img.At(x, y).RGBA()
			avg := (r*299 + g*587 + b*114) / 1000
			index := int(avg>>8) / 25
			if index > 9 {
				index = 9
			}
			fmt.Fprint(&buffer, pixelChar[index])
			if x == width-1 {
				fmt.Fprint(&buffer, "\n")
			}
		}
	}
	fmt.Println(buffer.String())
}
