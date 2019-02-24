class Solution {
    class Point {
        int x;
        int y;

        @Override
        public boolean equals(Object o) {
            if (this == o) return true;
            if (o == null || getClass() != o.getClass()) return false;

            Point point = (Point) o;

            if (x != point.x) return false;
            return y == point.y;
        }

        @Override
        public int hashCode() {
            int result = x;
            result = 31 * result + y;
            return result;
        }
    }

    public void addMap(HashMap<Integer, Integer> map, int i) {
        int origin = map.getOrDefault(i, 0);
        map.put(i, origin+1);
    }

    public void subMap(HashMap<Integer, Integer> map, int i) {
        if (map.containsKey(i)) {
            int origin = map.get(i);
            if (origin > 0) {
                map.put(i, origin-1);
            }
        }
    }

    public int[] gridIllumination(int N, int[][] lamps, int[][] queries) {
//        boolean[][] lamp = new boolean[N][N];
        HashMap<Point, Boolean> lamp = new HashMap<>();
        // 记录每一行 有多少灯
        HashMap<Integer, Integer> rowCount = new HashMap<>();
        // 记录每一列 有多少灯
        HashMap<Integer, Integer> colCount = new HashMap<>();
        //  [N-1, 0]-[1, 0],[0,0] [0,1]-[0, N-1]  编号规则：(N-x+y)
        HashMap<Integer, Integer> diag1Count = new HashMap<>();
        //  [0, 0]- [0,N-2], [0,N-1], [N-1, 1]-[N-1, N-1]  编号规则：(y+x+1)
        HashMap<Integer, Integer> diag2Count = new HashMap<>();
        int n1 = queries.length;
        int[] ans = new int[n1];
        // light on and update count;
        for (int i = 0; i < lamps.length; i++) {
            int x = lamps[i][0];
            int y = lamps[i][1];
            Point point = new Point();
            point.x = x;
            point.y = y;
//            lamp[x][y] = true;
            lamp.put(point, true);
            addMap(rowCount, x);
            addMap(colCount, y);
            addMap(diag1Count, diag1(x, y, N));
            addMap(diag2Count, diag2(x, y, N));
        }

        for (int i = 0; i < queries.length; i++) {
            int x = queries[i][0];
            int y = queries[i][1];
            if (rowCount.getOrDefault(x, 0) > 0) {
                ans[i] = 1;
            } else if (colCount.getOrDefault(y, 0)> 0) {
                ans[i] = 1;
            } else if (diag1Count.getOrDefault(diag1(x, y, N),0) > 0) {
                ans[i] = 1;
            } else if (diag2Count.getOrDefault(diag2(x, y, N), 0) > 0) {
                ans[i] = 1;
            } else {
                ans[i] = 0;
            }
            // light off x,y
            for (int p = -1; p <= 1; p++) {
                for (int q = -1; q <= 1; q++) {
                    if ((p + x) >= 0 && (p + x) < N && (q + y) >= 0 && (q + y) < N) {
                        Point point = new Point();
                        point.x = p + x;
                        point.y = q + y;
                        if (lamp.containsKey(point)) {
                            lamp.remove(point);
                            lightOff(p + x, q + y, rowCount, colCount, diag1Count, diag2Count, N);
                        }
                    }
                }
            }

        }
        return ans;
    }

    public int diag1(int x1, int y1, int N) {
        if (x1 > y1) {
            x1 = x1 - y1;
            y1 = 0;
        } else {
            y1 = y1 - x1;
            x1 = 0;
        }
        return N - x1 + y1 - 1;
    }

    public int diag2(int x2, int y2, int N) {
        if (x2 > (N - 1 - y2)) {
            x2 = x2 - (N - 1 - y2);
            y2 = N - 1;
        } else {
            y2 = y2 + x2;
            x2 = 0;
        }
        return y2 + x2;
    }

    public void lightOff(int x, int y, HashMap<Integer, Integer> rowCount, HashMap<Integer, Integer> colCount, HashMap<Integer, Integer> diag1Count, HashMap<Integer, Integer> diag2Count, int N) {
        subMap(rowCount, x);
        subMap(colCount, y);
        subMap(diag1Count, diag1(x, y, N));
        subMap(diag2Count, diag2(x, y, N));
    }
}

