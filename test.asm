// Assuming: 
// X0 = g, X1 = h, X2 = i, X3 = j
// X9, X10 are used as temporary registers
epstein:
def:
asdf:
leaf_example:  SUBI    SP, SP, #16
    // 1. Save old register states on the stack (Stack Pointer allocation)
            // Adjust stack pointer down by 16 bytes (2 doublewords)
hef:    STUR    X10, [SP, #8]       // Save temporary register X10 on the stack
    STUR    X9, [SP, #0]        // Save temporary register X9 on the stack

    // 2. Perform the actual calculation
    ADD     X9, X0, X1          // X9 = g + h
    ADD     X10, X2, X3         // X10 = i + j
    SUB     X0, X9, X10         // X0 = X9 - X10 (f = (g + h) - (i + j))
                                // Note: result is placed in X0 for return

    // 3. Restore old register states from the stack
    LDUR    X9, [SP, #0]        // Restore X9 for the calling function
    LDUR    X10, [SP, #8]       // Restore X10 for the calling function
    ADDI    SP, SP, #16         // Pop 16 bytes back off the stack

    // 4. Return to the caller
    BR      X30                 // Branch back to the address in Link Register