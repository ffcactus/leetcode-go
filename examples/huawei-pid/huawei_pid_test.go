package huawei_pid

import (
	"fmt"
	"testing"
)

func TestTerminateProcess_Case1(t *testing.T) {
	pid := []int{1}
	ppid := []int{0}
	kill := 1
	child := terminateProcess(pid, ppid, kill)
	fmt.Println(child)
}

func TestTerminateProcess_Case2(t *testing.T) {
	pid := []int{1, 3, 10, 5}
	ppid := []int{3, 0, 5, 3}
	kill := 5
	child := terminateProcess(pid, ppid, kill)
	fmt.Println(child)
}

func TestTerminateProcess_Case3(t *testing.T) {
	pid := []int{3}
	ppid := []int{0}
	kill := 5
	child := terminateProcess(pid, ppid, kill)
	fmt.Println(child)
}

func TestTerminateProcess_Case4(t *testing.T) {
	pid := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	ppid := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	kill := 1
	child := terminateProcess(pid, ppid, kill)
	fmt.Println(child)
}

func TestTerminateProcess_Case5(t *testing.T) {
	pid := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	ppid := []int{1, 1, 1, 1, 1, 1, 2, 2, 0, 0}
	kill := 2
	child := terminateProcess(pid, ppid, kill)
	fmt.Println(child)
}
