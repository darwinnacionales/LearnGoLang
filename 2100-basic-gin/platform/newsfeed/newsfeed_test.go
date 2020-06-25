package newsfeed

import "testing"

func TestAdd(t *testing.T) {
	feed := New()
	feed.Add(Item{"An Item", "with Body"})

	if len(feed.Items) != 1 {
		t.Errorf("Item was not added")
	}
}

func TestGetAll(t *testing.T) {
	feed := New()
	feed.Add(Item{})
	results := feed.GetAll()

	if len(results) != 1 {
		t.Errorf("Item was not added")
	}
}

/*
	To run in console:
		- go test ./platform/newsfeed
		- go test ./...
		- go test -cover ./...
*/
