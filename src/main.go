package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

var VERSION = "0.0.1"
var allowedArgs = []string{"get-env", "save", "get-account", "version", "help"}
var choices = []string{"Yes", "No"}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit

		case "enter":
			m.choice = choices[m.cursor]
			return m, tea.Quit

		case "down", "j":
			m.cursor++
			if m.cursor >= len(choices) {
				m.cursor = 0
			}

		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(choices) - 1
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	s := strings.Builder{}
	s.WriteString("Would you like to save the current AWS environment?\n\n")

	for i := 0; i < len(choices); i++ {
		if m.cursor == i {
			s.WriteString("(•) ")
		} else {
			s.WriteString("( ) ")
		}
		s.WriteString(choices[i])
		s.WriteString("\n")
	}
	s.WriteString("\n(press q to quit)\n")

	return s.String()
}

func startInteractiveMode() {
	p := tea.NewProgram(model{})

	m, err := p.Run()
	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}

	if m, ok := m.(model); ok && m.choice != "" {
		if m.choice == "Yes" {
			err, awsEnv := getAWSEnv()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			writeIniFile(awsEnv)
			fmt.Println("Saving...")
		} else {
			fmt.Println("Not saving")
		}
	}
}

func main() {
	invaldArgsCheck()

	mode := os.Args[1]

	// Make sure the config directory exists
	// It should exist, if you have the aws cli installed,
	// but just in case...
	createDirectoryIfNotExists(getAWSConfigPathDir())

	switch mode {
	case "get-account":
		getAWSCallerIdentity()
	case "get-env":
		outputEnvironmentExports()
	case "save":
		startInteractiveMode()
	case "version":
		fmt.Println(VERSION)
	case "help":
		printHelp()
	}
}
