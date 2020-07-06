//插入排序
package main

import "fmt"

func main() {
	arr := []int{1, 5, 3}
	fmt.Println(canMakeArithmeticProgression(arr))
}

func canMakeArithmeticProgression(arr []int) bool {
	if len(arr) == 2 {
		return true
	} else {
		arrSort := InsertSort(arr)
		gap := arrSort[1] - arrSort[0]
		for i := 0; i < len(arrSort)-1; i++ {
			if gap != arrSort[i+1]-arrSort[i] {
				return false
			}
		}
		return true
	}
}

func InsertSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	} else {
		for i := 1; i < len(arr); i++ {
			backup := arr[i]
			j := i - 1
			for j >= 0 && backup < arr[j] {
				arr[j+1] = arr[j]
				j--
			}
			arr[j+1] = backup
		}
		return arr
	}
}
