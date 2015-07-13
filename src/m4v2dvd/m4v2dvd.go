package main

import (
	"io/ioutil"
	"strings"
)

type M4V2DVD struct{}

func (m M4V2DVD) getFiles() map[int]string {

	files, _ := ioutil.ReadDir("./")

	videoFiles := make(map[int]string)

	i := 0

	for _, f := range files {
		file := f.Name()
		extArr := strings.Split(file, ".")
		extension := extArr[len(extArr)-1]
		if extension != "m4v" {
			continue
		}

		videoFiles[i] = file
		i = i + 1
	}

	return videoFiles
}

func (m M4V2DVD) getTotalDuration() float64 {
	ffprobe := new(FFProbe)
	files := m.getFiles()
	duration := float64(0)
	for _, file := range files {
		duration += ffprobe.getDuration(file)
	}

	return duration
}
