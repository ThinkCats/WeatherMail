// Emails
package main

import (
	"fmt"
	process "lib"
)

func main() {
	fmt.Println("Hello World!")
	str := process.Get7DaysWeather()
	fmt.Println(str)
	email := process.NewEmail("798207330@qq.com","新消息来啦",str)
	process.SendEmail(email)
}
