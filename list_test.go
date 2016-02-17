package gondb

import "testing"

//ListType: f = food , n = all nutrients, ns = speciality nutrients, nr = standard release nutrients only,g = food group.

//Sort order: n=name or id (Meaning of id varies by list type: nutrient number
//for a nutrient list, NDBno for a foods list ,food group id for a food group list.

func TestGetList(t *testing.T) {
	api := NewClient(nil, APIKEY)

	p := &Parameters{
		ListType: "f",
		Sort:     "n",
		Max:      "12",
	}

	_, err := api.GetList(p)
	if err != nil {
		t.Fatal(err)
	}

	p.ListType = "n"
	p.Sort = "id"
	_, err = api.GetList(p)
	if err != nil {
		t.Fatal(err)
	}

	p.ListType = "ns"
	p.Sort = "n"
	_, err = api.GetList(p)
	if err != nil {
		t.Fatal(err)
	}

	p.ListType = "nr"
	p.Sort = "id"
	_, err = api.GetList(p)
	if err != nil {
		t.Fatal(err)
	}

	p.ListType = "g"
	p.Sort = "n"
	_, err = api.GetList(p)
	if err != nil {
		t.Fatal(err)
	}

}
