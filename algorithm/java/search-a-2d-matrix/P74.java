class Solution {
    public boolean searchMatrix(int[][] matrix, int target) {
        if (matrix == null || matrix.length == 0 || matrix[0] == null ||  matrix[0].length == 0) {
            return false;
        }
        int m = matrix.length;
        int n = matrix[0].length;
        int up = 0;
        int down = m-1;
        while(up <= down) {
            int mid = (up+down)/2;
            if (matrix[mid][0] > target) {
                down = mid - 1;
            } else {
                up = mid + 1;
            }
        }
        // down up
        if (down == -1) {
            return false;
        }
        int left = 0;
        int right = n-1;
        while (left <= right) {
            int mid = (left+right)/2;
            if (matrix[down][mid] == target) {
                return true;
            } else if (matrix[down][mid] > target) {
                right = mid -1;
            } else {
                left = mid +1;
            }
        }
        return false;
    }
}
