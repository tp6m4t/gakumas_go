package main

import (
	field "LocalProject/pkg/CardGame/Field"
	skillCard "LocalProject/pkg/core/Card/SkillCard"
	"fmt"
)

func main() {
	fmt.Println("Card Game Main")
	var f field.Field
	//f.State.Turns = {field.Turn{}}
	f.ExtraTurn = 10
	f.Cars.Hand = append(f.Cars.Hand, skillCard.AppealBasics)
	f.Health = 0
	f.Energy = 11
	f.MaxHealth = 5
	for !f.IsEnd() {
		//顯示當前狀態
		fmt.Println(f.String())
		nextTrun := UserAction(&f)
		if !nextTrun {
			break
		}
	}
	fmt.Println("遊戲結束")
}

func UserAction(f *field.Field) bool {
	tips := "\n使用卡片: c\n使用飲料: d\n回復體力: s\n退出: e\n\n請輸入指令:"
	in := UserInput("\033[H\033[2J"+f.String()+"\n"+tips, []string{"c", "d", "e", "s"})
	switch in {
	case "c":
		UseCard(f)
	case "d":
		UseDrink(f)
	case "s":
		f.SkipTurn()
	case "e":
		return false
	}
	return true
}

func UseCard(f *field.Field) {
	var in int
	fmt.Print("請輸入要與使用的卡片編號:")
	fmt.Scan(&in)
	_, err := f.UseCard(in)
	if err != nil {
		fmt.Println(err)
	}
}

func UseDrink(f *field.Field) {
	var in int
	fmt.Scan(&in)
	f.UseDrink(in)
}

func UserInput(tips string, Allow []string) string {
	var in string
	var end bool = false
	for !end {
		fmt.Println(tips)
		fmt.Scan(&in)
		for _, v := range Allow {
			if in == v {
				end = true
				break
			}
		}
		if !end {
			fmt.Printf("輸入錯誤: %s\n", in)
		}
	}
	return in
}
