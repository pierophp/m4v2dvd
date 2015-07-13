package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func getTempDir() string {
	tmpDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return tmpDir + "/tmp/"
}

func main() {

	ffmpeg := new(FFMpeg)
	m4v2dvd := new(M4V2DVD)
	files := m4v2dvd.getFiles()

	total := len(files)
	fmt.Println(fmt.Sprintf("Arquivos Encontrados: %d", total))
	totalDuration := m4v2dvd.getTotalDuration()
	fmt.Println(totalDuration)
	for _, file := range files {
		fmt.Println("Convertendo: " + file)
		ffmpeg.toM2v(file)
	}
}
