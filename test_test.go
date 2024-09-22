package move_zero_to_the_end

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	var nums = []int{0, 1, 0, 3, 12, 0, 14}         // Input array
	var expectedNums = []int{1, 3, 12, 14, 0, 0, 0} // The expected answer with correct length.
	moveZeroes(nums)                                // Calls your implementation

	assert.Equal(t, nums, expectedNums)
}
func moveZeroes(nums []int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			for j := i + 1; j < len(nums); j++ {
				if nums[j] != 0 {
					nums[i] = nums[j]
					nums[j] = 0
					break
				}
			}
		}
	}
}
