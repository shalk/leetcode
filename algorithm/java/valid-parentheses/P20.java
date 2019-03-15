class Solution {
    public boolean isValid(String s) {
        if (s == null) {
            return false;
        }
        int n = s.length();
        if (n==0) {
            return true;
        }
        char[] stack = new char[n];
        int top = -1;
        for (char ch: s.toCharArray()) {
            if (ch == '(' || ch == '[' || ch == '{') {
                stack[++top] = ch;
            } else if (ch == ')') {
                if (top == -1 || stack[top] != '(') {
                    return false;
                }
                top--;
            } else if (ch == ']') {
                if (top == -1 || stack[top] != '[') {
                    return false;
                }
                top--;
            } else if (ch == '}') {
                if (top == -1 || stack[top] != '{') {
                    return false;
                }
                top--;
            } else {
                
            }
        }
        if (top == -1) {
            return true;
        } else {
            return false;
        }
    }
}
