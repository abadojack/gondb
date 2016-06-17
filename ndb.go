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
//
//Queries
//
//Executing queries is simple.
//
//	result, _ := api.Search("cheese", nil)
//	for _, item := range result.Items {
//		fmt.Println(item.Ndbno)
//	}
//
//The endpoints allow separate optional parameter; if desired, these can be passed as the final parameter.
//	v := url.Values{}
//	v.Set("ndbno", "01009")
//	v.Set("type", "f")
//
//	nutrientIDs := []string{"204", "205", "269"}
//
//	report, _ := api.GetNutrientReport(nutrientIDs, v)
//
//Endpoints
//
//gondb implements the endpoints defined in the documentation: http://ndb.nal.usda.gov/ndb/doc/.
//
//More detailed information about the behavior of each particular endpoint can be found at the official documentation.
package gondb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

//BaseURL represents the base URL for API requests.
const (
	BaseURL = "http://api.nal.usda.gov/ndb/"
)

//Client represents an NDB API client.
type Client struct {
	HTTPClient *http.Client

	//APIKey required to use the NDB API. Must be a data.gov registered API key.
	APIKey string
}

//APIError represents an code/message error pair returned by the API.
type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

//NewClient returns a new NDB API client. http.Default will be used if no httpClient is provided.
func NewClient(httpClient *http.Client, apiKey string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	if len(apiKey) == 0 {
		panic("No API key was supplied. Get one at http://api.nal.usda.gov")
	}

	return &Client{
		APIKey:     apiKey,
		HTTPClient: httpClient,
	}
}

func (c *Client) apiGet(endpoint string, form url.Values, data interface{}) error {
	if form == nil {
		form = url.Values{}
	}

	form.Set("api_key", c.APIKey)
	form.Set("format", "json")

	urlStr := BaseURL + endpoint + form.Encode()

	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		var apiError map[string]APIError
		json.NewDecoder(resp.Body).Decode(&apiError) //This ain't cool.
		return apiError["error"]
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(data)
}

func (e APIError) Error() string {
	return fmt.Sprintf("%s : %s", e.Code, e.Message)
}
