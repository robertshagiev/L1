package main

import (
	"fmt"
)

func binarySearch(nums []int, x int) int {
	min := 0
	max := len(nums) - 1 // Задаем границы интервала поиска

	for min < max { // Задаем цикл: пока в интервале больше одного элемента
		mid := (max + min) / 2 // Смотрим в середину
		switch {
		case nums[mid] == x: // В середине искомый элемент - прекращаем поиск
			return mid
		case nums[mid] > x: // Определяемся с какой стороны искомый элемент - сокращаем интервал
			max = mid - 1 // ищем слева
		case nums[mid] < x:
			min = mid + 1 // ищем справа
		}
	}
	return -1
}

func main() {
	sortedArr := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(binarySearch(sortedArr, 5))
}
