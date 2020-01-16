package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

func main() {
	res, err := http.Get("https://www.58.com/changecity.html")

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", res.StatusCode)
		return
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	e := determineEncoding(bytes.NewReader((b)))
	utf8Reader := transform.NewReader(bytes.NewReader(b), e.NewDecoder())
	all, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}
	printCityList(all)
}

func printCityList(contents []byte) {
	provinceList := make([]string, 0)

	re := regexp.MustCompile(`provinceList = {([^}]*)`)
	matches := re.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		for _, subMatch := range m[1:] {
			str := strings.Replace(string(subMatch), " ", "", -1)
			str = strings.Replace(str, "\n", "", -1)
			for _, sub := range strings.FieldsFunc(str, splitByComma) {
				provinceList = append(provinceList, strings.Trim(strings.FieldsFunc(sub, splitBySemi)[0], `"`))
			}
		}
	}
	fmt.Println(provinceList)
	cityList := make(map[string]string)
	re = regexp.MustCompile(`independentCityList = {([^}]*)`)
	matches = re.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		for _, subMatch := range m[1:] {
			str := strings.Replace(string(subMatch), " ", "", -1)
			str = strings.Replace(str, "\n", "", -1)
			for _, sub := range strings.FieldsFunc(str, splitByComma) {
				independentCityStr := strings.FieldsFunc(sub, splitBySemi)
				independentCity := strings.Trim(independentCityStr[0], `"`)
				independentCityAb := strings.FieldsFunc(strings.Trim(independentCityStr[1], `"`), splitByVertical)[0]
				cityList[independentCity] = independentCityAb
			}
		}
	}

	for _, province := range provinceList {
		re = regexp.MustCompile(province + `":{([^}]*)`)
		matches = re.FindAllSubmatch(contents, -1)
		for _, m := range matches {
			for _, subMatch := range m[1:] {
				str := strings.Replace(string(subMatch), " ", "", -1)
				str = strings.Replace(str, "\n", "", -1)
				for _, sub := range strings.FieldsFunc(str, splitByComma) {
					cityStr := strings.FieldsFunc(sub, splitBySemi)
					city := strings.Trim(cityStr[0], `"`)
					cityAb := strings.FieldsFunc(strings.Trim(cityStr[1], `"`), splitByVertical)[0]
					cityList[city] = cityAb
				}
			}
		}
	}
	fmt.Println(cityList)
}

func splitByComma(s rune) bool {
	if s == ',' {
		return true
	}
	return false
}

func splitBySemi(s rune) bool {
	if s == ':' {
		return true
	}
	return false
}

func splitByVertical(s rune) bool {
	if s == '|' {
		return true
	}
	return false
}

func determineEncoding(r io.Reader) encoding.Encoding {
	bs, err := bufio.NewReader(r).Peek(1024)

	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bs, "")
	return e
}
