package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/aimof/yomuRSS/domain"
	"github.com/aimof/yomuRSS/interfaces/clowler"
	"github.com/aimof/yomuRSS/interfaces/view"
)

var yomudir = os.Getenv("YOMUDIR")

func main() {
	if len(os.Args) < 2 {
		tui()
	} else if os.Args[1] == "get" {
		get()
	}
}

func get() {
	targetFile, err := os.Open(yomudir + "/target.txt")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	s, err := ioutil.ReadAll(targetFile)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	tu := strings.Split(string(s), "\n")
	c := clowler.NewClowler(tu)
	a, err := c.GetArticles()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	file, err := os.Create(yomudir + "/articles/" + time.Now().Format("2006-01-02-15:04:05") + ".json")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	err = json.NewEncoder(file).Encode(a)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func tui() {
	files, err := ioutil.ReadDir(yomudir + "/articles/")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	if len(files) == 0 {
		log.Println("No article file.")
		os.Exit(1)
	}
	f, err := os.Open("./articles/" + files[len(files)-1].Name())
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	decoder := json.NewDecoder(f)
	a := new(domain.Articles)
	err = decoder.Decode(a)
	f.Close()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	v := view.NewView()
	v.AddArticles(*a)
	if err := v.Run(); err != nil {
		panic(err)
	}
}
