package gondb

import (
	"errors"
	"net/url"
)

//GetNutrientReport returns a list of foods and their nutrient values for a set of specified nutrients.
func (c Client) GetNutrientReport(nutrientsID []string, v url.Values) (NutrientReport, error) {
	if nutrientsID == nil {
		return NutrientReport{}, errors.New("GetNutrientReport requires at least one nutrient id in the nutrients parameter.")
	}

	if v == nil {
		v = url.Values{}
	}

	for _, n := range nutrientsID {
		v.Add("nutrient", n)
	}

	var report map[string]NutrientReport

	err := c.apiGet("nutrients/?", v, &report)
	if err != nil {
		return NutrientReport{}, err
	}

	return report["report"], nil
}

//NutrientReport represents basic information about the nutrient report.
type NutrientReport struct {
	StdReleaseVersion string      `json:"sr"` //Standard Release version of the data being reported
	Groups            interface{} `json:"groups"`
	Subset            string      `json:"subset"`
	Start             int         `json:"start"`
	End               int         `json:"end"`
	Total             int         `json:"total"`
	Foods             []Food      `json:"foods"`
}
