package common

import (
	"image/color"

	"github.com/afocus/captcha"
)

/*
* 定义一个json的返回值类型(首字母必须大写)
 */
type JsonIn struct {
	Status int
	Msg    string
	Data   map[string]interface{}
}

/**
* 生成验证码
* @width   int   验证码图片的宽度
* @height  int   验证码图片的高度
* @font   string   图片的引入的字体文件 ("xx.ttf")
* @text    string   验证码内容
* @frontColor  color.Color  验证码的字体颜色
* @bkgColor    color.Color  验证码图片的背景色
 */
func CreateCaptcha(width, height int, font, text string, frontColor, bkgColor color.Color) *captcha.Image {
	caps := captcha.New()
	// 可以设置多个字体 或使用cap.AddFont("xx.ttf")追加
	caps.SetFont(font)
	// 设置验证码大小
	caps.SetSize(width, height)
	// 设置干扰强度
	caps.SetDisturbance(captcha.MEDIUM)
	// 设置前景色 可以多个 随机替换文字颜色 默认黑色
	caps.SetFrontColor(frontColor)
	// 设置背景色 可以多个 随机替换背景色 默认白色
	caps.SetBkgColor(bkgColor)
	img := caps.CreateCustom(text)
	return img
}
