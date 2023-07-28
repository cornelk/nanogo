package regalloc2

import (
	"fmt"
	"os"
	"strings"

	"github.com/rj45/nanogo/ir2"
)

// WriteGraphvizCFG emits a Graphviz dot file of the CFG
func WriteGraphvizCFG(ra *RegAlloc) {
	dot, _ := os.Create(ra.fn.Name + ".cfg.dot")
	defer dot.Close()

	fmt.Fprintln(dot, "digraph G {")
	fmt.Fprintln(dot, "labeljust=l;")
	fmt.Fprintln(dot, "node [shape=record, fontname=\"Noto Mono\", labeljust=l];")

	for bn := 0; bn < ra.fn.NumBlocks(); bn++ {
		blk := ra.fn.Block(bn)

		// info := &ra.blockInfo[blk.ID()]

		for i := 0; i < blk.NumPreds(); i++ {
			pred := blk.Pred(i)
			// pinfo := &ra.blockInfo[pred.ID()]
			// outs := maptolist(pinfo.liveOuts) + " - " + maptolist(pinfo.phiOuts[blk])
			// ins := maptolist(info.liveIns) + " - " + maptolist(info.phiIns[pred])
			fmt.Fprintf(dot, "%s -> %s;\n", pred, blk)
		}

		liveInKills := ""
		label := fmt.Sprintf("%s:\\l", blk)
		for i := 0; i < blk.NumInstrs(); i++ {
			instr := blk.Instr(i)

			kills := ""

			// for i, kill := range info.kills[instr] {
			// 	if i != 0 {
			// 		kills += " "
			// 	}
			// 	kills += kill.IDString()
			// }

			label += fmt.Sprintf("%s [%s]\\l", instr.LongString(), kills)
		}

		// first := true
		// kills := ""
		// for kill := range info.blkKills {
		// 	if !first {
		// 		kills += " "
		// 	}
		// 	first = false
		// 	kills += kill.IDString()
		// }

		// label += fmt.Sprintf("%s [%s]\\l" /* blk.OpString()*/, "", kills)

		label = strings.ReplaceAll(label, "\"", "\\\"")
		label = strings.ReplaceAll(label, "{", "\\{")
		label = strings.ReplaceAll(label, "}", "\\}")

		fmt.Fprintf(dot, "%s [label=\"%s\"];\n", blk, label)
		fmt.Fprintln(dot, liveInKills)
	}

	fmt.Fprintln(dot, "}")
}

// WriteGraphvizInterferenceGraph emits a Graphviz dot file of the Interference Graph
func WriteGraphvizInterferenceGraph(ra *RegAlloc) {
	dot, _ := os.Create(ra.fn.Name + ".igraph.dot")
	defer dot.Close()

	fmt.Fprintln(dot, "graph G {")
	// fmt.Fprintln(dot, "labeljust=l;")
	// fmt.Fprintln(dot, "node [shape=record, fontname=\"Noto Mono\", labeljust=l];")

	edges := map[string]bool{}

	for _, nodeA := range ra.iGraph.nodes {
		label := fmt.Sprintf("%s\n%d:%d", nodeA.val.IDString(), nodeA.order, nodeA.colour)
		fmt.Fprintf(dot, "%s [label=%q];\n", nodeA.val.IDString(), label)

		for _, nodeBid := range nodeA.interferences {
			nodeB := ra.iGraph.nodes[nodeBid]
			if !edges[nodeB.val.IDString()+"--"+nodeA.val.IDString()] {
				edge := nodeA.val.IDString() + "--" + nodeB.val.IDString()
				fmt.Fprintf(dot, "%s;\n", edge)
				edges[edge] = true
			}
		}

	}

	fmt.Fprintln(dot, "}")
}

func DumpLivenessChart(ra *RegAlloc) {
	html, _ := os.Create(ra.fn.Name + ".liveness.html")
	defer html.Close()

	fmt.Fprintln(html, `
<html>
<head>
	<style>
		td {
			border: 1px solid black;
			font-family: monospace;
			padding: 1px 5px;
		}
	</style>
</head>
<body>
	<table>`)

	it := ra.fn.InstrIter()
	var blk *ir2.Block
	num := uint32(0)

	for ; it.HasNext(); it.Next() {
		newblock := false
		if it.Block() != blk {
			if blk != nil {
				fmt.Fprintf(html, "<tr><td colspan=\"%d\">out:\n", 4)
				for id := range ra.info[blk.Index()].liveOuts {
					fmt.Fprint(html, " ", id.ValueIn(ra.fn).String())
				}
				fmt.Fprintln(html, "</td></tr>")
			}
			blk = it.Block()
			fmt.Fprintf(html, "<tr><td colspan=\"%d\">in:\n", 4)
			for id := range ra.info[blk.Index()].liveIns {
				fmt.Fprint(html, " ", id.ValueIn(ra.fn).String())
			}
			fmt.Fprintln(html, "</td></tr>")

			fmt.Fprintf(html, "<tr><td colspan=\"%d\">%s(", 4, blk)
			for i, val := range blk.Defs() {
				if i != 0 {
					fmt.Fprint(html, ", ")
				}
				fmt.Fprintf(html, "%s", val)
			}
			fmt.Fprintln(html, "):</td></tr>")

			blk.Args()

			newblock = true
		}

		fmt.Fprintln(html, "<tr>")

		if newblock {
			fmt.Fprintln(html, "<td rowspan=\"", blk.NumInstrs()*2, "\">", blk.String(), "</td>")
		}
		fmt.Fprintln(html, "<td>", num, "</td>")

		fmt.Fprintln(html, "<td rowspan=\"2\">")
		fmt.Fprintln(html, it.Instr().LongString())

		fmt.Fprintln(html, "</td>")
		fmt.Fprintln(html, "</tr>")

		num++

		fmt.Fprintln(html, "<tr><td>", num, "</td>")
		fmt.Fprintln(html, "</tr>")
		num++
	}

	fmt.Fprintln(html, "</table>")
	fmt.Fprintln(html, "</body>")
	fmt.Fprintln(html, "</html>")
}
