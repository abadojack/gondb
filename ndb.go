//Package gondb provides structs and functions for accessing the National Nutrient Database for Standard Reference API.
//
//Successful API queries return native Go structs that can be used immediately,
//with no need for type assertions.
//
//Authentication
//
//A data.gov API key must be incorporated into each API request. Sign up (http://ndb.nal.usda.gov/ndb/doc/#) now if you do not have a key.
//
//	api := NewClient(nil, "your-api-key")
//
//Parameters
//
//The Parameters struct contains all parameters required by the endpoints.
//
//	p := &Parameters{
//		NdbNo: "01009",
//		Type:  "f",
//	}
//
//Check the NDB documentation for the various parameters for each endpoint.
//
//
//Queries
//
//Executing queries is simple.
//
//Example 0:
//
//	p := &Parameters{
//		Query: "raw mangoes",
//	}
//	s, err := api.Search(p)
//
//Example 1:
//
//	p := &Parameters{
//		Max:       "12",
//		Nutrients: []int{208, 204, 205, 269},
//		NdbNo:     "01009",
//	}
//	r, err := api.GetNutrientReport(p)
//
//Endpoints
//
//gondb implements the endpoints defined in the documentation: http://ndb.nal.usda.gov/ndb/doc/.
//
//More detailed information about the behavior of each particular endpoint can be found at the official documentation.
package gondb

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

const (
	baseURL = "http://api.nal.usda.gov/ndb/"
)

//Client represents an NDB API client.
type Client struct {
	HTTPClient *http.Client

	//BaseURL for API requests.
	BaseURL *url.URL

	//APIKey required to use the NDB API. Must be a data.gov registered API key.
	APIKey string
}

//Parameters represents API request parameters.
type Parameters struct {
	Format    string   `url:"format,omitempty"` //json or xml ... always json
	Type      string   `url:"type,omitempty"`   //[b]asic or [f]ull or [s]tats
	NdbNo     string   `url:"ndbno,omitempty"`
	APIKey    string   `url:"api_key,omitempty"`
	Query     string   `url:"q,omitempty"`
	FoodGroup []string `url:"fg,omitempty"`
	Sort      string   `url:"sort,omitempty"`
	Max       string   `url:"max,omitempty"`
	Offset    string   `url:"offset,omitempty"`
	Nutrients []int    `url:"nutrients,omitempty"`
	Subset    int      `url:"subset,omitempty"`
	ListType  string   `url:"lt"`
}

//NewClient returns a new NDB API client. http.Default will be used if no httpClient is provided.
func NewClient(httpClient *http.Client, apiKey string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	
	//Use DEMO_KEY if no api key is provided.
	if len(apiKey) == 0 {
		apiKey = "DEMO_KEY"
	}

	baseURL, _ := url.Parse(baseURL)

	c := &Client{
		APIKey:     apiKey,
		BaseURL:    baseURL,
		HTTPClient: httpClient,
	}

	return c
}

func (c *Client) apiGet(endpoint string, param *Parameters, data interface{}) error {
	if param == nil {
		param = new(Parameters)
	}
	param.APIKey = c.APIKey
	param.Format = "json" // Must always be json.

	v, err := query.Values(param)
	if err != nil {
		return err
	}

	rel, err := url.Parse(endpoint + v.Encode())
	if err != nil {
		return err
	}

	u := c.BaseURL.ResolveReference(rel)

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(data)
}
