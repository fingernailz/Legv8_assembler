package encoder

import "legv8_assembler/internal/types"

type RFormat struct {
	Opcode      string
	Rm          string
	Shamt       string
	Rn          string
	Rd          string
	Instruction types.InstructionInformation
}

type IFormat struct {
	Opcode      string
	Immediate   string
	Rn          string
	Rd          string
	Instruction types.InstructionInformation
}

type DFormat struct {
	Opcode      string
	Address     string
	Opcode2     string
	Rn          string
	Rd          string
	Instruction types.InstructionInformation
}

type BFormat struct {
	Opcode        string
	BranchAddress string
	Instruction   types.InstructionInformation
}

type CBFormat struct {
	Opcode        string
	BranchAddress string
	Rt            string
	Instruction   types.InstructionInformation
}

type IWFormat struct {
	Opcode       string
	MovImmediate string
	Rd           string
	Instruction  types.InstructionInformation
}
