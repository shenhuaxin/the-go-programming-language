package main

type Expr interface {
	Eval(env Env) float64
}

type Var string

type literal float64

type unary struct {
	op rune
	x  Expr
}
type binary struct {
	op   rune
	x, y Expr
}

type Env map[Var]float64

type call struct {
	fn   string
	args []Expr
}

func main() {

}