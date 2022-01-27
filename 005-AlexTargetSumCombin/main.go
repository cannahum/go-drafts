package main

import "fmt"

func getTargetSumCombin(arr []int, target int, result [][]int) [][]int {
	if result == nil {
		result = [][]int{}
	}
	if target <= 0 || len(arr) == 0 {
		return result
	}
	//fmt.Println(arr)

	for i := len(arr) - 1; i >= 0; i-- {
		nextMax := arr[i]
		if nextMax > target {
			continue
		}
		if nextMax == target {
			result = append(result, []int{nextMax})
			continue
		}

		remainingArr := arr[0:i]
		//fmt.Println(remainingArr)
		for _, y := range getTargetSumCombin(remainingArr, target-nextMax, nil) {
			comb := append(y, nextMax)
			result = append(result, comb)
		}

		//nextMid := nextMax
		//for j := 0; nextMid < target; j++ {
		//	factor := j + 1
		//	nextMid = nextMax * factor
		//	// combinations of nextMaxes? 4 = 1 + 1 + 2 for example...
		//	for _, z := range getTargetSumCombin(remainingArr, target-nextMid) {
		//		for k := 0; k < factor; k++ {
		//			z = append(z, nextMax, nextMax)
		//		}
		//		result = append(result, z)
		//	}
		//}
	}

	return result
}

func main() {
	//result := getTargetSumCombin([]int{1, 2}, 2, nil)
	result := getTargetSumCombin([]int{1, 2, 3, 4, 5, 6, 7}, 7, nil)
	fmt.Println(result)
}
