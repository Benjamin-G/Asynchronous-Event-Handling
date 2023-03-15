package main

import (
	"fmt"
	"os"
)

func main() {
	schema := `
			{
              "type": "record",
              "name": "LongList",
              "fields" : [
	      			{
						"name": "next", 
						"type": ["null", "LongList", {"type": "long", "logicalType": "timestamp-millis"}], 
						"default": null
					}
              ]
            }
	`
	// Create a file
	f, err := os.Create("schemas/schema.avsc")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	// Write to file
	n2, err := f.Write([]byte(schema))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("wrote %d bytes\n", n2)
}
