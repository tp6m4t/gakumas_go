package field

import (
	"LocalProject/pkg/CardGame/core"
	"fmt"
)

type Turn struct {
	Color *int
	Skill []core.Skill
}

type CarDeck struct {
	Deck         []core.SkillCard //山扎
	Hand         []core.SkillCard //手扎
	Discard      []core.SkillCard //捨扎
	BanishedPile []core.SkillCard //除外
}

type State struct {
	Cars      CarDeck     //牌堆
	Score     int         //得分
	Health    int         //當前體力
	MaxHealth int         //最大體力
	Shield    int         //當前護盾
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
	str += fmt.Sprintf("回和數:%d/%d+%d", len(f.Turns), f.Turn, f.ExtraTurn)
	return str
}

func (f *Field) IsEnd() bool {
	return f.State.Turn == len(f.State.Turns) && f.State.ExtraTurn == 0
}

func (f *Field) NextTrue() {
	if f.State.Turn == len(f.State.Turns) {
		f.State.ExtraTurn--
	} else {
		f.State.Turn++
	}
}

func (f *Field) UseCard(idx int) (State, error) {
	var err error
	if len(f.State.Cars.Hand) <= idx {
		err = fmt.Errorf("卡片只有%d張", len(f.State.Cars.Hand))
		return f.State, err
		//errors.New(fmt.Sprintf("卡片只有%d張", len(f.State.Cars.Hand)))
	}
	for _, skil := range f.State.Cars.Hand[idx].SkilList {
		//檢查符合發動條件
		if f.checkConditions(skil.Condition) {
			f.useEffects(skil.Effect)
		}
	}
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
