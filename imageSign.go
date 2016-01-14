package main

import (
	"fmt"
	"github.com/golang/freetype"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"os"
)

const (
	dx       = 100
	dy       = 40
	fontFile = "/Users/baike/Downloads/kaishu.ttf" // 需要使用的字体文件
	fontSize = 20                                  // 字体尺寸
	fontDPI  = 72                                  // 屏幕每英寸的分辨率
)

func main() {
	imgcounter := 123
	imgfile, _ := os.Create(fmt.Sprintf("%03d.png", imgcounter))

	defer imgfile.Close()

	img := image.NewNRGBA(image.Rect(0, 0, dx, dy))

	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			img.Set(x, y, color.RGBA{uint8(x), uint8(y), 0, 255})
		}
	}

	fontBytes, err := ioutil.ReadFile(fontFile) //读取字体数据
	if err != nil {
		fmt.Println(err)
		return
	}
	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		fmt.Print(err)
		return
	}
	c := freetype.NewContext()
	c.SetDPI(fontDPI)
	c.SetFont(font)
	c.SetFontSize(fontSize)
	c.SetClip(img.Bounds())
	c.SetDst(img)
	c.SetSrc(image.White)
	pt := freetype.Pt(10, 10+int(c.PointToFixed(fontSize)>>8)) // 字出现的位置

	_, err = c.DrawString("ABCDE", pt)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = png.Encode(imgfile, img)
	if err != nil {
		fmt.Println(err)
	}

}
