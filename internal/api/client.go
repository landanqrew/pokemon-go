package api

import (
	"fmt"

	"github.com/landanqrew/pokemon-go/internal/pokecache"
	"github.com/landanqrew/pokemon-go/internal/web"
)


type Client struct {
	Cache *pokecache.Cache
	BaseURL string
}

func NewClient(cache *pokecache.Cache, baseURL string) *Client {
	return &Client{
		Cache: cache,
		BaseURL: baseURL,
	}
}

func (c *Client) GetResponse(relPath string) ([]byte, error) {
	url := fmt.Sprintf("%s%s", c.BaseURL, relPath)
	response, err := c.Cache.Get(url)
	if err != nil {
		response, err = web.GetResponseBytesBaseUrl(url)
		if err != nil {
			return nil, fmt.Errorf("error getting response: %v", err)
		}
	}
	return response, nil
}

