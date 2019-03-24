class Solution {
    public int maxScoreSightseeingPair(int[] A) {
        int n = A.length;
        int[] A2 = new int[n];
        for (int i = 0; i < n; i++) {
            A2[i] = A[i]-i;
        }
        for (int i = n-2; i >= 0; i--) {
            A2[i] = Math.max(A2[i],A2[i+1]);
        }
        
        int score = Integer.MIN_VALUE;
        for (int i = 0; i < A.length -1; i++) {
            score = Math.max(score, A[i]+i+A2[i+1]);
        }
        return score;
    }
}
