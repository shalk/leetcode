class Solution {
    public void setZeroes(int[][] matrix) {
        int m = matrix.length;
        int n = matrix[0].length;
        // -1 is up to down visited
        // -2 is left to right visited
        // -3 is both
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (matrix[i][j] == 0) {
                    // try up to down;
                    if (matrix[0][j] != -10000 && matrix[0][j] != -30000) {
                        // -10000 => keep
                        // 0 => keep
                        // -30000 => keep
                        // -20000 => -30000
                        // other => -10000
                        for ( int k = 0; k < m; k++) {
                            if (matrix[k][j]  == -10000 || matrix[k][j] == 0) {
                            } else if (matrix[k][j] == -20000) {
                                matrix[k][j] = -30000;
                            } else {
                                matrix[k][j] = -10000;
                            }
                        }
                    }
                    // try left to right;
                    if (matrix[i][0] != -20000 && matrix[i][0] != -30000) {
                        for (int k = 0; k < n; k++) {
                            if (  matrix[i][k] == -20000 || matrix[i][k] == 0) {
                            } else if (matrix[i][k] == -10000) {
                                matrix[i][k] = -30000;
                            } else {
                                matrix[i][k] = -20000;
                            }
                        }
                    }
                    matrix[i][j] = -30000;
                }
            }
        }
        
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (matrix[i][j] == -10000 || matrix[i][j] == -20000 || matrix[i][j] == -30000) {
                    matrix[i][j] = 0;
                }
            }
        }
    }
}
