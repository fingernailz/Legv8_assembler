package main

import (
	"context"
	"fmt"
	"legv8_assembler/internal/lexer"
	"legv8_assembler/internal/types"
	"log"
	"os"
	"sync"
)

func worker(jobs <-chan types.BinaryConversioninter, wg *sync.WaitGroup) {
	defer wg.Done()
	for instruction := range jobs {
		if instruction == nil {
			continue
		}
		err := instruction.BinaryConversion()
		if err != nil {
			log.Fatal(err)
		}
		instruction.Assemble()
	}
}

// func worker(jobs <- `j`)

// type Registers map[string]string /*just create a fun to dynamic find the bin instead of this*/o

func someinterface(w types.InstructionInformation) string {
	return w.BinaryInstruction
}

func main() {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	ctx = context.WithValue(ctx, "labellocations", types.Labels{})
	defer cancel()
	file, err := os.ReadFile("test.asm")
	if err != nil {
		log.Fatal(err)
	}
	ctx = context.WithValue(ctx, "filesize", []any{})

	instructionSlice := lexer.LexerInit(file, &ctx)
	wg := sync.WaitGroup{}

	jobchannel := make(chan types.BinaryConversioninter)

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(jobchannel, &wg)
	}

	// go homieWorker()

	for _, y := range instructionSlice {
		fmt.Println(y)
		jobchannel <- y
	}
	close(jobchannel)
	//3 worker would be fine

	wg.Wait()

	cntToWrite := ""
	for _, y := range instructionSlice {
		if y == nil {
			continue
		}
		cntToWrite += y.GetBinary()
	}

	smt := os.WriteFile("output.o", []byte(cntToWrite), 0644)

	if smt != nil {
		log.Fatal(smt)
	}

	fmt.Println("CHECK SOMEHTING")
}
