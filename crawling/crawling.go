package main

import (
	"log"
	"context"
	"time"
	"fmt"
	"strings"
	"github.com/chromedp/chromedp"
	db "crawling/db"
)


func main() {
	db.SetDB()


	linklist := getLinkList()
	if len(linklist) <1{
		return
	}

	contextVar, cancelFunc := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancelFunc()
	contextVar, cancelFunc = context.WithTimeout(contextVar, 10000*time.Second)	// timeout 값을 설정 
	defer cancelFunc()

	for i := 0; i<len(linklist)-1; i++{

		getDescription(contextVar, linklist[i])
		
	}

}

func getDescription(contextVar context.Context, url string){
	log.Println(url)

	
	var strVar string
	var title string
	err := chromedp.Run(contextVar,		
		chromedp.Navigate("https://www.youtube.com"+url),
		chromedp.Click("#primary div#primary-inner div#below ytd-watch-metadata div#above-the-fold div#bottom-row div#description tp-yt-paper-button#expand-sizer", chromedp.ByID ),
		chromedp.Text("#primary div#primary-inner div#below ytd-watch-metadata div#above-the-fold div#bottom-row div#description div#description-inner ytd-text-inline-expander yt-formatted-string", &strVar,chromedp.ByID ),
		chromedp.Text("#primary div#primary-inner div#below ytd-watch-metadata div#above-the-fold div#title", &title,chromedp.ByID ),
		//chromedp.Text("#primary div#primary-inner div#below ytd-watch-metadata div#above-the-fold div#bottom-row div#description tp-yt-paper-button#expand-sizer", &attr,chromedp.ByQueryAll ),
	)
	if err != nil {
		panic(err)
	}
	strVar = strings.Replace(strVar, "\"", "\\\"", -1)
	title = strings.Replace(title, "\"", "\\\"", -1)
	param := make(map[string]string)
	param["url"] = url
	param["description"] = strVar
	param["title"] = title
	db.InsertBase(param)


}





func getLinkList() []string {

	contextVar, cancelFunc := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancelFunc()

	contextVar, cancelFunc = context.WithTimeout(contextVar, 10000*time.Second)	// timeout 값을 설정 
	defer cancelFunc()

	err := chromedp.Run(contextVar,		
		chromedp.Navigate("https://www.youtube.com/@paik_jongwon/videos"),
	)
	if err != nil {
		panic(err)
	}

	var oldHeight int
	var newHeight int
	for {

		err = chromedp.Run(contextVar,		
			chromedp.Evaluate(`window.scrollTo(0,document.querySelector("body ytd-app div#content").clientHeight); document.querySelector("body ytd-app div#content").clientHeight;`, &newHeight),
			chromedp.Sleep(700*time.Millisecond),
		)
		if err != nil {
			panic(err)
		}
		if(oldHeight == newHeight){
			break
		}
		oldHeight = newHeight
	}
	//var strVar string
	//var strTitle string
	attr := make([]map[string]string, 0)
	//var nodes []cdp.NodeID
	err = chromedp.Run(contextVar,		

		chromedp.AttributesAll("#primary ytd-rich-grid-renderer div#contents ytd-rich-grid-row div#contents ytd-rich-item-renderer #video-title-link", &attr,chromedp.ByQueryAll ),

	)
	if err != nil {
		panic(err)
	}

	var linklist []string
	for _, val := range attr {
		linklist = append(linklist, val["href"])
	}
	fmt.Println(len(linklist))
	return linklist

	
}
