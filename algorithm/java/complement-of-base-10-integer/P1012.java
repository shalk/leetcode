class Solution {
    public int bitwiseComplement(int N) {
        if (N == 0) {
            return 1;
        }
        StringBuffer sb = new StringBuffer();
        while(N > 0) {
            if (N % 2 == 1) {
                sb.append('1');
            } else {
                sb.append('0');
            }
            N = N/2;
        }
        sb.reverse();
        int ret = 0;
        for (char c: sb.toString().toCharArray()) {
            if (c == '0') {
                ret = ret * 2 + 1;
            } else {
                ret = ret * 2 + 0;
            }
        }
        return ret;
    }
}
