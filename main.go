package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/taipi/kakumasu_CLI/pages"
)

func main() {
	// 初始化應用程式，從首頁開始
	p := tea.NewProgram(pages.NewHomeModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("程式出錯: %v", err)
		os.Exit(1)
	}
}
