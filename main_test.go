package main

import (
	"fmt"
	"math"
	"testing"
)

func getShiftedRange(inArray []int, offset int) []int {
	curArray := append(inArray[offset:], inArray[:offset]...)
	return curArray
}
func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func BenchmarkBinarySearch(b *testing.B) {
	const minPow = 6
	const maxPow = 14
	maxArraySize := int(math.Pow(2,maxPow))
	benchArraySize := int(math.Pow(2,minPow))
	for run := minPow; run < maxPow; run++ {
		benchArraySize = benchArraySize * 2
		benchArray := makeRange(0, benchArraySize)
		name := fmt.Sprintf("Bench for length %v", benchArraySize)
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N * maxArraySize * maxArraySize / (benchArraySize * benchArraySize); i++ {
				for shift := 0; shift < benchArraySize; shift++ {
					tempArray := getShiftedRange(benchArray, shift)
					for offset := 0; offset < benchArraySize; offset++ {
						res := binSearch(tempArray, tempArray[offset])
						if res != offset {
							b.Errorf("Expected: %v  got: %v", offset, res)
						}
					}
				}

			}
		})
	}
}
func TestBinarySearch(t *testing.T) {
	const testSetSize = 20
	testSet := make([][]int, testSetSize)
	for n := 0; n < testSetSize; n++ {
		testSet[n] = makeRange(0, n)
	}
	t.Logf("Testing search each item:")
	t.Logf("Array, value")
	for i := 0; i < len(testSet); i++ {
		for shift := 0; shift < len(testSet[i]); shift++ {
			tempArray := getShiftedRange(testSet[i], shift)
			for offset := 0; offset < len(testSet[i]); offset++ {
				t.Logf("%v, %v",tempArray,tempArray[offset])
				res := binSearch(tempArray, tempArray[offset])
				if res != offset {
					t.Errorf("Expected: %v  got: %v", offset, res)
				}
			}
		}

	}
}
