#include "textflag.h"

// func mulSIMD(lhs, rhs *Vec4)
TEXT ·mulSIMD(SB),NOSPLIT,$0
  // load vector
  MOVQ    lhs+0(FP), R8
  MOVUPS  (R8), X0
  MOVQ    rhs+8(FP), R9
  MOVUPS  (R9), X1

  // multiply
  MULPS   X1, X0

  // save result back into vector
  MOVUPS  X0, (R8)
  RET
