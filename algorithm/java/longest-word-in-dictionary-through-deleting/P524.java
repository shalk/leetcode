class Solution {
    
    public String findLongestWord(String s, List<String> d) {
        TreeSet<String> set = new TreeSet<>();
        int maxLen = -1;
        for (String word: d) {
            if (check(s, word)) {
                set.add(word);
                maxLen = Math.max(maxLen, word.length());
            }
        }
        if (maxLen == -1) {
            return "";
        }
        for (Iterator<String> it =  set.iterator(); it.hasNext();  ) {
            String word = it.next();
            if (word.length() != maxLen) {
                it.remove();
            }
        }
        return set.first();
    }
    
    public boolean check(String s, String word) {
        if (word.length() > s.length()) {
            return false;
        }
        int j = 0;
        int i = 0;
        while (i < word.length() && j < s.length()) {
            if (word.charAt(i) == s.charAt(j)) {
                j++;
                i++;
            } else {
                j++;
            }
        }
        if ( i == word.length() ) {
            return true;
        }
        return false;
    }
}
