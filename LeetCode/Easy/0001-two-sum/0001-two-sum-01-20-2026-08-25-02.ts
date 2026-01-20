function twoSum(nums: number[], target: number): number[] {
    for (let i = 0; i <= nums.length;i++) {
        if (nums[i + 1]) {
            const toFindNum = target - nums[i];
            for (let j = i + 1; j <= nums.length; j++) {
                if (nums[j] === toFindNum) {
                    return [i, j]
                }
                continue;
            }
        }
    }
};