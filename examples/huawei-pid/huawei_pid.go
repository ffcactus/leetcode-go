package huawei_pid

func terminateProcess(pid []int, ppid []int, kill int) []int {
	ret := make([]int, 0)
	ret = append(ret, kill)

	parent := ret
	for {
		newChild := make([]int, 0)
		for _, v := range parent {
			newChild = append(newChild, child(pid, ppid, v)...)
		}
		ret = append(ret, newChild...)
		if len(newChild) == 0 {
			break
		} else {
			parent = newChild
		}
	}
	return ret

}

func child(pid []int, ppid []int, kill int) []int {
	ret := make([]int, 0)
	for i, v := range ppid {
		if v == kill {
			ret = append(ret, pid[i])
		}
	}
	return ret
}
