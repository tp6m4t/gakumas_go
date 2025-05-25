package buff

type GoodImpression struct {
	BaseBuff
}

func (b *GoodImpression) OnTurnEnd(f field) {
	f.AddScore(b.count)
	b.count--
}

func (b *GoodImpression) AddCount(count int) {
	b.count += count
}

func (b *GoodImpression) GetCount() int {
	return b.count
}

func (b *GoodImpression) Subscribe(f field) {
	f.Subscribe("turnEnd", func(field interface{}) { b.OnTurnEnd(f) })
}

func NewGoodImpression(count int) *GoodImpression {
	return &GoodImpression{
		BaseBuff{
			id:           0,
			name:         "好印象",
			descriptions: "每回合結束加上對應層數的分數",
			count:        count,
		},
	}
}
