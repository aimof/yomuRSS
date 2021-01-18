package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/aimof/yomuRSS/domain"
	"github.com/aimof/yomuRSS/interfaces/clowler"
	"github.com/aimof/yomuRSS/interfaces/view"
)

var yomudir = os.Getenv("YOMUDIR")

func main() {
	if yomudir == "" {
		fmt.Println("Please set an environment value YOMUDIR.")
		os.Exit(0)
	}
	if len(os.Args) < 2 {
		tui()
	} else if os.Args[1] == "get" {
		get()
	}
}

func get() {
	targetFile, err := os.Open(yomudir + "/target.txt")
	check(err, 1)
	s, err := ioutil.ReadAll(targetFile)
	check(err, 1)
	tu := strings.Split(string(s), "\n")
	c := clowler.NewClowler(tu)
	a, err := c.GetArticles()
	check(err, 1)
	file, err := os.Create(yomudir + "/articles/" + time.Now().Format("2006-01-02-15:04:05") + ".json")
	check(err, 1)
	defer file.Close()
	err = json.NewEncoder(file).Encode(a)
	check(err, 1)
}

func tui() {
	files, err := ioutil.ReadDir(yomudir + "/articles/")
	check(err, 1)
	if len(files) == 0 {
		log.Println("No article file.")
		os.Exit(1)
	}
	f, err := os.Open(yomudir + "/articles/" + files[len(files)-1].Name())
	check(err, 1)
	decoder := json.NewDecoder(f)
	a := new(domain.Articles)
	err = decoder.Decode(a)
	f.Close()
	check(err, 1)
	sort.Sort(a)
	v := view.NewView()
	v.AddArticles(*a)
	if err := v.Run(); err != nil {
		check(err, 1)
	}
}
