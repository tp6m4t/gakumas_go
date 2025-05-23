package field

import (
	"LocalProject/pkg/core/Card/SkillCard"
	"fmt"
)

// 檢查是否符合出牌條件
func (f *Field) isCardUsable(c SkillCard.SkillCard) bool {
	return c.IsUsable(f)
}

// 使用卡片: 扣除花費/提升分數/調用技能Use
func (f *Field) UseCard(idx int) (State, error) {
	var err error
	if len(f.State.Cars.Hand) <= idx {
		err = fmt.Errorf("卡片只有%d張", len(f.State.Cars.Hand))
		return f.State, err
		//errors.New(fmt.Sprintf("卡片只有%d張", len(f.State.Cars.Hand)))
	} else if !f.isCardUsable(f.State.Cars.Hand[idx]) {
		err = fmt.Errorf("不符合該卡牌使用條件")
		return f.State, err
	}
	/*
		else if check Cost
	*/

	f.Cars.Hand[idx].Use(f)

	if f.Cars.Hand[idx].GetIsBanished() {
		f.Cars.BanishedPile.Push(f.Cars.Hand[idx])
	} else {
		f.Cars.Discard.Push(f.Cars.Hand[idx])
	}
	f.Cars.Hand.Remove(idx)
	if true { //檢查回合是否結束
		f.NextTrue()
	}
	return f.State, err
}
