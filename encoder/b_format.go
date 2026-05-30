package encoder

import (
	"legv8_assembler/errors"
	"legv8_assembler/isa"
	"legv8_assembler/types"
	"strconv"
	"strings"
)

// something is wrong idk FIX plz
func Branch_format(instruction string, label_locations types.Labels) (string, error) {
	// check for b if there then good

	// instruction_slice = slices.Delete(instruction_slice, 0, 1)
	// label := strings.TrimSpace(strings.Join(instruction_slice, ""))
	instruction_slice, after, comma := strings.Cut(instruction, " ")

	if !comma {
		return "", errors.Invalid_syntax
	}

	label := strings.TrimSpace(after)
	location, available := label_locations[label]
	if !available {
		return "", errors.Invalid_label
	}

	location_binary := strconv.FormatUint(uint64(location), 2)

	if len(location_binary) > 26 {
		return "", errors.Label_location_out_of_bonds_error
	}

	for i := 0; i+len(location_binary) < 26; i++ {
		location_binary = "0" + location_binary
	}

	// 6 opcode 26 Address
	return isa.Instructions[strings.ToUpper(instruction_slice)]["op-code"] + location_binary, nil
}
