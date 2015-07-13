package main

import (
	//"fmt"
	"os/exec"
	"strings"
)

type Command struct{}

func (c Command) exec(cmd string) (string, error) {

	cmd = strings.Replace(cmd, "\"", "", -1)
	parts := strings.Fields(cmd)

	head := parts[0]
	parts = parts[1:len(parts)]

	outByte, err := exec.Command(head, parts...).Output()

	out := string(outByte[:])

	//fmt.Println(out)
	//fmt.Println(err)

	return out, err
}
