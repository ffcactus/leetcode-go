package huawei_depends

import "fmt"

func assignTaskOrder(numTasks int, predependents [][]int) []int {
	if numTasks == 1 {
		return []int{predependents[0][1], predependents[0][0]}
	}
	solutions := make([][]int, 0)
	for _, eachTask := range predependents {
		if len(solutions) == 0 {
			solutions = append(solutions, []int{eachTask[1], eachTask[0]})
			continue
		}
		newSolutions := make([][]int, 0)
		for i := range solutions {
			newSolution, ok := merge(&solutions[i], eachTask)
			if !ok && newSolution == nil {
				return []int{}
			}
			if ok {
				newSolutions = make([][]int, 0)
				break
			}
			newSolutions = append(newSolutions, newSolution)
			fmt.Printf("newSolutions = %v\n", newSolutions)
		}
		solutions = append(solutions, newSolutions...)
		fmt.Println(solutions)
	}

	finalMerge(solutions)
	return []int{}
}

func merge(solution *[]int, task []int) ([]int, bool) {
	first := task[1]
	second := task[0]

	// 找到first, second的位置。
	firstIndex := -1
	sencondIndex := -1
	for i, v := range *solution {
		if v == first {
			firstIndex = i
		}
		if v == second {
			sencondIndex = i
		}
	}

	// 如果都有且 firstIndex < secondIndex, 则融入
	if firstIndex >= 0 && sencondIndex >= 0 && sencondIndex > firstIndex {
		return nil, true
	}

	// 如果都有且 firstIndex > secondIndex, 冲突
	if firstIndex >= 0 && sencondIndex >= 0 && sencondIndex < firstIndex {
		return nil, false
	}
	// 如果 secondIndex = 0 且 first < 0, 这合入。
	if sencondIndex == 0 && firstIndex < 0 {
		*solution = append([]int{first}, *solution...)
		return nil, true
	}
	if firstIndex == 0 && sencondIndex < 0 {
		*solution = append(*solution, second)
		return nil, true
	}
	return []int{first, second}, false
}

func mergeSolution(a, b *[]int) bool {
	lenA := len(*a)
	lenB := len(*b)
	if (*a)[lenA-1] == (*b)[0] {
		*a = append(*a, (*b)[1:]...)
		return true
	}
	if (*b)[lenB-1] == (*a)[0] {
		*a = append(*b, (*a)[1:]...)
		return true
	}
	return false
}

func finalMerge(solutions [][]int) []int {
	if len(solutions) == 1 {
		return solutions[0]
	}

	for i := 0; i < len(solutions); i++ {
		mg := solutions[i]
		for j := 0; j < len(solutions); j++ {
			if mergeSolution(&solutions[j], &mg) {
				solutions = append(solutions[:i], solutions[i+1:]...)
				break
			}
		}
	}
	if len(solutions) > 1 {
		return []int{}
	}
	return solutions[0]
}
