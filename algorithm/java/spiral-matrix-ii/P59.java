class Solution {
    public int[][] generateMatrix(int n) {
        int[][] ret = new int[n][n];
        if (n == 0) {
            return ret;
        }
        // n >= 2
        // loop by square, every square have 4 side;
        // if for n * n matrix
        // first square length is n
        // second square length is n-2
        // if n - 2 < 0 , there are not square;
        // first square from ret[0][0]
        // i-th square from ret[i-1][i-1]
        // let's simulate it
        for(int i = 0,len = n,val=1; len > 0; i++,len-=2) {
            // if len is 1, there are one point
            if (len == 1) {
                ret[i][i] = val;
                break;
            }
            // from (i, i) to  (i, i+len-2)
            for (int k =0; k <= len-2; k++) {
                ret[i][i+k] = val;
                val++;
            }            
            
            // from (i, i+len-1) to (i+len-2, i+len-1)
            for (int k=0; k <= len-2; k++) {
                ret[i+k][i+len-1]=val;
                val++;
            }
            
            // from (i+len-1, i+len-1)  to (i+len-1, i+1)
            for (int k=0; k <= len-2; k++) {
                ret[i+len-1][i+len-1-k] = val;
                val++;
            }
            // from (i+len-1, i) to (i+1, i)
            for (int k=0; k <= len-2; k++) {
                ret[i+len-1-k][i]=val;
                val++;
            }
        }
        
        
        return ret;
    }
}
