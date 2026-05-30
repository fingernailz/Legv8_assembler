package encoder

import (
	"fmt"
	"legv8_assembler/isa"
	"legv8_assembler/types"
	"strconv"
	"strings"
)

// something is wrong idk FIX plz
func call_b_format(instruction string, label_locations types.Labels) string {
	// check for b if there then good

	// instruction_slice = slices.Delete(instruction_slice, 0, 1)
	// label := strings.TrimSpace(strings.Join(instruction_slice, ""))
	instruction_slice, _, _ := strings.Cut(instruction, " ")
	label := instruction_slice
	location, available := label_locations[label]
	if !available {
		fmt.Println("Invalid label ", label)
	}

	location_binary := strconv.FormatUint(uint64(location), 2)

	if len(location_binary) > 26 {
		fmt.Println("Error with location log")
	}

	for i := 0; i+len(location_binary) < 26; i++ {
		location_binary = "0" + location_binary
	}

	fmt.Println("here in b, ", isa.Instructions[strings.ToUpper(instruction_slice)]["op-code"])
	// 6 opcode 26 Address
	return isa.Instructions[strings.ToUpper(instruction_slice)]["op-code"] + location_binary
}
