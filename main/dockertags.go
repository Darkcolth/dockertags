package main

import (
	"../model"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	client := http.Client{}
	resp, err := client.Get(cmdLineParse())
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	tags := model.Tags{}
	json.Unmarshal(bytes, &tags)
	if len(tags.Results) <= 0 {
		fmt.Println("No tags. Please confirm that the container' name is correct.")
		os.Exit(0)
	}

	var spaceName int
	var spaceSize int
	for _, tag := range tags.Results {
		if len(tag.Name) > spaceName {
			spaceName = len(tag.Name)
		}
		if len(tag.Size()) > spaceSize {
			spaceSize = len(tag.Size())
		}
	}

	fmt.Println(title(spaceName, spaceSize))
	for _, tag := range tags.Results {
		fmt.Println(tag.ToString(spaceName, spaceSize))
	}
}

func cmdLineParse() string {
	container := os.Args[1]
	size := flag.Int("s", 25, "page size")
	page := flag.Int("p", 1, "page number")
	flag.Parse()
	return fmt.Sprintf("https://registry.hub.docker.com/v2/repositories/%s/tags?page_size=%v&page=%v",
		container, *size, *page)
}

func title(spaceName int, spaceSize int) string {
	result := "name"
	for i := 0; i < spaceName; i++ {
		result += " "
	}
	result += "size"
	for i := 0; i < spaceSize; i++ {
		result += " "
	}
	return result + "update time"
}
