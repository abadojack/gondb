package gondb

import "net/url"

//List represents a list of foods, nutrients or food groups.
type List struct {
	Type              string `json:"lt"`
	Start             int    `json:"start"`
	End               int    `json:"end"`
	Total             int    `json:"total"`
	StdReleaseVersion string `json:"sr"`
	Sort              string `json:"sort"`
	Items             []Item `json:"item"`
}

//GetList returns metadata about your request and a list of names and id's of foods or nutrients depending on you request.
//With no other parameters the request will return 50 foods items sorted by food name beginning with the first item (0 offset) in JSON format.
func (c Client) GetList(v url.Values) (List, error) {
	var list map[string]List

	err := c.apiGet("list/?", v, &list)

	return list["list"], err
}
