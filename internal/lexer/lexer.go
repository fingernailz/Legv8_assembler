package lexer

import (
	"bufio"
	"bytes"
	"context"
	"legv8_assembler/internal/parser"
	"log"
	"strings"
)

func Lexer(instruction string, sliceptr *[]any, ctx *context.Context) {
	//to check if it a comment
	if strings.HasPrefix(instruction, "//") {
		return
	}

	if len(instruction) == 0 {
		return
	}

	instruction = strings.TrimSpace(instruction)
	instruction2, _, _ := strings.Cut(instruction, "//")
	instructionSlice := strings.Fields(strings.ToUpper(instruction2))

	out, err := parser.ParseInstruction(instructionSlice, ctx)
	if err != nil {
		// context.Background().Done()u
		log.Fatal(err)
	}

	*sliceptr = append(*sliceptr, out)
	// fmt.Println(len(instructionSlice))
	//call the parser

}

func LexerInit(file []byte, ctx *context.Context) []any {
	sliceptr := make([]any, bytes.Count(file, []byte{'\n'}))
	reader := bytes.NewReader(file)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		Lexer(scanner.Text(), &sliceptr, ctx)
	}

	return sliceptr
}
