package parser

import (
	"context"
	"errors"
	"legv8_assembler/internal/encoder"
	"legv8_assembler/internal/isa"
	"legv8_assembler/internal/labels"
	"legv8_assembler/internal/types"
	"log"
	"strings"
)

// try creating a context thing it would be easier
var instructionCount int = 0
var instructionCounter int = 0

func ParseInstruction(instructionSlice []string, ctx *context.Context) (types.BinaryConversioninter, error) {
	if len(instructionSlice) > 4 {
		return nil, errors.New("SHUM ERROR")
	}

	mnemonic := instructionSlice[0]
	ins, ok := isa.Instructions[mnemonic]

	//check for labels here
	if !ok {
		err := checkForLabels(instructionSlice, ctx)
		if err != nil {
			return nil, err
		}

		instructionCount += 1
		return nil, nil
	}

	switch ins["format"] {
	case isa.R_FORMAT:
		rformat := encoder.RFormat{
			Opcode: ins["op-code"],
			Instruction: types.InstructionInformation{
				StringInstruction: strings.Join(instructionSlice, " "),
				BinaryInstruction: "",
			},
		}
		instructionCount += 1
		return &rformat, nil
	case isa.I_FORMAT:
		iformat := encoder.RFormat{
			Opcode: ins["op-code"],
			Instruction: types.InstructionInformation{
				StringInstruction: strings.Join(instructionSlice, " "),
				BinaryInstruction: "",
			},
		}
		instructionCount += 1
		return &iformat, nil
	case isa.D_FORMAT:
		dformat := encoder.DFormat{
			Opcode: ins["op-code"],
			Instruction: types.InstructionInformation{
				StringInstruction: strings.Join(instructionSlice, " "),
				BinaryInstruction: "",
			},
		}
		instructionCount += 1
		return &dformat, nil

	case isa.B_FORMAT:
		bformat := encoder.BFormat{
			Opcode: ins["op-code"],
			Instruction: types.InstructionInformation{
				StringInstruction: strings.Join(instructionSlice, " "),
				BinaryInstruction: "",
			},
		}
		instructionCount += 1
		return &bformat, nil

	case isa.CB_FORMAT:
		cbformat := encoder.CBFormat{
			Opcode: ins["op-code"],
			Instruction: types.InstructionInformation{
				StringInstruction: strings.Join(instructionSlice, " "),
				BinaryInstruction: "",
			},
		}
		instructionCount += 1
		return &cbformat, nil

	case isa.IW_FORMAT:
		instructionCount += 1
		return nil, nil

	default:
		log.Fatal("ERROR")
	}

	return nil, errors.New("SHUM ERROR")
}

func checkForLabels(instructionSlice []string, ctx *context.Context) error {
	if !strings.HasSuffix(strings.TrimSpace(instructionSlice[0]), ":") {
		return errors.New("Invalid syntax for label")
	}

	if len(instructionSlice) > 1 {
		ins := instructionSlice[1:]
		if !strings.HasPrefix(strings.Join(ins, " "), "//") {
			return errors.New("Label parsing error")
		}
	}

	if _, ok := labels.LabelLocation[strings.ReplaceAll(instructionSlice[0], ":", "")]; ok {
		return errors.New("Label already exists")
	}

	labels.LabelLocation[strings.ReplaceAll(instructionSlice[0], ":", "")] = instructionCount
	return nil
}
