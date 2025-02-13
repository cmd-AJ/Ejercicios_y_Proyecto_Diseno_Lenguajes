package Minimal

type Symbol = string

type Table struct {
	Table_2D [][]string
	Y_index  []string
	X_index  []string
	Finals   map[string]bool
	Initial  string
}

type Tuple struct {
	OuterKey string
	InnerKey string
}
