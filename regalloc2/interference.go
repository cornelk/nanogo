package regalloc2

import (
	"github.com/rj45/nanogo/ir2"
)

type iNodeID uint32

type iGraph struct {
	nodes   []iNode
	valNode map[ir2.ID]iNodeID

	maxColour uint16
}

type iNode struct {
	val           ir2.ID
	interferences []iNodeID
	interferes    map[iNodeID]struct{}
	moves         []iNodeID

	colour uint16
	order  uint16
}

// buildInterferenceGraph takes the liveness information and builds a
// graph where nodes in the graph represent variables, and edges between
// the nodes represent variables that are live at the same time, in other
// words, variables that interfere with one another. This is done in order
// to aide in colouring the graph with non-interfering registers.
func (ra *RegAlloc) buildInterferenceGraph() {
	ig := &ra.iGraph

	ig.nodes = nil
	ig.valNode = make(map[ir2.ID]iNodeID)

	addNode := func(id ir2.ID) iNodeID {
		nodeID, found := ig.valNode[id]

		if !id.ValueIn(ra.fn).NeedsReg() {
			panic("attempt to add non reg value: " + id.ValueIn(ra.fn).IDString())
		}

		if !found {
			nodeID = iNodeID(len(ig.nodes))
			ig.nodes = append(ig.nodes, iNode{
				val: id,
			})
			ig.valNode[id] = nodeID
		}
		return nodeID
	}

	addEdge := func(var1 ir2.ID, var2 ir2.ID) {
		node1ID := addNode(var1)
		node2ID := addNode(var2)

		if var1 == var2 {
			// don't add edges between ourself
			return
		}

		for _, pair := range [2][2]iNodeID{{node1ID, node2ID}, {node2ID, node1ID}} {
			node := &ig.nodes[pair[0]]
			neighbor := pair[1]
			if _, found := node.interferes[neighbor]; !found {
				// add it to the interferences list
				node.interferences = append(node.interferences, neighbor)

				if node.interferes == nil {
					node.interferes = make(map[iNodeID]struct{})
				}

				// and the interferes map
				node.interferes[neighbor] = struct{}{}
			}
		}
	}

	addMove := func(var1 ir2.ID, var2 ir2.ID) {
		node1ID := addNode(var1)
		node2ID := addNode(var2)

		if var1 == var2 {
			// don't add moves between ourself
			return
		}

		for _, pair := range [2][2]iNodeID{{node1ID, node2ID}, {node2ID, node1ID}} {
			node := &ig.nodes[pair[0]]
			neighbor := pair[1]

			found := false
			for _, id := range node.moves {
				if id == neighbor {
					found = true
				}
			}

			if !found {
				// add it to the moves list
				node.moves = append(node.moves, neighbor)
			}
		}
	}

	fn := ra.fn
	for i := 0; i < fn.NumBlocks(); i++ {
		blk := fn.Block(i)
		info := ra.info[blk.Index()]

		live := make(map[ir2.ID]struct{})
		for k := range info.liveOuts {
			live[k] = struct{}{}
		}

		// block args are live immediately before leaving the block
		// and there is an implicit move between them and the defs of
		// succ blocks
		offset := 0
		for s := 0; s < blk.NumSuccs(); s++ {
			succ := blk.Succ(s)
			for d := 0; d < succ.NumDefs(); d++ {
				def := succ.Def(d)
				arg := blk.Arg(offset + d)

				if def.NeedsReg() {
					// arg is live now
					live[arg.ID] = struct{}{}

					if arg.NeedsReg() {
						// mark a move from the arg to the def
						addMove(arg.ID, def.ID)
					}
				}
			}

			offset += succ.NumDefs()
		}

		// all currently live variables interfere
		for id1 := range live {
			for id2 := range live {
				if id1 != id2 {
					addEdge(id1, id2)
				}
			}
		}

		for j := blk.NumInstrs() - 1; j >= 0; j-- {
			instr := blk.Instr(j)

			// all defs interfere with one another, so removing it from the
			// live set should be done after adding edges
			for d := 0; d < instr.NumDefs(); d++ {
				def := instr.Def(d)
				if def.NeedsReg() {
					// make sure the node is in the graph, even if there's no
					// other live values at the time
					addNode(def.ID)

					// make sure all live vars are marked as interfering
					for id := range live {
						addEdge(def.ID, id)
					}

					// if it's a move (aka copy)
					if instr.Op.IsCopy() && instr.Arg(d).NeedsReg() {
						// add the move between the corresponding defs and args
						addMove(def.ID, instr.Arg(d).ID)
					}
				}
			}

			// now we can remove each def from the live set
			for d := 0; d < instr.NumDefs(); d++ {
				def := instr.Def(d)
				if def.NeedsReg() {
					// def is now no longer live
					delete(live, def.ID)
				}
			}

			// todo: for calls, make sure all live variables at the call site interfere
			// with caller saved registers... that is, add some fake pre-coloured nodes
			// and edges to all live variables

			// mark each used arg as now live
			for u := 0; u < instr.NumArgs(); u++ {
				use := instr.Arg(u)
				if use.NeedsReg() {
					live[use.ID] = struct{}{}
				}
			}
		}
	}
}

// findPerfectEliminationOrder finds the perfect elimination order by
// using the max cardinality search algorithm. This is done because
// the graph should be chordal thanks to SSA. Chordal graphs can
// be optimally coloured in reverse perfect elimination order.
// There are other algorithms that could find the PEO as well,
// such as lexicographic breadth first search. This seemed simpler
// though it may be slower (not sure).
func (ig *iGraph) findPerfectEliminationOrder() []iNodeID {
	marked := make(map[iNodeID]struct{})
	output := make([]iNodeID, 0, len(ig.nodes))
	unmarked := make([]iNodeID, len(ig.nodes))
	for i := range ig.nodes {
		unmarked[i] = iNodeID(i)
	}

	// for each unmarked node
	for len(unmarked) > 0 {
		// find the unmarked node with the most marked neighbors
		maxNode := unmarked[0]
		maxI := 0
		maxCard := -1
		for i, cand := range unmarked {
			card := 0
			for _, neighbor := range ig.nodes[cand].interferences {
				if _, found := marked[neighbor]; found {
					card++
				}
			}
			if card > maxCard {
				maxI = i
				maxNode = cand
				maxCard = card
			}
		}

		// remove node from unmarked list. Order doesn't matter
		// so the faster way of removing an item from the slice works.
		unmarked[maxI] = unmarked[len(unmarked)-1]
		unmarked = unmarked[:len(unmarked)-1]

		// mark the node
		marked[maxNode] = struct{}{}

		// add node to output
		output = append(output, maxNode)

		ig.nodes[maxNode].order = uint16(len(output) - 1)
	}

	return output
}

func (ig *iGraph) pickColours() {
	order := ig.findPerfectEliminationOrder()

	// pick colours in reverse perfect elimination order
	for i := len(order) - 1; i >= 0; i-- {
		nodeID := order[i]
		node := &ig.nodes[nodeID]
		node.pickColour(ig)
	}
}

const noColour uint16 = 0

func (nd *iNode) pickColour(ig *iGraph) {
	if nd.colour != noColour {
		// already coloured
		return
	}

	// first try to pick a move colour if that colour doesn't
	// interfere with any others
	for _, mv := range nd.moves {
		moveColour := ig.nodes[mv].colour

		// skip if the move node has not already been assigned a colour
		if moveColour == noColour || moveColour == dontColour {
			continue
		}

		// check if that colour interferes with any neighbors
		interferes := false
		for nb := range nd.interferes {
			if ig.nodes[nb].colour == moveColour {
				interferes = true
				break
			}
		}

		// if it doesn't interfere, then choose that colour
		if !interferes {
			nd.colour = moveColour
			return
		}
	}

	// find the lowest numbered colour that doesn't interfere
	for colour := uint16(1); ; colour++ {
		interferes := false

		// for each neighbour in the interferences
		for nb := range nd.interferes {
			// if the neighbour already has this colour
			if ig.nodes[nb].colour == colour {
				// then it interferes and we can't use it
				interferes = true
				break
			}
		}

		// if it doesn't interfere then
		if !interferes {
			// choose the colour
			nd.colour = colour

			// keep track of the largest chosen colour
			if ig.maxColour < colour {
				ig.maxColour = colour
			}

			return
		}
	}
}
