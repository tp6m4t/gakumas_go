package field

import (
	"LocalProject/pkg/CardGame/core"
	"LocalProject/pkg/core/Card/SkillCard"
	"fmt"
)

type Turn struct {
	Color *int
	Skill []core.Skill
}

type CarDeck struct {
	Deck         []SkillCard.SkillCard //山扎
	Hand         []SkillCard.SkillCard //手扎
	Discard      []SkillCard.SkillCard //捨扎
	BanishedPile []SkillCard.SkillCard //除外
}

type State struct {
	Cars      CarDeck     //牌堆
	Score     int         //得分
	Health    int         //當前體力
	MaxHealth int         //最大體力
	Energy    int         //當前能量
	Drinks    []int       //可用飲料
	Buffs     []core.Buff //當前Buff DeBuff
	Items     []int       //當前道具
	Turns     []Turn      //本局回合資訊
	ExtraTurn int         //額外回合
	Turn      int         //當前回合
}

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
		str += v.Name
	}
	str += "\n"
	str += "可用飲料:"
	for _, v := range f.Drinks {
		str += fmt.Sprintf("%d ", v)
	}
	str += "\n"

	return str
}

func (f *Field) IsEnd() bool {
	return f.State.Turn == len(f.State.Turns) && f.State.ExtraTurn == 0
}

func (f *Field) SkipTurn() {
	if f.Health+2 > f.MaxHealth {
		f.Health = f.MaxHealth
	} else {
		f.Health += 2
	}
	f.NextTrue()
}
func (f *Field) NextTrue() {
	if f.State.Turn == len(f.State.Turns) {
		f.State.ExtraTurn--
	} else {
		f.State.Turn++
	}
}

func (f *Field) AddScore(value int) {
	f.Score += value
}

func (f *Field) SubEnergy(value int) {
	if f.Energy >= value {
		f.Energy -= value
	} else {
		value -= f.Energy
		f.Energy = 0
		f.SubHealth(value)
	}
}

func (f *Field) SubHealth(value int) {
	f.Health -= value
}

func (f *Field) isCardUsable(c *SkillCard.SkillCard) bool {
	switch c.Cost.Type {
	case "Energy":
		if f.Energy+f.Health < c.Cost.Value {
			return false
		}

	}
	return true
}

func (f *Field) UseCard(idx int) (State, error) {
	var err error
	if len(f.State.Cars.Hand) <= idx {
		err = fmt.Errorf("卡片只有%d張", len(f.State.Cars.Hand))
		return f.State, err
		//errors.New(fmt.Sprintf("卡片只有%d張", len(f.State.Cars.Hand)))
	} else if !f.isCardUsable(&f.State.Cars.Hand[idx]) {
		err = fmt.Errorf("不符合該卡牌使用條件")
		return f.State, err
	}
	/*
		else if check Cost
	*/

	switch f.State.Cars.Hand[idx].Cost.Type {
	case "Energy":
		f.SubEnergy(f.State.Cars.Hand[idx].Cost.Value)
	default:
		err = fmt.Errorf("未定義消耗類型:%s", f.State.Cars.Hand[idx].Cost.Type)
		return f.State, err
	}
	f.Cars.Hand[idx].Use(f)

	f.NextTrue()
	return f.State, err
}

func (f *Field) UseDrink(idx int) State {
	f.NextTrue()
	return f.State
}

func (f *Field) checkConditions(c []core.Condition) bool {
	for _, v := range c {
		if !f.checkCondition(v) {
			return false
		}
	}
	return true
}

func (f *Field) checkCondition(c core.Condition) bool {
	switch c {

	}
	return true
}

func (f *Field) useEffects(e []core.Effect) State {
	for _, v := range e {
		f.useEffect(v)
	}
	return f.State
}

func (f *Field) useEffect(e core.Effect) State {
	switch e {

	}
	return f.State
}
