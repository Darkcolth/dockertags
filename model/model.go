package model

import (
	"fmt"
	"time"
)

var UNIT = [6]string{"B", "KB", "MB", "GB", "TB", "PB"}

type Tag struct {
	Name        string    `json:"name"`
	FullSize    int       `json:"full_size"`
	LastUpdated time.Time `json:"last_updated"`
	sizeStr     string
}

type Tags struct {
	Count   int   `json:"count"`
	Results []Tag `json:"results"`
}

func (tag *Tag) Size() string {
	if tag.sizeStr == "" {
		var size = float32(tag.FullSize)
		var i = 0
		for ; size >= 1024; i++ {
			size /= 1024
		}
		tag.sizeStr = fmt.Sprintf("%.2f%v", size, UNIT[i])
	}
	return tag.sizeStr
}

func (tag *Tag) LastUpdatedTime() string {
	return tag.LastUpdated.Format("2006-01-02 15:04:05")
}

func (tag *Tag) ToString(spaceName int, spaceSize int) string {
	result := tag.Name
	time1 := spaceName - len(tag.Name) + 4
	for i := 0; i < time1; i++ {
		result += " "
	}
	result += tag.Size()
	time2 := spaceSize - len(tag.Size()) + 4
	for i := 0; i < time2; i++ {
		result += " "
	}
	return result + tag.LastUpdatedTime()
}
