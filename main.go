package main

import (
	"fmt"

	// isa "legv8_assembler/isa"
	"os"
	"strings"
)

type ISA map[string]string

// type Registers map[string]string /*just create a fun to dynamic find the bin instead of this*/

var Instructions ISA = ISA{
	"ADD":  "10001011000",
	"SUB":  "11001011000",
	"ADDS": "10101011000",
	"SUBS": "11101011000",
	"AND":  "10001010000",
	"ANDS": "11101010000",
	"ORR":  "10101010000",
	"EOR":  "11101010000",
	"LSL":  "11010011011",
	"LSR":  "11010011010",
	"BR":   "11010110000",

	// I format
	"ADDI":  "1001000100",
	"SUBI":  "1101000100",
	"ADDIS": "1011000100",
	"SUBIS": "1111000100",
	"ANDI":  "1001001000",
	"ANDIS": "1111001000",
	"ORRI":  "1011001000",
	"EORI":  "1101001000",

	// D format
	"LDUR":   "11111100010",
	"STUR":   "11111100000",
	"LDURSW": "10111000100",
	"STURW":  "10111000000",
	"LDURH":  "01111000010",
	"STURH":  "01111000000",
	"LDURB":  "00111000010",
	"STURB":  "00111000000",
	"LDXR":   "11001000010",
	"STXR":   "11001000000",

	// IM format
	"MOVZ": "110100101",
	"MOVK": "111100101",

	// CB format

	"CBZ":  "10110100",
	"CBNZ": "10110101",
	// branch with condition, idk how to do, do later! opcode 01010100

	// B format
	"B":  "000101",
	"BL": "100101",
}

func init() {

}

// func find_labels(s string) map[string]int {
// 	loc := strings.Split(s, "\n")
// 	var locations map[string]int

// 	if loc == nil {
// 		return map[string]int
// 	}

// 	for x, y := range loc {
// 		// if strings.HasPrefix(strings.TrimSpace(y), "//")
// 	}
// }

func main() {
	/*;)*/ cnt, err := os.ReadFile("test.asm")
	if err != nil {
		fmt.Println("OH FUCK")
	}
	var test_var map[string]int = make(map[string]int)

	var loc int = 0

	//label part
	var final_cut []string = make([]string, len(strings.Split(string(cnt), "\n"))) // labels with instructions add more length for that

	for _, y := range strings.Split(string(cnt), "\n") {

		if y == "\n" || strings.HasPrefix(strings.TrimSpace(y), "//") {
			continue
		}

		// "  ADD x1, x1, x3" first remove the space from the sentance then split wrt space and trimspace for safety to get the instruction or label
		z := strings.TrimSpace(strings.Split(strings.TrimSpace(y), " ")[0])

		// there is an error with empty line where hassuffix function returns false fix it

		_, m := Instructions[z]
		// fmt.Print(strings.TrimSpace(z))
		if m {
			// if strings.Contains(y, "//") {
			// 	// usera := strings.Split(y, "//")[0]
			inst, _, _ := strings.Cut(y, "//")
			final_cut = append(final_cut, strings.TrimSpace(inst))
			// }

			// loc += 1
			// inal_cut = append(final_cut, splited_instruct[0])
			// // final_cut = append(final_cut, )
			// if strings.HasPrefix(splited_instruct[1], "//") {
			// 	continue
			// }
			// // splited_instruct =
			// final_cut = append*(final_cut, strings.Join(slices.Delete(splited_instruct, 0, 1), ""))
			loc += 1
			continue
		}

		if !strings.HasSuffix(z, ":") {
			fmt.Println("Invalid instruction ", z)
			// os.Exit(1)
			continue
		}
		prod, ok := Instructions[strings.Trim(z, ":")]

		if ok {
			fmt.Println("Invalid label ", z, " ", prod)
			continue
		}

		test_var[z] = loc

		//I still have this problem "labelsomethign: instruction_next_to_it" send the instruction to the next line so that There won't be an location problems

		//this could be the solutioon
		// splited_instruct := strings.Split(strings.TrimSpace(y), " ")
		// if len(splited_instruct) > 1 {
		// 	// final_cut = append(final_cut, splited_instruct[0])
		// 	// // final_cut = append(final_cut, )
		// 	// if strings.HasPrefix(splited_instruct[1], "//") {
		// 	// 	continue
		// 	// }
		// 	// // splited_instruct =
		// 	// final_cut = append(final_cut, strings.Join(slices.Delete(splited_instruct, 0, 1), ""))
		// 	// loc += 1

		// }

		label, instruction, found := strings.Cut(strings.TrimSpace(y), " ")
		final_cut = append(final_cut, label)
		loc += 1
		if found {
			instruction, _, _ := strings.Cut(strings.TrimSpace(instruction), "//")
			final_cut = append(final_cut, strings.TrimSpace(instruction))
			loc += 1
		}

	}
	fmt.Println(strings.Join(final_cut, "\n"))
	fmt.Println(test_var)

	// fmt.Println(cnt)
}
