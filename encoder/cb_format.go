package encoder

import (
	"legv8_assembler/errors"
	"legv8_assembler/isa"
	"legv8_assembler/registers"
	"legv8_assembler/types"
	"strconv"
	"strings"
)

func Conditional_branch_format(instruction string, label_locations types.Labels) (string, error) {
	// opcode 8 location 19 condition register 5
	// I've zero clue how to implement cb.<condition format>, REMIND LATER
	instruction_slice, after, _ := strings.Cut(instruction, " ")
	after = strings.TrimSpace(after)
	test_space := strings.Split(after, ",")

	if len(test_space) > 2 {
		return "", errors.Invalid_Number_of_Operands
	}

	// have a seperate one fo B.cond

	opcode, _ := isa.Instructions[instruction_slice]
	rd, rd_available := registers.RegistersBin[strings.ToUpper(strings.TrimSpace(test_space[0]))]
	if !rd_available {
		return "", errors.Invalid_register
	}

	_, invalid_label := registers.RegistersBin[strings.ToUpper(strings.TrimSpace(test_space[1]))]

	if invalid_label {
		return "", errors.Label_register_conflict
	}

	label, label_a := label_locations[strings.TrimSpace(test_space[1])]
	if !label_a {
		return "", errors.Invalid_label
	}

	binary_label_location := strconv.FormatUint(uint64(label), 2)
	if len(binary_label_location) > 19 {
		return "", errors.Label_location_out_of_bonds_error
	}

	for x, y := 0, len(binary_label_location); x+y < 19; x++ {
		binary_label_location = "0" + binary_label_location
	}

	return opcode["op-code"] + binary_label_location + rd, nil
}
