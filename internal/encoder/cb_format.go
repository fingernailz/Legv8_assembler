package encoder

import (
	"legv8_assembler/internal/errors"
	"legv8_assembler/internal/labels"
	"legv8_assembler/internal/registers"
	"strconv"
	"strings"
)

func (cb *CBFormat) BinaryConversion() error {
	// opcode 8 location 19 condition register 5
	// I've zero clue how to implement cb.<condition format>, REMIND LATER
	_, after, _ := strings.Cut(cb.Instruction.StringInstruction, " ")
	after = strings.TrimSpace(after)
	test_space := strings.Split(after, ",")

	if len(test_space) > 2 {
		return errors.Invalid_Number_of_Operands
	}

	// have a seperate one fo B.cond

	rt, rt_available := registers.RegistersBin[strings.ToUpper(strings.TrimSpace(test_space[0]))]
	if !rt_available {
		return errors.Invalid_register
	}

	_, invalid_label := registers.RegistersBin[strings.ToUpper(strings.TrimSpace(test_space[1]))]

	if invalid_label {
		return errors.Label_register_conflict
	}

	label, label_a := labels.LabelLocation[strings.TrimSpace(test_space[1])]
	if !label_a {
		return errors.Invalid_label
	}

	binary_label_location := strconv.FormatUint(uint64(label), 2)
	if len(binary_label_location) > 19 {
		return errors.Label_location_out_of_bonds_error
	}

	for x, y := 0, len(binary_label_location); x+y < 19; x++ {
		binary_label_location = "0" + binary_label_location
	}

	cb.BranchAddress = binary_label_location
	cb.Rt = rt

	return nil
}

func (instruction *CBFormat) Assemble() {
	instruction.Instruction.BinaryInstruction =
		instruction.Opcode +
			instruction.BranchAddress +
			instruction.Rt
}

func (instruction *CBFormat) GetBinary() string {
	return instruction.Instruction.BinaryInstruction
}
