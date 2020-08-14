package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func cmdLineParse() string {
	set := flag.NewFlagSet("dockerTags", flag.ExitOnError)
	size := set.Int("s", 25, "page size")
	page := set.Int("p", 1, "page number, greater than zero")

	if len(os.Args) <= 1 {
		set.Usage()
		os.Exit(2)
	}

	container := &os.Args[1]

	if strings.Index(*container, "-") == 0 {
		set.Usage()
		os.Exit(2)
	}

	_ = set.Parse(os.Args[2:])
	return fmt.Sprintf("https://registry.hub.docker.com/v2/repositories/%v/tags?page_size=%v&page=%v",
		*container, *size, *page)
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
