package twosum;

import java.util.Arrays;
import java.util.HashMap;


class Solution {


    public int[] twoSum(int[] nums, int target) {

        HashMap<Integer, Integer> map = new HashMap<>();

        for (int i = 0; i < nums.length; i++) {
            map.put(nums[i], i);
        }

        for (int i = 0; i < nums.length; i++) {
            int other = target - nums[i];
            if (map.containsKey(other) && i != map.get(other)) {
                return new int[]{i, map.get(other)};
            }
        }
        return null;
    }
}

public class TwoSum {
    public static void main(String[] args) {
        Solution s = new Solution();
        int[] result = s.twoSum(new int[]{3, 2, 4}, 6);
        System.out.println(Arrays.toString(result));
    }
}
