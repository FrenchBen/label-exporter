package exporter

import (
	"context"

	"github.com/google/go-github/v28/github"
)

// Label interface for fetching github labels
type Label struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Color       string `json:"color"`
}

// ListLabels will fetch the labels from Github, with a specific format
func (c *Client) ListLabels(ctx context.Context, owner, repo string) ([]*Label, error) {
	opt := &github.ListOptions{PerPage: 10}
	var labels []*Label
	for {
		ghLabels, resp, err := c.client.Issues.ListLabels(ctx, owner, repo, opt)
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
