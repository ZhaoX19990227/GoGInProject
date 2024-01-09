package main

import (
	"fmt"
	"github.com/mojocn/base64Captcha"
	"image/color"
)

var stores = base64Captcha.DefaultMemStore

func CaptMake() (id, b64s string, code string, err error) {
	var driver base64Captcha.Driver
	var driverString base64Captcha.DriverString

	// 配置验证码信息
	captchaConfig := base64Captcha.DriverString{
		Height:          60,
		Width:           200,
		NoiseCount:      0,
		ShowLineOptions: 2 | 4,
		Length:          4,
		Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
		BgColor: &color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
		Fonts: []string{"wqy-microhei.ttc"},
	}
	driverString = captchaConfig
	driver = driverString.ConvertFonts()
	captcha := base64Captcha.NewCaptcha(driver, stores)
	lid, lb64s, _, err := captcha.Generate()
	code = stores.Get(lid, true)
	fmt.Println(code)
	return lid, lb64s, code, err
}

// 这个解析方式只是上面介绍的内存存储器的方法调用，使用起来看情况是否使用
func CaptVerify(id string, capt string) bool {
	if stores.Verify(id, capt, false) {
		return true
	} else {
		return false
	}
}

func main() {
	captMake, s, s2, err := CaptMake()
	if err != nil {
		return
	}
	fmt.Println(captMake, s, s2, err)
}
