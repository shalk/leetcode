class Solution {
    public boolean canThreePartsEqualSum(int[] A) {
        if (A == null || A.length == 0) {
            return false;
        }
        int sum = 0;
        for (int num: A) {
            sum+=num;
        }
        if (sum % 3 != 0) {
            return false;
        }
        int sum1 = 0;
        int part = 0;
        for (int i = 0; i < A.length; i++) {
            sum1 += A[i];
            if (sum1 == sum /3) {
                part++;
                sum1 = 0;
            }
            if (part == 2) {
                break;
            }
        }
        return part == 2;
        
    }
}
