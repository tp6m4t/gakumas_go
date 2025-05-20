package main

import (
	"fmt"
	"os"
	"strings"

	textinput "github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var mode string

// 定義選單節點
var menu = map[string][]string{
	"0": {"家"},
	"1": {"劇情", "a.主線", "b.個人", "c.支援卡", "d.活動"},
	"2": {"偶像", "a.角色卡", "b.回憶卡", "c.支援卡", "d.編成"},
	"3": {"競賽", "a.競技場", "b.偶像之路"},
	"4": {"抽卡", "a.常駐", "a.1.支援卡", "a.1.角色卡", "b.活動a", "b.1.支援卡", "b.1.角色卡"},
}

// 分頁顯示的最大項目數
const pageSize = 4

// TUI model
type model struct {
	pos      string
	quitting bool
	msg      string
	input    textinput.Model
	page     int
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "輸入選項後按 Enter"
	ti.Focus()
	return model{pos: "0", input: ti, page: 0}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.Type == tea.KeyEnter {
			input := strings.ToLower(strings.TrimSpace(m.input.Value()))
			switch input {
			case "ctrl+c", "q", "..", "exit":
				if m.pos != "0" {
					m.pos = "0"
					m.page = 0
					m.msg = "返回主選單"
				} else {
					m.quitting = true
					return m, tea.Quit
				}
			case "next":
				maxPage := (len(menu[m.pos]) - 2) / pageSize
				if m.page < maxPage {
					m.page++
				}
			case "prev":
				if m.page > 0 {
					m.page--
				}
			case "0", "1", "2", "3", "4":
				m.pos = input
				m.page = 0
				m.msg = ""
			case "a", "b", "c", "d", "a.1", "b.1":
				m.msg = fmt.Sprintf("選擇了子選項 [%s]", input)
			default:
				m.msg = fmt.Sprintf("未知指令：%s", input)
			}
			m.input.SetValue("") // 清除輸入框
		}
	}
	return m, cmd
}

func (m model) View() string {
	if m.quitting {
		return "再見！\n"
	}

	output := fmt.Sprintf("當前位置：[%s] %s\n", m.pos, menu[m.pos][0])
	output += fmt.Sprintf("\n選單 (第 %d 頁)：\n", m.page+1)
	items := menu[m.pos][1:]
	start := m.page * pageSize
	end := start + pageSize
	if end > len(items) {
		end = len(items)
	}
	for _, item := range items[start:end] {
		output += fmt.Sprintf("[%s]\n", item)
	}
	if len(items) > pageSize {
		output += "\n[prev] 上一頁	[next] 下一頁\n"
	}
	output += "\n數字轉跳主選單、字母選子選項，q/.. 離開\n"
	output += m.input.View() + "\n"
	if m.msg != "" {
		output += "\n→ " + m.msg + "\n"
	}
	return output
}

func runTUI() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Println("錯誤：", err)
		os.Exit(1)
	}
}

func runSimulate() {
	fmt.Println("模擬模式啟動：自動訓練中...")
	fmt.Println(`{"day": 1, "action": "dance", "result": {"dance": +3}}`)
}

func main() {
	rootCmd := &cobra.Command{
		Use:   "idol",
		Short: "偶像育成 CLI 工具",
		Run: func(cmd *cobra.Command, args []string) {
			switch mode {
			case "tui":
				runTUI()
			case "simulate":
				runSimulate()
			default:
				fmt.Println("請指定模式: --mode [tui|simulate]")
				os.Exit(1)
			}
		},
	}

	rootCmd.Flags().StringVar(&mode, "mode", "", "執行模式: tui / simulate")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
