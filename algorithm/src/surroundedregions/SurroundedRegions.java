package surroundedregions;

// https://leetcode.com/problems/surrounded-regions/description/
// 思路： 靠着四边的O以及连接着的O是可以保留的。
// 1. 先沿着4边找到所有的0,设置成Z。
// 2. 深度优先遍历每个Z，Z的周围有0，就把0设置为Z。
// 3. 最后完全遍历一遍整个数组，非Z的设置为X， Z设置为0。
//  41%

class Solution {
    public void solve(char[][] board) {
        int row = board.length;
        if (row == 0) {
            return;
        }
        int col = board[0].length;

        // first row
        for (int j = 0 ; j < col; j++) {
            if (board[0][j] == 'O') {
                bfs_visit(board, row, col,0, j);
            }
        }

        // last row
        for (int j = 0 ; j < col; j++) {
            if (board[row-1][j] == 'O') {
                bfs_visit(board, row, col,row-1, j);
            }
        }

        // first col
        for (int i = 0 ; i < row; i++) {
            if (board[i][0] == 'O') {
                bfs_visit(board, row, col,i, 0);
            }
        }
        // last col
        for (int i = 0 ; i < row; i++) {
            if (board[i][col-1] == 'O') {
                bfs_visit(board, row, col,i, col-1);
            }
        }
        for (int i = 0 ; i < row; i++) {
            for (int j = 0; j < col; j++) {
                if (board[i][j] == 'z') {
                    board[i][j] = 'O';
                } else {
                    board[i][j] = 'X';
                }
            }
        }

    }
    void bfs_visit(char[][] board, int row, int col, int i, int j) {
        //判断越界
        if (i < 0 || i >= row) {
            return;
        }
        if (j < 0 || j >= col) {
            return;
        }

        //判断是否为o,为o时，继续广度遍历。
        if (board[i][j] != 'O') {
            return;
        }
        board[i][j] = 'z';
        //访问上下左右
        bfs_visit(board, row, col, i-1, j);
        bfs_visit(board, row, col, i+1, j);
        bfs_visit(board, row, col, i, j-1);
        bfs_visit(board, row, col, i, j+1);
    }
}

public class SurroundedRegions {
    public static void main(String[] args) {
    }
}
