package Buff

type Buff int

const (
	GoodCondition Buff = iota // 好調
	PerfectForm               // 絕好調
	Concentration             // 集中
	Impression                // 好印象
	Motivation                // 幹勁
	Other                     // 其他
)

func (b Buff) String() string {
	return [...]string{
		"好調",
		"絕好調",
		"集中",
		"好印象",
		"幹勁",
		"其他",
	}[b]
}
