package main

import (
	"../model"
	"encoding/json"
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
	err = json.Unmarshal(bytes, &tags)
	if err != nil {
		panic(err)
	}
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
