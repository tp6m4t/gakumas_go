package field

import (
	"LocalProject/pkg/CardGame/Field/deck"
	"LocalProject/pkg/CardGame/core"
	"fmt"
)

type Turn struct {
	Color *int
	Skill []core.Skill
}

type CarDeck struct {
	Deck         deck.Deck //山扎
	Hand         deck.Deck //手扎
	Discard      deck.Deck //捨扎
	BanishedPile deck.Deck //除外
}

// 場地資訊相關變數
type State struct {
	Cars      CarDeck     //牌堆
	Score     int         //得分
	Health    int         //當前體力
	MaxHealth int         //最大體力
	Energy    int         //當前能量
	Drinks    []int       //可用飲料
	Buffs     []core.Buff //當前Buff DeBuff
	Items     []int       //當前道具
	Turns     []Turn      //本局回	合資訊
	ExtraTurn int         //額外回合
	Turn      int         //當前回合
}

// 行為計算相關變數
type Field struct {
	State
}

func (f Field) String() string {
	str := ""
	str += fmt.Sprintf("分數:%d\n", f.Score)
	if f.Energy > 0 {
		str += fmt.Sprintf("能量:%d   ", f.Energy)
	}
	str += fmt.Sprintf("體力:%d/%d\n", f.Health, f.MaxHealth)
	str += fmt.Sprintf("回和數:%d/%d+%d\n", len(f.Turns), f.Turn, f.ExtraTurn)
	//當前buff
	//當前道具
	str += "可用卡牌:"
	for _, v := range f.Cars.Hand {
		str += v.GetName() + "/"
	}
	str += "\n"
	str += "可用飲料:"
	for _, v := range f.Drinks {
		str += fmt.Sprintf("%d ", v)
	}
	str += "\n"

	return str
}

// 所有回和/額外回合結束?
func (f *Field) IsEnd() bool {
	return f.State.Turn == len(f.State.Turns) && f.State.ExtraTurn == 0
}
