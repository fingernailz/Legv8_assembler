package errors

import (
	"errors"
)

var (
	Invalid_register                  = errors.New("Invalid register")
	Invalid_Number_of_Operands        = errors.New("Not enough operands to work with")
	Immediate_syntax_error            = errors.New("Invalid syntax for immediate")
	Shamt_parsing_error               = errors.New("Error with parsing shamt")
	Shamt_value_error                 = errors.New("Invalid shamt value")
	Signed_immediate_error            = errors.New("Immediate has been provided with signed integer")
	Immediate_parsing_error           = errors.New("Error while parsing immediate")
	Immediate_value_error             = errors.New("Invalid immediate value")
	Invalid_label                     = errors.New("Couldn't identify the label")
	Label_location_out_of_bonds_error = errors.New("Label location exceeds more than the allocated bits")
	Label_register_conflict           = errors.New("Labels cannot be named after labels")
	Invalid_syntax                    = errors.New("Invalid syntax")
	Invalid_instruction               = errors.New("Invalid instruction")
	Illegal_label                     = errors.New("Illegal label") //ik I have the same thing here but istg I wrote it just cuz the word illegal sounded cool
)
