
class Solution {
    public boolean isMatch(String s, String p) {
        if (s == null || p == null) {
            return false;
        }
        return helper(s.toCharArray(), s.length()-1, p.toCharArray(), p.length()-1);
    }
    public boolean helper(char[] s, int end1, char[] p, int end2) {
        if (end2 == -1) {
            if (end1 == -1) {
                return true;
            } else {
                return false;
            }
        }
        if (end1 == -1) {// p[0..end2] should like a*b*c*
            if ((end2+1) % 2 == 1) {
                return false;
            }
            for (int i = 1 ; i <= end2; i+=2) {
                if (p[i] != '*') {
                    return false;
                }
            }
            return true;
        }
        if (p[end2] == '.') {
            return helper(s, end1-1, p, end2-1);
        } else if (p[end2] >= 'a' && p[end2] <= 'z' ) {
            if (p[end2] != s[end1]) {
                return false;
            } else {
                return  helper(s, end1-1, p, end2-1);
            }
        } else {
            // p[end2] == '*'
            if (end2 == 0) {
                return false;
            }
            if (p[end2-1] == '*') {
                return false;
            } else if (p[end2-1] >= 'a' && p[end2-1] <= 'z') {
                int count = 0;
                for (int i = end1; i >= 0;i--) {
                    if (s[i] == p[end2-1]) {
                        count++;
                    } else {
                        break;
                    }
                }
                for (int i = count; i >= 0; i--) {
                    if (helper(s, end1-i, p, end2-2)) {
                        return true;
                    }
                }
                return false;
            } else {
                for (int i = end1+1; i >= 0; i--) {
                    if (helper(s, end1-i, p, end2-2)) {
                        return true;
                    }
                }
                return false;
            }

        }

    }
}
