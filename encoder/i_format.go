package encoder

import (
	"fmt"
	"legv8_assembler/isa"
	"legv8_assembler/registers"
	"strconv"
	"strings"
)

func call_i_format(instruction string) string {
	fmt.Println("I format")
	instruction_slice, after, _ := strings.Cut(instruction, " ")
	after = strings.TrimSpace(after)
	test_space := strings.Split(after, ",")
	if len(test_space) != 3 {
		fmt.Println("error, number of arguments do not match")
	}
	opcode, _ := isa.Instructions[instruction_slice]
	rn, rn_available := registers.RegistersBin[strings.ToUpper(strings.TrimSpace(test_space[0]))]
	rd, rd_avaiable := registers.RegistersBin[strings.ToUpper(strings.TrimSpace(test_space[1]))]

	if !(rn_available && rd_avaiable) {
		fmt.Println("Invalid register")
	}

	string_imm := strings.Replace(strings.TrimSpace(test_space[2]), "#", "", 1)
	integer_imm, err := strconv.Atoi(string_imm)

	if integer_imm < 0 {
		fmt.Println("error with immediate, non unsigned integer value")
	}

	if err != nil {
		fmt.Println("error parsing and converting string to integer")
	}

	binary_imm := strconv.FormatUint(uint64(integer_imm), 2)

	if len(binary_imm) > 12 {
		fmt.Println("Problem with immediate, value too high")
	}

	for x, y := 0, len(binary_imm); x+y < 12; x++ {
		binary_imm = "0" + binary_imm
	}

	return opcode["op-code"] + binary_imm + rn + rd
}
