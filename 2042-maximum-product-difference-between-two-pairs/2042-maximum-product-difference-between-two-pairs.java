class Solution {
    public int maxProductDifference(int[] nums) {
        Arrays.sort(nums);
        int l0 = nums[nums.length - 1];
        int l1 = nums[nums.length - 2];
        int f0 = nums[0];
        int f1 = nums[1];
        return (l0 * l1) - (f0 * f1);
    }
}