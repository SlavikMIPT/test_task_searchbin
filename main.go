package main

func binSearch(arr []int, item int) int {
	var m int
	l := -1
	r := len(arr)
	for l < r-1{
		m = (l + r) / 2
		if arr[m] == item{
			return m
		}
		if item > arr[len(arr)-1] { //искомый элемент  в завернутом хвосте упорядоченного массива
			if arr[m] > arr[len(arr)-1] { //если точка тоже в хвосте - проверяем левее она или правее
				if arr[m] > item {
					r = m
				} else {
					l = m
				}
			} else { //если точка не в хвосте - она однозначно правее
				r = m
			}
		} else { //искомый элемент в смещенной голове упорядоченного массива
			if arr[m] < arr[len(arr)-1] { //если точка тоже в голове - проверяем левее она или правее
				if arr[m] > item {
					r = m
				} else {
					l = m
				}
			} else { //если точка не в голове - она однозначно левее
				l = m
			}
		}
	}
	return l
}
func main() {

}
