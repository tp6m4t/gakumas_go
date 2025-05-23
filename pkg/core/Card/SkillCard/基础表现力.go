package SkillCard

type expressionBasics struct {
	SkillCardBase
}

func NewExpressionBasics() *expressionBasics {
	c := &expressionBasics{}
	c.SkillCardBase.Set(
		0, 0, 0, 0,
		"基础表现力表演",
		"M", []int{}, 4, "", 0, "活力+4  课程中限1次", true,
	)
	return c
}

func (c *expressionBasics) Use(f field) {
	f.SubEnergy(c.BaseCost)
}

func (c *expressionBasics) Upgrade() {}
