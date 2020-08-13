// Исследуйте работу последовательного и параллельного сканера веб-сайтов, задав большое (не менее 10) количество URL. Какие выводы можно сделать?

// Действительно, при параллельных запросах к сайтам общее время выполнения не превышает
// времени выполнения самого длительного запроса.
// При больших количествах параллельное выполнение запросов может существенно сократить
// общее время выполнения, причём, чем больше количество запросов, тем больше выигрыш.

// Какие практические варианты применения сканера веб-сайтов вы можете предложить?
// Примеры: агрегатор новостей, мониторинг цен на разных торговых площадках, автоматизация получения  и парсинга данных из различных интернет ресурсов
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var urls = []string{
	"https://ya.ru",
	"https://google.ru",
	"https://2ip.ru",
	"https://mail.ru",
	"https://yahoo.com",
	"https://apple.com",
	"https://kinopoisk.ru",
	"https://vk.com",
	"https://ok.ru",
	"https://golang.org",
}

func main() {
	sequence()
	parallel()
}

func sequence() {
	start := time.Now()
	for _, url := range urls {
		sequenceFetch(url)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func sequenceFetch(url string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	//bytes, err := ioutil.ReadAll(resp.Body)
	nbytes, err := io.Copy(ioutil.Discard, resp.Body) // to devNull
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%.2fs %7d %s\n", time.Since(start).Seconds(), nbytes, url)
}

func parallel() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range urls {
		go parallelFetch(url, ch)
	}
	for range urls {
		fmt.Print(<-ch) // receive from channel
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func parallelFetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body) // to devNull
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	ch <- fmt.Sprintf("%.2fs %7d %s\n", time.Since(start).Seconds(), nbytes, url)
}
