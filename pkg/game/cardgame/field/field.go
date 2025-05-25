package field

import (
	"LocalProject/pkg/game/cardgame/field/deck"
	"LocalProject/pkg/game/cardgame/field/turn"
	"LocalProject/pkg/game/data/buff"
	"LocalProject/pkg/tool/event"

	"fmt"
)

type CarDeck struct {
	Deck         deck.Deck //山扎
	Hand         deck.Deck //手扎
	Discard      deck.Deck //捨扎
	BanishedPile deck.Deck //除外
}

// 場地資訊相關變數
type State struct {
	Cars        CarDeck           //牌堆
	Score       int               //得分
	Health      int               //當前體力
	MaxHealth   int               //最大體力
	Energy      int               //當前能量
	Drinks      []int             //可用飲料
	Buffs       map[int]buff.Buff //當前Buff DeBuff
	Items       []int             //當前道具
	Turns       []turn.Turn       //本局回	合資訊
	ExtraTurn   int               //額外回合
	CurrentTurn int               //當前回合
}

// 行為計算相關變數
type Field struct {
	State
	eventBus             *event.EventBus //事件總線
	ScoreBonus           int             //得分加成
	IsScoreAddInvalid    bool            //分數上升無效
	IsScoreDebuffInvalid bool            //debuff無效
}

func (f Field) String() string {
	str := ""
	str += fmt.Sprintf("分數:%d\n", f.Score)
	if f.Energy > 0 {
		str += fmt.Sprintf("能量:%d   ", f.Energy)
	}
	str += fmt.Sprintf("體力:%d/%d\n", f.Health, f.MaxHealth)
	str += fmt.Sprintf("回和數:%d/%d+%d\n", len(f.Turns), f.CurrentTurn, f.ExtraTurn)
	//當前buff
	str += "當前buff:\n"
	for _, b := range f.Buffs {
		str += fmt.Sprintf("%s:%d\n", b.GetName(), b.GetCount())
	}
	str += "\n"
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
	return f.State.CurrentTurn == len(f.State.Turns) && f.State.ExtraTurn == 0
}

func NewField(Deck deck.Deck, Health int, MaxHealth int, ExtraTurn int, Drinks []int, Items []int, Turns []turn.Turn) *Field {
	f := &Field{}
	f.Score = 0                       //當前分數
	f.Energy = 0                      //當前能量
	f.CurrentTurn = 0                 //當前回合
	f.Buffs = make(map[int]buff.Buff) //當前Buff DeBuff
	f.eventBus = event.NewEventBus()  //事件總線

	f.Cars.Deck.Push(Deck...) //牌堆
	f.Health = Health         //當前體力
	f.MaxHealth = MaxHealth   //最大體力
	f.ExtraTurn = ExtraTurn   //額外回合
	f.Drinks = []int{}        //可用飲料
	f.Items = []int{}         //當前道具
	f.Turns = []turn.Turn{}   //本局回合資訊
	f.DrawCards(3)

	f.eventBus.Publish("startGame", f)
	return f
}

func LoadField(Deck deck.Deck, Score int, Health int, MaxHealth int, Energy int, ExtraTurn int, Turn int, Drinks []int, Buffs map[int]buff.Buff, Items []int, Turns []turn.Turn) *Field {
	f := &Field{}
	f.eventBus = event.NewEventBus() //事件總線

	f.Cars.Deck.Push(Deck...) //牌堆
	f.Score = Score           //當前分數
	f.Health = Health         //當前體力
	f.MaxHealth = MaxHealth   //最大體力
	f.Energy = Energy         //當前能量
	f.ExtraTurn = ExtraTurn   //額外回合
	f.CurrentTurn = Turn      //當前回合
	f.Drinks = Drinks         //可用飲料
	f.Buffs = Buffs           //當前Buff DeBuff
	f.Items = Items           //當前道具
	f.Turns = Turns           //本局回合資訊

	return f
}
