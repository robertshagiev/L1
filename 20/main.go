package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "snow dog sun"
	fmt.Println(s)
	fmt.Println(turnString(s))
}

func turnString(s string) string {
	w := strings.Fields(s) //Разбиваем строку на массив, состоящий из слов (точнее из объектов)
	n := len(w)            // Получаем количество слов
	for i := 0; i < n/2; i++ {
		w[i], w[n-i-1] = w[n-i-1], w[i] // Меняем местами слова
	}
	return strings.Join(w, " ") // Объединяем элементы для создания единной строки
}
