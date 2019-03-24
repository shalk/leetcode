class Solution {
    public boolean queryString(String S, int N) {
        if (N == 1) {
            return S.contains("1");
        }
        if (N == 2) {
            return S.contains("10") && S.contains("1");
        }
        int lenN = 0;
        int N1 = N;
        while(N1 > 0) {
            lenN++;
            N1 = N1/2;
        }
        HashSet<Integer> mark = new HashSet<>();
        for (int i = 0; i < S.length(); i++) {
            if (S.charAt(i) == '0') {
                continue;
            }
            int val = 0;
            for (int j = i; j < i+lenN && j < S.length(); j++) {
                val = val * 2 + (S.charAt(j) - '0');
                if (val > 0 && val <= N) {
                    mark.add(val);
                }
            }
            
        }
        return mark.size() == N;
    }    
}
