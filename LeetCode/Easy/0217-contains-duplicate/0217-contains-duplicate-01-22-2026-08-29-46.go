func containsDuplicate(nums []int) bool {
    m := make(map[int]int)
    for i, num := range nums {
        _, isValInM := m[num]
        if isValInM {
            return true
        }
        m[num] = i
    }

    return false
}