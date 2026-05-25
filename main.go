package main

import (
	"fmt"
	"legv8_assembler/isa"
	"os"
	"strings"
)

// type Registers map[string]string /*just create a fun to dynamic find the bin instead of this*/

func init() {

}

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

		// "  ADD x1, x1, x3" first remove the space from the sentance then split wrt space and trimspace for safety to get the instruction or label
		z := strings.TrimSpace(strings.Split(strings.TrimSpace(y), " ")[0])

		// there is an error with empty line where hassuffix function returns false fix it

		_, m := isa.Instructions[z]

		if m {
			inst, _, _ := strings.Cut(y, "//")
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
			fmt.Println("Invalid label ", z, " ", prod)
			continue
		}

		label_locations[z] = loc

		//I still have this problem "labelsomethign: instruction_next_to_it" send the instruction to the next line so that There won't be an location problems

		//this could be the solutioon
		label, instruction, found := strings.Cut(strings.TrimSpace(y), " ")
		final_cut = append(final_cut, label)
		loc += 1
		if found {
			instruction, _, _ := strings.Cut(strings.TrimSpace(instruction), "//")
			final_cut = append(final_cut, strings.TrimSpace(instruction))
			loc += 1
		}

	}

	// for index, ins := range final_cut {

	// }

	fmt.Println(strings.Join(final_cut, "\n"))
	fmt.Println(label_locations)

}
