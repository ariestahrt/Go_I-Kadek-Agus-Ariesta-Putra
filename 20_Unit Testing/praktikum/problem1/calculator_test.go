package main

import "testing"

var (
    A	float64 = 10
    B	float64 = 5
	expectedAdd float64 = 15
	expectedSubstract float64 = 5
	expectedMultiply float64 = 50
	expectedDivide float64 = 2
)

func TestAdd(t *testing.T) {
    t.Logf("A : %.2f", A)
    t.Logf("B : %.2f", B)
	result := Add(A,B)
    t.Logf("Add(%.2f,%.2f) : %.2f", A, B, result)

    if result != expectedAdd {
        t.Errorf("FAILED! expected %.2f instead of %.2f", expectedAdd, result)
    }
}

func TestSubstract(t *testing.T) {
    t.Logf("A : %.2f", A)
    t.Logf("B : %.2f", B)
	result := Substract(A,B)
    t.Logf("Substract(%.2f,%.2f) : %.2f", A, B, result)

    if result != expectedSubstract {
        t.Errorf("FAILED! expected %.2f instead of %.2f", expectedAdd, result)
    }
}

func TestMultiply(t *testing.T) {
    t.Logf("A : %.2f", A)
    t.Logf("B : %.2f", B)
	result := Multiply(A,B)
    t.Logf("Multiply(%.2f,%.2f) : %.2f", A, B, result)

    if result != expectedMultiply {
        t.Errorf("FAILED! expected %.2f instead of %.2f", expectedAdd, result)
    }
}

func TestDivide(t *testing.T) {
    t.Logf("A : %.2f", A)
    t.Logf("B : %.2f", B)
	result := Divide(A,B)
    t.Logf("Divide(%.2f,%.2f) : %.2f", A, B, result)

    if result != expectedDivide {
        t.Errorf("FAILED! expected %.2f instead of %.2f", expectedAdd, result)
    }
}