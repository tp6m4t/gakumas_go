package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Page struct {
	Name        string
	Description string
	Options     map[string]*Page // 子頁面或同級頁面轉跳
	Parent      *Page            // 指向父頁面（可退回用）
	BackOption  string           // 退回選項名稱 (例如 "b" 退回)
}

func (p *Page) Show() {
	fmt.Println("=== " + p.Name + " ===")
	fmt.Println(p.Description)
	fmt.Println("選項：")
	for k := range p.Options {
		fmt.Println(" -", k)
	}
	if p.Parent != nil {
		fmt.Printf("輸入 '%s' 退回上一頁\n", p.BackOption)
	}
	fmt.Print("請輸入選項: ")
}

func main() {
	// 建立頁面
	root := &Page{
		Name:        "主頁面",
		Description: "這是主頁面",
		Options:     make(map[string]*Page),
		Parent:      nil,
	}

	sub1 := &Page{
		Name:        "子頁面1",
		Description: "這是子頁面1",
		Options:     make(map[string]*Page),
		Parent:      root,
		BackOption:  "b",
	}

	sub2 := &Page{
		Name:        "子頁面2",
		Description: "這是子頁面2",
		Options:     make(map[string]*Page),
		Parent:      root,
		BackOption:  "b",
	}

	// 同級頁面互相轉跳
	sub1.Options["goto2"] = sub2
	sub2.Options["goto1"] = sub1

	// 主頁面指向子頁面
	root.Options["1"] = sub1
	root.Options["2"] = sub2

	currentPage := root

	scanner := bufio.NewScanner(os.Stdin)

	for {
		currentPage.Show()

		if !scanner.Scan() {
			break
		}
		input := strings.TrimSpace(scanner.Text())

		// 退回上一頁
		if input == currentPage.BackOption && currentPage.Parent != nil {
			currentPage = currentPage.Parent
			continue
		}

		// 跳轉選項
		nextPage, ok := currentPage.Options[input]
		if ok {
			currentPage = nextPage
		} else {
			fmt.Println("無效選項，請重新輸入。")
		}
	}
}
