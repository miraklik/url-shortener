package shortener

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const UserID = "e0dba740-fc4b-4977-872c-d360239e6b1a"

func TestShortLink(t *testing.T) {
	initialLink := "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html"
	shortLink := GenerateShortLink(initialLink, UserID)

	initialLink_2 := "https://www.eddywm.com/lets-build-a-url-shortener-in-go-with-redis-part-2-storage-layer/"
	shortLink_2 := GenerateShortLink(initialLink_2, UserID)

	assert.Equal(t, shortLink, "jTa4L57P")
	assert.Equal(t, shortLink_2, "d66yfx7N")
}
