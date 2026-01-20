func twoSum(nums []int, target int) []int {
    length := len(nums)

    for i:=0; i < length; i++ {
        diff := target - nums[i]
        for j:=i+1; j < length; j++ {
            if (nums[j] == diff) {
                return []int{i, j}
            }
        }
    }

    return nil
}