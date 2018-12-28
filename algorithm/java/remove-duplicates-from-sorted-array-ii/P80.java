class Solution {
    public int removeDuplicates(int[] nums) {
        if (nums == null || nums.length == 0) {
            return 0;
        }
        // skip first value;
        int count = 1;
        int p = 1;
        int currentCount = 1;
        for (int i = 1; i < nums.length; i++) {
            if (nums[i] == nums[i-1]) {
                currentCount++;
                if (currentCount <= 2) {
                    count++;
                    nums[p] = nums[i];
                    p++;
                }
            } else {
                currentCount = 1;
                count++;
                nums[p] = nums[i];
                p++;
            }
        }
        return count;
        
    }
}
