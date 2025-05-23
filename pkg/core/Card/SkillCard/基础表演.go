package SkillCard

type appealBasics struct {
	SkillCardBase
}

func NewAppealBasics() *appealBasics {
	c := &appealBasics{}
	c.SkillCardBase.Set(
		1, 0, 0, 0,
		"基础表演",
		"A", []int{9}, 0, "", 4, "数值+9", false,
	)
	return c
}

func (c *appealBasics) Use(f field) {
	f.AddScore(c.BaseScore[0])
}

func (c *appealBasics) Upgrade() {}
