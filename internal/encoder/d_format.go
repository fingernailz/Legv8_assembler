package encoder

import (
	"fmt"
	"legv8_assembler/internal/errors"
	"legv8_assembler/internal/isa"
	"legv8_assembler/internal/registers"
	"strconv"
	"strings"
)

func Load_store_format(instruction string) (string, error) {

	fmt.Println("D format")
	instruction_slice, after, _ := strings.Cut(instruction, " ")

	//opcode 11, address 9, op2 2, base rn 5, source rd 5
	after = strings.TrimSpace(after)
	register, bracket_value, present := strings.Cut(after, ",")

	opcode, _ := isa.Instructions[instruction_slice]

	if !present {
		return "", errors.Invalid_syntax
	}

	rd, avialable := registers.RegistersBin[strings.ToUpper(strings.TrimSpace(register))]
	if !avialable {
		return "", errors.Invalid_register
	}

	bracket_value = strings.TrimSpace(bracket_value)
	if !(strings.HasPrefix(bracket_value, "[") && strings.HasSuffix(bracket_value, "]")) {
		return "", errors.Invalid_syntax
	}

	bracket_value = strings.Replace(bracket_value, "[", "", 1)
	bracket_value = strings.Replace(bracket_value, "]", "", 1)
	register2, immediate, present := strings.Cut(bracket_value, ",")
	if !present {
		return "", errors.Invalid_syntax
	}

	rn, available := registers.RegistersBin[strings.ToUpper(strings.TrimSpace(register2))]
	if !available {
		return "", errors.Invalid_register
	}

	immediate = strings.TrimSpace(immediate)
	if !strings.HasPrefix(immediate, "#") {
		return "", errors.Immediate_syntax_error
	}

	immediate = strings.Replace(immediate, "#", "", 1)
	integer_immediate, err := strconv.Atoi(immediate)
	if err != nil {
		return "", errors.Immediate_parsing_error
	}

	binary_immediate := strconv.FormatUint(uint64(integer_immediate), 2)
	if len(binary_immediate) > 9 { // obv you wouldn't have less than 0 but anyways. len does not return negatives ;)
		return "", errors.Immediate_value_error
	}

	for x, z := 0, len(binary_immediate); x+z < 9; x++ {
		binary_immediate = "0" + binary_immediate
	}

	return opcode["op-code"] + binary_immediate + "00" /*opcode 2*/ + rn + rd, nil
}
