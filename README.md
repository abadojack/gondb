# gondb provides structs and functions for accessing the National Nutrient Database for Standard Reference API.

[Documentation](https://godoc.org/github.com/abadojack/gondb)

## Installation

go get -u github.com/abadojack/gondb

## Usage Example

        ```go
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
