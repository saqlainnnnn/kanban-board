package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type status int

const divisor = 4 

const (
	todo status = iota
	inProgress
	done
)

//encapsulates data related to the task
type Task struct {
	status      status
	title       string
	description string
	

}

// implementing the list.Item interface
func (t Task) FilterValue() string {
	return t.title
}

func (t Task) Title() string {
	return t.title
}

func (t Task) Description() string {
	return t.description
}

//  MAIN MODEL

type Model struct {
	focused status
	lists 	[]list.Model
	err  	error
	loaded 	bool
}

func New() *Model {
	return &Model{}
}

func (m *Model) initLists(width, height int) {
	defaultList := list.New([]list.Item{}, list.NewDefaultDelegate(), width/divisor, height)
	defaultList.SetShowHelp(false)
	m.lists = []list.Model{defaultList, defaultList, defaultList}
	
	//INIT TO DO
	m.lists[todo].Title = "TO DO"
	m.lists[todo].SetItems([]list.Item{
		Task{status: todo, title: "buy milk", description: "stawberry milk"},
		Task{status: todo, title: "eat sushi", description: "idk what roll, some soup"},
		Task{status: todo, title: "fold laundry", description: "wear clothes"},
	})

	//INIT IN PROGRESS
	m.lists[inProgress].Title = "In Progress"
	m.lists[inProgress].SetItems([]list.Item{
		Task{status: inProgress, title: "write code", description: "dont worry it's just go"},
	})

	//INIT DONE
	m.lists[done].Title = "DONE"
	m.lists[done].SetItems([]list.Item{
		Task{status: done, title: "stay cool", description: "like a cucumber😭😭"},
	})
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if !m.loaded {
			m.initLists(msg.Width, msg.Height)
			m.loaded = true
		}
		
	}
	var cmd tea.Cmd
	m.lists[m.focused], cmd = m.lists[m.focused].Update(msg)
	return m, cmd
}

func (m Model) View() string {

	if m.loaded {
		
		return lipgloss.JoinHorizontal(
			lipgloss.Left,
			m.lists[todo].View(),
			m.lists[inProgress].View(),
			m.lists[done].View(),
		)	
	} else {
		return "loading..."
	}
}

func main()  {
	m := New()
	p := tea.NewProgram(m)

	if err := p.Start(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}