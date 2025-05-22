package pages

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Routing struct {
	Name string
	Open func() tea.Model
}

// SharedNavItems 共享導航項目列表

var SharedNavItems = []Routing{
	{Name: "劇情", Open: NewPlotModel},
	{Name: "偶像", Open: NewIdolModel},
	{Name: "主選單", Open: NewHomeModel},
	{Name: "競技場", Open: NewArenaModel},
	{Name: "抽卡", Open: NewPage4Model},
}
