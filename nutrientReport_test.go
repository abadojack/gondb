package gondb

import "testing"

//Parameters used for testing
//Nutrients : A list of up to a maximum of 20 nutrient_id's to include in the report.
//FoodGroup : limit your nutrients to one or more food groups by providing a list of food group ID's via the fg parameter.
//			The default is a blank list meaning no food group filtering will be applied. Up to 10 food groups may be specified.
//NdbNo : Report the nutrients for a single food identified by it's unique id -- nutrient number.

func Test_NutrientReport(t *testing.T) {
	api := NewClient(nil, APIKEY)

	p := &Parameters{
		Max:       "12",
		Nutrients: []int{208, 204, 205, 269},
		NdbNo:     "01009",
	}
	_, err := api.GetNutrientReport(p)
	if err != nil {
		t.Fatal(err)
	}

	p.FoodGroup = []string{"0500"}
	_, err = api.GetNutrientReport(p)
	if err != nil {
		t.Fatal(err)
	}

	p.NdbNo = "01009"
	_, err = api.GetNutrientReport(p)
	if err != nil {
		t.Fatal(err)
	}
}
