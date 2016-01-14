package main

import (
	"fmt"
	"github.com/golang/freetype"
	"image"
	//"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"os"
	"strings"
)

type Signer struct {
	fontPath string
	fontSize float64
	dpi      float64
}

func (this *Signer) drawStringImage(text string) (image.Image, error) {
	fontBytes, err := ioutil.ReadFile(this.fontPath)
	if err != nil {
		fmt.Println(err)
	}

	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		fmt.Println(err)
	}

	rgba := image.NewRGBA(image.Rect(0, 0, 900, 900))

	draw.Draw(rgba, rgba.Bounds(), image.Black, image.ZP, draw.Src)
	c := freetype.NewContext()
	c.SetDPI(this.dpi)
	c.SetFontSize(this.fontSize)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(image.White)
	c.SetFont(font)

	pt := freetype.Pt(100, 100+int(c.PointToFixed(this.fontSize)>>8))
	for _, s := range strings.Split(text, "\r\n") {
		_, err = c.DrawString(s, pt)
		pt.Y += c.PointToFixed(12 * 1.5)

	}

	return rgba, nil

}

func main() {
	text := "测试"
	s := Signer{}
	s.dpi = 80
	s.fontSize = 30
	s.fontPath = "/Users/baike/Downloads/kaishu.ttf"

	img, err := s.drawStringImage(text)
	if err != nil {
		fmt.Println(err)
	}

	file, _ := os.Create("test.png")

	err = png.Encode(file, img)
	if err != nil {
		fmt.Println(err)
	}

}
