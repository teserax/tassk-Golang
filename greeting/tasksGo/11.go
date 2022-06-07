package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type User struct {
	ID        int64  `validate:"required"`
	Email     string `validate:"required,email"`
	FirstName string `validate:"required"`
}

func main() {
	// создали пустую структуру, чтобы проверить валидацию
	u := User{}

	// создаем валидатор
	v := validator.New()

	// метод Struct валидирует переданную структуру и возвращает ошибку `error`, если какое-то поле некорректно
	fmt.Println(v.Struct(u))
}
