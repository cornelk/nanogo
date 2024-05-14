package xform

import (
	"fmt"
	"github.com/rj45/nanogo/ir"
	"github.com/rj45/nanogo/ir/op"
)

func AddPhiCopies(val *ir.Value) int {
	if val.Op != op.Phi {
		return 0
	}

	changes := 0

	for i := 0; i < val.NumArgs(); i++ {
		src := val.Arg(i)

		if src.Op == op.PhiCopy {
			continue
		}

		copy := val.Func().NewValue(op.PhiCopy, src.Type, src)
		pred := val.Block().Pred(i)

		pred.InsertInstr(-1, copy)
		val.ReplaceArg(i, copy)
		changes++
		fmt.Println("debug: AddPhiCopies")
	}

	return changes
}

var _ = addToPass(Lowering, AddPhiCopies)
