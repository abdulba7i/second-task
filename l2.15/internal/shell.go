package internal

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
)

// currentCmd нужен для того, чтобы можно могли отправить Ctrl+C только процессу, не завершая сам при этом shell
var currentCmd *exec.Cmd

// RunShell основной цикл обработки командной оболочки
func RunShell() {
	// Канал для перехвата Ctrl+C
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT)
	go func() {
		// Если сейчас запущен внешний процесс — отправляем ему SIGINT
		for range sigCh {
			if currentCmd != nil && currentCmd.Process != nil {
				currentCmd.Process.Signal(syscall.SIGINT)
			}
		}
		// Если никакой процесс не запущен shell просто продолжает работу
	}()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("minishell> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			// Если пользователь нажал Ctrl+D (EOF), то shell завершает работу
			if err == io.EOF {
				fmt.Println()
				break
			}
			fmt.Fprintln(os.Stderr, "Ошибка чтения:", err)
			continue
		}
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// подстановка переменных окружения
		line = os.ExpandEnv(line)

		// если есть |, разбиваем на команды и соединяем их через os.Pipe
		if strings.Contains(line, "|") {
			parts := strings.Split(line, "|")
			cmds := make([]*exec.Cmd, 0)
			for _, p := range parts {
				args := strings.Fields(strings.TrimSpace(p))
				if len(args) == 0 {
					continue
				}
				cmds = append(cmds, exec.Command(args[0], args[1:]...))
			}
			var lastStdout *os.File
			for i := 0; i < len(cmds)-1; i++ {
				r, w, _ := os.Pipe()
				cmds[i].Stdout = w
				cmds[i+1].Stdin = r
				if i == 0 {
					cmds[i].Stdin = os.Stdin
				}
				lastStdout = w
			}
			if len(cmds) > 0 {
				cmds[len(cmds)-1].Stdout = os.Stdout
			}
			for _, c := range cmds {
				c.Stderr = os.Stderr
			}
			for _, c := range cmds {
				currentCmd = c
				c.Start()
			}
			for _, c := range cmds {
				c.Wait()
			}
			if lastStdout != nil {
				lastStdout.Close()
			}
			currentCmd = nil
			continue
		}

		// редиректы ввода/вывода: > и <
		if strings.Contains(line, ">") || strings.Contains(line, "<") {
			args := strings.Fields(line)
			var (
				cmdArgs         []string
				inFile, outFile string
			)
			for i := 0; i < len(args); i++ {
				if args[i] == ">" && i+1 < len(args) {
					outFile = args[i+1]
					i++
				} else if args[i] == "<" && i+1 < len(args) {
					inFile = args[i+1]
					i++
				} else {
					cmdArgs = append(cmdArgs, args[i])
				}
			}
			if len(cmdArgs) == 0 {
				continue
			}
			cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
			if inFile != "" {
				f, err := os.Open(inFile)
				if err == nil {
					cmd.Stdin = f
					defer f.Close()
				}
			}
			if outFile != "" {
				f, err := os.Create(outFile)
				if err == nil {
					cmd.Stdout = f
					defer f.Close()
				}
			}
			if inFile == "" {
				cmd.Stdin = os.Stdin
			}
			if outFile == "" {
				cmd.Stdout = os.Stdout
			}
			cmd.Stderr = os.Stderr
			currentCmd = cmd
			cmd.Run()
			currentCmd = nil
			continue
		}

		// условные операторы: если есть && или ||
		if strings.Contains(line, "&&") {
			parts := strings.SplitN(line, "&&", 2)
			first := strings.TrimSpace(parts[0])
			second := strings.TrimSpace(parts[1])
			args1 := strings.Fields(first)
			args2 := strings.Fields(second)
			if len(args1) == 0 || len(args2) == 0 {
				continue
			}
			cmd := exec.Command(args1[0], args1[1:]...)
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if cmd.Run() == nil {
				cmd2 := exec.Command(args2[0], args2[1:]...)
				cmd2.Stdin = os.Stdin
				cmd2.Stdout = os.Stdout
				cmd2.Stderr = os.Stderr
				cmd2.Run()
			}
			continue
		}
		if strings.Contains(line, "||") {
			parts := strings.SplitN(line, "||", 2)
			first := strings.TrimSpace(parts[0])
			second := strings.TrimSpace(parts[1])
			args1 := strings.Fields(first)
			args2 := strings.Fields(second)
			if len(args1) == 0 || len(args2) == 0 {
				continue
			}
			cmd := exec.Command(args1[0], args1[1:]...)
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if cmd.Run() != nil {
				cmd2 := exec.Command(args2[0], args2[1:]...)
				cmd2.Stdin = os.Stdin
				cmd2.Stdout = os.Stdout
				cmd2.Stderr = os.Stderr
				cmd2.Run()
			}
			continue
		}

		// команда встроенная или внешняя
		args := strings.Fields(line)
		if len(args) == 0 {
			continue
		}

		// если до этого момент все ок, то вызываем HandleBuiltin, где работае со встроенными командамми
		if HandleBuiltin(args) {
			continue
		}

		// Внешняя команда
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		currentCmd = cmd
		cmd.Run()
		currentCmd = nil
	}
}
