class Solution {
    public List<List<Integer>> permute(int[] nums) {
        List<List<Integer>> ret = new ArrayList<>();
        Deque<Integer> candi = new LinkedList<>();
        for (int num : nums) {
            candi.add(num);
        }
        helper(ret, new ArrayList<>(), candi);
        return ret;
    }
    public void helper(List<List<Integer>> ret, List<Integer> sub, Deque<Integer> candi) {
        if (candi.size() == 0) {
            ret.add(new ArrayList<>(sub));
            return;
        }
        
        int n = candi.size();
        for (int i = 0; i < n; i++) {
            int num = candi.poll();
            sub.add(num);
            helper(ret, sub, candi);
            sub.remove(sub.size()-1);
            candi.offer(num);
        }
        
    }
  
}
