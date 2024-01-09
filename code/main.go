package main

import (
	"fmt"
	"github.com/mojocn/base64Captcha"
	"image/color"
)

// 库中提供的默认内存存储器。
// 内存存储器用于存储和验证生成的验证码信息。
// 它将验证码的标识符、验证码图片和相关的验证数据存储在内存中。
var stores = base64Captcha.DefaultMemStore

func CaptMake() (id, b64s string, code string, err error) {
	var driver base64Captcha.Driver
	var driverString base64Captcha.DriverString

	// 配置验证码信息
	captchaConfig := base64Captcha.DriverString{
		Height:     60,
		Width:      200,
		NoiseCount: 0,
		// 控制显示在验证码图片中的线条的选项。在这个例子中，1: 直线  2: 曲线4: 点线8: 虚线16: 中空直线32: 中空曲线
		ShowLineOptions: 2 | 4,
		Length:          4,
		//验证码的字符源，用于生成验证码的字符。在这个例子中，使用数字和小写字母作为字符源。
		Source: "1234567890qwertyuioplkjhgfdsazxcvbnm",
		BgColor: &color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
		//用于绘制验证码文本的字体文件。在这个例子中，使用名为"wqy-microhei.ttc"的字体文件。
		Fonts: []string{"wqy-microhei.ttc"},
	}
	driverString = captchaConfig
	//将driverString中指定的字体文件转换为驱动程序所需的字体格式，并将结果赋值给driver变量。这个步骤是为了将字体文件转换为正确的格式，以便在生成验证码时使用正确的字体。
	driver = driverString.ConvertFonts()
	//使用driver和stores参数创建一个新的验证码实例，并将其赋值给captcha变量。这里的stores参数表示验证码存储器，用于存储和验证验证码。
	captcha := base64Captcha.NewCaptcha(driver, stores)
	//调用captcha实例的Generate方法生成验证码。lid是生成的验证码的唯一标识符，lb64s是生成的验证码图片的Base64编码字符串，lerr是生成过程中的任何错误。
	lid, lb64s, _, err := captcha.Generate()
	// 从验证码存储器中获取验证码的值
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
