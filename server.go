package main

import (
	"fmt"
	"net/http"
	"log"
	"os"
	"io/ioutil"
	"encoding/json"
	"encoding/base64"
	"golang.org/x/net/html"
	"strings"
	"bytes"
	"io"
	"text/template"
	"github.com/ConvertAPI/convertapi-go"
	"github.com/ConvertAPI/convertapi-go/config"
	"github.com/ConvertAPI/convertapi-go/param"
	"strconv"
)

type Data struct {
	Url string //`json:"url"`
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Listening on port 9602 ...")
	logError(http.ListenAndServe(":9602", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	err := os.Mkdir("served-files", 0777)
	logError(err)
	urldata := Data{}
	rbody, err := ioutil.ReadAll(r.Body)
	logError(err)
	err = json.Unmarshal(rbody, &urldata)
	logError(err)
	writePage(urldata.Url)
	_, err = os.Stat("served-files/new-nairaland-page.pdf")
	if os.IsNotExist(err) {
		fmt.Println("Foreign API couldn't generate PDF file, and a default file generated from 'https://www.nairaland.com/6602040/how-successful-people-achieve-career' is returned.")
		fmt.Println("To get desired results, please try for threads with lower number of posts.")
		nrlnd_page, err := ioutil.ReadFile("default-served-files/new-nairaland-page.pdf")
		logError(err)
		page1_b64 := base64.StdEncoding.EncodeToString([]byte(nrlnd_page))
		w.Write([]byte(page1_b64))
		err = os.RemoveAll("served-files")
		logError(err)
	} else {
		fmt.Println("Foreign API successfully generated PDF file.")
		nrlnd_page, err := ioutil.ReadFile("served-files/new-nairaland-page.pdf")
		logError(err)
		page1_b64 := base64.StdEncoding.EncodeToString([]byte(nrlnd_page))
		fmt.Println("Writing PDF data to the connection...")
		w.Write([]byte(page1_b64))
		err = os.RemoveAll("served-files")
		logError(err)
	}
	fmt.Println("***************************************************************************************")
  fmt.Println("                                   SESSION ENDS")
	fmt.Println("***************************************************************************************")
	fmt.Println("Listening on port 9602 ...")
}

func logError(err error) {
	if err != nil {
		log.Fatal("Error encountered:- ", err)
	}
}

var webpage = template.Must(template.New("webpage").Parse(`
<style type="text/css">
  *, *::after, *::before {
    box-sizing: border-box;
  }

  html {
    font-size: 62.5%;
  }

  body {
    font-size: 1.6rem;
    font-family: 'Verdana'!important;
    color: #25283D;
    background-color: #f1f7f9;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
  }

  a {
    color: #428bca;
    text-decoration: none;
  }

 .comment{
   margin-bottom: 20px;
   position: relative;
   z-index: 0;
 }

 .comment .comment-avatar{
   border: 2px solid #fff;
   border-radius: 50%;
   box-shadow: 0 1px 2px rgba(0, 0, 0, .2);
   height: 40px;
   left: 0;
   overflow: hidden;
   position: absolute;
   top: 0;
   width: 40px;
 }

 .comment .comment-avatar img{
   display: block;
   height: 100%;
   width: 100%;
 }

 .comment .comment-box{
   background-color: #fcfcfc;
   border-radius: 4px;
   box-shadow: 0 1px 1px rgba(0, 0, 0, .50);
   margin-left: 50px;
   min-height: 60px;
   position: relative;
   padding: 5px;
   padding-bottom: 5px;
 }

  .comment .comment-main{
   background-color: #fcfcfc;
   border-radius: 4px;
   box-shadow: 0 1px 1px rgba(0, 0, 0, .15);
   margin-left: 0px;
   position: relative;
   padding: 5px;
 }

  .comment blockquote{
   background-color: #dfe5ed;
   border-radius: 4px;
   box-shadow: 0 1px 1px rgba(0, 0, 0, .50);
   min-height: 60px;
   position: relative;
   padding: 5px;
 }

 .comment .comment-box:before,
 .comment .comment-box:after{
   border-width: 10px 10px 10px 0;
   border-style: solid;
   border-color: transparent #FCFCFC;
   content: "";
   left: -10px;
   position: absolute;
   top: 20px;
 }


 .comment .comment-text{
   color: #555f77;
   font-size: 12px;
   padding-left:10px;
   padding-top: 5px;
   padding-bottom: 15px;
 }

 .comment .comment-footer{
   font-size: 12px;
   padding: 5px;
   background-color: #dfe5ed;
   color: #555f77;
 }

  .comment .comment-header{
    line-height:20px;
    font-weight:700;
    font-size:12px;
    padding-bottom:8px;
    padding-top:8px;
    padding-left:10px;
    background-color: #dfe5ed;
 }

  .comment-main .comment-header{
    line-height:50px;
    font-weight:1500;
    font-size:15px;
    padding:8px;
    background-color: #dfe5ed;
 }

 .comment .comment-footer:after{
   content: "";
   display: table;
   clear: both;
 }

 #pstby{
 	font-weight: normal;
 }
</style>

<div class="comment">
  <div class="comment-main">
    <div class="comment-header"><a href="https://htmlcolors.com/user/Antonios" style="color:#428bca">{{.Topic}}<span id="pstby"> posted by </span>{{.Main.Name}}</a></div>
    <div style="line-height:20px;white-space: pre-wrap;" class="comment-text">{{.Main.Comment}}</div>
    <div class="comment-footer">
      <div class="comment-info">
        <span style="line-height:18px" class="comment-date"><span id="pstby">Posted </span>{{.Time}}</span>
      </div>
    </div>
  </div>
</div>

{{range .Posts}}
<div class="comment">
  <div class="comment-avatar">
  	<a href="https://htmlcolors.com/user/Antonios"><img class="mobilecommentimage" style="object-fit:cover" src="https://i.ibb.co/kgHjcs6/change-user.png" border="0"></a>
  </div>

  <div class="comment-box">
    <div class="comment-header"><a href="https://htmlcolors.com/user/Antonios" style="color:#428bca">{{.Name}}</a></div>
    <div style="line-height:20px;white-space: pre-wrap;" class="comment-text">{{.Comment}}</div>
  </div>
</div>
{{end}}

`))

var divs []Post
var ind = 0
var cnt, x, y, z = 1, 1, 1, 1
var elmntcnt = 0
var pageposts Webpage

type Post struct {
	Comment string
	Name string
	//Time string
	//Date string
	//Likes string
}

type Webpage struct {
	Posts []Post
	Main Post
	Topic string
	Time string
	Date string
}

func writePage(pagelink string) {
	var link string
	var buf bytes.Buffer
  w := io.Writer(&buf)
  var threads = []string{"https://www.nairaland.com/6602040/how-successful-people-achieve-career",
  											 "https://www.nairaland.com/7229653/court-orders-upward-review-judges",
												 "https://www.nairaland.com/7106890/software-engineering-journey-alx-africa"}
  fmt.Println("***************************************************************************************")
  fmt.Println("                                  SESSION BEGINS")
  fmt.Println("***************************************************************************************")
	link = pagelink
	if link == "" {
		link = threads[0]
		fmt.Println("The default link", link, "was entered.")
	} else {
		fmt.Println("The link", link, "was entered.")
	}
	fmt.Println("***************************************************************************************")
	fmt.Println("Retrieving web pages...")
	fmt.Println("***************************************************************************************")
	link0 := link
	var page *http.Response
	var pageTrack *http.Response
	var err error
	var pages []*http.Response
	x := 1
	for {
		page, err = http.Get(link)
		logError(err)
		if pageTrack == nil {
			// do nothing
		} else if page.Request.URL.Path == pageTrack.Request.URL.Path || x == 10000 {
			break
		}
		fmt.Printf("The HTTP request for https://www.nairaland.com%s was successful\n", page.Request.URL.Path)
		fmt.Println("***************************************************************************************")
		link = link0 + "/" + strconv.Itoa(x)
		pages = append(pages, page)
		pageTrack = page
		x++
	}
	for i, wpage := range pages {
		pagetext, err := ioutil.ReadAll(wpage.Body)
		logError(err)
		fmt.Printf("The content of the webpage at https://www.nairaland.com%s was successfully retrieved\n", wpage.Request.URL.Path)
		err = ioutil.WriteFile("served-files/webpage.html", pagetext, 0644)
		logError(err)
		text := string(pagetext)
		doc, err := html.Parse(strings.NewReader(text))
		logError(err)
		procNode(doc)
		fmt.Println("[The retrieved webpage indexed -", i+1, "was successfully parsed]")
		fmt.Println("***************************************************************************************")
		wpage.Body.Close()
	}
	page.Body.Close()
	pageTrack.Body.Close()
	divs = cleanDivs(divs)
	if len(divs) > 1 {
		pageposts.Posts = divs[1: ]
	} else if len(divs) == 0 {
		log.Fatal("No posts in the thread!")
	}
	pageposts.Main = divs[0]
	err = webpage.Execute(w, pageposts)
	logError(err)
	err = ioutil.WriteFile("served-files/new-nairaland-page.html", []byte(buf.String()), 0644)
	logError(err)
	divs = []Post{}; ind = 0; cnt, x, y, z = 1, 1, 1, 1; elmntcnt = 0; pageposts = Webpage{}
	//fmt.Println("An html version of the new webpage was saved as 'new-nairaland-page.html' in the directory 'served-files'")
	//fmt.Println("***************************************************************************************")
	fmt.Println("Generating PDF version (this may take a while)...")
	fmt.Println("***************************************************************************************")
	//config.Default.Secret = "Uhy0MidCpF8ZmoUT"
	config.Default.Secret = "lcyiBRQKAkSUEMFD"
	convertapi.ConvDef("html", "pdf",
		param.NewPath("File", "served-files/new-nairaland-page.html", nil)).ToPath("served-files/new-nairaland-page.pdf")
	//fmt.Println("The PDF presentation of the webpage was successfully generated, and saved as 'new-nairaland-page.pdf' in the directory 'served-files'")
}

func renderNode(n *html.Node) string {
  var buf bytes.Buffer
  w := io.Writer(&buf)
  html.Render(w, n)
  return buf.String()
}

func procNode(node *html.Node) {
	if node.Type == html.ElementNode  && node.Data == "b" {
		elmntcnt++
	}
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, val := range node.Attr {
			if val.Key == "class" && val.Val == "user" {
				divs = append(divs, Post{"", ""})
				divs[ind].Name = renderNode(node)
			}
		}
	} else if node.Type == html.ElementNode && node.Data == "div" {
		if node.Attr[0].Key == "class" && node.Attr[0].Val == "narrow" {
			divs = append(divs, Post{"", ""})
			divs[ind].Comment = renderNode(node)
			ind = ind + 1
		}
	} else if elmntcnt == 8 && x == cnt {
		pageposts.Time = renderNode(node)
		x++
	} else if elmntcnt == 9 && y == cnt {
		pageposts.Date = renderNode(node)
		y++
	} else if node.Type == html.ElementNode && node.Data == "title" && z == cnt {
		pageposts.Topic = renderNode(node.FirstChild)
		z++
	}
	for i := node.FirstChild; i != nil; i = i.NextSibling {
		procNode(i)
	}
}

func cleanDivs(divs []Post) []Post {
	pst := Post{}
	for {
		if divs[len(divs)-1] == pst {
			divs = divs[ :len(divs)-1]
		} else {
			break
		}
	}
	return divs
}

















