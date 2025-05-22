package field

import (
	"LocalProject/CardGame/Buff"
	"LocalProject/CardGame/Card"
)

type Turn struct {
	color int
	buff  Buff.Buff
}

type CarDeck struct {
	Deck         []Card.Card //山扎
	Hand         []Card.Card //手扎
	Discard      []Card.Card //捨扎
	BanishedPile []Card.Card //除外
}

type State struct {
	Cars      CarDeck     //牌堆
	Score     int         //得分
	Health    int         //當前體力
	MaxHealth int         //最大體力
	Shield    int         //當前護盾
	Drinks    []int       //可用飲料
	Buffs     []Buff.Buff //當前Buff DeBuff
	Items     []int       //當前道具
	Turns     []Turn      //本局回合資訊
	ExtraTurn int         //額外回合
	Turn      int         //當前回合
}

type Effect struct {
}
type Condition struct {
}

type Field struct {
	State State
}

func (f *Field) UseCard(idx int) State {
	if len(f.State.Cars.Hand) >= idx {
		// 產生error
	}
	for _, skil := range f.State.Cars.Hand[idx].SkilList {
		//檢查符合發動條件
		if f.checkConditions(skil.Condition) {
			f.useEffects(skil.Effect)
		}

	}
	return f.State
}

func (f *Field) UseDrink(idx int) State {
	return f.State
}

func (f *Field) checkConditions(c []Condition) bool {
	for _, v := range c {
		if !f.checkCondition(v) {
			return false
		}
	}
	return true
}

func (f *Field) checkCondition(c Condition) bool {
	switch c {

	}
	return true
}

func (f *Field) useEffects(e []Effect) State {
	for _, v := range e {
		f.useEffect(v)
	}
	return f.State
}

func (f *Field) useEffect(e Effect) State {
	switch e {

	}
	return f.State
}
