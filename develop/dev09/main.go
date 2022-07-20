package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	//Флаги для урла и глубины загрузки
	url := flag.String("url", "", "url")
	fullSite := flag.Bool("o", false, "download full site")

	flag.Parse()

	//Создаем путь для загрузки страницы
	pathArr := strings.Split(*url, "/")
	outputPath := "test/" + pathArr[len(pathArr)-1]

	//Ориентируясь на флаг смотрим, загружать ли сайт целиком
	if !*fullSite {
		err := downloadOnePage(outputPath, *url)
		if err != nil {
			log.Fatalln(err)
		}
		return
	}

	downloadAllPages(*url)
}

//Получаем массив всех ссылок сайта
func LinkParser(url string) []string {
	//Загружаем ссылку
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	//обрабатываем чтобы можно было делать поиск по тегам
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	//срез для результатов поиска ссылок
	var links []string

	//Ищем и записываем все ссылки на странице
	doc.Find("body a").Each(func(index int, item *goquery.Selection) {
		linkTag := item
		link, _ := linkTag.Attr("href")
		links = append(links, link)
	})
	return links
}

//Скачиваем все файлы с сайта
func downloadAllPages(url string) {
	//Получаем срез ссылок
	links := LinkParser(url)

	//Для каждой из них создаем свой файл
	for _, l := range links {
		fmt.Println(l)
		pathArr := strings.Split(l, "/")
		outputPath := "test/" + pathArr[len(pathArr)-1]
		if len(pathArr) > 2 {
			continue
		}

		resp, err := http.Get(url + l)
		if err != nil {
			fmt.Println("Ошибка скачивания страницы")

		}
		defer func() {
			err := resp.Body.Close()
			if err != nil {
				log.Fatalln(err)
			}
		}()

		f, err := os.Create(outputPath)
		if err != nil {
			fmt.Println("Ошибка создания файла")
		}
		defer func() {
			err := f.Close()
			if err != nil {
				log.Fatalln(err)
			}
		}()

		_, err = io.Copy(f, resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

	}
}

func downloadOnePage(filepath string, url string) error {

	//Загружаем данные страницы
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	//Создаем файл куда будем записывать данные страницы
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Записываем то что загрузили в файл
	_, err = io.Copy(out, resp.Body)
	return err
}
