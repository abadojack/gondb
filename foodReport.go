package gondb

import "errors"

//GetFoodReport obtains nutrient reports on individual foods.
func (c Client) GetFoodReport(param *Parameters) (FoodReport, error) {
	if param == nil {
		return FoodReport{}, errors.New("Parameter NdbNo must not be empty")
	}

	var report map[string]FoodReport
	err := c.apiGet("reports/?", param, &report)

	return report["report"], err
}

//FoodReport represents basic information about the report.
type FoodReport struct {
	StdReleaseVersion string     `json:"sr"` //Standard Release version of the data being reported
	Type              string     `json:"type"`
	FoodDetails       Food       `json:"food"`
	Sources           []Source   `json:"sources"`
	FootNotes         []Footnote `json:"footnotes"`
	Languals          []Langual  `json:"language"`
}

//Food represents metadata elements for a food.
type Food struct {
	NdbNo              string     `json:"nodno"`
	Name               string     `json:"name"`
	FoodGroup          string     `json:"fg"`
	ScientificName     string     `json:"sn"`
	CommercialName     string     `json:"cn"`
	Manufacturer       string     `json:"manu"`
	NitrogenFactor     float64    `json:"nf"`
	CarbohydrateFactor float64    `json:"cf"`
	FatFactor          float64    `json:"ff"`
	ProteinFactor      float64    `json:"pf"`
	Refuse             string     `json:"r"`
	RefuseDescription  string     `json:"rd"`
	Nutrients          []Nutrient `json:"nutrients"`
	Weight             float64    `json:"weight"`
	FoodMeasure        string     `json:"measures"`
}

//Nutrient represents metadata elements for each nutrient.
type Nutrient struct {
	ID            interface{} `json:"nutrient_id"` //Can either be string or int
	Name          string      `json:"name"`
	Group         string      `json:"group"`
	Unit          string      `json:"unit"`
	Value         interface{} `json:"value"` //Can either be string or float64
	SourceCode    interface{} `json:"sourcecode"`
	Dp            interface{} `json:"dp"`
	StandardError string      `json:"se"`
	Measures      []Measure   `json:"measures"`
	Gm            interface{} `json:"gm"` //Can either be float64 or string i.e "--"
}

//Measure represents list of measures reported for a nutrient.
type Measure struct {
	Label      string      `json:"label"`
	Equivalent float64     `json:"eqv"`
	Value      interface{} `json:"value"` //Can either be string or float64
	Quantity   float64     `json:"qty"`
}

//Source represents reference source, usually a bibliographic citation, for the food.
type Source struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Authors string `json:"authors"`
	Volume  string `json:"vol"`
	Iss     string `json:"iss"`
	Year    string `json:"year"`
	Start   string `json:"start"`
	End     string `json:"end"`
}

type Footnote struct {
	Idv         string `json:"idv"`
	Description string `json:"desc"`
}

//Langual represents LANGUAL codes assigned to the food.
type Langual struct {
	Code        string `json:"code"`
	Description string `json:"code"`
}
