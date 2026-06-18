package encoder

type InstructionInterface interface {
	InstructionInformation
}

//type BinaryConversion interface {
//BinaryConversion() error
//}

type InstructionInformation struct {
	StringInstruction string
	BinaryInstruction string
	//BinaryConversion  BinaryConversion
}

type RFormat struct {
	Opcode      string
	Rm          string
	Shamt       string
	Rn          string
	Rd          string
	Instruction InstructionInformation
}

type IFormat struct {
	Opcode      string
	Immediate   string
	Rn          string
	Rd          string
	Instruction InstructionInformation
}

type DFormat struct {
	Opcode      string
	Address     string
	Opcode2     string
	Rn          string
	Rd          string
	Instruction InstructionInformation
}

type BFormat struct {
	Opcode        string
	BranchAddress string
	Instruction   InstructionInformation
}

type CBFormat struct {
	Opcode        string
	BranchAddress string
	Rt            string
	Instruction   InstructionInformation
}

type IWFormat struct {
	Opcode       string
	MovImmediate string
	Rd           string
	Instruction  InstructionInformation
}
