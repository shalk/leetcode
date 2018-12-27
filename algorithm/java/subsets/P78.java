class Solution {
    public List<List<Integer>> subsets(int[] nums) {
        List<List<Integer>> ret = new ArrayList<>();
        helper(ret, nums, new ArrayList<Integer>(), 0);
        return ret;
    }
    
    public void helper(List<List<Integer>> ret, int[] nums, List<Integer> sub, int pos) {
        if (pos == nums.length) {
            ret.add(new ArrayList<Integer>(sub));
            return;
        }
        //
        helper(ret, nums, sub, pos+1);
        
        sub.add(nums[pos]);
        helper(ret, nums, sub, pos+1);
        sub.remove(sub.size()-1);
    }
}
