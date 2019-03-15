class Solution {
    public List<String> generateParenthesis(int n) {
        List<String> ret = new ArrayList<>();
        helper(ret, "", 0, 0, n);
        return ret;
    }
    public void helper(List<String> ret, String s, int open, int close, int max) {
        if (open > max || close > max) {
            return;
        }
        if (open == max && close == max) {
            ret.add(s);
        }
        if (open == close) {
            // chose open
            helper(ret, s+"(", open+1, close, max);
        } else if (open > close) {
            helper(ret, s+"(", open+1, close, max);
            helper(ret, s+")", open, close+1, max);
        } else {
            // open < close
            return;
        }
    }
}
