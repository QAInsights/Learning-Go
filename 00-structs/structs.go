package main

import "fmt"

func main() {

	type testingTools struct {
		toolName         string
		toolVersion      int
		toolPrice        float32
		toolTestingTypes []string
	}

	jmeter := testingTools{
		toolName:    "JMeter",
		toolVersion: 5,
		toolPrice:   0.00,
		toolTestingTypes: []string{
			"load testing",
			"automation",
			"ci/cd",
		},
	}

	fmt.Println(jmeter.toolTestingTypes) // Prints everything

	// One line struct define
	automationTool := struct{ name string }{name: "Selenium"}
	fmt.Println(automationTool.name) // Prints Selenium

	// Copying
	automationToolA := automationTool
	automationToolA.name = "UFT"

	fmt.Println(automationToolA.name) // Prints UFT

	automationToolB := &automationTool
	automationToolB.name = "Postman"

	fmt.Println(automationTool.name)  // Prints Postman
	fmt.Println(automationToolB.name) // Prints Postman

}
