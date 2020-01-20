package sort_colors

import "fmt"

const (
	red   = 0
	white = 1
	blue  = 2
)

func sortColors(nums []int) {
	redHead := 0
	blueHead := len(nums) - 1
	for i := 0; i <= blueHead; {
		switch nums[i] {
		case red:
			nums[redHead], nums[i] = nums[i], nums[redHead]
			redHead++
			i++
		case blue:
			nums[blueHead], nums[i] = nums[i], nums[blueHead]
			blueHead--
		case white:
			i++
		}
	}
	fmt.Println(nums)
}
