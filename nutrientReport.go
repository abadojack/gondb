package gondb

import "errors"

//GetNutrientReport returns a list of foods and their nutrient values for a set of specified nutrients.
func (c Client) GetNutrientReport(param *Parameters) (NutrientReport, error) {
	if param == nil || len(param.Nutrients) == 0 {
		return NutrientReport{}, errors.New("GetNutrientReport requires at least one nutrient_id in the nutrients parameter.")
	}

	var report map[string]NutrientReport
	err := c.apiGet("nutrients/?", param, &report)
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
