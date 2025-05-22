package cardgame

import (
	"fmt"
)

func main() {
	fmt.Println("Card Game Main")
}

type CarDeck struct {
	Deck         []int
	Hand         []int
	Discard      []int
	BanishedPile []int
}

type Turn struct {
	color int
	buff  Buff
}

type State struct {
	Cars      CarDeck //牌堆
	Score     int     //得分
	Health    int     //當前體力
	MaxHealth int     //最大體力
	Shield    int     //當前護盾
	Drinks    []int   //可用飲料
	Buffs     []Buff  //當前Buff DeBuff
	Items     []int   //當前道具
	Turns     []Turn  //本局回合資訊
	Turn      int     //當前回合
}

/* 效果發動
type Effect struct{
	Target int //對象
	Value  int //效果值
}

type ConditionEffect interface{
		Condition struct{} //條件
		Effect //效果
	}*/
