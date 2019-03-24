class Solution {
    public int smallestRepunitDivByK(int K) {
        if (K == 1) {
            return 1;
        }
        int lastNum = K % 10;
        if (lastNum != 1 && lastNum != 3 && lastNum != 7 && lastNum != 9) {
            return -1;
        }
        
        
        int[] solution = new int[10];
        for (int i = 0; i < 10; i++) {
            solution[i] = -1;
            for (int j = 0; j < 10; j++) {
                if ((j * lastNum + i  ) % 10 == 1) {
                    solution[i] = j;
                    break;
                }
            }
        }
        int N = 0;
        int carry = 0;
        int count = 0;
        do {
            count++;
            int val = solution[carry];
            if (val == -1) {
                return -1;
            }
            N += val * K;
            N  /= 10;
            carry = N % 10;
        } while (N != 0);
        return count;
    }
}
