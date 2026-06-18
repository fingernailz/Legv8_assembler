package encoder

import (
	"legv8_assembler/internal/errors"
	"legv8_assembler/internal/labels"
	"strconv"
	"strings"
)

// something is wrong idk FIX plz
func (b *BFormat) BinaryConversion() error {
	// check for b if there then good

	// instruction_slice = slices.Delete(instruction_slice, 0, 1)
	// label := strings.TrimSpace(strings.Join(instruction_slice, ""))
	_, after, comma := strings.Cut(b.Instruction.StringInstruction, " ")

	if !comma {
		return errors.Invalid_syntax
	}

	label := strings.TrimSpace(after)
	location, available := labels.LabelLocation[label]
	if !available {
		return errors.Invalid_label
	}

	location_binary := strconv.FormatUint(uint64(location), 2)

	if len(location_binary) > 26 {
		return errors.Label_location_out_of_bonds_error
	}

	for i := 0; i+len(location_binary) < 26; i++ {
		location_binary = "0" + location_binary
	}

	b.BranchAddress = location_binary

	// 6 opcode 26 Address
	return nil
}

func (instruction *BFormat) Assemble() {
	instruction.Instruction.BinaryInstruction =
		instruction.Opcode +
			instruction.BranchAddress
}
func (instruction *BFormat) GetBinary() string {
	return instruction.Instruction.BinaryInstruction
}
