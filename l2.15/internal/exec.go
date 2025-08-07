package internal

import (
	"fmt"
	"os"
	"os/exec"
)

func ExecExternal(args []string) {
	// создаём процесс, с переданными командами и аргументами args
	cmd := exec.Command(args[0], args[1:]...)

	// сохраняем команду в переменную, в случае если захотим остановить процесс
	currentCmd = cmd

	// имитация терминала:
	cmd.Stdin = os.Stdin   // ввод данных
	cmd.Stdout = os.Stdout // вывод команды в терминал
	cmd.Stderr = os.Stderr // вывод по ошибка
	if err := cmd.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка запуска:", err)
	}

	// после завершения работы очищаем переменную, чтобы никакой внешний процесс больше не выполнялся
	currentCmd = nil
}
