package types

type Labels map[string]int
type Registers map[string]string
type ISA map[string]map[string]string

type BinaryConversioninter interface {
	BinaryConversion() error
	Assemble()
	GetBinary() string
}

type InstructionInterface interface {
	InstructionInformation
}

type InstructionInformation struct {
	StringInstruction string
	BinaryInstruction string
}
