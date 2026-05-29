package feed

import (
	"errors"
	"net/url"

	"github.com/mmcdole/gofeed"
)

func validateFeedLink(link string) error {
	if link == "" {
		return errors.New("feed link is required")
	}

	parsed, err := url.ParseRequestURI(link)
	if err != nil {
		return errors.New("invalid feed link")
	}

	if parsed.Scheme != "http" && parsed.Scheme != "https" {
		return errors.New("feed link must use http or https")
	}

	fp := gofeed.NewParser()
	feedContent, err := fp.ParseURL(link)
	if err != nil {
		return errors.New("feed link is not a valid web feed")
	}

	if len(feedContent.Items) == 0 {
		return errors.New("feed has no items")
	}

	return nil
}
