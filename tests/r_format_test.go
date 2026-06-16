package tests

import (
	"legv8_assembler/encoder"
	"strings"
	"testing"
)

// something needs to be done

// ADD X1, x2, X4
// 10001011000001000000000001000001
func TestRAdd(t *testing.T) {
	//
	bin, err := encoder.Register_format("ADD X1, x2, x4")

	if err != nil {
		t.Errorf("Error, test was not supposed to produce an error: %s", err)
	}

	if strings.Compare(bin, "10001011000001000000000001000001") != 0 {
		t.Errorf("Wrong binary translation\nShould produce 10001011000001000000000001000001\ninstead produced %s", bin)
	}

}

func TestRSUB(t *testing.T) {
	bin, err := encoder.Register_format("SUB X0, X1, X2")

	if err != nil {
		t.Errorf("Error, test was not supposed to produce an error: %s", err)
	}

	if strings.Compare(bin, strings.ReplaceAll("11001011000 00010 000000 00001 00000", " ", "")) != 0 {
		t.Errorf("Wrong binary translation\nShould produce 11001011000 00010 000000 00001 00000\ninstead produced %s", bin)
	}

}

func TestRADDS(t *testing.T) {
	bin, err := encoder.Register_format("ADDS X0, X1, X2")

	if err != nil {
		t.Errorf("Error, test was not supposed to produce an error: %s", err)
	}

	if strings.Compare(bin, strings.ReplaceAll("10101011000 00010 000000 00001 00000", " ", "")) != 0 {
		t.Errorf("Wrong binary translation\nShould produce 10101011000 00010 000000 00001 00000\ninstead produced %s", bin)
	}
}

func TestRSUBS(t *testing.T) {
	bin, err := encoder.Register_format("SUBS X0, X1, X2")

	if err != nil {
		t.Errorf("Error, test was not supposed to produce an error: %s", err)
	}

	if strings.Compare(bin, strings.ReplaceAll("11101011000 00010 000000 00001 00000", " ", "")) != 0 {
		t.Errorf("Wrong binary translation\nShould produce 11101011000 00010 000000 00001 00000\ninstead produced %s", bin)
	}
}

func TestRAND(t *testing.T) {
	bin, err := encoder.Register_format("AND X0, X1, X2")

	if err != nil {
		t.Errorf("Error, test was not supposed to produce an error: %s", err)
	}

	if strings.Compare(bin, strings.ReplaceAll("10001010000 00010 000000 00001 00000", " ", "")) != 0 {
		t.Errorf("Wrong binary translation\nShould produce 10001010000 00010 000000 00001 00000\ninstead produced %s", bin)
	}
}

func TestRANDS(t *testing.T) {
	bin, err := encoder.Register_format("ANDS X0, X1, X2")

	if err != nil {
		t.Errorf("Error, test was not supposed to produce an error: %s", err)
	}

	if strings.Compare(bin, strings.ReplaceAll("11101010000 00010 000000 00001 00000", " ", "")) != 0 {
		t.Errorf("Wrong binary translation\nShould produce 11101010000 00010 000000 00001 00000\ninstead produced %s", bin)
	}
}

func TestRORR(t *testing.T) {
	bin, err := encoder.Register_format("ORR X0, X1, X2")

	if err != nil {
		t.Errorf("Error, test was not supposed to produce an error: %s", err)
	}

	if strings.Compare(bin, strings.ReplaceAll("10101010000 00010 000000 00001 00000", " ", "")) != 0 {
		t.Errorf("Wrong binary translation\nShould produce 10101010000 00010 000000 00001 00000\ninstead produced %s", bin)
	}
}

func TestREOR(t *testing.T) {
	bin, err := encoder.Register_format("EOR X0, X1, X2")

	if err != nil {
		t.Errorf("Error, test was not supposed to produce an error: %s", err)
	}

	if strings.Compare(bin, strings.ReplaceAll("11101010000 00010 000000 00001 00000", " ", "")) != 0 {
		t.Errorf("Wrong binary translation\nShould produce 11101010000 00010 000000 00001 00000\ninstead produced %s", bin)
	}
}

func TestRLSL(t *testing.T) {
	bin, err := encoder.Register_format("LSL X0, X1, #4")

	if err != nil {
		t.Errorf("Error, test was not supposed to produce an error: %s", err)
	}

	if strings.Compare(bin, strings.ReplaceAll("11010011011 00000 000100 00001 00000", " ", "")) != 0 {
		t.Errorf("Wrong binary translation\nShould produce 11010011011 00000 000100 00001 00000\ninstead produced %s", bin)
	}
}

func TestRLSR(t *testing.T) {
	bin, err := encoder.Register_format("LSR X0, X1, #4")

	if err != nil {
		t.Errorf("Error, test was not supposed to produce an error: %s", err)
	}

	if strings.Compare(bin, strings.ReplaceAll("11010011010 00000 000100 00001 00000", " ", "")) != 0 {
		t.Errorf("Wrong binary translation\nShould produce 11010011010 00000 000100 00001 00000\ninstead produced %s", bin)
	}
}

func TestRBR(t *testing.T) {
	bin, err := encoder.Register_format("BR X1")

	if err != nil {
		t.Errorf("Error, test was not supposed to produce an error: %s", err)
	}

	if strings.Compare(bin, strings.ReplaceAll("11010110000 00000 000000 00001 00000", " ", "")) != 0 {
		t.Errorf("Wrong binary translation\nShould produce 11010110000 00000 000000 00001 00000\ninstead produced %s", bin)
	}
}

func TestR2(t *testing.T) {
	bin, err := encoder.Register_format("ADD asd, asd, asd")

	if err == nil {
		t.Errorf("error, test was supposed to produce an error")
	}

	if bin != "" {
		t.Errorf("Bin should be equal to nil")
	}
}
