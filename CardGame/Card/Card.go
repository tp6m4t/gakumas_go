package Card

type Card struct {
	Name       string
	Cost       int
	CostHealth int
	Score      int
	SkilList   []struct {
		Condition []int
		Effect    []int
	}
}
