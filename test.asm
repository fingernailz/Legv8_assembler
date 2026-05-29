START:
    // --- I-format Tests (Immediate operations) ---
    ADDI    X0, X31, #5      // X0 = 0 + 5    (Set up loop counter)
    ADDI    X1, X31, #10     // X1 = 0 + 10   (Base value)
    SUBI    X2, X31, #2      // X2 = 0 - 2    (Step value)

LOOP:
    // --- R-format Tests (Register operations) ---
    ADD     X1, X1, X0       // X1 = X1 + X0  (Accumulate)
    SUB     X0, X0, X2       // X0 = X0 - X2  (Modify loop counter)
    AND     X3, X1, X0       // X3 = X1 AND X0
    ORR     X4, X1, X3       // X4 = X1 OR X3

    // --- B-format Tests (Branching) ---
    B       LOOP             // Unconditional jump back to LOOP
