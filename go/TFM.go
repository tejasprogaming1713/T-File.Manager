package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func listFiles() {
	files, _ := os.ReadDir(".")
	for _, f := range files {
		fmt.Println(f.Name())
	}
}

func download(url string) {
	name := url[strings.LastIndex(url, "/")+1:]
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Download error:", err)
		return
	}
	defer resp.Body.Close()

	file, err := os.Create(name)
	if err != nil {
		fmt.Println("File create error:", err)
		return
	}
	defer file.Close()

	io.Copy(file, resp.Body)
	fmt.Println("Downloaded:", name)
}

func zipFolder(folder string) {
	err := exec.Command("zip", "-r", folder+".zip", folder).Run()
	if err != nil {
		fmt.Println("Zip error:", err)
		return
	}
	fmt.Println("Zipped:", folder+".zip")
}

func unzipFile(file string) {
	err := exec.Command("unzip", file).Run()
	if err != nil {
		fmt.Println("Unzip error:", err)
		return
	}
	fmt.Println("Unzipped:", file)
}

func deleteFileOrFolder(target string) {
	err := os.RemoveAll(target)
	if err != nil {
		fmt.Println("Delete error:", err)
		return
	}
	fmt.Println("Deleted:", target)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: tfm <command> [args]")
		return
	}

	cmd := os.Args[1]

	switch cmd {
	case "list":
		listFiles()
	case "download":
		if len(os.Args) < 3 {
			fmt.Println("Usage: tfm download <url>")
			return
		}
		download(os.Args[2])
	case "zip":
		if len(os.Args) < 3 {
			fmt.Println("Usage: tfm zip <folder>")
			return
		}
		zipFolder(os.Args[2])
	case "unzip":
		if len(os.Args) < 3 {
			fmt.Println("Usage: tfm unzip <file.zip>")
			return
		}
		unzipFile(os.Args[2])
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: tfm delete <file/folder>")
			return
		}
		deleteFileOrFolder(os.Args[2])
	default:
		fmt.Println("Unknown command:", cmd)
		fmt.Println("Available commands: list, download, zip, unzip, delete")
	}
}
