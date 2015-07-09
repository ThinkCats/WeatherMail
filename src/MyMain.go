// Emails
package main

import (
	"fmt"
	process "lib"
	"github.com/robfig/cron"
)

func doSend() {
	str := process.Get7DaysWeather()
	fmt.Println(str)
	email := process.NewEmail("798207330@qq.com","新消息来啦",str)
	process.SendEmail(email)
}

func main(){
	fmt.Println("Begin Send")
	c := cron.New()
	spec := "0 2 13 * * ?"
	c.AddFunc(spec,doSend)
	c.Start()
	select{}
}
