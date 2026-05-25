package isa

type ISA map[string]string
// type Registers map[string]string /*just create a fun to dynamic find the bin instead of this*/

var Instructions ISA = {
	// R format
	"ADD": "10001011000",
	"SUB": "11001011000",
	"ADDS": "10101011000",
	"SUBS": "11101011000",
	"AND": "10001010000",
	"ANDS": "11101010000",
	"ORR": "10101010000",
	"EOR": "11101010000",
	"LSL": "11010011011",
	"LSR": "11010011010",
	"BR": "11010110000",

	// I format
	"ADDI": "1001000100",
	"SUBI": "1101000100",
	"ADDIS": "1011000100",
	"SUBIS": "1111000100",
	"ANDI": "1001001000",
	"ANDIS": "1111001000",
	"ORRI": "1011001000",
	"EORI": "1101001000",

	// D format
	"LDUR": "11111100010",
	"STUR": "11111100000",
	"LDURSW": "10111000100",
	"STURW": "10111000000",
	"LDURH": "01111000010",
	"STURH": "01111000000",
	"LDURB": "00111000010",
	"STURB": "00111000000",
	"LDXR": "11001000010",
	"STXR": "11001000000",

	// IM format
	"MOVZ": "110100101",
	"MOVK": "111100101",

	// CB format

	"CBZ": "10110100",
	"CBNZ": "10110101",
	// branch with condition, idk how to do, do later! opcode 01010100

	// B format
	"B": "000101",
	"BL": "100101"
}

/*
R format 
	1. 3 operands
	2. 2 operands, one shmt
	3. BR with one operand mostly x30
*/

/* D format, only prob is parsing the brackets */

/* B has no prob afaics sw CB ig*/

/* two pass assemblers so first find labels and their addresses*/