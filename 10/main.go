package main

import (
	"fmt"
	"math"
)

func TGroup(tmp []float64) map[int][]float64 {
	groups := make(map[int][]float64) // Инициализация карты для группировки температур

	for _, val := range tmp {
		key := int(math.Floor(val/10) * 10)    // Определение ключа группы путем округления вниз
		groups[key] = append(groups[key], val) // Добавление температуры в соответствующую группу
	}

	return groups
}

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := TGroup(temperatures)
	fmt.Println(groups)
}
