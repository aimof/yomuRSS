package clowler

import (
	"os"
	"testing"

	"github.com/aimof/yomuRSS/domain"
)

func TestClowler(t *testing.T) {
	c := NewClowler(domain.TargetURL{})
	file0, err := os.Open("./test0.xml")
	if err != nil {
		t.Errorf("error: could not open test0.xml: %v", err)
	}
	feed0, err := c.parse(file0)
	file0.Close()
	if err != nil {
		t.Fatal("error: could not parse file.")
	}
	if feed0.Title != "blog.aimof.net feed" {
		t.Error("Title is wrong")
	}
	if len(feed0.Items) != 2 {
		t.Fatal("error: feed items must be 2")
	}
	if feed0.Items[0].Title != "gatsbyのprismjsにテーマを適用する" {
		t.Error("error: title is different")
	}

	c.addArticles(feed0)
	if len(c.articles) != 2 {
		t.Fatal("error: items must be 2")
	}

	file1, err := os.Open("./test1.xml")
	if err != nil {
		t.Errorf("error: could not open test1.xml: %v", err)
	}
	feed1, err := c.parse(file1)
	file1.Close()
	if err != nil {
		t.Fatal("error: could not parse file.")
	}

	c.addArticles(feed1)
	if len(c.articles) != 4 {
		t.Fatal("error: items must be 4")
	}
}
