//go:build e2e
// +build e2e

package test

import (
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetArticles(t *testing.T) {
	client := resty.New()
	resp, err := client.R().Get(BASE_URL + "/api/article")
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, 200, resp.StatusCode())
}

func TestPostArticle(t *testing.T) {
	client := resty.New()
	resp, err := client.R().
		SetBody(`{"Title": "test title", "Body": "12345"}`).
		Post(BASE_URL + "/api/article")
	assert.NoError(t, err)

	assert.Equal(t, 200, resp.StatusCode())
}