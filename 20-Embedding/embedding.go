package main

import (
	"fmt"
	"reflect"
)

func main() {
	type TestingTool struct {
		Name 		string `required max: "50"`
		OpenSource 	bool
	}

	type PerformanceTesting struct {
		TestingTool
		LoadTesting bool
	}

	type Automation struct {
		TestingTool
		Automation bool
	}

	// Composition or Embedding aka Inheritance

	myTool := PerformanceTesting{}
	myTool.Name = "JMeter"
	myTool.OpenSource = true
	myTool.LoadTesting = true

	fmt.Println(myTool)

	myAutoTool := Automation{}
	myAutoTool.Name = "Selenium"
	myAutoTool.OpenSource = true
	myAutoTool.Automation = true
	fmt.Println(myAutoTool)

	// Embedding

	myAnotherPerfTool := PerformanceTesting{
		TestingTool: TestingTool{Name: "Locust",OpenSource: true},
			LoadTesting: true,
	}

	fmt.Println(myAnotherPerfTool)

	// tags using reflect package
	t := reflect.TypeOf(TestingTool{})
	f, _ := t.FieldByName("Name")
	fmt.Println(f.Name, f.Tag, f.Type)
}
