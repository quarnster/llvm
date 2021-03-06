package ir

import (
	"github.com/llir/llvm/types"
	"github.com/llir/llvm/values"
)

// An Instruction performs a non-branching operation and belongs to one of the
// following groups:
//
//    * binary instructions [1]
//    * bitwise binary instructions [2]
//    * memory instructions [3]
//    * other instructions [4]
//
//    [1]: http://llvm.org/docs/LangRef.html#binaryops
//    [2]: http://llvm.org/docs/LangRef.html#bitwiseops
//    [3]: http://llvm.org/docs/LangRef.html#memoryops
//    [4]: http://llvm.org/docs/LangRef.html#otherops
type Instruction interface {
	// isInst ensures that only non-terminator instructions can be assigned to
	// the Instruction interface.
	isInst()
}

// =============================================================================
// Binary Operations
//
//    ref: http://llvm.org/docs/LangRef.html#binaryops
// =============================================================================

// The AddInst returns the sum of its two operands, which may be integers or
// vectors of integer values.
//
// Syntax:
//    <Result> = add <Type> <Op1>, <Op2>
//
// Semantics:
//    Result = Op1 + Op2;
//
// References:
//    http://llvm.org/docs/LangRef.html#i-add
type AddInst struct {
	// Operand type.
	Type types.Type
	// Operands.
	Op1, Op2 values.Value
}

// The FaddInst returns the sum of its two operands, which may be floating point
// values or vectors of floating point values.
//
// Syntax:
//    <Result> = fadd <Type> <Op1>, <Op2>
//
// Semantics:
//    Result = Op1 + Op2;
//
// References:
//    http://llvm.org/docs/LangRef.html#i-fadd
type FaddInst struct {
	// Operand type.
	Type types.Type
	// Operands.
	Op1, Op2 values.Value
}

// The SubInst returns the difference of its two operands, which may be integers
// or vectors of integer values.
//
// Syntax:
//    <Result> = sub <Type> <Op1>, <Op2>
//
// Semantics:
//    Result = Op1 - Op2;
//
// References:
//    http://llvm.org/docs/LangRef.html#sub-instruction
type SubInst struct {
	// Operand type.
	Type types.Type
	// Operands.
	Op1, Op2 values.Value
}

// The FsubInst returns the difference of its two operands, which may be
// floating point values or vectors of floating point values.
//
// Syntax:
//    <Result> = fsub <Type> <Op1>, <Op2>
//
// Semantics:
//    Result = Op1 - Op2;
//
// References:
//    http://llvm.org/docs/LangRef.html#i-fsub
type FsubInst struct {
	// Operand type.
	Type types.Type
	// Operands.
	Op1, Op2 values.Value
}

// The MulInst returns the product of its two operands, which may be integers or
// vectors of integer values.
//
// Syntax:
//    <Result> = mul <Type> <Op1>, <Op2>
//
// Semantics:
//    Result = Op1 * Op2;
//
// References:
//    http://llvm.org/docs/LangRef.html#mul-instruction
type MulInst struct {
	// Operand type.
	Type types.Type
	// Operands.
	Op1, Op2 values.Value
}

// The FmulInst returns the product of its two operands, which may be floating
// point values or vectors of floating point values.
//
// Syntax:
//    <Result> = fmul <Type> <Op1>, <Op2>
//
// Semantics:
//    Result = Op1 * Op2;
//
// References:
//    http://llvm.org/docs/LangRef.html#fmul-instruction
type FmulInst struct {
	// Operand type.
	Type types.Type
	// Operands.
	Op1, Op2 values.Value
}

// The UdivInst returns the unsigned integer quotient of its two operands, which
// may be integers or vectors of integer values.
//
// Syntax:
//    <Result> = udiv <Type> <Op1>, <Op2>
//
// Semantics:
//    Result = Op1 / Op2;
//
// References:
//    http://llvm.org/docs/LangRef.html#udiv-instruction
type UdivInst struct {
	// Operand type.
	Type types.Type
	// Operands.
	Op1, Op2 values.Value
}

// The SdivInst returns the signed integer quotient of its two operands, which
// may be integers or vectors of integer values.
//
// Syntax:
//    <Result> = sdiv <Type> <Op1>, <Op2>
//
// Semantics:
//    Result = Op1 / Op2;
//
// References:
//    http://llvm.org/docs/LangRef.html#sdiv-instruction
type SdivInst struct {
	// Operand type.
	Type types.Type
	// Operands.
	Op1, Op2 values.Value
}

// The FdivInst returns the quotient of its two operands, which may be floating
// point values or vectors of floating point values.
//
// Syntax:
//    <Result> = fdiv <Type> <Op1>, <Op2>
//
// Semantics:
//    Result = Op1 / Op2;
//
// References:
//    http://llvm.org/docs/LangRef.html#fdiv-instruction
type FdivInst struct {
	// Operand type.
	Type types.Type
	// Operands.
	Op1, Op2 values.Value
}

// The UremInst returns the unsigned integer remainder of a division between its
// two operands, which may be integers or vectors of integers.
//
// Syntax:
//    <Result> = urem <Type> <Op1>, <Op2>
//
// Semantics:
//    Result = Op1 % Op2;
//
// References:
//    http://llvm.org/docs/LangRef.html#urem-instruction
type UremInst struct {
	// Operand type.
	Type types.Type
	// Operands.
	Op1, Op2 values.Value
}

// The SremInst returns the signed integer remainder of a division between its
// two operands, which may be integers or vectors of integers.
//
// Syntax:
//    <Result> = srem <Type> <Op1>, <Op2>
//
// Semantics:
//    Result = Op1 % Op2;
//
// References:
//    http://llvm.org/docs/LangRef.html#srem-instruction
type SremInst struct {
	// Operand type.
	Type types.Type
	// Operands.
	Op1, Op2 values.Value
}

// The FremInst returns the remainder of a division between its two operands,
// which may be floating point values or vectors of floating point values.
//
// Syntax:
//    <Result> = frem <Type> <Op1>, <Op2>
//
// Semantics:
//    Result = Op1 % Op2;
//
// References:
//    http://llvm.org/docs/LangRef.html#frem-instruction
type FremInst struct {
	// Operand type.
	Type types.Type
	// Operands.
	Op1, Op2 values.Value
}

// =============================================================================
// Bitwise Binary Operations
//
//    ref: http://llvm.org/docs/LangRef.html#bitwiseops
// =============================================================================

// The ShlInst returns the first operand shifted to the left a specified number
// of bits. The arguments may be integers or vectors of integer values.
//
// Syntax:
//    <Result> = shl <Type> <Op1>, <Op2>
//
// Semantics:
//    Result = Op1 << Op2;
//
// References:
//    http://llvm.org/docs/LangRef.html#shl-instruction
type ShlInst struct {
	// Operand type.
	Type types.Type
	// Operands.
	Op1, Op2 values.Value
}

// The LshrInst (logical shift right) returns the first operand shifted to the
// right a specified number of bits with zero fill. The arguments may be
// integers or vectors of integer values.
//
// Syntax:
//    <Result> = lshr <Type> <Op1>, <Op2>
//
// Semantics:
//    Result = Op1 >> Op2;
//
// References:
//    http://llvm.org/docs/LangRef.html#lshr-instruction
type LshrInst struct {
	// Operand type.
	Type types.Type
	// Operands.
	Op1, Op2 values.Value
}

// The AshrInst (arithmetic shift right) returns the first operand shifted to
// the right a specified number of bits with sign extension. The arguments may
// be integers or vectors of integer values.
//
// Syntax:
//    <Result> = ashr <Type> <Op1>, <Op2>
//
// Semantics:
//    Result = Op1 >> Op2; // right shift with sign extension.
//
// References:
//    http://llvm.org/docs/LangRef.html#ashr-instruction
type AshrInst struct {
	// Operand type.
	Type types.Type
	// Operands.
	Op1, Op2 values.Value
}

// The AndInst returns the bitwise logical and of its two operands, which may be
// integers or vectors of integer values.
//
// Syntax:
//    <Result> = and <Type> <Op1>, <Op2>
//
// Semantics:
//    Result = Op1 & Op2;
//
// References:
//    http://llvm.org/docs/LangRef.html#and-instruction
type AndInst struct {
	// Operand type.
	Type types.Type
	// Operands.
	Op1, Op2 values.Value
}

// The OrInst returns the bitwise logical inclusive or of its two operands,
// which may be integers or vectors of integer values.
//
// Syntax:
//    <Result> = or <Type> <Op1>, <Op2>
//
// Semantics:
//    Result = Op1 | Op2;
//
// References:
//    http://llvm.org/docs/LangRef.html#or-instruction
type OrInst struct {
	// Operand type.
	Type types.Type
	// Operands.
	Op1, Op2 values.Value
}

// The XorInst returns the bitwise logical exclusive or of its two operands,
// which may be integers or vectors of integer values.
//
// Syntax:
//    <Result> = xor <Type> <Op1>, <Op2>
//
// Semantics:
//    Result = Op1 ^ Op2;
//
// References:
//    http://llvm.org/docs/LangRef.html#xor-instruction
type XorInst struct {
	// Operand type.
	Type types.Type
	// Operands.
	Op1, Op2 values.Value
}

// =============================================================================
// Vector Operations
//
//    ref: http://llvm.org/docs/LangRef.html#vector-operations
// =============================================================================

// TODO: Add the following instructions:
//    - extractelement
//    - insertelement
//    - shufflevector

// =============================================================================
// Aggregate Operations
//
//    ref: http://llvm.org/docs/LangRef.html#aggregate-operations
// =============================================================================

// TODO: Add the following instructions:
//    - extractvalue
//    - insertvalue

// =============================================================================
// Memory Access and Addressing Operations
//
//    ref: http://llvm.org/docs/LangRef.html#memoryops
// =============================================================================

// The AllocaInst allocates memory on the stack frame of the currently executing
// function, to be automatically released when this function returns to its
// caller.
//
// Syntax:
//    <Result> = alloca <Type> [, <Type> <NumElems> ] [, align <Align> ]
//
// Semantics:
//    f() {
//       Type Result;
//       Type[NumElems] Result;
//    }
//
// References:
//    http://llvm.org/docs/LangRef.html#alloca-instruction
type AllocaInst struct {
	// Underlying type of the pointer.
	Type types.Type
	// Number of elements to allocate; defaults to 1.
	NumElems int
	// Memory alignment.
	Align int
}

// The LoadInst reads from memory.
//
// Syntax:
//    <Result> = load <Type>* <Addr> [, align <Align> ]
//
// Semantics:
//    Result = *(Type *)Addr;
//
// References:
//    http://llvm.org/docs/LangRef.html#load-instruction
type LoadInst struct {
	// Underlying type of the pointer.
	Type types.Type
	// Memory address to load.
	Addr values.Value
	// Memory alignment.
	Align int
}

// The StoreInst writes to memory.
//
// Syntax:
//    store <Type> <Val>, <Type>* <Addr> [, align <Align> ]
//
// Semantics:
//    *(Type *)Addr = Val;
//
// References:
//    http://llvm.org/docs/LangRef.html#store-instruction
type StoreInst struct {
	// Value type.
	Type types.Type
	// Value to store.
	Val values.Value
	// Memory address to store at.
	Addr values.Value
	// Memory alignment.
	Align int
}

// TODO(u): Add the following memory access and addressing operations:
//    - fence
//    - cmpxchg
//    - atomicrmw

// TODO(u): Read up about the syntax and semantics of getelementptr when used
// with vectors of pointers.

// The GetelementptrInst gets the address of a subelement of an aggregate data
// structure. It performs address calculation only and does not access memory.
//
// Syntax:
//    <Result> = getelementptr <Type>* <Ptr> {, <Type> <Idx>}*
//
// Semantics:
//    Result = &Ptr[Idx1];
//    Result = &Ptr.Field1; // Field1 is identified by Idx1.
//    Result = &Ptr[Idx1].Field2[Idx3];
//
// References:
//    http://llvm.org/docs/LangRef.html#getelementptr-instruction
type GetelementptrInst struct {
	// Underlying type of the pointer.
	Type types.Type
	// Pointer to the aggregate data structure.
	Ptr values.Value
	// Element indicies.
	Indicies []int
}

// =============================================================================
// Conversion Operations
//
//    ref: http://llvm.org/docs/LangRef.html#conversion-operations
// =============================================================================

// TODO: Add the following instructions:
//    - trunc
//    - zext
//    - sext
//    - fptrunc
//    - fpext
//    - fptoui
//    - fptosi
//    - uitofp
//    - sitofp
//    - ptrtoint
//    - inttoptr
//    - bitcast
//    - addrspacecast

// =============================================================================
// Other Operations
//
//    ref: http://llvm.org/docs/LangRef.html#other-operations
// =============================================================================

// The IcmpInst compares integer values.
//
// Syntax:
//    <Result> = icmp <Pred> <Type> <Op1>, <Op2>
//
// Semantics:
//    Result = (Op1 Pred Op2); // Where Pred is ==, !=, >, >=, < or <=.
//
// References:
//    http://llvm.org/docs/LangRef.html#icmp-instruction
type IcmpInst struct {
	// Comparison operation.
	Pred IntPredicate
	// TODO: Restrict to IntsType and IntsValue?

	// Value type.
	Type types.Type
	// Operands.
	Op1, Op2 values.Value
}

// IntPredicate specifies a comparison operation to perform between two integer
// values.
type IntPredicate int

// Integer comparison operations.
const (
	IntEq  IntPredicate = iota // equal
	IntNe                      // not equal
	IntUgt                     // unsigned greater than
	IntUge                     // unsigned greater or equal
	IntUlt                     // unsigned less than
	IntUle                     // unsigned less or equal
	IntSgt                     // signed greater than
	IntSge                     // signed greater or equal
	IntSlt                     // signed less than
	IntSle                     // signed less or equal
)

// The FcmpInst compares floating point values.
//
// Syntax:
//    <Result> = fcmp <Pred> <Type> <Op1>, <Op2>
//
// Semantics:
//    Result = (Op1 Pred Op2); // Where Pred is ==, !=, >, >=, < or <=.
//
// References:
//    http://llvm.org/docs/LangRef.html#fcmp-instruction
type FcmpInst struct {
	// Comparison operation.
	Pred FloatPredicate
	// TODO: Restrict to FloatsType and FloatsValue?

	// Value type.
	Type types.Type
	// Operands.
	Op1, Op2 values.Value
}

// FloatPredicate specifies a comparison operation to perform between two
// floating point values.
type FloatPredicate int

// Floating point comparison operations.
const (
	FloatFalse FloatPredicate = iota // no comparison, always returns false
	FloatOeq                         // ordered and equal
	FloatOgt                         // ordered and greater than
	FloatOge                         // ordered and greater than or equal
	FloatOlt                         // ordered and less than
	FloatOle                         // ordered and less than or equal
	FloatOne                         // ordered and not equal
	FloatOrd                         // ordered (no nans)
	FloatUeq                         // unordered or equal
	FloatUgt                         // unordered or greater than
	FloatUge                         // unordered or greater than or equal
	FloatUlt                         // unordered or less than
	FloatUle                         // unordered or less than or equal
	FloatUne                         // unordered or not equal
	FloatUno                         // unordered (either nans)
	FloatTrue                        // no comparison, always returns true

)

// The PhiInst is used to implement φ nodes in the SSA graph representation of a
// function.
//
// Syntax:
//    <Result> = phi <Type> [ <Val0>, <Label0> ], ...
//
// Semantics:
//    Result = Val0 // if the previously executed basic block was Label0, etc.
//
// References:
//    http://llvm.org/docs/LangRef.html#phi-instruction
type PhiInst struct {
	// Value type.
	Type types.Type
	// Predecessor basic block labels and their corresponding values.
	Preds map[string]values.Value
}

// TODO: Add the following instructions:
//    - select
//    - call
//    - va_arg
//    - landingpad

// isInst ensures that only non-terminator instructions can be assigned to the
// Instruction interface.
func (AddInst) isInst()  {}
func (FaddInst) isInst() {}
func (SubInst) isInst()  {}
func (FsubInst) isInst() {}
func (MulInst) isInst()  {}
func (FmulInst) isInst() {}
func (UdivInst) isInst() {}
func (SdivInst) isInst() {}
func (FdivInst) isInst() {}
func (UremInst) isInst() {}
func (SremInst) isInst() {}
func (FremInst) isInst() {}
