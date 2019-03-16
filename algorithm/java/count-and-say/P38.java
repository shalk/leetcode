class Solution {
    public String countAndSay(int n) {
        String s = "1";
        StringBuffer sb = new StringBuffer();
        while(--n > 0) {
            sb.setLength(0);
            for (int i = 0; i < s.length(); i++) {
                char ch = s.charAt(i);
                int count =  1;
                while(i+1 < s.length() && s.charAt(i+1) == ch) {
                    count++;
                    i++;
                }
                sb.append(count).append(ch);
            }
            s = sb.toString();
        }
        return s;
    }
}
