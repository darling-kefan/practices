package toy

type Toy struct {
	Name   string
	Weight int
	onHand int
	sold   int
}

func New(name string, weight int) *Toy {
	return &Toy{
		Name:   name,
		Weight: weight,
	}
}

func (t *Toy) OnHand() int {
	return t.onHand
}

func (t *Toy) UpdateOnHand(onHand int) int {
	t.onHand = onHand
	return t.onHand
}

func (t *Toy) Sold() int {
	return t.sold
}

func (t *Toy) UpdateSold(sold int) int {
	t.sold = sold
	return t.sold
}
