package Minimal

type Symbol = string

type Table struct {
	Table_2D [][]string
	y_index  []string
	x_index  []string
	finals   map[string]bool
	Initial  string
}

type Tuple struct {
	OuterKey string
	InnerKey string
}
