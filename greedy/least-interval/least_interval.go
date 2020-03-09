package leastinterval

import (
	"fmt"
	"math"
	"sort"
)

func leastInterval(tasks []byte, n int) int {
	allTask := make([]task, 26)
	count := 0
	for i := range allTask {
		allTask[i].name = fmt.Sprintf("%c", 'A'+byte(i))
		allTask[i].countdown = 0
	}
	for _, v := range tasks {
		count++
		allTask[v-'A'].count++
	}
	for i := range allTask {
		if allTask[i].count == 0 {
			allTask[i].countdown = math.MaxInt64
		}
	}
	// pick up a task
	time := 0
	for count > 0 {
		sortTask(allTask)
		time++
		if allTask[0].countdown > 0 {
			fmt.Printf("* ")
		} else {
			fmt.Printf("%s ", allTask[0].name)
			count--
			allTask[0].count--
			if allTask[0].count == 0 {
				allTask[0].countdown = math.MaxInt64
			} else {
				allTask[0].countdown = n + 1
			}

		}
		countdownTask(allTask)
	}
	return time
}

type task struct {
	name      string
	count     int
	countdown int
}

func sortTask(tasks []task) {
	sort.Slice(tasks, func(i, j int) bool {
		if tasks[i].countdown == tasks[j].countdown {
			return tasks[i].count > tasks[j].count
		}
		return tasks[i].countdown < tasks[j].countdown
	})
}

func countdownTask(allTask []task) {
	for i := range allTask {
		if allTask[i].countdown > 0 {
			allTask[i].countdown--
		}
	}
}
