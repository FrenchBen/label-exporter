package exporter

import (
	"context"
	"errors"
	"net/url"
	"os"

	"github.com/google/go-github/v28/github"
	"golang.org/x/oauth2"
)

// Client is the struct wrapper for a github client
type Client struct {
	client *github.Client
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// NewClient provides a new github client based on the provided env vars
func NewClient() (*Client, error) {
	gitURL := getEnv("GITHUB_BASEURL", "https://api.github.com/")
	baseURL, _ := url.Parse(gitURL)
	token := getEnv("GITHUB_TOKEN", "")
	if token == "" {
		return nil, errors.New("missing GITHUB_TOKEN")
	}
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: token,
	})
	tc := oauth2.NewClient(ctx, ts)
	gc := github.NewClient(tc)
	gc.BaseURL = baseURL
	return &Client{
		client: gc,
	}, nil
}
