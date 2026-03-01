package tree_sitter_vue

// #cgo CFLAGS: -std=c11 -fPIC -I../../src
// #cgo CXXFLAGS: -std=c++11 -fPIC -I../../src
// #cgo LDFLAGS: -lstdc++
// #include "../../src/parser.c"
// #include "../../src/scanner.cc"
import "C"

import "unsafe"

// Get the tree-sitter Language for this grammar.
func Language() unsafe.Pointer {
	return unsafe.Pointer(C.tree_sitter_vue())
}
