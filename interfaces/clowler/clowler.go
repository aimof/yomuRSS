package clowler

import (
	"io"
	"log"
	"net/http"

	"github.com/aimof/yomuRSS/domain"
	"github.com/mmcdole/gofeed"
)

type Clowler interface {
	GetArticles() (domain.Articles, error)
}

type clowler struct {
	targetURL domain.TargetURL
	parser    *gofeed.Parser
	articles  domain.Articles
}

func NewClowler(tu domain.TargetURL) *clowler {
	return &clowler{
		targetURL: tu,
		parser:    gofeed.NewParser(),
		articles:  make(domain.Articles, 0, 10*len(tu)),
	}
}

func (c *clowler) GetArticles() (domain.Articles, error) {
	for _, url := range c.targetURL {
		resp, err := http.Get(url)
		if err != nil {
			log.Printf("Could not read feed: %s", url)
			continue
		}
		feed, err := c.parse(resp.Body)
		resp.Body.Close()
		if err != nil {
			log.Printf("Could not parse feed: %s", url)
			continue
		}
		c.addArticles(feed)
	}
	return c.articles, nil
}

func (c *clowler) parse(r io.Reader) (*gofeed.Feed, error) {
	feed, err := c.parser.Parse(r)
	if err != nil {
		return nil, err
	}
	return feed, nil
}

func (c *clowler) addArticles(f *gofeed.Feed) {
	for _, item := range f.Items {
		c.articles = append(c.articles, domain.Article{
			ID:          item.GUID,
			Title:       item.Title,
			Description: item.Description,
			Content:     item.Content,
			Link:        item.Link,
			PublishedAt: item.PublishedParsed.Format("2006-01-02-15-04:05"),
		})
	}
}
