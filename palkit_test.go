package palkit

import (
	"testing"
)

func TestAtoi64(t *testing.T) {
	intVal, err := Atoi64("25")

	if err != nil {
		t.Fatal(err)
	}

	if intVal != 25 {
		t.Fatal("Got wrong number back")
	}
}

func TestJsonNumberToInt(t *testing.T) {
	var number float64 = 10.0
	returnedNumber, err := JsonNumberToInt(number)

	if err != nil {
		t.Fatal(err)
	}

	if returnedNumber != 10 {
		t.Fatal("Wrong number returned", returnedNumber)
	}

}
