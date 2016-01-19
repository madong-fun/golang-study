package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func read(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	slice := strings.Split(string(bytes), "\n")
	return slice

}

func main() {
	slice := read("/data/goworkspace/text.txt")
	var s string
	fmt.Println(len(slice))
	slices := make([]string, 0)
	for i := 0; i < len(slice); i++ {
		s = slice[i]
		ss := strings.Split(s, "\t")
		if len(ss) == 2 {
			s = "insert into apps.t_wiki_app_iframe (type_id,product_code,object_id,iframe_value,iframe_state) values(6,'doc','" + strings.TrimSpace(ss[0]) + "','" + strings.TrimSpace(ss[1]) + "',1);"
			slices = append(slices, s)

		}
	}
	sql := strings.Join(slices, "\n")

	fmt.Println(len(slices))

	ioutil.WriteFile("/data/goworkspace/iframe.sql", []byte(sql), 0666)

}
