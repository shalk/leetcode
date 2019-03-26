class Solution {
    public List<List<String>> solveNQueens(int n) {
        List<List<String>> ret = new ArrayList<>();
        if (n <= 0) {
            return ret;
        }
        boolean[] row = new boolean[n];
        boolean[] x1 = new boolean[n+n-1];
        boolean[] x2 = new boolean[n+n-1];
        helper(ret, new ArrayList<String>(), row, col, x1, x2);       
        return ret;
    }
    public void helper(List<List<String>> ret, List<String> sub, boolean[] col, boolean[] x1, boolean[] x2) {
        int n = col.length;
        if (sub.size() == n) {
            ret.add(new ArrayList<>(sub));
            return;
        }
        int i = sub.size();
        for (int j = 0 ; j < n; j++) {
            if (col[j]) {
                continue;
            }
            if (x1[i+j]) {
                continue;
            }
            if (x2[n-1-i+j]) {
                continue;
            }
            col[j] = true;
            x1[i+j] = true;
            x2[n-1-i+j] = true;
            sub.add(getQ(n,j));
            helper(ret, sub, row, col, x1, x2);
            col[j] = false;
            x1[i+j] = false;
            x2[n-1-i+j] = false;
            sub.remove(sub.size()-1);
            
        }
    }
    
    private  String getQ(int n , int i) {
        char[] arr = new char[n];
        for (int k = 0 ; k < n; k++) {
            if (k != i) {
                arr[k] = '.';
            } else {
                arr[k] = 'Q';
            }
        }
        return new String(arr);
    }
}
