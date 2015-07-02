
package lib

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func Get7DaysWeather()  string {
	doc,_ := goquery.NewDocument("http://m.weather.com.cn/mweather/101210101.shtml")
	//log.Println(doc.Html())
	s := doc.Find("div.days7")
	//fmt.Println(s.Html())
	var result  = []string{}
	s.Find("li").Each(func(i int, data *goquery.Selection){
		date := data.Find("b")
		result = append(result,date.Text(),"|")
		//fmt.Println("Date :",date.Text())
		data.Find("img").Each(func(j int,idata *goquery.Selection){
			T,_ := idata.Attr("alt")
			//fmt.Println("Weather:",T,j)
			if(j==0){
				result = append(result,T)
			}
			if(j==1){
				//fmt.Println("before data:",result[ len(result)-1])
				pre := result[ len(result)-1]
				if (pre == T){
					//fmt.Println("pre:",pre,"now:",T)
				}else{
					result = append(result,"转 ",T)
				}
			}
			
		})
		tmp := data.Find("span")
		//fmt.Println("tmp:",tmp.Text())
		result = append(result,"|温度:",tmp.Text(),"|")
	})
	//fmt.Println("list:",result)
	var str string
	for _,v := range result{
		str = str + v;
	}
	
	//fmt.Println(str)
	fin := strings.Replace(str,"|","\r\n",-1)
	fmt.Println(fin)
	return fin
//	for i ,_ := range result{
//		fmt.Println("content:",result[i])
//	}
}
