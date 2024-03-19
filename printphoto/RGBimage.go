package printphoto

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/gookit/color"
	"github.com/nfnt/resize"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"log"
	"os"
)

func PrintRGBimage(Width uint, Height uint) {
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
	var buffer bytes.Buffer
	for y := range height {
		for x := range width {
			r, g, b, _ := img.At(x, y).RGBA()
			c := color.RGB(uint8(r>>8), uint8(g>>8), uint8(b>>8), true)
			c.Print(&buffer, " ")
			if x == width-1 {
				c.Println(&buffer, "")
			}
		}
	}
	fmt.Println(buffer.String())
}
