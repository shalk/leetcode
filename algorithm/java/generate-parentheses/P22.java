class Solution {
    public List<String> generateParenthesis(int n) {
        // (f(0)) f(n-1)
        // (f(1))f(n-2)
            // ...
        // ( f(n-1) ) f(0)
        List<List<String>> ans = new ArrayList<>();
        List<String> ret1 = new ArrayList<>();
        ret1.add("");
        ans.add(ret1);
        for (int i = 1 ; i <= n; i++) {
            List<String> ret = new ArrayList<>();
            for (int left = 0; left <= i-1; left++) {
                int right = i-1-left;
                for (String l: ans.get(left)) {
                    for (String r: ans.get(right)) {
                        ret.add("(" + l +")" + r);
                    }
                }
            }
            ans.add(ret);
        }
        return ans.get(n);
    }
}
