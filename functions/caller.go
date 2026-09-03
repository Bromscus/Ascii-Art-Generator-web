package functions

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func CloneStandard() {
	fileURL := "https://learn.reboot01.com/git/root/public/raw/branch/master/subjects/ascii-art/standard.txt"

	_, err := os.Stat("formats/standard.txt")
	if err == nil {
		return
	}
	if !os.IsNotExist(err) {
		fmt.Println(err)
		return
	}

	resp, err := http.Get(fileURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	file, err := os.Create("formats/standard.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func CloneThinktertoy() {
	fileURL := "https://learn.reboot01.com/git/root/public/raw/branch/master/subjects/ascii-art/thinkertoy.txt"

	_, err := os.Stat("formats/thinktertoy.txt")
	if err == nil {
		return
	}
	if !os.IsNotExist(err) {
		fmt.Println(err)
		return
	}

	resp, err := http.Get(fileURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	file, err := os.Create("formats/thinkertoy.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func CloneShadow() {
	fileURL := "https://learn.reboot01.com/git/root/public/raw/branch/master/subjects/ascii-art/shadow.txt"

	_, err := os.Stat("formats/shadow.txt")
	if err == nil {
		return
	}
	if !os.IsNotExist(err) {
		fmt.Println(err)
		return
	}

	resp, err := http.Get(fileURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	file, err := os.Create("formats/shadow.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
}
