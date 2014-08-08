package httpcache

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func mustParseUrl(u string) *url.URL {
	ru, err := url.Parse(u)
	if err != nil {
		panic(err)
	}
	return ru
}

func TestKeysDiffer(t *testing.T) {
	k1 := Key("GET", mustParseUrl("http://x.org/test"))
	k2 := Key("GET", mustParseUrl("http://y.org/test"))

	assert.NotEqual(t, k1, k2)
}

func TestHeadIsTheSameAsGet(t *testing.T) {
	k1 := Key("GET", mustParseUrl("http://x.org/test"))
	k2 := Key("HEAD", mustParseUrl("http://x.org/test"))

	assert.Equal(t, k1, k2)
}

func TestSecondaryKeysDiffer(t *testing.T) {
	k1 := Key("GET", mustParseUrl("http://x.org/test"))
	k2 := SecondaryKey(Key("GET", mustParseUrl("http://x.org/test")), http.Header{
		"X-Test": []string{"llamas"},
	})

	assert.NotEqual(t, k1, k2)
}