package gondb

import "net/url"

//Search request sends keyword queries and returns lists of foods which contain
//one or more of the keywords in the food description, scientific name, or commerical name fields.
func (c Client) Search(searchTerm string, v url.Values) (SearchResultList, error) {
	var list map[string]SearchResultList

	if v == nil {
		v = url.Values{}
	}

	v.Set("q", searchTerm)

	err := c.apiGet("search/?", v, &list)

	return list["list"], err
}

//SearchResultList represents information about the items returned.
type SearchResultList struct {
	Query                  string `json:"q"`     //terms requested and used in the search
	StandardReleaseVersion string `json:"sr"`    //Standard Release version of the data being reported
	Start                  int    `json:"start"` //beginning item in the list
	End                    int    `json:"end"`   //last item in the list
	Total                  int    `json:"total"` //total # of items returned by the search
	Group                  string `json:"group"`
	Sort                   string `json:"sort"`   //requested sort order (r=relevance or n=name)
	Offset                 int    `json:"offset"` //beginning offset into the results list for the items in the list requested
	FoodGroup              string `json:"fg"`     //food group filter
	Items                  []Item `json:"item"`
}

//Item represents individual items on the list.
type Item struct {
	Offset int    `json:"offset"`
	Group  string `json:"group"`
	Name   string `json:"name"`
	Ndbno  string `json:"ndbno"`
	ID     string `json:"id"`
}
