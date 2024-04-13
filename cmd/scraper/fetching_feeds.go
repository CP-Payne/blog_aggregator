package scraper

import (
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/CP-Payne/blog_aggregator/internal/database"
	"github.com/CP-Payne/blog_aggregator/pkg/helper"
	"github.com/CP-Payne/blog_aggregator/pkg/models"
	"github.com/google/uuid"
)

type Rss struct {
	XMLName xml.Name `xml:"rss"`
	Channel struct {
		Item []struct {
			Title       string `xml:"title"`
			Link        string `xml:"link"`
			PubDate     string `xml:"pubDate"`
			Description string `xml:"description"`
		} `xml:"item"`
	} `xml:"channel"`
}

type Scraper struct {
	util       helper.Util
	DB         *database.Queries
	httpClient *http.Client
}

func NewScraper(util helper.Util, db *database.Queries) *Scraper {
	return &Scraper{
		util:       util,
		DB:         db,
		httpClient: &http.Client{},
	}
}

func (s *Scraper) StartScraper() {
	for {
		feeds := s.GetOldestFetchedFeeds(10)
		var wg sync.WaitGroup

		for _, feed := range feeds {
			wg.Add(1)
			go func(feedData models.Feed) {
				defer wg.Done()
				// Fetch rss data
				feedRssData, err := s.FetchDataFromRSS(feedData.Url)
				if err != nil {
					s.util.ErrorLog.Printf("Failed to fetch rss data for %s. FeedID: %v", feed.Name, feed.ID)
				}
				s.processRSSData(feedRssData)
				// TODO Return error from the above function when implementing database
				// If error is not nil then makr feed as fetched
				s.MarkFeedFetched(feed.ID)
			}(feed)
		}
		wg.Wait()
		time.Sleep(60 * time.Second)
	}
}

func (s *Scraper) processRSSData(rssData *Rss) {
	items := rssData.Channel.Item
	for _, item := range items {
		fmt.Printf("Item Title: %s\n", item.Title)
	}
}

func (s *Scraper) GetOldestFetchedFeeds(numToFetch int) []models.Feed {
	feeds, err := s.DB.GetNextFeedsToFetch(context.Background(), int32(numToFetch))
	if err != nil {
		s.util.ErrorLog.Print(err)
	}

	feedsConv := models.DatabaseFeedsToFeeds(feeds)

	return feedsConv
}

func (s *Scraper) MarkFeedFetched(feedId uuid.UUID) {
	err := s.DB.MarkFeedFetched(context.Background(), feedId)
	if err != nil {
		s.util.ErrorLog.Printf("Failed to mark feed as fetched: %v", err)
		return
	}
	fmt.Print("Feed marked as fetched.\n**********************\n")
}

func (s *Scraper) FetchDataFromRSS(url string) (*Rss, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		s.util.ErrorLog.Print("Failed to create request: ", err)
		return nil, err
	}

	resp, err := s.httpClient.Do(req)
	if err != nil {
		s.util.ErrorLog.Print("Failed to make request: ", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		s.util.ErrorLog.Fatalf("unexpected status code: %d", resp.StatusCode)
		return nil, errors.New("unexpected status code")
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		s.util.ErrorLog.Fatal(err)
		return nil, err
	}
	rssData := Rss{}

	err = xml.Unmarshal(data, &rssData)
	if err != nil {
		s.util.ErrorLog.Fatalf("xml.Unmarshal failed: %v", err)
		return nil, err
	}

	return &rssData, nil
}
