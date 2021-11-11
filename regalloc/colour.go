package regalloc

import (
	"fmt"
	"log"

	"github.com/rj45/nanogo/ir"
	"github.com/rj45/nanogo/ir/reg"
)

func (ra *regAlloc) colour() {
	ra.Func.Blocks[0].VisitSuccessors(func(blk *ir.Block) bool {
		info := &ra.blockInfo[blk.ID]
		var used reg.Reg

		for id := range info.liveIns {
			val := ra.Func.ValueForID(id)
			used |= val.Reg
		}

		for _, val := range blk.Instrs {
			for _, id := range info.kills[val.ID] {
				free := ra.Func.ValueForID(id).Reg
				used &^= free
			}

			if val.Reg == reg.None {
				val.Reg = ra.chooseReg(info, val, used)
			}

			if val.Reg == reg.None {
				fmt.Println(blk.LongString())
				log.Fatal("Ran out of registers, spilling not implemented")
			}

			used |= val.Reg
		}

		fmt.Println(blk.LongString())

		return true
	})
}

var savedRegs = []reg.Reg{reg.S0, reg.S1, reg.S2, reg.S3}
var tempRegs = []reg.Reg{reg.T0, reg.T1, reg.T2, reg.T3, reg.T4, reg.T5}
var argRegs = []reg.Reg{reg.A0, reg.A1, reg.A2}

func (ra *regAlloc) chooseReg(info *blockInfo, val *ir.Value, used reg.Reg) reg.Reg {
	var chosen reg.Reg
	if len(ra.affinities[val.ID]) > 0 {
		votes := make(map[reg.Reg]int)
		for _, v := range ra.affinities[val.ID] {
			if v.Reg != reg.None && (used&v.Reg) == 0 {
				votes[v.Reg]++
			}
		}
		max := 0
		for reg, votes := range votes {
			if votes > max {
				max = votes
				chosen = reg
			}
		}
		if chosen != reg.None {
			return chosen
		}
	}

	sets := [][]reg.Reg{tempRegs, argRegs, savedRegs}
	if info.liveOuts[val.ID] {
		sets = [][]reg.Reg{savedRegs, tempRegs, argRegs}
	}

	for _, set := range sets {
		for _, reg := range set {
			if (used & reg) == 0 {
				return reg
			}
		}
	}

	return reg.None
}

// func safeToUse(val *ir.Value, info *blockInfo, val *ir.Value, used reg.Reg) bool {
// 	if info.liveIns[val.ID] &&
// }
