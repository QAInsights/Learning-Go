package main

import "fmt"

func main() {

	type testingTools struct {
		toolName string
		toolVersion int
		toolPrice float32
		toolTestingTypes []string

	}

	jmeter := testingTools{
		toolName: "JMeter",
		toolVersion: 5,
		toolPrice: 0.00,
		toolTestingTypes: []string {
			"load testing",
			"automation",
			"ci/cd",
		},
	}

	fmt.Println(jmeter.toolTestingTypes)
}