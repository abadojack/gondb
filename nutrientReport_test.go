package gondb

import (
	"net/url"
	"testing"
)

//Queries used for testing
//Nutrients : A list of up to a maximum of 20 nutrient_id's to include in the report.
//FoodGroup : limit your nutrients to one or more food groups by providing a list of food group ID's via the fg parameter.
//			The default is a blank list meaning no food group filtering will be applied. Up to 10 food groups may be specified.
//NdbNo : Report the nutrients for a single food identified by it's unique id -- nutrient number.

func Test_NutrientReport(t *testing.T) {

	nutrientsIDs := []string{"204", "205", "208", "269"}
	v := url.Values{}
	v.Set("max", "12")

	//All foods
	_, err := api.GetNutrientReport(nutrientsIDs, v)
	if err != nil {
		t.Fatal(err)
	}

	//For food groups Dairy and Egg Products (id = 0100) and Poultry Products (id=0500)
	v.Add("fg", "0100")
	v.Add("fg", "0500")
	_, err = api.GetNutrientReport(nutrientsIDs, v)
	if err != nil {
		t.Fatal(err)
	}

	v.Del("fg")

	//For chedder cheese (ndbno 01009) only:
	v.Set("ndbno", "01009")
	_, err = api.GetNutrientReport(nutrientsIDs, v)
	if err != nil {
		t.Fatal(err)
	}
}
