class Solution {
    public int strStr(String haystack, String needle) {
        if (haystack == null || needle == null || needle.length() > haystack.length()) {
            return -1;
        }
        if (needle.equals("")) {
            return 0;
        }
        
        int l1 = haystack.length();
        int l2 = needle.length();
        
        for (int i = 0; i < l1-l2+1; i++) {
            int j = 0;
            while (j < l2) {
                if (haystack.charAt(i+j) != needle.charAt(j)) {
                    break;
                } else {
                    j++;
                }
            }
            if (j == l2) {
                return i;
            }
        }
        return -1;
    }
    
}
