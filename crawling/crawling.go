package main

import (
	"log"
	"context"
	"time"
	"fmt"
	"strings"
	"sync"
	"github.com/chromedp/chromedp"
	db "crawling/db"
	//"github.com/chromedp/cdproto/cdp"
	//"github.com/chromedp/cdproto/runtime"
)


func main() {
	db.SetDB()


	linklist := getLinkList()
	if len(linklist) <1{
		return
	}
	var wg sync.WaitGroup
	remainder := len(linklist)
	forMax := 5
	for i := 0; i<len(linklist)-1; i++{
		log.Println(i)
		if i % forMax ==0{
			log.Print("==============")
		
			wg.Wait()
			log.Print("--------------")
			if remainder >= forMax {
				remainder -= forMax
				log.Printf("add")	
				log.Print(forMax)				
				wg.Add(forMax)
			}else{
				log.Printf("add2")	
				log.Print(remainder)	
				wg.Add(remainder)
			}
		}
	
		getDescription(linklist[i],&wg)
		
		
	
	}

	wg.Wait()



}

func getDescription(url string, wg *sync.WaitGroup){
	go func(){
		
		defer func(){
			log.Printf("Done")	
			wg.Done()
		}()
		contextVar, cancelFunc := chromedp.NewContext(
			context.Background(),
			chromedp.WithLogf(log.Printf),
		)
		defer cancelFunc()
		contextVar = context.WithValue(contextVar, url, url)
		contextVar, cancelFunc = context.WithTimeout(contextVar, 600*time.Second)	// timeout 값을 설정 
		defer cancelFunc()
		
		var strVar string
		err := chromedp.Run(contextVar,		
			chromedp.Navigate("https://www.youtube.com"+url),
			chromedp.Click("#primary div#primary-inner div#below ytd-watch-metadata div#above-the-fold div#bottom-row div#description tp-yt-paper-button#expand-sizer", chromedp.ByID ),
			chromedp.Text("#primary div#primary-inner div#below ytd-watch-metadata div#above-the-fold div#bottom-row div#description", &strVar,chromedp.ByID ),
			//chromedp.Text("#primary div#primary-inner div#below ytd-watch-metadata div#above-the-fold div#bottom-row div#description tp-yt-paper-button#expand-sizer", &attr,chromedp.ByQueryAll ),
		)
		if err != nil {
			panic(err)
		}
		strVar = strings.Replace(strVar, "\"", "\\\"", -1)
		param := make(map[string]string)
		param["url"] = url
		param["description"] = strVar
		db.InsertBase(param)

	}()

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
