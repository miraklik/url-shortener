package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testStoreService = &StorageService{}
)

func init() {
	testStoreService = NewStore()
}

func TestStore(t *testing.T) {
	assert.True(t, testStoreService.redisClient != nil)
}

func TestInsertionAndRetrieval(t *testing.T) {
	initialURL := "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html"
	userID := "e0dba740-fc4b-4977-872c-d360239e6b1a"
	shortURL := "Jsz4k57oAX"

	SaveURLMapping(shortURL, initialURL, userID)

	retrievalURL := RetrieveInitialUrl(shortURL)

	assert.Equal(t, initialURL, retrievalURL)
}
