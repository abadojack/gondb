package gondb

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
func (c Client) GetList(param *Parameters) (List, error) {
	var list map[string]List

	err := c.apiGet("list/?", param, &list)
	if err != nil {
		return List{}, err
	}

	return list["list"], nil
}
