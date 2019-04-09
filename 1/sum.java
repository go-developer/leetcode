class Solution {
    public int[] twoSum(int[] nums, int target) {
        int index = 0;
        int len = nums.length;
        int tmp = 0;
        int[] result = new int[2];
        boolean isHasResult = false;
        for (index = 0; index < len; index++) {
            for (tmp = index + 1; tmp < len; tmp++) {
                if (nums[index] + nums[tmp] == target) {
                    result[0] = index;
                    result[1] = tmp;
                    isHasResult = true;
                    break;
                }
            }
            if (isHasResult) {
                break;
            }
        }
        return result;
    }
}