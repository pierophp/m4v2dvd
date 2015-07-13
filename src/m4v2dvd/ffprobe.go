package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	//  "time"
)

type FFProbe struct{}

func (f FFProbe) getBinPath() string {

	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	dir = dir + "/bin/"

	if runtime.GOOS == "windows" {
		bin := dir + "ffprobe.exe"
		if _, err := os.Stat(bin); os.IsNotExist(err) {
			fmt.Println("ffprobe.exe not found")
		}

		return bin
	}

	bin := dir + "ffprobe"
	if _, err := os.Stat(bin); os.IsNotExist(err) {
		fmt.Println("ffprobe not found")
	}

	return bin
}

func (f FFProbe) getStreams(file string) []map[string]interface{} {

	//tmpName := getTempDir() + time.Now().Format("20060102150405") + ".txt"

	cmd := f.getBinPath() + " -v quiet -print_format json -show_streams " + file

	command := new(Command)
	out, _ := command.exec(cmd)

	var streams map[string][]map[string]interface{}
	json.Unmarshal([]byte(out), &streams)

	return streams["streams"]

}

func (f FFProbe) getDuration(file string) float64 {
	streams := f.getStreams(file)
	for _, stream := range streams {
		if !f.isVideo(stream["codec_name"].(string)) {
			continue
		}
		duration, _ := strconv.ParseFloat(stream["duration"].(string), 32)
		return duration
	}

	return 0
}

func (f FFProbe) isVideo(codec string) bool {
	videoCodecs := map[string]bool{
		"h264": true,
		"vp8":  true,
	}

	return videoCodecs[codec]
}
