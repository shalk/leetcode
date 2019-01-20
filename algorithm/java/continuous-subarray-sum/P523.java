class Solution {
    public boolean checkSubarraySum(int[] nums, int k) {
        if(nums == null || nums.length < 2) {
            return false;
        }
        if (k == 0) {
            for (int i =1; i < nums.length; i++) {
                if (nums[i-1] == 0 && nums[i] == 0) {
                    return true;
                }
            }
            return false;
        }
        if(k < 0) {
            k = -k;
        }
        for (int i =0; i < nums.length; i++) {
            nums[i] = nums[i] % k;
            if (i != 0 && nums[i] == 0 && nums[i-1] == 0) {
                return true;
            }
        }
        
        HashSet<Integer> set = new HashSet<>();
        int sum = nums[0];
        int pre = sum % k;
        for (int j =1; j < nums.length; j++) {
            sum += nums[j];
            if (sum % k == 0) {
                return true;
            } else {
               if (set.contains(sum%k)) {
                   return true;
               }  else {
                   set.add(pre);
                   pre = sum % k;
               }
            }
        }
        return false;
    }
}
