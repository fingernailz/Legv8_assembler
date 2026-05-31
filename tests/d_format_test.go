package tests

import (
	"legv8_assembler/encoder"
	"strings"
	"testing"
)

func TestDLDUR(t *testing.T) {
	bin, err := encoder.Load_store_format("LDUR X3, [X5, #16]")

	if err != nil {
		t.Errorf("Error, test was not supposed to produce an error: %s", err)
	}

	if strings.Compare(bin, strings.ReplaceAll("11111100010 000010000 00 00101 00011", " ", "")) != 0 {
		t.Errorf("Wrong binary translation\nShould produce 11111100010 000010000 00 00101 00011\ninstead produced %s", bin)
	}
}

func TestDSTUR(t *testing.T) {
	bin, err := encoder.Load_store_format("STUR X4, [X10, #8]")

	if err != nil {
		t.Errorf("Error, test was not supposed to produce an error: %s", err)
	}

	if strings.Compare(bin, strings.ReplaceAll("11111100000 000001000 00 01010 00100", " ", "")) != 0 {
		t.Errorf("Wrong binary translation\nShould produce 11111100000 000001000 00 01010 00100\ninstead produced %s", bin)
	}
}

func TestDLDURSW(t *testing.T) {
	bin, err := encoder.Load_store_format("LDURSW X9, [X20, #4]")

	if err != nil {
		t.Errorf("Error, test was not supposed to produce an error: %s", err)
	}

	if strings.Compare(bin, strings.ReplaceAll("10111000100 000000100 00 10100 01001", " ", "")) != 0 {
		t.Errorf("Wrong binary translation\nShould produce 10111000100 000000100 00 10100 01001\ninstead produced %s", bin)
	}
}

func TestDSTURW(t *testing.T) {
	bin, err := encoder.Load_store_format("STURW X3, [X10, #16]")

	if err != nil {
		t.Errorf("Error, test was not supposed to produce an error: %s", err)
	}

	if strings.Compare(bin, strings.ReplaceAll("10111000000 000010000 00 01010 00011", " ", "")) != 0 {
		t.Errorf("Wrong binary translation\nShould produce 10111000000 000010000 00 01010 00011\ninstead produced %s", bin)
	}
}

func TestDLDURH(t *testing.T) {
	bin, err := encoder.Load_store_format("LDURH X4, [X5, #4]")

	if err != nil {
		t.Errorf("Error, test was not supposed to produce an error: %s", err)
	}

	if strings.Compare(bin, strings.ReplaceAll("01111000010 000000100 00 00101 00100", " ", "")) != 0 {
		t.Errorf("Wrong binary translation\nShould produce 01111000010 000000100 00 00101 00100\ninstead produced %s", bin)
	}
}

func TestDSTURH(t *testing.T) {
	bin, err := encoder.Load_store_format("STURH X9, [X20, #8]")

	if err != nil {
		t.Errorf("Error, test was not supposed to produce an error: %s", err)
	}

	if strings.Compare(bin, strings.ReplaceAll("01111000000 000001000 00 10100 01001", " ", "")) != 0 {
		t.Errorf("Wrong binary translation\nShould produce 01111000000 000001000 00 10100 01001\ninstead produced %s", bin)
	}
}

func TestDLDURB(t *testing.T) {
	bin, err := encoder.Load_store_format("LDURB X3, [X20, #16]")

	if err != nil {
		t.Errorf("Error, test was not supposed to produce an error: %s", err)
	}

	if strings.Compare(bin, strings.ReplaceAll("00111000010 000010000 00 10100 00011", " ", "")) != 0 {
		t.Errorf("Wrong binary translation\nShould produce 00111000010 000010000 00 10100 00011\ninstead produced %s", bin)
	}
}

func TestDSTURB(t *testing.T) {
	bin, err := encoder.Load_store_format("STURB X4, [X5, #8]")

	if err != nil {
		t.Errorf("Error, test was not supposed to produce an error: %s", err)
	}

	if strings.Compare(bin, strings.ReplaceAll("00111000000 000001000 00 00101 00100", " ", "")) != 0 {
		t.Errorf("Wrong binary translation\nShould produce 00111000000 000001000 00 00101 00100\ninstead produced %s", bin)
	}
}

func TestDLDXR(t *testing.T) {
	bin, err := encoder.Load_store_format("LDXR X3, [X5, #0]")

	if err != nil {
		t.Errorf("Error, test was not supposed to produce an error: %s", err)
	}

	if strings.Compare(bin, strings.ReplaceAll("11001000010 000000000 00 00101 00011", " ", "")) != 0 {
		t.Errorf("Wrong binary translation\nShould produce 11001000010 000000000 00 00101 00011\ninstead produced %s", bin)
	}
}

func TestDSTXR(t *testing.T) {
	bin, err := encoder.Load_store_format("STXR X9, X4, [X10]")

	if err != nil {
		t.Errorf("Error, test was not supposed to produce an error: %s", err)
	}

	if strings.Compare(bin, strings.ReplaceAll("11001000000 000000000 00 01010 01001", " ", "")) != 0 {
		t.Errorf("Wrong binary translation\nShould produce 11001000000 000000000 00 01010 01001\ninstead produced %s", bin)
	}
}
