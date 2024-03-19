package main

import (
	"flag"
	"fmt"
	"main/printphoto"
	"main/video"
	"strconv"
)

type ResizeParams struct {
	Width  uint
	Height uint
}

func main() {
	resizeParams := ResizeParams{
		Width:  150,
		Height: 40,
	}
	v := flag.Bool("v", false, "版本信息")
	c := flag.Bool("c", false, "彩色输出")
	f := flag.String("f", "null", "视频文件")
	r := flag.Bool("r", false, "调整大小")
	flag.Parse()
	if *f != "null" {
		if *r {
			args := flag.Args()
			if len(args) != 2 {
				fmt.Println("⽤⼾输⼊错误")
				return
			}
			width, _ := strconv.Atoi(args[0])
			height, _ := strconv.Atoi(args[1])
			resizeParams.Width = uint(width)
			resizeParams.Height = uint(height)
		}
		if *v {
			fmt.Println("dian-player v1.2")
			return
		} else {
			if *c {
				video.Videoprintcolor(resizeParams.Width, resizeParams.Height, *f)
				return
			} else {
				video.Videoprintgray(resizeParams.Width, resizeParams.Height, *f)
				return
			}
		}
	} else {
		if *r {
			args := flag.Args()
			if len(args) != 2 {
				fmt.Println("⽤⼾输⼊错误")
				return
			}
			width, _ := strconv.Atoi(args[0])
			height, _ := strconv.Atoi(args[1])
			resizeParams.Width = uint(width)
			resizeParams.Height = uint(height)
		}
		if *v {
			fmt.Println("dian-player v1.2")
			return
		} else {
			if *c {
				printphoto.PrintRGBimage(resizeParams.Width, resizeParams.Height)
				return
			} else {
				printphoto.Printgrayscaleimage(resizeParams.Width, resizeParams.Height)
				return
			}
		}
	}

}
