package main

var reserved = map[string]int{
	"and":       tAND,
	"array":     tARRAY,
	"begin":     tBEGIN,
	"boolean":   tBOOLEAN,
	"break":     tBREAK,
	"char":      tCHAR,
	"continue":  tCONTINUE,
	"div":       tDIV,
	"do":        tDO,
	"downto":    tDOWNTO,
	"else":      tELSE,
	"end":       tEND,
	"false":     tFALSE,
	"for":       tFOR,
	"function":  tFUNCTION,
	"goto":      tGOTO,
	"if":        tIF,
	"integer":   tINTEGER,
	"mod":       tMOD,
	"not":       tNOT,
	"of":        tOF,
	"or":        tOR,
	"procedure": tPROCEDURE,
	"program":   tPROGRAM,
	"real":      tREAL,
	"record":    tRECORD,
	"repeat":    tREPEAT,
	"string":    tSTRING,
	"then":      tTHEN,
	"to":        tTO,
	"true":      tTRUE,
	"type":      tTYPE,
	"until":     tUNTIL,
	"var":       tVAR,
	"while":     tWHILE,
}
