package errors

import (
	"errors"
)

var (
	Invalid_register           = errors.New("Invalid register")
	Invalid_Number_of_Operands = errors.New("Not enough operands to work with")
	Immediate_syntax_error     = errors.New("Invalid syntax for immediate")
	Shamt_parsing_error        = errors.New("Error with parsing shamt")
	Shamt_value_error          = errors.New("Invalid shamt value")
	Signed_immediate_error     = errors.New("Immediate has been provided with signed integer")
	Immediate_parsing_error    = errors.New("Error while parsing immediate")
	Immediate_value_error      = errors.New("Invalid immediate value")
)
