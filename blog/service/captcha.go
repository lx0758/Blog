package service

import (
	"blog/res"
	"bytes"
	"encoding/gob"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
	"math/rand"
	"strings"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

var (
	captchaRandom = rand.New(rand.NewSource(time.Now().UnixNano()))
	captchaChars  = []byte{
		'a', 'c', 'd', 'e', 'f', 'g',
		'h', 'j', 'k', 'm', 'n',
		'p', 'r', 's',
		'u', 'v', 'w', 'x', 'y', 'z',
		'A', 'B', 'C', 'D', 'E', 'F', 'G',
		'H', 'J', 'K', 'M', 'N',
		'P', 'Q', 'R', 'S', 'T',
		'U', 'V', 'W', 'X', 'Y', 'Z',
		'2', '3', '4', '5', '6', '7', '8', '9',
	}
	captchaLength     = 4
	captchaWidth      = 150
	captchaHeight     = 50
	captchaBgColor    = color.White
	captchaSessionKey = "captcha_type_"
	captchaFont       *truetype.Font
)

const (
	CAPTCHA_TYPE_LOGIN   = 1
	CAPTCHA_TYPE_COMMENT = 2
)

type CaptchaService struct {
	Service
}

type Captcha struct {
	Text  string
	Valid time.Time
}

func (s *CaptchaService) OnInitService() {
	gob.Register(&Captcha{})
	captchaFont, _ = freetype.ParseFont(res.FontFile)
}

func (s *CaptchaService) Generate(session sessions.Session, captchaType int, validDuration time.Duration) bytes.Buffer {
	key := fmt.Sprintf("%s%d", captchaSessionKey, captchaType)
	text := s.generateRandomText(captchaLength)
	session.Set(key, &Captcha{Text: text, Valid: time.Now().Add(validDuration)})
	_ = session.Save()
	return s.generateCaptchaImage(captchaWidth, captchaHeight, text)
}

func (s *CaptchaService) Verify(session sessions.Session, captchaType int, text string) bool {
	key := fmt.Sprintf("%s%d", captchaSessionKey, captchaType)
	captcha := session.Get(key).(*Captcha)
	if captcha == nil {
		return false
	}

	result := strings.ToLower(captcha.Text) == strings.ToLower(text) && captcha.Valid.After(time.Now())
	session.Delete(key)
	_ = session.Save()
	return result

}

func (s *CaptchaService) generateCaptchaImage(width int, height int, text string) bytes.Buffer {
	nRgba := image.NewNRGBA(image.Rect(0, 0, width, height))

	// 画背景
	draw.Draw(nRgba, nRgba.Bounds(), &image.Uniform{C: captchaBgColor}, image.ZP, draw.Src)
	// 画字符
	var x, y int
	context := freetype.NewContext()
	context.SetDPI(72)
	context.SetDst(nRgba)
	context.SetClip(nRgba.Bounds())
	context.SetFont(captchaFont)
	context.SetHinting(font.HintingFull)
	fontWidth := width / len(text)
	for i, ch := range text {
		fontSize := float64(height) * (0.85 + 0.15*captchaRandom.Float64())
		context.SetSrc(image.NewUniform(s.generateRandomColor(true)))
		context.SetFontSize(fontSize)
		x = fontWidth*i + captchaRandom.Intn(fontWidth/5)
		y = height/7*5 + captchaRandom.Intn(height/5)
		pt := freetype.Pt(x, y)
		_, _ = context.DrawString(string(ch), pt)
	}
	// 画干扰线(Bresenham算法)
	var x1, x2, y1, y2, dx, dy int
	var flag bool
	var lineColor color.RGBA
	for range (int)((float64(len(text))) * 1.5) {
		x1 = captchaRandom.Intn(width)
		y1 = captchaRandom.Intn(height)
		x2 = captchaRandom.Intn(width)
		y2 = captchaRandom.Intn(height)
		lineColor = s.generateRandomColor(false)
		dx, dy, flag = int(math.Abs(float64(x2-x1))),
			int(math.Abs(float64(y2-y1))),
			false
		if dy > dx {
			flag = true
			x1, y1 = y1, x1
			x2, y2 = y2, x2
			dx, dy = dy, dx
		}
		ix, iy := sign(x2-x1), sign(y2-y1)
		n2dy := dy * 2
		n2dydx := (dy - dx) * 2
		d := n2dy - dx
		for x1 != x2 {
			if d < 0 {
				d += n2dy
			} else {
				y1 += iy
				d += n2dydx
			}
			if flag {
				nRgba.Set(y1-1, x1, lineColor)
				nRgba.Set(y1, x1, lineColor)
				nRgba.Set(y1+1, x1, lineColor)
			} else {
				nRgba.Set(x1, y1-1, lineColor)
				nRgba.Set(x1, y1, lineColor)
				nRgba.Set(x1, y1+1, lineColor)
			}
			x1 += ix
		}
	}
	// 画噪点
	var pointColor color.RGBA
	for range len(text) * 10 {
		x = captchaRandom.Intn(width)
		y = captchaRandom.Intn(height)
		pointColor = s.generateRandomColor(false)
		nRgba.Set(x, y, pointColor)
		nRgba.Set(x+1, y, pointColor)
		nRgba.Set(x, y+1, pointColor)
		nRgba.Set(x+1, y+1, pointColor)
	}

	buffer := bytes.Buffer{}
	_ = png.Encode(&buffer, nRgba)
	return buffer
}

func (s *CaptchaService) generateRandomText(length int) string {
	var builder strings.Builder
	for range length {
		index := captchaRandom.Intn(len(captchaChars))
		builder.WriteByte(captchaChars[index])
	}
	return builder.String()
}

func (s *CaptchaService) generateRandomColor(dark bool) color.RGBA {
	threshold := 255
	if dark {
		threshold = 180
	}
	red := captchaRandom.Intn(threshold)
	green := captchaRandom.Intn(threshold)
	blue := captchaRandom.Intn(threshold)
	return color.RGBA{R: uint8(red), G: uint8(green), B: uint8(blue), A: uint8(255)}
}

func sign(x int) int {
	if x > 0 {
		return 1
	}
	return -1
}
