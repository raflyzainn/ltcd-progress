/**
 * @param {number[]} nums
 * @return {number[]}
 */
var findErrorNums = function(nums) {
    let duplicate;
    let missing;
    let result = [];

    // manual sorting
    for (let i = 0; i < nums.length; i++) {
        for (let j = i + 1; j < nums.length; j++) {
            if (nums[i] > nums[j]) {
                let temp = nums[i];
                nums[i] = nums[j];
                nums[j] = temp;
            }
        }
    }

    // cari duplicate
    for (let i = 0; i < nums.length - 1; i++) {
        if (nums[i] == nums[i + 1]) {
            duplicate = nums[i];
        }
    }

    // cari missing
    for (let i = 1; i <= nums.length; i++) {
        if (!nums.includes(i)) {
            missing = i;
        }
    }

    result[0] = duplicate;
    result[1] = missing;

    return result;
};