package buff

import "fmt"

type field interface {
	AddScore(value int)
	SubEnergy(value int)
	SubHealth(value int)
	AddEnergy(value int)
	AddHealth(value int)
	Subscribe(event string, f func(interface{}))
}

//Buff interface
type Buff interface {
	GetCount() int
	AddCount(count int)
	GetID() int
	GetName() string
	Subscribe(f field)
}

//BaseBuff
type BaseBuff struct {
	id           int
	name         string
	descriptions string
	count        int
}

func (b *BaseBuff) GetID() int {
	return b.id
}

func (b *BaseBuff) GetName() string {
	return b.name
}
func (b *BaseBuff) GetCount() int {
	return b.count
}
func (b *BaseBuff) GetDescription() string {
	return b.descriptions
}
func (b *BaseBuff) Subscribe(f field) {}

func (BaseBuff) IsDebuff() bool {
	return false
}

//Debuff
type DeBuff struct {
}

func (DeBuff) IsDebuff() bool {
	return true
}

//buffBuid
var buffBuild = make(map[int]func(value int) Buff)
var buffMap = make(map[string]int)

func init() {
	fmt.Println("Buff init")
	BuffBuildAdd(func(value int) Buff { return NewGoodImpression(value) })
}

func BuffBuildAdd(BuildFunc func(value int) Buff) {
	Card := BuildFunc(0)
	_, ok := buffBuild[Card.GetID()]
	if ok {
		fmt.Printf("%s,%s技能牌ID:%d重複\n", Card.GetName(), buffBuild[Card.GetID()](0).GetName(), Card.GetID())
	} else {
		buffMap[Card.GetName()] = Card.GetID()
		buffBuild[Card.GetID()] = BuildFunc
	}
}

func NewBuffByID(id int, value int) Buff {
	_, ok := buffBuild[id]
	if !ok {
		fmt.Printf("buff ID:%d不存在\n", id)
		return nil
	}
	return buffBuild[id](value)
}
func NewBuffByName(name string, value int) Buff {
	id, ok := buffMap[name]
	if !ok {
		fmt.Printf("buff詞條:%s 不存在\n", name)
		return nil
	}
	return NewBuffByID(id, value)
}

func GetBuffID(name string) int {
	id, ok := buffMap[name]
	if !ok {
		fmt.Printf("buff詞條:%s 不存在\n", name)
		return -1
	}
	return id
}

/*
const (
	//職業相關buff
	GoodCondition   // 好調(分數上升量提升50%)
	PerfectForm     // 絕好調(分數上升量提升 好調*10%)
	Concentration   // 集中(數值上升時而外增加數值)
	Impression      // 好印象(回合結束依數值上升分數)
	Motivation      // 幹勁(能量上升時而外增加能量)
	FullPowerPoints // 全力
	Strength        // 堅決
	Preservation    // 溫存
	//悠閒
	//其他bufff
	//體力消費減少
	//卡牌可使用次數數增加
	//免消耗體力
	//應援
	//debuff
	//體力消費增加
	//活力增加無效
	//狀態不佳
	//使用體力追加
	//禁用部分手牌
	//手牌減少
	//心情浮躁 手牌消耗體力隨機變化
	//膽怯 能量增加量-1
	//精神萎靡 數值無法上升
	//不安 能量增加量-33%
	Other // 其他
)

func (b Buff) String() string {
	return [...]string{
		"好調",
		"絕好調",
		"集中",
		"好印象",
		"幹勁",
		"全力",
		"堅決",
		"溫存",
		"其他",
	}[b]
}
*/
