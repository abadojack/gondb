package gondb

import (
	"net/url"
	"testing"
)

func Test_Search(t *testing.T) {
	v := url.Values{}
	v.Set("max", "16")

	queries := []string{"", "butter", "raw mango"}
	sort := []string{"r", "n"}

	for key, query := range queries {
		v.Set("sort", sort[key%2]) // alternate btn "n" and "r"
		_, err := api.Search(query, v)
		if err != nil {
			t.Fatal(err)
		}
	}
}
