func twoSum(nums []int, target int) []int {
    m := make(map[int]int)

    for i, num := range nums {
        diff := target - num
        val, ok := m[diff]
        if (ok) {
            return []int{i, val}
        }
        m[num] = i
    }

    return nil
}