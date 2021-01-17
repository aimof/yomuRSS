package domain

// Article is an interface of an article
type Article struct {
	ID          string `json:"id"`
	PublishedAt string `json:"published"`
	Link        string `json:"link"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
}

// Articles is a collection of Articles
type Articles []Article

// Len in sort.
func (a Articles) Len() int { return len(a) }

// Less in sort.
func (a Articles) Less(i, j int) bool { return a[i].PublishedAt < a[j].PublishedAt }

// Swap in sort.
func (a Articles) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

type TargetURL []string
