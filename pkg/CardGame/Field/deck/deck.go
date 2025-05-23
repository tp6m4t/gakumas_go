package deck

import (
	"LocalProject/pkg/core/Card/SkillCard"
	"fmt"
	"math/rand"
	"time"
)

// 牌堆不保證其中的順序
type Deck []SkillCard.SkillCard

// 移除特定牌
func (d *Deck) Remove(index int) {
	(*d)[index], (*d)[len(*d)-1] = (*d)[len(*d)-1], (*d)[index]
	*d = (*d)[:len(*d)-1]
}

// 抽牌
func (d *Deck) DrawCards(n int) (Deck, error) {
	var err error
	if n > len(*d) {
		n = len(*d)
		err = fmt.Errorf("卡牌數量不足抽取數量")
	}
	var cards Deck
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		idx := r.Intn(len(*d))
		cards = append(cards, (*d)[idx])
		d.Remove(idx)
	}
	return cards, err
}

// 放入牌堆
func (d *Deck) Push(cards ...SkillCard.SkillCard) {
	*d = append(*d, cards...)
}

// 清空牌堆
func (d *Deck) Clear() {
	(*d) = (*d)[:0]
}
