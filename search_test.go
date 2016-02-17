package gondb

import "testing"

func Test_Search(t *testing.T) {
	api := NewClient(nil, APIKEY)

	p := &Parameters{
		Query: "raw mango",
		Max:   "16",
		Sort:  "r",
	}
	_, err := api.Search(p)
	if err != nil {
		t.Fatal(err)
	}

	p.Query = "butter"
	p.Sort = "n"
	_, err = api.Search(p)
	if err != nil {
		t.Fatal(err)
	}
}
