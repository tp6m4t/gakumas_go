package skillcard

import "fmt"

type field interface {
	AddScore(value int)
	SubEnergy(value int)
	SubHealth(value int)
	AddEnergy(value int)
	AddHealth(value int)
	AddBuff(name string, value int)
}

// 只記錄圖鑑需顯示的牌面資料
type SkillCardBase struct {
	ID          int    //卡片唯一識別碼（ID
	Plvl        int    //解鎖等級 -1表示來源為支援卡
	Plan        int    //ALL(0) 感性(1) 理性(2) 非凡(3)
	Rarity      int    //稀有度 N(0) R(1) SR(2) SSR(3)
	Name        string //卡片名稱
	Type        string //"M"	,"A"
	BaseScore   []int  //分數
	BaseEnergy  int    //可獲得能量
	BaseCost    int    //消耗多少
	CostType    string //消耗甚麼 "Energy"(優先消耗能量),"Health"(直接消耗體力)
	Description string //技能或效果描述
	IsBanished  bool   //使用後是否除外
	BuffIcon    []int  //buff圖標
}

func (c *SkillCardBase) GetID() int {
	return c.ID
}

func (c *SkillCardBase) GetName() string {
	return c.Name
}

func (c *SkillCardBase) GetPlan() int {
	return c.Plan
}

func (c *SkillCardBase) GetRarity() int {
	return c.Rarity
}

func (c *SkillCardBase) GetPlvl() int {
	return c.Plvl
}

func (c *SkillCardBase) GetType() string {
	return c.Type
}

func (c *SkillCardBase) GetBaseScore() []int {
	return c.BaseScore
}

func (c *SkillCardBase) GetBaseEnergy() int {
	return c.BaseEnergy
}

func (c *SkillCardBase) GetBaseCost() int {
	return c.BaseCost
}

func (c *SkillCardBase) GetCostType() string {
	return c.CostType
}

func (c *SkillCardBase) GetDescription() string {
	return c.Description
}

func (c *SkillCardBase) GetIsBanished() bool {
	return c.IsBanished
}

func (c *SkillCardBase) GetBuffIcon() []int {
	return c.BuffIcon
}

func (c *SkillCardBase) Set(
	id int, Plvl int, Plan int, Rarity int,
	Name string, Type string, BaseScore []int,
	BaseEnergy int, CostType string, BaseCost int,
	Description string, IsBanished bool) {
	c.ID = id
	c.Plvl = Plvl
	c.Plan = Plan
	c.Rarity = Rarity
	c.Name = Name
	c.Type = Type
	c.BaseScore = BaseScore
	c.BaseEnergy = BaseEnergy
	c.CostType = CostType
	c.BaseCost = BaseCost
	c.Description = Description
	c.IsBanished = IsBanished
}

func (c *SkillCardBase) IsUsable(f field) bool {
	return true
}

type SkillCard interface {
	Use(f field)           //使用卡牌
	Upgrade()              //升級卡牌
	IsUsable(f field) bool //是否可用
	//Defind in SkillCardBase
	GetID() int             //獲得ID
	GetName() string        //獲得名稱
	GetPlan() int           //獲得類型 ALL(0) 感性(1) 理性(2) 非凡(3)
	GetRarity() int         //獲得稀有度 N(0) R(1) SR(2) SSR(3)
	GetPlvl() int           //獲得解鎖等級 -1表示來源為支援卡
	GetType() string        //獲得類型"M"	,"A"
	GetBaseScore() []int    //獲得分數
	GetBaseEnergy() int     //獲得可獲得能量
	GetBaseCost() int       //獲得消耗多少
	GetCostType() string    //獲得消耗甚麼 "Energy"(優先消耗能量),"Health"(直接消耗體力)
	GetDescription() string //獲得技能或效果描述
	GetIsBanished() bool    //使用後是否除外
	//GetBuffIcon() //獲得buff圖標
}

var skillCardBuild = make(map[int]func() SkillCard)
var skillCardMap = make(map[string]int)

func init() {
	fmt.Println("SkillCard init")
	SkillCardBuildAdd(func() SkillCard { return NewAppealBasics() })
	SkillCardBuildAdd(func() SkillCard { return NewExpressionBasics() })
	SkillCardBuildAdd(func() SkillCard { return NewPoseBasics() })
	SkillCardBuildAdd(func() SkillCard { return NewTrouble() })
	SkillCardBuildAdd(func() SkillCard { return NewCuteGestures() })
}

func SkillCardBuildAdd(BuildFunc func() SkillCard) {
	Card := BuildFunc()
	_, ok := skillCardBuild[Card.GetID()]
	if ok {
		fmt.Printf("%s,%s技能牌ID:%d重複\n", Card.GetName(), skillCardBuild[Card.GetID()]().GetName(), Card.GetID())
	} else {
		skillCardMap[Card.GetName()] = Card.GetID()
		skillCardBuild[Card.GetID()] = BuildFunc
	}
}

func NewSkillCardByID(id int) SkillCard {
	_, ok := skillCardBuild[id]
	if !ok {
		fmt.Printf("技能牌ID:%d不存在\n", id)
		return nil
	}
	return skillCardBuild[id]()
}
func NewSkillCardByName(name string) SkillCard {
	id, ok := skillCardMap[name]
	if !ok {
		fmt.Printf("技能牌名稱:%s 不存在\n", name)
		return nil
	}
	return NewSkillCardByID(id)
}
