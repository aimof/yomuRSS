package view

import (
	"bytes"
	"io/ioutil"
	"strings"

	"github.com/aimof/yomuRSS/domain"
	"github.com/mattn/godown"
	"github.com/rivo/tview"
)

type View interface {
	AddArticles(a domain.Articles)
	Run() error
}

type view struct {
	flex     *tview.Flex
	list     *tview.List
	textview *tview.TextView
	app      *tview.Application
	articles domain.Articles
}

func NewView() *view {
	return &view{
		flex:     tview.NewFlex(),
		list:     tview.NewList(),
		textview: tview.NewTextView(),
		app:      tview.NewApplication(),
	}
}

func (v *view) AddArticles(articles domain.Articles) {
	v.articles = articles
	v.list.AddItem("quit", "Press to quit.", 'q', func() { v.app.Stop() })
	for _, a := range v.articles {
		v.list.AddItem(a.Title, a.PublishedAt, 's', func() {})
		v.list.SetSelectedFunc(func(i int, _ string, _ string, _ rune) {
			v.textview.Clear()
			if len(v.articles[i].Content) != 0 {
				b := make([]byte, 0, 10000)
				buf := bytes.NewBuffer(b)
				err := godown.Convert(buf, strings.NewReader(v.articles[i].Content), nil)
				if err != nil {
					v.textview.SetText(v.articles[i].Content)
					return
				}
				md, err := ioutil.ReadAll(buf)
				if err != nil {
					v.textview.SetText(v.articles[i].Content)
				}
				v.textview.SetText(string(md))
			} else {
				v.textview.SetText(v.articles[i].Description)
			}
		})
	}
}

func (v *view) Run() error {
	v.flex.AddItem(v.list, 0, 2, true).AddItem(v.textview, 0, 3, false)
	v.app.SetRoot(v.flex, true)
	if err := v.app.Run(); err != nil {
		return err
	}
	return nil
}
