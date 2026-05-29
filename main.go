package main

import (
	"fmt"
	"legv8_assembler/isa"
	"os"
	"slices"
	"strings"
)

func convert_numbers_to_binary(num int) string {
	// var inital_string string = "b"
	// var num_float float64 = float64(num)
	// for num > 0 {
	// 	inital_string += string(int(num_float) % 2)
	// 	num_float = num_float / 2
	// 	num = num / 2
	// }

	// return inital_string
	if num == 0 || num == 1 {
		return string(num)
	}

	return "" + convert_numbers_to_binary(num/2)
}

// type Registers map[string]string /*just create a fun to dynamic find the bin instead of this*/

// const (

// )

// func

// func init() {

// }

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

		_, m := isa.Instructions[z]

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
		prod, ok := isa.Instructions[strings.Trim(z, ":")]

		if ok {
			fmt.Println("Illegal label ", z, " ", prod)
			continue
		}

		label_locations[z] = loc
		fmt.Println(z)

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
		instruction_slice := strings.Split(ins, " ")
		// what am I doibg here???????????????????????????????
		// instruction_slice2 := slices.DeleteFunc(instruction_slice, func(n string) bool {
		// 	// return (instruction_slice[n] == " " || instruction_slice[n] == "\n")
		// 	return strings.TrimSpace(n) == "" || n == " " || n == "\n"
		// })

		// Invalid instructions has been covered in labeling, don't do it again.
		// fmt.Println(strings.Join(ins, ""))
		// final_binary = final_binary + isa.Instructions[instruction_slice[0]]
		// fuck it check again if the instuction is there or not
		switch isa.Instructions[strings.ToUpper(instruction_slice[0])]["format"] {

		case isa.R_FORMAT:
			// 3 different r formats, normal, with immediate and LR. considering there are only 3 odds, just use if
			fmt.Println("R format")

			if strings.EqualFold(strings.TrimSpace(instruction_slice[0]), "BR") {
				// do for this
			}

			if strings.EqualFold(strings.TrimSpace(instruction_slice[0]), "LSL") ||
				strings.EqualFold(strings.TrimSpace(instruction_slice[0]), "LSR") {
				// do something
			}

		case isa.I_FORMAT:
			fmt.Println("I format")

		case isa.D_FORMAT:
			fmt.Println("D format")
		case isa.CB_FORMAT:
			fmt.Println("CB format")
		case isa.IW_FORMAT:
			fmt.Println("IW FORMAT")
		case isa.B_FORMAT:
			// check for b if there then good

			// 6 opcode 26 Address
			final_binary += isa.Instructions[strings.ToUpper(instruction_slice[0])]["op-code"]
			instruction_slice = slices.Delete(instruction_slice, 0, 1)
			label := strings.TrimSpace(strings.Join(instruction_slice, ""))
			label = label
			// location, available := label_locations[label]
			// if !available {
			// 	fmt.Println("Invalid label ", label)
			// }

			final_binary += "0000000000000000000000"
		}
	}

	fmt.Println("converted number to binary", convert_numbers_to_binary(100))
	fmt.Println(strings.Join(final_cut, ""))
	fmt.Println(len(final_cut))
	for x, y := range final_cut {
		fmt.Println(x, y)
	}
	fmt.Println(label_locations)
	fmt.Println(
		"final binary\n", final_binary,
	)

}
