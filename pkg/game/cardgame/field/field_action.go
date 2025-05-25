package field

import (
	"LocalProject/pkg/game/data/buff"
)

// 跳過當前回合,且體力+2
func (f *Field) SkipTurn() {
	f.AddHealth(2)
	f.NextTrue()
}

// 進入下一回合
// 需結算各項buff回合結束效果
func (f *Field) NextTrue() {
	//觸發回合結束事件
	f.eventBus.Publish("turnEnd", f)

	//回收手牌
	f.Cars.Discard.Push(f.Cars.Hand...)
	f.Cars.Hand.Clear()

	//更新回合資訊
	if f.State.CurrentTurn == len(f.State.Turns) {
		f.State.ExtraTurn--
	} else {
		f.State.CurrentTurn++
	}

	f.DrawCards(3)

}

// 對外增加分數的接口 內部計算buff疊加效果
func (f *Field) AddScore(value int) {
	f.Score += value
}

// 回復能量
func (f *Field) AddEnergy(value int) {
	f.Energy += value
}

// 消耗體力(能量優先)
func (f *Field) SubEnergy(value int) {
	if f.Energy >= value {
		f.Energy -= value
	} else {
		value -= f.Energy
		f.Energy = 0
		f.SubHealth(value)
	}
}

func (f *Field) AddHealth(value int) {
	if f.Health+2 > f.MaxHealth {
		f.Health = f.MaxHealth
	} else {
		f.Health += 2
	}
}

// 直接消耗體力(無視能量)
func (f *Field) SubHealth(value int) {
	f.Health -= value
}

// 抽卡到手牌
func (f *Field) DrawCards(value int) {
	// 牌堆不夠時 先將所有牌加入手牌 再重新自山扎取牌
	if len(f.Cars.Deck) < value {
		value -= len(f.Cars.Deck)

		f.Cars.Hand.Push(f.Cars.Deck...)
		f.Cars.Deck.Clear()
		f.Cars.Deck.Push(f.Cars.Discard...)
		f.Cars.Discard.Clear()
	}
	Cards, _ := f.Cars.Deck.DrawCards(value)
	f.Cars.Hand.Push(Cards...)
}

func (f *Field) AddBuff(name string, value int) {
	if buff.GetBuffID(name) != -1 {
		if f.Buffs[buff.GetBuffID(name)] != nil {
			f.Buffs[buff.GetBuffID(name)].AddCount(value)
		} else {
			f.Buffs[buff.GetBuffID(name)] = buff.NewBuffByName(name, value)
			//註冊監聽事件
			f.Buffs[buff.GetBuffID(name)].Subscribe(f)
		}
	}
}
