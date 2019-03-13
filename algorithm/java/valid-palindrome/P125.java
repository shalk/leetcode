class Solution {
    public boolean isPalindrome(String s) {
        if (s == null) {
            return false;
        }
        int left = 0;
        int right = s.length()-1;
        while(left < right) {
            char ch1 = s.charAt(left);
            char ch2 = s.charAt(right);
            if (!isLetter(ch1)) {
                left++;
                continue;
            }
            if (!isLetter(ch2)) {
                right--;
                continue;
            }
            ch2 = toLow(ch2);
            ch1 = toLow(ch1);
            if (ch1 != ch2) {
                return false;
            } else {
                left++;
                right--;
            }
        }
        return true;
    }
    public boolean isLetter(char c) {
        return (c >= 'a' && c <='z') || (c >='A' && c <= 'Z') || (c >= '0' && c <='9');
    }
    
    public char toLow(char c) {
        if (c >='A' && c <= 'Z') {
            return (char)(c - 'A' + 'a');
        }
        return c;
    }
}
