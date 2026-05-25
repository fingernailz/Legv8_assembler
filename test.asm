SUBI SP,SP,#24    // move the stack pointer up 24 
STUR X10,[SP,#16] // store at botom of stack
STUR X9,[SP,#8]   // store at middle of stack
STUR X19,[SP,#0]  // store at top of stack
ADD X9,X0,X1
ADD X10,X2,X3
SUB X19,X9,X10
ADD X0,X19,XZR
LDUR X10,[SP,#16] // get from bottom of stack
LDUR X9,[SP,#8]   // get from middle of stack
LDUR X19,[SP,#0]  // get from top of stack
ADDI SP,SP,#24    // move the stack pointer down 24
BR LR