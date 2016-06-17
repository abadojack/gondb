package gondb

import (
	"net/url"
	"testing"
)

//ListType: f = food , n = all nutrients, ns = speciality nutrients, nr = standard release nutrients only,g = food group.

//Sort order: n=name or id (Meaning of id varies by list type: nutrient number
//for a nutrient list, NDBno for a foods list ,food group id for a food group list.)

func TestGetList(t *testing.T) {
	listType := []string{"f", "n", "ns", "nr", "g"}
	sort := []string{"id", "n"}

	v := url.Values{}
	v.Set("max", "12")
	for key, lt := range listType {
		v.Set("lt", lt)
		v.Set("sort", sort[key%2]) // alternate sorting btn "id" and "n".

		_, err := api.GetList(v)
		if err != nil {
			t.Fatal(err)
		}
	}
}
