package main

import "testing"

type List Patient

func TestFilter(t *testing.T) {
	ListOfPacient := []List{ //условные данне дла проверки
		{1, 15, "333", "test", FullName{FirstName: "Petrov", LastName: "Serghei", Patronymic: "qww"}},
		{2, 25, "qq", "est", FullName{FirstName: "Ivanov", LastName: "Andrei", Patronymic: "qggg"}},
		{3, 45, "ee", "top", FullName{FirstName: "Sidorov", LastName: "Petar", Patronymic: "qhhh"}},
		{4, 65, "00", "test", FullName{FirstName: "Kozlov", LastName: "Ivan", Patronymic: "qkkk"}},
	}
	tests := []Options{
		{1, 15, "333", "test"},
		{2, 25, "qq", "est"},
		{3, 45, "ee", "top"},
		{4, 65, "00", "test"},
	}
	for _, val := range tests {
		if col := ListOfPacient.Filter(val); col != test.expected {
			t.Errorf("%d.Expected '%s', but got '%s'", i, test.expected, col)
		}

	}
}
