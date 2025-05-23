package field

//跳過當前回合,且體力+2
func (f *Field) SkipTurn() {
	f.AddHealth(2)
	f.NextTrue()
}

//進入下一回合
//需結算各項buff回合結束效果
func (f *Field) NextTrue() {
	f.Cars.Discard.Push(f.Cars.Hand...)
	for len(f.Cars.Hand) != 0 {
		f.Cars.Hand.Remove(0)
	}
	if f.State.Turn == len(f.State.Turns) {
		f.State.ExtraTurn--
	} else {
		f.State.Turn++
	}
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

//
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
