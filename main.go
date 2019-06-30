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
		if item > arr[len(arr)-1] { //required element in a wrapped tail of an ordered array
			if arr[m] > arr[len(arr)-1] { //if the point is also in the tail - we check it to the left or to the right
				if arr[m] > item {
					r = m
				} else {
					l = m
				}
			} else { //if the point is not in the tail - it is definitely to the right
				r = m
			}
		} else { //required element in the displaced head of an ordered array
			if arr[m] < arr[len(arr)-1] { //if the point is also in the head - check it to the left or to the right
				if arr[m] > item {
					r = m
				} else {
					l = m
				}
			} else { //if the point is not in the head - it is definitely to the left
				l = m
			}
		}
	}
	return l
}
func main() {

}
