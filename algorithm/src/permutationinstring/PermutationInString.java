package permutationinstring;

// https://leetcode.com/problems/permutation-in-string/description/
// s1的一个排列 是s2的字串；
// 1. 暴力查全排列，查子串
// 2. 把全排列的字母数量进行统计，在s2 建立一个长度等于字串的窗口， 统计数量，移动窗口，一直到相等为止

class Solution {
    public boolean checkInclusion(String s1, String s2) {
        char[] a1 = s1.toCharArray();
        char[] a2 = s2.toCharArray();
        int[] alphabet = getArrayInfo(a1, 0, s1.length());
        int wstart = 0;
        int wend = a1.length;
        int[] current = getArrayInfo(a2, wstart, wend);
        while (wend <= a2.length) {
            if (ArrayEquals(alphabet, current)) {
                return true;
            }

            wstart++;
            wend++;
        }
        return false;
    }

    public boolean ArrayEquals(int[] a, int[] b) {
        if ( a == null  || b == null) {
            return  false;
        }
        if ( a.length != b.length) {
            return  false;
        }
        for (int i = 0 ; i < a.length; i++ ) {
            if (a[i] != b[i]) {
                return false;
            }
        }
        return true;
    }

    public int[] getArrayInfo(char[] a, int start, int end) {
        int[] alpha = new int[26];
        for (int i = start ; i < end ; i++) {
            alpha[a[i] - 'a']++;
        }
        return alpha;
    }

}

class Solution1 {
    public boolean checkInclusion(String s1, String s2) {
        char[] a1 = s1.toCharArray();
        char[] a2 = s2.toCharArray();

        int[] alphabet = getArrayInfo(a1, 0, s1.length());

        int wstart = 0;
        int wend = a1.length ;
        // start from  [ 0  , s2.length )
        while(wend <= a2.length) {
            int[] current = getArrayInfo(a2, wstart, wend);
            if (ArrayEquals(current, alphabet)) {
                return true;
            }
            wstart++;
            wend++;
        }
        return false;
    }

    public int[] getArrayInfo(char[] a, int start, int end) {
        int[] alpha = new int[26];
        for (int i = start ; i < end ; i++) {
            alpha[a[i] - 'a']++;
        }
        return alpha;
    }

    public boolean ArrayEquals(int[] a, int[] b) {
        if ( a == null  || b == null) {
            return  false;
        }
        if ( a.length != b.length) {
            return  false;
        }
        for (int i = 0 ; i < a.length; i++ ) {
            if (a[i] != b[i]) {
                return false;
            }
        }

        return true;

    }

}

public class PermutationInString {
    public static void main(String[] args) {
        Solution1 s = new Solution1();
        System.out.println(s.checkInclusion("hello", "ooolleoooleh"));
    }
}
