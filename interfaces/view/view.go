package view

import (
	"github.com/aimof/yomuRSS/domain"
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
	v.list.AddItem("quit", "Press to quit.", 'q', func() { v.app.Stop() })
	for _, a := range articles {
		v.list.AddItem(a.Title, a.PublishedAt, 's', func() {
			v.textview.Clear()
			if len(a.Content) != 0 {
				v.textview.SetText(a.Content)
			} else {
				v.textview.SetText(a.Description)
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
