package main

import (
	"fmt"
	"legv8_assembler/isa"
	"os"
	"strconv"
	"strings"
)

// type Registers map[string]string /*just create a fun to dynamic find the bin instead of this*/

func main() {
	/*;)*/ cnt, err := os.ReadFile("test.asm")
	if err != nil {
		fmt.Println("OH FUCK")
	}
	var label_locations map[string]int = make(map[string]int)

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
			fmt.Println("1 ", m)
			inst, _, _ := strings.Cut(y, "//")
			fmt.Println("val:", strings.TrimSpace(inst))
			final_cut = append(final_cut, strings.TrimSpace(inst))
			loc += 1
			continue
		}

		if !strings.HasSuffix(z, ":") {
			fmt.Println("Invalid instruction ", z)
			// os.Exit(1)
			continue
		}
		prod, ok := isa.Instructions[strings.Trim(strings.ToUpper(z), ":")]

		if ok {
			fmt.Println("Illegal label ", z, " ", prod)
			continue
		}

		label_locations[z] = loc

		//I still have this problem "labelsomethign: instruction_next_to_it" send the instruction to the next line so that There won't be an location problems

		//this could be the solutioon
		// problem with y thats y new line always
		_, instruction, found := strings.Cut(strings.TrimSpace(y), " ")
		// final_cut = append(final_cut, label)
		loc += 1
		if found {
			fmt.Println("the found ", instruction, len(instruction))
			instruction, _, _ := strings.Cut(strings.TrimSpace(instruction), "//")
			final_cut = append(final_cut, strings.TrimSpace(instruction))
			loc += 1
		}

	}

	var final_binary string = ""

	for index, ins := range final_cut {
		if ins == "" {
			fmt.Println("Empty")
			continue
		}
		fmt.Println(index)
		ins = strings.TrimSpace(ins)
		instruction_slice, after, _ := strings.Cut(ins, " ")

		// what am I doibg here???????????????????????????????
		// instruction_slice2 := slices.DeleteFunc(instruction_slice, func(n string) bool {
		// 	// return (instruction_slice[n] == " " || instruction_slice[n] == "\n")
		// 	return strings.TrimSpace(n) == "" || n == " " || n == "\n"
		// })

		// Invalid instructions has been covered in labeling, don't do it again.
		// fmt.Println(strings.Join(ins, ""))
		// final_binary = final_binary + isa.Instructions[instruction_slice[0]]
		// fuck it check again if the instuction is there or not
		switch isa.Instructions[strings.ToUpper(instruction_slice)]["format"] {

		case isa.R_FORMAT:
			fmt.Println("here in r")
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
				register_number, valid := isa.RegistersBin[after]
				if !valid {
					// throw error
					fmt.Println("Invalid register for the instruction BR")
				}

				opcode, _ := isa.Instructions[instruction_slice]

				final_binary += opcode["op-code"] + testVar["rm"] + testVar["shamt"] + register_number + testVar["rd"]
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
				rd, available := isa.RegistersBin[strings.ToUpper(strings.TrimSpace(test_space[0]))]
				if !available {
					// throw error
					fmt.Println(strings.ToUpper(strings.TrimSpace(test_space[0])))
					fmt.Println("invalid register")
				}

				rn, available := isa.RegistersBin[strings.ToUpper(strings.TrimSpace(test_space[1]))]
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
				// what
				final_binary += isa.Instructions[instruction_slice]["op-code"] + "00000" + binary_shamt + rn + rd
			}

			temp := strings.Split(strings.TrimSpace(after), ",")

			if len(temp) != 3 {
				// throw error where we find there are not enough arguments fuck
				fmt.Println("Not enough arguments")
				break
			}

			//change name

			//change variable names
			for x, z := range temp {
				a, b := isa.RegistersBin[strings.ToUpper(strings.TrimSpace(z))]
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
			final_binary += opcode["op-code"] + testVar["rm"] + testVar["shamt"] + testVar["rn"] + testVar["rd"]
			break
		case isa.I_FORMAT:
			//opcode 10 immediate 12 rn 5 rd 5
			fmt.Println("I format")
			after = strings.TrimSpace(after)
			test_space := strings.Split(after, ",")
			if len(test_space) != 3 {
				fmt.Println("error, number of arguments do not match")
			}
			opcode, _ := isa.Instructions[instruction_slice]
			rn, rn_available := isa.RegistersBin[strings.ToUpper(strings.TrimSpace(test_space[0]))]
			rd, rd_avaiable := isa.RegistersBin[strings.ToUpper(strings.TrimSpace(test_space[1]))]

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

			final_binary += opcode["op-code"] + binary_imm + rn + rd

		case isa.D_FORMAT:
			fmt.Println("D format")
		case isa.CB_FORMAT:
			fmt.Println("CB format")
		case isa.IW_FORMAT:
			fmt.Println("IW FORMAT")
		case isa.B_FORMAT:
			// check for b if there then good

			// instruction_slice = slices.Delete(instruction_slice, 0, 1)
			// label := strings.TrimSpace(strings.Join(instruction_slice, ""))
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
			final_binary += isa.Instructions[strings.ToUpper(instruction_slice)]["op-code"] + location_binary
		}
	}

	fmt.Println(
		"final binary\n", final_binary,
	)

}
