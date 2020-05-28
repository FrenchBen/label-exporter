package exporter

import (
	"context"
	"fmt"

	"github.com/google/go-github/v28/github"
)

// Label interface for fetching github labels
type Label struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Color       string `json:"color"`
}

// ListLabels will fetch the labels from Github, with a specific format
func (c *Client) ListLabels(ctx context.Context, user, repo string) ([]*Label, error) {
	opt := &github.ListOptions{PerPage: 10}
	var labels []*Label
	for {
		ghLabels, resp, err := c.client.Issues.ListLabels(ctx, user, repo, opt)
		if err != nil {
			return nil, err
		}
		for _, l := range ghLabels {
			labels = append(labels, &Label{
				Name:        l.GetName(),
				Description: l.GetDescription(),
				Color:       l.GetColor(),
			})
		}
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	return labels, nil
}

// GetRepo will get details about a repo for an org/user
func (c *Client) GetRepo(ctx context.Context, user, repo string) {
	repos, _, err := c.client.Repositories.Get(context.Background(), user, repo)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Recently updated repositories by %q: %v", user, github.Stringify(repos))
}
