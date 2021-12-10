// Code generated by "enumer -type=Op -transform title-lower"; DO NOT EDIT.

package op

import (
	"fmt"
	"strings"
)

const _OpName = "invalidbuiltincallcallBuiltinchangeInterfacechangeTypeconstconvertcopyextractfieldfieldAddrfreeVarfuncglobalindexindexAddrinlineAsmlocallookupmakeInterfacemakeSlicenextnewparameterphiphiCopyrangeregslicesliceToArrayPointerstoreswapInswapOuttypeAssertaddsubmuldivremandorxorshiftLeftshiftRightandNotequalnotEquallesslessEqualgreatergreaterEqualnotnegateloadinvertjump2if2return2panic2ifEqual2ifNotEqual2ifLess2ifLessEqual2ifGreater2ifGreaterEqual2numOps"

var _OpIndex = [...]uint16{0, 7, 14, 18, 29, 44, 54, 59, 66, 70, 77, 82, 91, 98, 102, 108, 113, 122, 131, 136, 142, 155, 164, 168, 171, 180, 183, 190, 195, 198, 203, 222, 227, 233, 240, 250, 253, 256, 259, 262, 265, 268, 270, 273, 282, 292, 298, 303, 311, 315, 324, 331, 343, 346, 352, 356, 362, 367, 370, 377, 383, 391, 402, 409, 421, 431, 446, 452}

const _OpLowerName = "invalidbuiltincallcallbuiltinchangeinterfacechangetypeconstconvertcopyextractfieldfieldaddrfreevarfuncglobalindexindexaddrinlineasmlocallookupmakeinterfacemakeslicenextnewparameterphiphicopyrangeregsliceslicetoarraypointerstoreswapinswapouttypeassertaddsubmuldivremandorxorshiftleftshiftrightandnotequalnotequallesslessequalgreatergreaterequalnotnegateloadinvertjump2if2return2panic2ifequal2ifnotequal2ifless2iflessequal2ifgreater2ifgreaterequal2numops"

func (i Op) String() string {
	if i < 0 || i >= Op(len(_OpIndex)-1) {
		return fmt.Sprintf("Op(%d)", i)
	}
	return _OpName[_OpIndex[i]:_OpIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _OpNoOp() {
	var x [1]struct{}
	_ = x[Invalid-(0)]
	_ = x[Builtin-(1)]
	_ = x[Call-(2)]
	_ = x[CallBuiltin-(3)]
	_ = x[ChangeInterface-(4)]
	_ = x[ChangeType-(5)]
	_ = x[Const-(6)]
	_ = x[Convert-(7)]
	_ = x[Copy-(8)]
	_ = x[Extract-(9)]
	_ = x[Field-(10)]
	_ = x[FieldAddr-(11)]
	_ = x[FreeVar-(12)]
	_ = x[Func-(13)]
	_ = x[Global-(14)]
	_ = x[Index-(15)]
	_ = x[IndexAddr-(16)]
	_ = x[InlineAsm-(17)]
	_ = x[Local-(18)]
	_ = x[Lookup-(19)]
	_ = x[MakeInterface-(20)]
	_ = x[MakeSlice-(21)]
	_ = x[Next-(22)]
	_ = x[New-(23)]
	_ = x[Parameter-(24)]
	_ = x[Phi-(25)]
	_ = x[PhiCopy-(26)]
	_ = x[Range-(27)]
	_ = x[Reg-(28)]
	_ = x[Slice-(29)]
	_ = x[SliceToArrayPointer-(30)]
	_ = x[Store-(31)]
	_ = x[SwapIn-(32)]
	_ = x[SwapOut-(33)]
	_ = x[TypeAssert-(34)]
	_ = x[Add-(35)]
	_ = x[Sub-(36)]
	_ = x[Mul-(37)]
	_ = x[Div-(38)]
	_ = x[Rem-(39)]
	_ = x[And-(40)]
	_ = x[Or-(41)]
	_ = x[Xor-(42)]
	_ = x[ShiftLeft-(43)]
	_ = x[ShiftRight-(44)]
	_ = x[AndNot-(45)]
	_ = x[Equal-(46)]
	_ = x[NotEqual-(47)]
	_ = x[Less-(48)]
	_ = x[LessEqual-(49)]
	_ = x[Greater-(50)]
	_ = x[GreaterEqual-(51)]
	_ = x[Not-(52)]
	_ = x[Negate-(53)]
	_ = x[Load-(54)]
	_ = x[Invert-(55)]
	_ = x[Jump2-(56)]
	_ = x[If2-(57)]
	_ = x[Return2-(58)]
	_ = x[Panic2-(59)]
	_ = x[IfEqual2-(60)]
	_ = x[IfNotEqual2-(61)]
	_ = x[IfLess2-(62)]
	_ = x[IfLessEqual2-(63)]
	_ = x[IfGreater2-(64)]
	_ = x[IfGreaterEqual2-(65)]
	_ = x[NumOps-(66)]
}

var _OpValues = []Op{Invalid, Builtin, Call, CallBuiltin, ChangeInterface, ChangeType, Const, Convert, Copy, Extract, Field, FieldAddr, FreeVar, Func, Global, Index, IndexAddr, InlineAsm, Local, Lookup, MakeInterface, MakeSlice, Next, New, Parameter, Phi, PhiCopy, Range, Reg, Slice, SliceToArrayPointer, Store, SwapIn, SwapOut, TypeAssert, Add, Sub, Mul, Div, Rem, And, Or, Xor, ShiftLeft, ShiftRight, AndNot, Equal, NotEqual, Less, LessEqual, Greater, GreaterEqual, Not, Negate, Load, Invert, Jump2, If2, Return2, Panic2, IfEqual2, IfNotEqual2, IfLess2, IfLessEqual2, IfGreater2, IfGreaterEqual2, NumOps}

var _OpNameToValueMap = map[string]Op{
	_OpName[0:7]:          Invalid,
	_OpLowerName[0:7]:     Invalid,
	_OpName[7:14]:         Builtin,
	_OpLowerName[7:14]:    Builtin,
	_OpName[14:18]:        Call,
	_OpLowerName[14:18]:   Call,
	_OpName[18:29]:        CallBuiltin,
	_OpLowerName[18:29]:   CallBuiltin,
	_OpName[29:44]:        ChangeInterface,
	_OpLowerName[29:44]:   ChangeInterface,
	_OpName[44:54]:        ChangeType,
	_OpLowerName[44:54]:   ChangeType,
	_OpName[54:59]:        Const,
	_OpLowerName[54:59]:   Const,
	_OpName[59:66]:        Convert,
	_OpLowerName[59:66]:   Convert,
	_OpName[66:70]:        Copy,
	_OpLowerName[66:70]:   Copy,
	_OpName[70:77]:        Extract,
	_OpLowerName[70:77]:   Extract,
	_OpName[77:82]:        Field,
	_OpLowerName[77:82]:   Field,
	_OpName[82:91]:        FieldAddr,
	_OpLowerName[82:91]:   FieldAddr,
	_OpName[91:98]:        FreeVar,
	_OpLowerName[91:98]:   FreeVar,
	_OpName[98:102]:       Func,
	_OpLowerName[98:102]:  Func,
	_OpName[102:108]:      Global,
	_OpLowerName[102:108]: Global,
	_OpName[108:113]:      Index,
	_OpLowerName[108:113]: Index,
	_OpName[113:122]:      IndexAddr,
	_OpLowerName[113:122]: IndexAddr,
	_OpName[122:131]:      InlineAsm,
	_OpLowerName[122:131]: InlineAsm,
	_OpName[131:136]:      Local,
	_OpLowerName[131:136]: Local,
	_OpName[136:142]:      Lookup,
	_OpLowerName[136:142]: Lookup,
	_OpName[142:155]:      MakeInterface,
	_OpLowerName[142:155]: MakeInterface,
	_OpName[155:164]:      MakeSlice,
	_OpLowerName[155:164]: MakeSlice,
	_OpName[164:168]:      Next,
	_OpLowerName[164:168]: Next,
	_OpName[168:171]:      New,
	_OpLowerName[168:171]: New,
	_OpName[171:180]:      Parameter,
	_OpLowerName[171:180]: Parameter,
	_OpName[180:183]:      Phi,
	_OpLowerName[180:183]: Phi,
	_OpName[183:190]:      PhiCopy,
	_OpLowerName[183:190]: PhiCopy,
	_OpName[190:195]:      Range,
	_OpLowerName[190:195]: Range,
	_OpName[195:198]:      Reg,
	_OpLowerName[195:198]: Reg,
	_OpName[198:203]:      Slice,
	_OpLowerName[198:203]: Slice,
	_OpName[203:222]:      SliceToArrayPointer,
	_OpLowerName[203:222]: SliceToArrayPointer,
	_OpName[222:227]:      Store,
	_OpLowerName[222:227]: Store,
	_OpName[227:233]:      SwapIn,
	_OpLowerName[227:233]: SwapIn,
	_OpName[233:240]:      SwapOut,
	_OpLowerName[233:240]: SwapOut,
	_OpName[240:250]:      TypeAssert,
	_OpLowerName[240:250]: TypeAssert,
	_OpName[250:253]:      Add,
	_OpLowerName[250:253]: Add,
	_OpName[253:256]:      Sub,
	_OpLowerName[253:256]: Sub,
	_OpName[256:259]:      Mul,
	_OpLowerName[256:259]: Mul,
	_OpName[259:262]:      Div,
	_OpLowerName[259:262]: Div,
	_OpName[262:265]:      Rem,
	_OpLowerName[262:265]: Rem,
	_OpName[265:268]:      And,
	_OpLowerName[265:268]: And,
	_OpName[268:270]:      Or,
	_OpLowerName[268:270]: Or,
	_OpName[270:273]:      Xor,
	_OpLowerName[270:273]: Xor,
	_OpName[273:282]:      ShiftLeft,
	_OpLowerName[273:282]: ShiftLeft,
	_OpName[282:292]:      ShiftRight,
	_OpLowerName[282:292]: ShiftRight,
	_OpName[292:298]:      AndNot,
	_OpLowerName[292:298]: AndNot,
	_OpName[298:303]:      Equal,
	_OpLowerName[298:303]: Equal,
	_OpName[303:311]:      NotEqual,
	_OpLowerName[303:311]: NotEqual,
	_OpName[311:315]:      Less,
	_OpLowerName[311:315]: Less,
	_OpName[315:324]:      LessEqual,
	_OpLowerName[315:324]: LessEqual,
	_OpName[324:331]:      Greater,
	_OpLowerName[324:331]: Greater,
	_OpName[331:343]:      GreaterEqual,
	_OpLowerName[331:343]: GreaterEqual,
	_OpName[343:346]:      Not,
	_OpLowerName[343:346]: Not,
	_OpName[346:352]:      Negate,
	_OpLowerName[346:352]: Negate,
	_OpName[352:356]:      Load,
	_OpLowerName[352:356]: Load,
	_OpName[356:362]:      Invert,
	_OpLowerName[356:362]: Invert,
	_OpName[362:367]:      Jump2,
	_OpLowerName[362:367]: Jump2,
	_OpName[367:370]:      If2,
	_OpLowerName[367:370]: If2,
	_OpName[370:377]:      Return2,
	_OpLowerName[370:377]: Return2,
	_OpName[377:383]:      Panic2,
	_OpLowerName[377:383]: Panic2,
	_OpName[383:391]:      IfEqual2,
	_OpLowerName[383:391]: IfEqual2,
	_OpName[391:402]:      IfNotEqual2,
	_OpLowerName[391:402]: IfNotEqual2,
	_OpName[402:409]:      IfLess2,
	_OpLowerName[402:409]: IfLess2,
	_OpName[409:421]:      IfLessEqual2,
	_OpLowerName[409:421]: IfLessEqual2,
	_OpName[421:431]:      IfGreater2,
	_OpLowerName[421:431]: IfGreater2,
	_OpName[431:446]:      IfGreaterEqual2,
	_OpLowerName[431:446]: IfGreaterEqual2,
	_OpName[446:452]:      NumOps,
	_OpLowerName[446:452]: NumOps,
}

var _OpNames = []string{
	_OpName[0:7],
	_OpName[7:14],
	_OpName[14:18],
	_OpName[18:29],
	_OpName[29:44],
	_OpName[44:54],
	_OpName[54:59],
	_OpName[59:66],
	_OpName[66:70],
	_OpName[70:77],
	_OpName[77:82],
	_OpName[82:91],
	_OpName[91:98],
	_OpName[98:102],
	_OpName[102:108],
	_OpName[108:113],
	_OpName[113:122],
	_OpName[122:131],
	_OpName[131:136],
	_OpName[136:142],
	_OpName[142:155],
	_OpName[155:164],
	_OpName[164:168],
	_OpName[168:171],
	_OpName[171:180],
	_OpName[180:183],
	_OpName[183:190],
	_OpName[190:195],
	_OpName[195:198],
	_OpName[198:203],
	_OpName[203:222],
	_OpName[222:227],
	_OpName[227:233],
	_OpName[233:240],
	_OpName[240:250],
	_OpName[250:253],
	_OpName[253:256],
	_OpName[256:259],
	_OpName[259:262],
	_OpName[262:265],
	_OpName[265:268],
	_OpName[268:270],
	_OpName[270:273],
	_OpName[273:282],
	_OpName[282:292],
	_OpName[292:298],
	_OpName[298:303],
	_OpName[303:311],
	_OpName[311:315],
	_OpName[315:324],
	_OpName[324:331],
	_OpName[331:343],
	_OpName[343:346],
	_OpName[346:352],
	_OpName[352:356],
	_OpName[356:362],
	_OpName[362:367],
	_OpName[367:370],
	_OpName[370:377],
	_OpName[377:383],
	_OpName[383:391],
	_OpName[391:402],
	_OpName[402:409],
	_OpName[409:421],
	_OpName[421:431],
	_OpName[431:446],
	_OpName[446:452],
}

// OpString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func OpString(s string) (Op, error) {
	if val, ok := _OpNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _OpNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Op values", s)
}

// OpValues returns all values of the enum
func OpValues() []Op {
	return _OpValues
}

// OpStrings returns a slice of all String values of the enum
func OpStrings() []string {
	strs := make([]string, len(_OpNames))
	copy(strs, _OpNames)
	return strs
}

// IsAOp returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Op) IsAOp() bool {
	for _, v := range _OpValues {
		if i == v {
			return true
		}
	}
	return false
}
