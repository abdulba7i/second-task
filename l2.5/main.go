package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	// ... do something
	return nil
}

func main() {
	var err error

	// Вызываем фукнцию, и получаем интерфейс, в котором тип *main.customError, а значение nil
	// тип можно узнать к примеру черех "fmt.Printf("Тип переменной err: %T\n", err)"
	err = test()

	// ответом будет false, так как тип не имеет nil
	if err != nil {
		println("error") // вывод будет error
		return           // завершаем программу, поэтому дальнейший код выводится не будет
	}
	println("ok")
}
