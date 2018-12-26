class Solution {
    public void swap(int[] nums, int i, int j) {
        int tmp = nums[i];
        nums[i] = nums[j];
        nums[j] = tmp;
    }
    public void sortColors(int[] nums) {
        if (nums == null || nums.length <= 1) {
            return;
        }
        int left = 0;
        int right = nums.length-1;
        for (int i = 0; i < nums.length && i <= right;) {
            if (nums[i] == 0) {
                swap(nums,left, i);
                left++;
                i++;
            } else if (nums[i] == 2) {
                swap(nums, right, i);
                right--;
            } else {
                i++;
            }
        }
        
    }
}
