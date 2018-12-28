class Solution {
    public boolean exist(char[][] board, String word) {
        if (board == null || board.length == 0 || board[0] == null || board[0].length == 0 ) {
            return false;
        }
        int m = board.length;
        int n = board[0].length;  
        boolean[][] v = new boolean[m][n];
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (check(i, j, board, word, 0, v)) {
                    return  true;
                }
            }
        }
        return false;
    }
    
    public boolean check(int i, int j, char[][]board, String word, int pos, boolean[][] v) {
        if (word.length() == 0) {
            return true;
        }        
        if (pos == word.length()) {
            return true;
        }  
        if (i >= board.length || i < 0) {
            return false;
        }
        if (j >= board[0].length || j < 0) {
            return false;
        }      
        char c = word.charAt(pos);
        if (board[i][j] != c || v[i][j]) {
            return false;
        } else {
            v[i][j] = true;
            if (check(i+1, j, board, word, pos+1, v)) {
                return true;
            }
            
            if (check(i-1, j, board, word, pos+1, v)) {
                return true;
            }
            
            if (check(i, j+1, board, word, pos+1, v)) {
                return true;
            }
            
            if (check(i, j-1, board, word, pos+1, v)) {
                return true;
            }
            v[i][j] = false;
            return false;
        }
    }

}
