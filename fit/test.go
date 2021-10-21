package main

// this clones an example from github.com/tormoder/fit

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/tormoder/fit"
)

func main() {
	// Read our FIT test file data
	testData, err := ioutil.ReadFile("example.fit")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Decode the FIT file data
	fit, err := fit.Decode(bytes.NewReader(testData))
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get the actual activity
	activity, err := fit.Activity()
	if err != nil {
		fmt.Println(err)
		return
	}

	prev := uint32(0)
	total := uint32(0)
	for _, record := range activity.Records {
		if record.Cadence > 100 {
			total += record.Distance - prev
		}
		prev = record.Distance
	}

	fmt.Println(total/160958, "mi above 100 rpm")
	fmt.Println(prev/160958, "mi total")
}
