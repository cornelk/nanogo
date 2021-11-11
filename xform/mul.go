package xform

import (
	"go/constant"
	"log"

	"github.com/rj45/nanogo/ir"
	"github.com/rj45/nanogo/ir/op"
	"github.com/rj45/nanogo/sizes"
)

func mulByConst(val *ir.Value) int {
	if val.Op != op.Mul {
		return 0
	}

	if val.Args[1].Op != op.Const {
		log.Println("op not const!")
		return 0
	}

	amt, ok := constant.Int64Val(val.Args[1].Value)
	if !ok {
		panic("expected int64 constant")
	}

	if amt == 1 {
		ir.SubstituteValue(val, val.Args[0])
		val.Block.RemoveInstr(val)
		return 1
	}

	if amt == 0 {
		ir.SubstituteValue(val, val.Args[1])
		val.Block.RemoveInstr(val)
		return 1
	}

	i := int64(1)
	n := int64(0)
	for i = 1; i < amt; i <<= 1 {
		n++
	}
	if i != amt {
		// TODO: can use multiple shifts and adds to calculate this
		return 0
	}

	val.Op = op.ShiftLeft
	val.Args[1] = val.Block.Func.Const(val.Args[1].Type, constant.MakeInt64(n))

	return 1
}

var _ = addToPass(Simplification, mulByConst)

func fixupConverts(val *ir.Value) int {
	if val.Op != op.Convert {
		return 0
	}

	if sizes.Sizeof(val.Args[0].Type) != sizes.Sizeof(val.Type) {
		log.Fatalf("Unable to convert %#v to %#v", val.Args[0].Type, val.Type)
	}

	ir.SubstituteValue(val, val.Args[0])
	val.Block.RemoveInstr(val)

	return 1
}

var _ = addToPass(Elaboration, fixupConverts)
