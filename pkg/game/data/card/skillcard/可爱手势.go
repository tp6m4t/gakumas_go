package skillcard

type cuteGestures struct {
	SkillCardBase
}

func NewCuteGestures() *cuteGestures {
	c := &cuteGestures{}
	c.SkillCardBase.Set(
		4, 0, 0, 0,
		"可爱手势",
		"A", []int{}, 4, "", 5, "好印象+2  提升好印象100%的数值  课程中限1次", true,
	)
	return c
}

func (c *cuteGestures) Use(f field) {
	f.SubEnergy(c.BaseCost)
	f.AddBuff("好印象", 2)
}

func (c *cuteGestures) Upgrade() {}
