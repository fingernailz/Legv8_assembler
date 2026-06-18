package encoder

import (
	"legv8_assembler/internal/errors"
	"legv8_assembler/internal/registers"
	"strconv"
	"strings"
)

// expect the argument to be alreadly trimmed of spaces
func (rf *RFormat) BinaryConversion() error {
	instruction_slice, after, _ := strings.Cut(rf.Instruction.StringInstruction, " ")
	var testVar map[string]string = map[string]string{
		"shamt": "000000",
		"rm":    "00000", //placeholders for a reason
		"rd":    "00000",
		"rn":    "00000",
	}
	// 3 different r formats, normal, with immediate and LR. considering there are only 3 odds, just use if

	if strings.EqualFold(strings.TrimSpace(instruction_slice), "BR") {

		// br register should be rn
		// do for this
		after = strings.TrimSpace(strings.ToUpper(after))
		//wat the fuck is this, fair engh
		rn, valid := registers.RegistersBin[after]
		if !valid {
			// throw error
			return errors.Invalid_register
		}

		rf.Rn = rn
		rf.Rm = testVar["rm"]
		rf.Rd = testVar["rd"]
		rf.Shamt = testVar["shamt"]

		return nil
	}

	if strings.EqualFold(strings.TrimSpace(instruction_slice), "LSL") ||
		strings.EqualFold(strings.TrimSpace(instruction_slice), "LSR") {
		after = strings.TrimSpace(after)
		test_space := strings.Split(after, ",")
		// should be op rm shamt and rd
		if len(test_space) != 3 {
			return errors.Invalid_Number_of_Operands
		}
		rd, available := registers.RegistersBin[strings.ToUpper(strings.TrimSpace(test_space[0]))]
		if !available {
			return errors.Invalid_register
		}

		rn, available := registers.RegistersBin[strings.ToUpper(strings.TrimSpace(test_space[1]))]
		if !available {
			return errors.Invalid_register
		}

		if !strings.HasPrefix(test_space[2], "#") {
			return errors.Immediate_syntax_error
		}

		string_shamt := strings.Replace(strings.TrimSpace(test_space[2]), "#", "", 1)
		integer_shamt, err := strconv.Atoi(string_shamt)

		if err != nil {
			return errors.Shamt_parsing_error
		}

		if integer_shamt > 31 || integer_shamt < 0 {
			return errors.Shamt_value_error
		}

		binary_shamt := strconv.FormatUint(uint64(integer_shamt), 2)
		for x, y := 0, len(binary_shamt); x+y < 6; x++ {
			binary_shamt = "0" + binary_shamt
		}

		rf.Rd = rd
		rf.Rn = rn
		rf.Rm = testVar["rm"]
		rf.Shamt = binary_shamt
		// what
		return nil
	}

	temp := strings.Split(strings.TrimSpace(after), ",")

	if len(temp) != 3 {
		// throw error where we find there are not enough arguments fuck
		return errors.Invalid_Number_of_Operands
	}

	//change name

	//change variable names
	for x, z := range temp {
		a, b := registers.RegistersBin[strings.ToUpper(strings.TrimSpace(z))]
		if !b {
			return errors.Invalid_register
			// throw an error
		}
		testVar[func() string {
			if x == 0 {
				return "rd"
			} else if x == 1 {
				return "rn"
			}
			return "rm"
		}()] = a
	}

	rf.Rn = testVar["rn"]
	rf.Rm = testVar["rm"]
	rf.Rd = testVar["rd"]
	rf.Shamt = testVar["shamt"]
	// opcode rm (second operand) shamt rn (first operand) rd (destination)
	return nil
}

func (instruction *RFormat) Assemble() {
	//opcode + rm + shamt + rn + rd
	instruction.Instruction.BinaryInstruction =
		instruction.Opcode +
			instruction.Rm +
			instruction.Shamt +
			instruction.Rn +
			instruction.Rd
}
