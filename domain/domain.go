package domain


type Holding struct {
	Name     string
	Holdings []SubHolding
}

type SubHolding struct {
	Holding *Holding
	Weight  float64
}