package encoder

import (
	"fmt"
	"legv8_assembler/isa"
	"legv8_assembler/registers"
	"legv8_assembler/types"
	"strconv"
	"strings"
)

func call_cb_format(instruction string, label_locations types.Labels) string {
	// opcode 8 location 19 condition register 5
	// I've zero clue how to implement cb.<condition format>, REMIND LATER
	fmt.Println("CB format")
	instruction_slice, after, _ := strings.Cut(instruction, " ")
	after = strings.TrimSpace(after)
	test_space := strings.Split(after, ",")

	if len(test_space) > 2 {
		fmt.Println("Error, invalid number of arguments")
	}

	// have a seperate one fo B.cond

	opcode, _ := isa.Instructions[instruction_slice]
	rd, rd_available := registers.RegistersBin[strings.ToUpper(strings.TrimSpace(test_space[0]))]
	if !rd_available {
		fmt.Println("Error with register, Invalid register")
	}

	_, invalid_label := registers.RegistersBin[strings.ToUpper(strings.TrimSpace(test_space[1]))]

	if invalid_label {
		fmt.Println("Registers cannot be labeled as labels")
	}

	label, label_a := label_locations[strings.TrimSpace(test_space[1])]
	if !label_a {
		fmt.Println("invalid label")
	}

	binary_label_location := strconv.FormatUint(uint64(label), 2)
	if len(binary_label_location) > 19 {
		fmt.Println("Label location too hard to find or sometshit fix this")
	}

	for x, y := 0, len(binary_label_location); x+y < 19; x++ {
		binary_label_location = "0" + binary_label_location
	}

	return opcode["op-code"] + binary_label_location + rd
}
