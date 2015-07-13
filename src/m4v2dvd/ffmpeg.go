package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type FFMpeg struct{}

func (f FFMpeg) getBinPath() string {

	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	dir = dir + "/bin/"

	fmt.Println(dir)
	if runtime.GOOS == "windows" {
		bin := dir + "ffmpeg.exe"
		if _, err := os.Stat(bin); os.IsNotExist(err) {
			fmt.Println("ffmpeg.exe not found")
		}

		return bin
	}

	bin := dir + "ffmpeg"
	if _, err := os.Stat(bin); os.IsNotExist(err) {
		fmt.Println("ffmpeg not found")
	}

	return bin
}

func (f FFMpeg) toM2v(s_file string) {

	return

	t_file := strings.Replace(s_file, "m4v", "mpg", -1)
	bitrate := "2000"

	cmd := f.getBinPath() + " -i " + s_file + " -target ntsc-dvd -vcodec mpeg2video -b:v " + bitrate + "k -maxrate " + bitrate + "k -bufsize " + bitrate + "k -y " + t_file + ""

	command := new(Command)
	_, _ = command.exec(cmd)

	//fmt.Println(cmd)
	//fmt.Println(err)
}
