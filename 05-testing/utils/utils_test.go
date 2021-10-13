package utils_test

import (
	"testing"
	"testing-app/utils"
)

/*
func Test_IsPrime_7(t *testing.T) {
	//Arrange
	no := 7
	expected := true

	//Act
	actual := utils.IsPrime(no)

	//Assert
	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func Test_IsPrime_9(t *testing.T) {
	//Arrange
	no := 9
	expected := false

	//Act
	actual := utils.IsPrime(no)

	//Assert
	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func Test_IsPrime_11(t *testing.T) {
	//Arrange
	no := 11
	expected := true

	//Act
	actual := utils.IsPrime(no)

	//Assert
	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

*/

type TestCase struct {
	name     string
	no       int
	expected bool
	actual   bool
}

func Test_IsPrime(t *testing.T) {
	testCases := []TestCase{
		TestCase{name: "Test_IsPrime_7", no: 7, expected: true},
		TestCase{name: "Test_IsPrime_9", no: 9, expected: false},
		TestCase{name: "Test_IsPrime_11", no: 11, expected: true},
		TestCase{name: "Test_IsPrime_12", no: 12, expected: false},
		TestCase{name: "Test_IsPrime_13", no: 13, expected: true},
		TestCase{name: "Test_IsPrime_15", no: 15, expected: false},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.actual = utils.IsPrime(testCase.no)
			if testCase.actual != testCase.expected {
				t.Errorf("Expected %v but got %v", testCase.expected, testCase.actual)
			}
		})
	}

}
