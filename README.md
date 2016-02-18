# gondb

[![GoDoc](https://godoc.org/github.com/abadojack/gondb?status.png)](http://godoc.org/github.com/abadojack/gondb)


gondb is a simple, transparent Go package for accessing the  [National Nutrient Database for Standard Reference](http://ndb.nal.usda.gov/ndb/doc/) API.

Successful API queries return native Go structs that can be used immediately, with no need for type assertions.

gondb implements the endpoints defined in the documentation: http://ndb.nal.usda.gov/ndb/doc/.
More detailed information about the behavior of each particular endpoint can be found at the official documentation of the API.

Examples
-------------

## Installation

	$ go get -u github.com/abadojack/gondb

## Authentication

A data.gov API key must be incorporated into each API request. [Sign up](http:ndb.nal.usda.gov/ndb/doc/#) now if you do not have a key.
```Go	
api := NewClient(nil, "your-api-key")
```

## Parameters

The Parameters struct contains all parameters required by the endpoints.

```Go
p := &Parameters{
	NdbNo: "01009",
	Type:  "f",
}
```

Check the NDB documentation for the various parameters for each endpoint.

## Usage Example

```Go
//A program to display the names and quantity of each nutrient in a raw mango.
package main

import (
        "fmt"
        "github.com/abadojack/gondb"
)

func main() {
api := gondb.NewClient(nil, "DEMO_KEY")

p := &gondb.Parameters{
        Query: "raw mango",
    }

    result, err := api.Search(p)

    if err != nil {
        panic(err)
    }

    if len(result.Items) > 0 {
        for _, item := range result.Items {
            p1 := &gondb.Parameters{
                	NdbNo: item.Ndbno,
              	}

            report, err := api.GetFoodReport(p1)
            if err != nil {
                panic(err)
            }

            for _, nutrient := range report.FoodDetails.Nutrients {
                fmt.Println(nutrient.Name, nutrient.Value, nutrient.Unit)
            }

        }
    }
}
```


## Licence
gondb is free software licensed under the GNU LGPL license. Details provided in the LICENSE file.
