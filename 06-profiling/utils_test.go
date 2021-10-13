package main

import (
	"testing"
)

type TestCase struct {
	name     string
	no       int
	expected bool
	actual   bool
}

func Test_IsPrime(t *testing.T) {
	testCases := []TestCase{
		TestCase{name: "Test_IsPrime_1", no: 1, expected: false},
		TestCase{name: "Test_IsPrime_2", no: 2, expected: true},
		TestCase{name: "Test_IsPrime_7", no: 7, expected: true},
		TestCase{name: "Test_IsPrime_9", no: 9, expected: false},
		TestCase{name: "Test_IsPrime_11", no: 11, expected: true},
		TestCase{name: "Test_IsPrime_12", no: 12, expected: false},
		TestCase{name: "Test_IsPrime_13", no: 13, expected: true},
		TestCase{name: "Test_IsPrime_15", no: 15, expected: false},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.actual = IsPrime(testCase.no)
			if testCase.actual != testCase.expected {
				t.Errorf("Expected %v but got %v", testCase.expected, testCase.actual)
			}
		})
	}

}

func Test_IsPrime_2(t *testing.T) {
	testCases := []TestCase{
		TestCase{name: "Test_IsPrime2_1", no: 1, expected: false},
		TestCase{name: "Test_IsPrime2_2", no: 2, expected: true},
		TestCase{name: "Test_IsPrime2_7", no: 7, expected: true},
		TestCase{name: "Test_IsPrime2_9", no: 9, expected: false},
		TestCase{name: "Test_IsPrime2_11", no: 11, expected: true},
		TestCase{name: "Test_IsPrime2_12", no: 12, expected: false},
		TestCase{name: "Test_IsPrime2_13", no: 13, expected: true},
		TestCase{name: "Test_IsPrime2_15", no: 15, expected: false},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.actual = IsPrime2(testCase.no)
			if testCase.actual != testCase.expected {
				t.Errorf("Expected %v but got %v", testCase.expected, testCase.actual)
			}
		})
	}

}

func Test_IsPrime_3(t *testing.T) {
	testCases := []TestCase{
		TestCase{name: "Test_IsPrime3_1", no: 1, expected: false},
		TestCase{name: "Test_IsPrime3_2", no: 2, expected: true},
		TestCase{name: "Test_IsPrime3_7", no: 7, expected: true},
		TestCase{name: "Test_IsPrime3_9", no: 9, expected: false},
		TestCase{name: "Test_IsPrime3_11", no: 11, expected: true},
		TestCase{name: "Test_IsPrime3_12", no: 12, expected: false},
		TestCase{name: "Test_IsPrime3_13", no: 13, expected: true},
		TestCase{name: "Test_IsPrime3_15", no: 15, expected: false},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.actual = IsPrime3(testCase.no)
			if testCase.actual != testCase.expected {
				t.Errorf("Expected %v but got %v", testCase.expected, testCase.actual)
			}
		})
	}

}

func Test_IsPrime_4(t *testing.T) {
	testCases := []TestCase{
		TestCase{name: "Test_IsPrime4_1", no: 1, expected: false},
		TestCase{name: "Test_IsPrime4_2", no: 2, expected: true},
		TestCase{name: "Test_IsPrime4_7", no: 7, expected: true},
		TestCase{name: "Test_IsPrime4_9", no: 9, expected: false},
		TestCase{name: "Test_IsPrime4_11", no: 11, expected: true},
		TestCase{name: "Test_IsPrime4_12", no: 12, expected: false},
		TestCase{name: "Test_IsPrime4_13", no: 13, expected: true},
		TestCase{name: "Test_IsPrime4_15", no: 15, expected: false},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.actual = IsPrime4(testCase.no)
			if testCase.actual != testCase.expected {
				t.Errorf("Expected %v but got %v", testCase.expected, testCase.actual)
			}
		})
	}

}

/* func BenchmarkIsPrime_73(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime(73)
	}
}

func BenchmarkIsPrime2_73(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime2(73)
	}
}

func BenchmarkIsPrime3_73(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime3(73)
	}
}

func BenchmarkIsPrime4_73(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime4(73)
	}
} */

func BenchmarkGeneratePrimes_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GeneratePrimes(3, 100)
	}
}
