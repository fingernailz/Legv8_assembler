package encoder

import (
	"legv8_assembler/internal/errors"
	"legv8_assembler/internal/registers"
	"strconv"
	"strings"
)

func (imm *IFormat) BinaryConversion() error {
	_, after, _ := strings.Cut(imm.Instruction.StringInstruction, " ")
	after = strings.TrimSpace(after)
	test_space := strings.Split(after, ",")
	if len(test_space) != 3 {
		return errors.Invalid_Number_of_Operands
	}
	rn, rn_available := registers.RegistersBin[strings.ToUpper(strings.TrimSpace(test_space[0]))]
	rd, rd_avaiable := registers.RegistersBin[strings.ToUpper(strings.TrimSpace(test_space[1]))]

	if !(rn_available && rd_avaiable) {
		return errors.Invalid_register
	}

	// check if # is present or not

	if !strings.HasPrefix(strings.TrimSpace(test_space[2]), "#") {
		return errors.Immediate_syntax_error
	}

	string_imm := strings.Replace(strings.TrimSpace(test_space[2]), "#", "", 1)
	integer_imm, err := strconv.Atoi(string_imm)

	if integer_imm < 0 {
		return errors.Signed_immediate_error
	}

	if err != nil {
		return errors.Immediate_parsing_error
	}

	binary_imm := strconv.FormatUint(uint64(integer_imm), 2)

	if len(binary_imm) > 12 {
		return errors.Immediate_value_error
	}

	for x, y := 0, len(binary_imm); x+y < 12; x++ {
		binary_imm = "0" + binary_imm
	}

	imm.Rn = rn
	imm.Rd = rd
	imm.Immediate = binary_imm

	return nil
}

func (instruction *IFormat) Assemble() {
	// opcode + immediate + rn + rd
	instruction.Instruction.BinaryInstruction =
		instruction.Opcode +
			instruction.Immediate +
			instruction.Rn +
			instruction.Rd
}
