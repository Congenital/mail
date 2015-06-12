package mail

import (
	"fmt"
	"testing"
)

func TestSend(t *testing.T) {
	user := "qiang.sheng@godinsec.com"
	password := "xiang19930805"
	host := "smtp.exmail.qq.com:25"
	to := "qiang.sheng@godinsec.com"

	subject := "Godinsec User Register"

	body := `
								        <html>
										            <head><title></title></head>
													            <body>
																                <a href="http://www.baidu.com">www.baidu.com</a>
																				            </body>
																							        </html>
																									    `

	fmt.Println("send email")
	err := Send(user, password, host, to, subject, body, "html")
	if err != nil {
		fmt.Println("send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("send mail success!", err)
	}
}
