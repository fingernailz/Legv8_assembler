package tests

import (
	"legv8_assembler/encoder"
	"strings"
	"testing"
)

func TestIADDI(t *testing.T) {
	bin, err := encoder.Immediate_format("ADDI X0, X1, #5")

	if err != nil {
		t.Errorf("Error, test was not supposed to produce an error: %s", err)
	}

	if strings.Compare(bin, strings.ReplaceAll("1001000100 000000000101 00001 00000", " ", "")) != 0 {
		t.Errorf("Wrong binary translation\nShould produce 1001000100 000000000101 00001 00000\ninstead produced %s", bin)
	}

}

func TestISUBI(t *testing.T) {
	bin, err := encoder.Immediate_format("SUBI X0, X1, #5")

	if err != nil {
		t.Errorf("Error, test was not supposed to produce an error: %s", err)
	}

	if strings.Compare(bin, strings.ReplaceAll("1101000100 000000000101 00001 00000", " ", "")) != 0 {
		t.Errorf("Wrong binary translation\nShould produce 1101000100 000000000101 00001 00000\ninstead produced %s", bin)
	}

}

func TestIADDIS(t *testing.T) {
	bin, err := encoder.Immediate_format("ADDIS X0, X1, #5")

	if err != nil {
		t.Errorf("Error, test was not supposed to produce an error: %s", err)
	}

	if strings.Compare(bin, strings.ReplaceAll("1011000100 000000000101 00001 00000", " ", "")) != 0 {
		t.Errorf("Wrong binary translation\nShould produce 1011000100 000000000101 00001 00000\ninstead produced %s", bin)
	}

}
func TestISUBIS(t *testing.T) {
	bin, err := encoder.Immediate_format("SUBIS X0, X1, #5")

	if err != nil {
		t.Errorf("Error, test was not supposed to produce an error: %s", err)
	}

	if strings.Compare(bin, strings.ReplaceAll("1111000100 000000000101 00001 00000", " ", "")) != 0 {
		t.Errorf("Wrong binary translation\nShould produce 1111000100 000000000101 00001 00000 00000\ninstead produced %s", bin)
	}

}
func TestIANDI(t *testing.T) {
	bin, err := encoder.Immediate_format("ANDI X0, X1, #5")

	if err != nil {
		t.Errorf("Error, test was not supposed to produce an error: %s", err)
	}

	if strings.Compare(bin, strings.ReplaceAll("1001001000 000000000101 00001 00000", " ", "")) != 0 {
		t.Errorf("Wrong binary translation\nShould produce 1001001000 000000000101 00001 00000\ninstead produced %s", bin)
	}

}
func TestIANDIS(t *testing.T) {
	bin, err := encoder.Immediate_format("ANDIS X0, X1, #5")

	if err != nil {
		t.Errorf("Error, test was not supposed to produce an error: %s", err)
	}

	if strings.Compare(bin, strings.ReplaceAll("1111001000 000000000101 00001 00000", " ", "")) != 0 {
		t.Errorf("Wrong binary translation\nShould produce 1111001000 000000000101 00001 00000\ninstead produced %s", bin)
	}

}
func TestIORRI(t *testing.T) {
	bin, err := encoder.Immediate_format("ORRI X0, X1, #5")

	if err != nil {
		t.Errorf("Error, test was not supposed to produce an error: %s", err)
	}

	if strings.Compare(bin, strings.ReplaceAll("1011001000 000000000101 00001 00000", " ", "")) != 0 {
		t.Errorf("Wrong binary translation\nShould produce 1011001000 000000000101 00001 00000\ninstead produced %s", bin)
	}

}
func TestIEORI(t *testing.T) {
	bin, err := encoder.Immediate_format("EORI X0, X1, #5")

	if err != nil {
		t.Errorf("Error, test was not supposed to produce an error: %s", err)
	}

	if strings.Compare(bin, strings.ReplaceAll("1101001000 000000000101 00001 00000", " ", "")) != 0 {
		t.Errorf("Wrong binary translation\nShould produce 1101001000 000000000101 00001 00000\ninstead produced %s", bin)
	}

}
