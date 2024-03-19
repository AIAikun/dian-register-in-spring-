package video

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/eiannone/keyboard"
	"github.com/nfnt/resize"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"log"
	"os"
	"os/exec"
	"time"
)

const (
	framerate = 60                                     // 帧率
	interval  = time.Second / time.Duration(framerate) //计算每帧的时间间隔
)

func Videoprintcolor(Width uint, Height uint, file string) {
	pauseCh := make(chan bool)
	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()
	go func() {
		for {
			char, _, err := keyboard.GetKey()
			if err != nil {
				panic(err)
				close(pauseCh)
				return
			}
			if char == '\x00' {
				pauseCh <- true // 如果处于暂停状态，则继续暂停
			}
		}
	}()
	for t := 0; ; t += 3 {
		select {
		case <-pauseCh:
			fmt.Println("Paused")
			<-pauseCh // 等待下一次暂停信号
			fmt.Println("Resumed")
		default:
		}
		start := time.Now()
		buffer0 := make(chan bool)
		buffer1 := make(chan bool)
		buffer2 := make(chan bool)
		buffer3 := make(chan bool)
		buffer4 := make(chan bool)
		buffer5 := make(chan bool)
		buffer6 := make(chan bool)
		buffer7 := make(chan bool)
		buffer8 := make(chan bool)
		buffer9 := make(chan bool)
		buf := bytes.NewBuffer(nil)
		ffmpeg.Input(file).
			Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", t)}).
			Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
			WithOutput(buf).
			Run()
		img, err := imaging.Decode(buf)
		buf.Reset()
		if err != nil {
			log.Fatal("生成视频失败：", err)
		}
		img = resize.Resize(Width, Height, img, resize.Lanczos3)
		bounds := img.Bounds()
		width, height := bounds.Dx(), bounds.Dy()
		var buffer [10]bytes.Buffer
		go func() {
			for y := 0; y < height/10; y++ {
				for x := range width {
					r, g, b, _ := img.At(x, y).RGBA()
					fmt.Fprintf(&(buffer[0]), "\x1b[48;2;%d;%d;%dm ", uint(r>>8), uint(g>>8), uint(b>>8))
				}
				fmt.Fprint(&(buffer[0]), "\x1b[0m\n")
			}
			buffer0 <- true
		}()
		go func() {
			for y := height / 10; y < height/5; y++ {
				for x := range width {
					r, g, b, _ := img.At(x, y).RGBA()
					fmt.Fprintf(&(buffer[1]), "\x1b[48;2;%d;%d;%dm ", uint(r>>8), uint(g>>8), uint(b>>8))
				}
				fmt.Fprint(&(buffer[1]), "\x1b[0m\n")
			}
			buffer1 <- true
		}()
		go func() {
			for y := height / 5; y < height*3/10; y++ {
				for x := range width {
					r, g, b, _ := img.At(x, y).RGBA()
					fmt.Fprintf(&(buffer[2]), "\x1b[48;2;%d;%d;%dm ", uint(r>>8), uint(g>>8), uint(b>>8))
				}
				fmt.Fprint(&(buffer[2]), "\x1b[0m\n")
			}
			buffer2 <- true
		}()
		go func() {
			for y := height * 3 / 10; y < height*2/5; y++ {
				for x := range width {
					r, g, b, _ := img.At(x, y).RGBA()
					fmt.Fprintf(&(buffer[3]), "\x1b[48;2;%d;%d;%dm ", uint(r>>8), uint(g>>8), uint(b>>8))
				}
				fmt.Fprint(&(buffer[3]), "\x1b[0m\n")
			}
			buffer3 <- true
		}()
		go func() {
			for y := height * 2 / 5; y < height/2; y++ {
				for x := range width {
					r, g, b, _ := img.At(x, y).RGBA()
					fmt.Fprintf(&(buffer[4]), "\x1b[48;2;%d;%d;%dm ", uint(r>>8), uint(g>>8), uint(b>>8))
				}
				fmt.Fprint(&(buffer[4]), "\x1b[0m\n")
			}
			buffer4 <- true
		}()
		go func() {
			for y := height / 2; y < height*3/5; y++ {
				for x := range width {
					r, g, b, _ := img.At(x, y).RGBA()
					fmt.Fprintf(&(buffer[5]), "\x1b[48;2;%d;%d;%dm ", uint(r>>8), uint(g>>8), uint(b>>8))
				}
				fmt.Fprint(&(buffer[5]), "\x1b[0m\n")
			}
			buffer5 <- true
		}()
		go func() {
			for y := height * 3 / 5; y < height*7/10; y++ {
				for x := range width {
					r, g, b, _ := img.At(x, y).RGBA()
					fmt.Fprintf(&(buffer[6]), "\x1b[48;2;%d;%d;%dm ", uint(r>>8), uint(g>>8), uint(b>>8))
				}
				fmt.Fprint(&(buffer[6]), "\x1b[0m\n")
			}
			buffer6 <- true
		}()
		go func() {
			for y := height * 7 / 10; y < height*4/5; y++ {
				for x := range width {
					r, g, b, _ := img.At(x, y).RGBA()
					fmt.Fprintf(&(buffer[7]), "\x1b[48;2;%d;%d;%dm ", uint(r>>8), uint(g>>8), uint(b>>8))
				}
				fmt.Fprint(&(buffer[7]), "\x1b[0m\n")
			}
			buffer7 <- true
		}()
		go func() {
			for y := height * 4 / 5; y < height*9/10; y++ {
				for x := range width {
					r, g, b, _ := img.At(x, y).RGBA()
					fmt.Fprintf(&(buffer[8]), "\x1b[48;2;%d;%d;%dm ", uint(r>>8), uint(g>>8), uint(b>>8))
				}
				fmt.Fprint(&(buffer[8]), "\x1b[0m\n")
			}
			buffer8 <- true
		}()
		go func() {
			for y := height * 9 / 10; y < height; y++ {
				for x := range width {
					r, g, b, _ := img.At(x, y).RGBA()
					fmt.Fprintf(&(buffer[9]), "\x1b[48;2;%d;%d;%dm ", uint(r>>8), uint(g>>8), uint(b>>8))
				}
				fmt.Fprint(&(buffer[9]), "\x1b[0m\n")
			}
			buffer9 <- true
		}()
		<-buffer0
		<-buffer1
		<-buffer2
		<-buffer3
		<-buffer4
		<-buffer5
		<-buffer6
		<-buffer7
		<-buffer8
		<-buffer9
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run() //清空终端字符
		for i := 0; i < 10; i++ {
			w := bufio.NewWriter(os.Stdout)
			w.WriteString(buffer[i].String())
			w.Flush()
		}
		end := time.Since(start)
		time.Sleep(3*interval - end)
	}
}
func Videoprintgray(Width uint, Height uint, file string) {
	pauseCh := make(chan bool)
	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()
	go func() {
		for {
			char, _, err := keyboard.GetKey()
			if err != nil {
				panic(err)
				close(pauseCh)
				return
			}
			if char == '\x00' {
				pauseCh <- true // 如果处于暂停状态，则继续暂停
			}
		}
	}()
	for t := 0; ; t += 3 {
		select {
		case <-pauseCh:
			fmt.Println("Paused")
			<-pauseCh // 等待下一次暂停信号
			fmt.Println("Resumed")
		default:
		}
		start := time.Now()
		buffer0 := make(chan bool)
		buffer1 := make(chan bool)
		buffer2 := make(chan bool)
		buffer3 := make(chan bool)
		buffer4 := make(chan bool)
		buffer5 := make(chan bool)
		buffer6 := make(chan bool)
		buffer7 := make(chan bool)
		buffer8 := make(chan bool)
		buffer9 := make(chan bool)
		buf := bytes.NewBuffer(nil)
		ffmpeg.Input(file).
			Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", t)}).
			Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
			WithOutput(buf).
			Run()
		img, err := imaging.Decode(buf)
		buf.Reset()
		if err != nil {
			log.Fatal("生成视频失败：", err)
		}
		img = resize.Resize(Width, Height, img, resize.Lanczos3)
		bounds := img.Bounds()
		width, height := bounds.Dx(), bounds.Dy()
		// 每个像素对应的字符
		pixelChar := []string{" ", ".", "-", ":", "*", "=", "+", "#", "%", "@"}
		var buffer [10]bytes.Buffer
		go func() {
			for y := 0; y < height/10; y++ {
				for x := range width {
					r, g, b, _ := img.At(x, y).RGBA()
					avg := (r*299 + g*587 + b*114) / 1000
					index := int(avg>>8) / 25
					if index > 9 {
						index = 9
					}
					fmt.Fprint(&(buffer[0]), pixelChar[index])
				}
				fmt.Fprint(&(buffer[0]), "\x1b[0m\n")
			}
			buffer0 <- true
		}()
		go func() {
			for y := height / 10; y < height/5; y++ {
				for x := range width {
					r, g, b, _ := img.At(x, y).RGBA()
					avg := (r*299 + g*587 + b*114) / 1000
					index := int(avg>>8) / 25
					if index > 9 {
						index = 9
					}
					fmt.Fprint(&(buffer[1]), pixelChar[index])
				}
				fmt.Fprint(&(buffer[1]), "\x1b[0m\n")
			}
			buffer1 <- true
		}()
		go func() {
			for y := height / 5; y < height*3/10; y++ {
				for x := range width {
					r, g, b, _ := img.At(x, y).RGBA()
					avg := (r*299 + g*587 + b*114) / 1000
					index := int(avg>>8) / 25
					if index > 9 {
						index = 9
					}
					fmt.Fprint(&(buffer[2]), pixelChar[index])
				}
				fmt.Fprint(&(buffer[2]), "\x1b[0m\n")
			}
			buffer2 <- true
		}()
		go func() {
			for y := height * 3 / 10; y < height*2/5; y++ {
				for x := range width {
					r, g, b, _ := img.At(x, y).RGBA()
					avg := (r*299 + g*587 + b*114) / 1000
					index := int(avg>>8) / 25
					if index > 9 {
						index = 9
					}
					fmt.Fprint(&(buffer[3]), pixelChar[index])
				}
				fmt.Fprint(&(buffer[3]), "\x1b[0m\n")
			}
			buffer3 <- true
		}()
		go func() {
			for y := height * 2 / 5; y < height*1/2; y++ {
				for x := range width {
					r, g, b, _ := img.At(x, y).RGBA()
					avg := (r*299 + g*587 + b*114) / 1000
					index := int(avg>>8) / 25
					if index > 9 {
						index = 9
					}
					fmt.Fprint(&(buffer[4]), pixelChar[index])
				}
				fmt.Fprint(&(buffer[4]), "\x1b[0m\n")
			}
			buffer4 <- true
		}()
		go func() {
			for y := height * 1 / 2; y < height*3/5; y++ {
				for x := range width {
					r, g, b, _ := img.At(x, y).RGBA()
					avg := (r*299 + g*587 + b*114) / 1000
					index := int(avg>>8) / 25
					if index > 9 {
						index = 9
					}
					fmt.Fprint(&(buffer[5]), pixelChar[index])
				}
				fmt.Fprint(&(buffer[5]), "\x1b[0m\n")
			}
			buffer5 <- true
		}()
		go func() {
			for y := height * 3 / 5; y < height*7/10; y++ {
				for x := range width {
					r, g, b, _ := img.At(x, y).RGBA()
					avg := (r*299 + g*587 + b*114) / 1000
					index := int(avg>>8) / 25
					if index > 9 {
						index = 9
					}
					fmt.Fprint(&(buffer[6]), pixelChar[index])
				}
				fmt.Fprint(&(buffer[6]), "\x1b[0m\n")
			}
			buffer6 <- true
		}()
		go func() {
			for y := height * 7 / 10; y < height*4/5; y++ {
				for x := range width {
					r, g, b, _ := img.At(x, y).RGBA()
					avg := (r*299 + g*587 + b*114) / 1000
					index := int(avg>>8) / 25
					if index > 9 {
						index = 9
					}
					fmt.Fprint(&(buffer[7]), pixelChar[index])
				}
				fmt.Fprint(&(buffer[7]), "\x1b[0m\n")
			}
			buffer7 <- true
		}()
		go func() {
			for y := height * 4 / 5; y < height*9/10; y++ {
				for x := range width {
					r, g, b, _ := img.At(x, y).RGBA()
					avg := (r*299 + g*587 + b*114) / 1000
					index := int(avg>>8) / 25
					if index > 9 {
						index = 9
					}
					fmt.Fprint(&(buffer[8]), pixelChar[index])
				}
				fmt.Fprint(&(buffer[8]), "\x1b[0m\n")
			}
			buffer8 <- true
		}()
		go func() {
			for y := height * 9 / 10; y < height; y++ {
				for x := range width {
					r, g, b, _ := img.At(x, y).RGBA()
					avg := (r*299 + g*587 + b*114) / 1000
					index := int(avg>>8) / 25
					if index > 9 {
						index = 9
					}
					fmt.Fprint(&(buffer[9]), pixelChar[index])
				}
				fmt.Fprint(&(buffer[9]), "\x1b[0m\n")
			}
			buffer9 <- true
		}()
		<-buffer0
		<-buffer1
		<-buffer2
		<-buffer3
		<-buffer4
		<-buffer5
		<-buffer6
		<-buffer7
		<-buffer8
		<-buffer9
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run() //清空终端字符
		for i := 0; i < 10; i++ {
			w := bufio.NewWriter(os.Stdout)
			w.WriteString(buffer[i].String())
			w.Flush()
		}
		end := time.Since(start)
		time.Sleep(3*interval - end)
	}
}
