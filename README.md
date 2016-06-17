# gondb

[![Build Status](https://travis-ci.org/abadojack/gondb.svg?branch=master)](https://travis-ci.org/abadojack/gondb)  [![GoDoc](https://godoc.org/github.com/abadojack/gondb?status.png)](http://godoc.org/github.com/abadojack/gondb)

gondb is a simple, transparent Go package for accessing the [National Nutrient Database for Standard Reference](http://ndb.nal.usda.gov/ndb/doc/) API.

Successful API queries return native Go structs that can be used immediately, with no need for type assertions.

gondb implements the endpoints defined in the documentation: http://ndb.nal.usda.gov/ndb/doc/.
More detailed information about the behavior of each particular endpoint can be found at the official documentation of the API.

Examples
-------------

## Installation

	$ go get -u github.com/abadojack/gondb

##Usage

```Go
	import "github.com/abadojack/gondb"
```

## Authentication

A data.gov API key must be incorporated into each API request. [Sign up](http:ndb.nal.usda.gov/ndb/doc/#) now if you do not have a key.
```Go
api := gondb.NewClient(nil, "your-api-key")
```

## Queries

Executing queries is simple.
```Go
	result, _ := api.Search("cheese", nil)
	for _, item := range result.Items {
		fmt.Println(item.Ndbno)
	}
```

The endpoints allow separate optional parameter; if desired, these can be passed as the final parameter.
```Go
	v := url.Values{}
	v.Set("ndbno", "01009")
	v.Set("type", "f")

	nutrientIDs := []string{"204", "205", "269"}

	report, _ := api.GetNutrientReport(nutrientIDs, v)
```

Check the NDB documentation for the various parameters for each endpoint.

## Usage Example

```Go
//A program to display the name and quantity of each nutrient in a raw mango.
package main

import (
	"fmt"

	"github.com/abadojack/gondb"
)

func main() {
	api := gondb.NewClient(nil, "DEMO_KEY")

	result, err := api.Search("fried chicken", nil)

	if err != nil {
		panic(err)
	}

	if len(result.Items) > 0 {
		for _, item := range result.Items {
			report, err := api.GetFoodReport(item.Ndbno, nil)
			if err != nil {
				panic(err)
			}

			for _, nutrient := range report.Food.Nutrients {
				fmt.Println(nutrient.Name, nutrient.Value, nutrient.Unit)
			}

		}
	}
}
```


## Licence
gondb is free software licensed under the GNU LGPL license. Details provided in the LICENSE file.
