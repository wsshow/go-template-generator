package main

import (
	"embed"
	"log"
	"os"
	"strings"
)

//go:embed templates
var templates embed.FS

func InitLogConfig() {
	log.SetPrefix("[go-template-generator] ")
	log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds)
}

func IsPathExist(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil
}

func CreatDir(dirPath string) error {
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return err
	}
	err = os.Chmod(dirPath, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func Search(dirPath string) {
	dirs, err := templates.ReadDir(dirPath)
	if err != nil {
		log.Fatalln(err)
	}
	for _, dir := range dirs {
		if dir.IsDir() {
			if !IsPathExist(dir.Name()) {
				err = CreatDir(dir.Name())
				if err != nil {
					log.Fatalln(err)
				}
			}
			Search(dirPath + "/" + dir.Name())
		} else {
			fullpath := dirPath + "/" + dir.Name()
			wpath := strings.TrimSuffix(fullpath[10:], "t")
			if !IsPathExist(wpath) {
				bs, err := templates.ReadFile(fullpath)
				if err != nil {
					log.Fatalln(err)
				}
				err = os.WriteFile(wpath, bs, os.ModePerm)
				if err != nil {
					log.Fatalln(err)
				}
			} else {
				log.Println(wpath, "已存在，跳过")
			}
		}
	}
}

func main() {
	InitLogConfig()
	Search("templates")
}
