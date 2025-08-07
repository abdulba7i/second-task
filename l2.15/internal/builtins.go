package internal

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

// HandleBuiltin проверяет, введенная команда из терминала пользователем является встроенной или нет. Если да, то возвращает true (cd, echo, pwd)
func HandleBuiltin(args []string) bool {
	// вытаскиваем аргумент команды
	switch args[0] {
	case "exit": // выход работы
		os.Exit(0)
	case "cd": // команда cd. Ниже ставим условие, в случае если путь не указан, вернем ошибку
		if len(args) < 2 {
			fmt.Fprintln(os.Stderr, "cd: путь не указан")
			return true
		}
		// если путь указан, то меняем директорию на указанную
		err := os.Chdir(args[1])
		if err != nil {
			// в случае, если директория не существует, вернем ошибку
			fmt.Fprintln(os.Stderr, "cd:", err)
		}
		return true
	case "pwd": // кейс с получением и выводом текущей директории
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Fprintln(os.Stderr, "pwd:", err)
		} else {
			fmt.Println(cwd)
		}
		return true
	case "echo":
		fmt.Println(strings.Join(args[1:], " "))
		return true
	case "kill":
		if len(args) < 2 {
			fmt.Fprintln(os.Stderr, "kill: pid не указан")
			return true
		}
		pid, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, "kill: неверный pid")
			return true
		}
		err = syscall.Kill(pid, syscall.SIGTERM)
		if err != nil {
			fmt.Fprintln(os.Stderr, "kill:", err)
		}
		return true
	case "ps": // вывод список запущенных процессоров
		cmd := exec.Command("ps", "-e", "-o", "pid,comm")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
		return true
	}

	// если команда не является встроенной, возвращаем ошибку
	return false
}
