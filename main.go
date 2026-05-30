package main

import (
	"errors"
	"fmt"
	"legv8_assembler/encoder"
	legv8_errors "legv8_assembler/errors"
	"legv8_assembler/isa"
	"legv8_assembler/types"
	"os"
	"strings"
)

// type Registers map[string]string /*just create a fun to dynamic find the bin instead of this*/

func main() {
	/*;)*/ cnt, err := os.ReadFile("test.asm")
	if err != nil {
		fmt.Println(err)
		return
	}
	var label_locations types.Labels = make(types.Labels)

	var loc int = 0

	// Labeling and formatting
	var final_cut []string = make([]string, len(strings.Split(string(cnt), "\n"))) // labels with instructions add more length for that

	for _, y := range strings.Split(string(cnt), "\n") {

		if /*(y == "\n"O ||*/ strings.HasPrefix(strings.TrimSpace(y), "//") {
			continue
		}

		if len(y) == 0 {
			continue
		}

		// "  ADD x1, x1, x3" first remove the space from the sentance then split wrt space and trimspace for safety to get the instruction or label
		z := strings.TrimSpace(strings.Split(strings.TrimSpace(y), " ")[0])

		// there is an error with empty line where hassuffix function returns false fix it

		_, m := isa.Instructions[strings.ToUpper(z)]

		if m {
			inst, _, _ := strings.Cut(y, "//")
			final_cut = append(final_cut, strings.TrimSpace(inst))
			loc += 1
			continue
		}

		if !strings.HasSuffix(z, ":") {
			fmt.Print(legv8_errors.Invalid_instruction)
			return
		}
		_, ok := isa.Instructions[strings.Trim(strings.ToUpper(z), ":")]

		if ok {
			fmt.Println(legv8_errors.Illegal_label)
		}

		label_locations[z] = loc

		//I still have this problem "labelsomethign: instruction_next_to_it" send the instruction to the next line so that There won't be an location problems

		//this could be the solutioon
		// problem with y thats y new line always
		_, instruction, found := strings.Cut(strings.TrimSpace(y), " ")
		// final_cut = append(final_cut, label)
		loc += 1
		if found {
			instruction, _, _ := strings.Cut(strings.TrimSpace(instruction), "//")
			final_cut = append(final_cut, strings.TrimSpace(instruction))
			loc += 1
		}
	}

	var final_binary string = ""

	for _, ins := range final_cut {
		if ins == "" {
			continue
		}
		ins = strings.TrimSpace(ins)
		instruction_slice, _, _ := strings.Cut(ins, " ")

		switch isa.Instructions[strings.ToUpper(instruction_slice)]["format"] {

		case isa.R_FORMAT:
			binary, err := encoder.Register_format(ins)

			if err != nil {
				fmt.Println(err)
				return
			}

			final_binary += binary
		case isa.I_FORMAT:
			binary, err := encoder.Immediate_format(ins)

			if err != nil {
				fmt.Println(err)
				return
			}

			final_binary += binary

		case isa.D_FORMAT:
			binary, err := encoder.Load_store_format(ins)

			if err != nil {
				fmt.Println(err)
				return
			}

			final_binary += binary
		case isa.CB_FORMAT:
			binary, err := encoder.Conditional_branch_format(ins, label_locations)

			if err != nil {
				fmt.Println(err)
				return
			}

			final_binary += binary
		case isa.IW_FORMAT:
			break
			// I'm not doing IW format fuck it
		case isa.B_FORMAT:
			binary, err := encoder.Branch_format(ins, label_locations)

			if err != nil {
				fmt.Println(err)
				return
			}

			final_binary += binary

		default:
			errors.New("Invalid Instuction").Error()
		}
	}

	fmt.Println(final_binary)

}
