package encoder

import (
	"fmt"
	"legv8_assembler/isa"
	"legv8_assembler/registers"
	"strconv"
	"strings"
)

// expect the argument to be alreadly trimmed of spaces
func Call_r_format(instuction string) string {
	fmt.Println("here in r")
	instruction_slice, after, _ := strings.Cut(instuction, " ")
	var testVar map[string]string = map[string]string{
		"shamt": "000000",
		"rm":    "00000", //placeholders for a reason
		"rd":    "00000",
		"rn":    "00000",
	}
	// 3 different r formats, normal, with immediate and LR. considering there are only 3 odds, just use if
	fmt.Println("R format")

	if strings.EqualFold(strings.TrimSpace(instruction_slice), "BR") {

		// br register should be rn
		// do for this
		after = strings.TrimSpace(strings.ToUpper(after))
		register_number, valid := registers.RegistersBin[after]
		if !valid {
			// throw error
			fmt.Println("Invalid register for the instruction BR")
		}

		opcode, _ := isa.Instructions[instruction_slice]

		return opcode["op-code"] + testVar["rm"] + testVar["shamt"] + register_number + testVar["rd"]
	}

	if strings.EqualFold(strings.TrimSpace(instruction_slice), "LSL") ||
		strings.EqualFold(strings.TrimSpace(instruction_slice), "LSR") {
		after = strings.TrimSpace(after)
		test_space := strings.Split(after, ",")
		// should be op rm shamt and rd
		if len(test_space) != 3 {
			// throw error
			fmt.Println("Invalid amount of arguments")
		}
		rd, available := registers.RegistersBin[strings.ToUpper(strings.TrimSpace(test_space[0]))]
		if !available {
			// throw error
			fmt.Println(strings.ToUpper(strings.TrimSpace(test_space[0])))
			fmt.Println("invalid register")
		}

		rn, available := registers.RegistersBin[strings.ToUpper(strings.TrimSpace(test_space[1]))]
		if !available {
			// throw error
			fmt.Println(test_space[1])
			fmt.Println("invalid register")
		}

		if !strings.HasPrefix(test_space[2], "#") {
			// throw error
			fmt.Println("error")
		}

		string_shamt := strings.Replace(strings.TrimSpace(test_space[2]), "#", "", 1)
		integer_shamt, err := strconv.Atoi(string_shamt)

		if err != nil {
			fmt.Println("error parsing and converting string to integer")
		}

		if integer_shamt > 31 || integer_shamt < 0 {
			fmt.Println("Invalid number")
		}

		binary_shamt := strconv.FormatUint(uint64(integer_shamt), 2)
		for x, y := 0, len(binary_shamt); x+y < 6; x++ {
			binary_shamt = "0" + binary_shamt
		}
		// what
		return isa.Instructions[instruction_slice]["op-code"] + "00000" + binary_shamt + rn + rd
	}

	temp := strings.Split(strings.TrimSpace(after), ",")

	if len(temp) != 3 {
		// throw error where we find there are not enough arguments fuck
		fmt.Println("Not enough arguments")

	}

	//change name

	//change variable names
	for x, z := range temp {
		a, b := registers.RegistersBin[strings.ToUpper(strings.TrimSpace(z))]
		if !b {
			fmt.Println("Invalid register")
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
	// opcode rm (second operand) shamt rn (first operand) rd (destination)
	opcode, _ := isa.Instructions[instruction_slice]
	return opcode["op-code"] + testVar["rm"] + testVar["shamt"] + testVar["rn"] + testVar["rd"]
}
