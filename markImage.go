package main

import (
	"fmt"
	"github.com/golang/freetype"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"strings"
)

type Signer struct {
	fontPath string
	dpi      float64
	fontSize float64
	dstPath  string
	srcImage string
}

func (this *Signer) mark(text string) (image.Image, error) {
	fontBytes, err := ioutil.ReadFile(this.fontPath)
	if err != nil {
		fmt.Println(err)
	}
	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		fmt.Println(err)
	}

	//读取图片
	file, err := os.Open(this.srcImage)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	img, err := jpeg.Decode(file)
	if err != nil {
		fmt.Println(err)
	}
	bounds := img.Bounds()
	nrgba := image.NewNRGBA(bounds)
	draw.Draw(nrgba, img.Bounds(), img, image.ZP, draw.Over)
	c := freetype.NewContext()
	c.SetDPI(this.dpi)
	c.SetFont(font)
	c.SetDst(nrgba)
	c.SetFontSize(this.fontSize)
	c.SetClip(nrgba.Bounds())
	c.SetSrc(image.White)

	pt := freetype.Pt(100, 100+int(c.PointToFixed(this.fontSize)>>8))
	for _, s := range strings.Split(text, "\r\n") {
		_, err = c.DrawString(s, pt)
		pt.Y += c.PointToFixed(12 * 1.5)

	}

	return nrgba, nil

}

func main() {
	text := "刘楚恬"
	s := Signer{fontPath: "/Users/baike/Downloads/kaishu.ttf",
		dpi:      90,
		fontSize: 20,
		dstPath:  "lct.png",
		srcImage: "/Users/baike/Downloads/lct.jpg"}

	img, err := s.mark(text)
	if err != nil {
		fmt.Println(err)
	}
	file, _ := os.Create(s.dstPath)
	err = png.Encode(file, img)
	if err != nil {
		fmt.Println(err)
	}
}
