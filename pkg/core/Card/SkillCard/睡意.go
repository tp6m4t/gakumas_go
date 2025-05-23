package SkillCard

type trouble struct {
	SkillCardBase
}

func NewTrouble() *trouble {
	c := &trouble{}
	c.SkillCardBase.Set(
		3, 0, 0, 0,
		"睡意",
		"T", []int{}, 0, "", 0, "", true,
	)
	return c
}

func (c trouble) Use(f field) {
}

func (c trouble) Upgrade() {}
