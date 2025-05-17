package hn

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

const HN_URL = "https://hacker-news.firebaseio.com/v0/"

type client struct {
	url string
}

type Story struct {
	By          string `json:"by"`
	Descendants int    `json:"descendants"`
	Id          int    `json:"id"`
	Kids        []int  `json:"kids"`
	Score       int    `json:"score"`
	Time        int    `json:"time"`
	Title       string `json:"title"`
	Url         string `json:"url"`
}

func NewClient() *client {
	return &client{
		url: HN_URL,
	}
}

func (c *client) TopStories(number int) ([]Story, error) {
	var result []int

	resp, err := c.getResource(c.url + "topstories.json")
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(resp) > number {
		result = result[:number]
	}

	return c.getStories(result[:number])
}

func (c *client) NewStories(number int) ([]Story, error) {
	var result []int

	resp, err := c.getResource(c.url + "newstories.json")
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(resp) > number {
		result = result[:number]
	}

	return c.getStories(result[:number])
}

func (c *client) getStories(ids []int) ([]Story, error) {
	stories := make([]Story, len(ids))
	var errSlice []error
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i, id := range ids {
		wg.Add(1)
		go func(i, id int) {
			defer wg.Done()
			story, err := c.getStory(id)
			if err != nil {
				mu.Lock()
				errSlice = append(errSlice, fmt.Errorf("failed to fetch story %d: %w", id, err))
				mu.Unlock()
				return
			}
			stories[i] = story
		}(i, id)
	}

	wg.Wait()

	if len(errSlice) > 0 {
		return nil, fmt.Errorf("multiple errors occurred: %v", errSlice)
	}

	return stories, nil
}

func (c *client) getStory(id int) (Story, error) {
	var s Story

	url := fmt.Sprintf("%s/item/%d.json", c.url, id)
	resp, err := c.getResource(url)
	if err != nil {
		return s, err
	}

	err = json.Unmarshal(resp, &s)
	return s, err
}

func (c *client) getResource(url string) ([]byte, error) {
	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Close = true

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return response, nil
}
