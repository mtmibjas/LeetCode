class Solution {
    public int maxAscendingSum(int[] nums) {

        int n = nums.length;
        int max = nums[0];
        int cursum = nums[0];

        for (int i = 1; i < n; i++) {
            if (nums[i - 1] < nums[i]) {
                cursum += nums[i];
            } else {
                if (max < cursum) {
                    max = cursum;
                }
                cursum = nums[i];
            }
        }

        if (cursum > max) {
            max = cursum;
        }

        return max;
    }
}