package gondb

import (
	"math/rand"
	"net/url"
	"os"
	"testing"
	"time"
)

var APIKEY = os.Getenv("NDB_API_KEY")
var api *Client

func init() {
	api = NewClient(nil, APIKEY)
}

func Test_GetFoodReport(t *testing.T) {

	reportType := []string{"f", "s", "b"}
	ndbno := []string{"01009", "18541", "04601", "19086", "36000", "23294", "23312"}
	v := url.Values{}
	r := rand.New(rand.NewSource(time.Now().Unix()))

	for _, rt := range reportType {
		v.Set("type", rt)
		_, err := api.GetFoodReport(ndbno[r.Intn(len(ndbno)-1)], v)
		if err != nil {
			t.Fatal(err)
		}
	}
}
