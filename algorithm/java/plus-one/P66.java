class Solution {
    public int[] plusOne(int[] digits) {
        // Collections.Sort();
        // Sort.Ints();
        int carry = 1;
        for (int i = digits.length - 1; i >= 0; i-- ) {
            digits[i] += carry;
            if (digits[i] >= 10) {
                digits[i] -= 10;
                carry = 1;
            } else {
                carry = 0;
            }
        }
        if (carry == 0) {
            return digits;
        } else {
            int[] ret = new int[digits.length+1];
            ret[0] = carry;
            for (int i=0; i<digits.length;i++) {
                ret[i+1] = digits[i];
            }
            return ret;
        }
    }
}
