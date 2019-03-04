class Solution {
    public int nthSuperUglyNumber(int n, int[] primes) {
        if (n == 1) {
            return 1;
        }
        int[] index = new int[primes.length];
        int[] nums = new int[n];
        int[] val = Arrays.copyOf(primes, primes.length);
        
        nums[0] = 1;
        // find every super ugly number
        for (int i=1; i< n;i++) {
            int ugly = Integer.MAX_VALUE;
            for (int k = 0; k < primes.length; k++) {
                if (val[k] == nums[i-1]) {
                    index[k]++;
                    val[k] = nums[index[k]] * primes[k];
                }
                ugly = Math.min(ugly, val[k]);
            }
            nums[i] = ugly;
        }
        return nums[n-1];
    }
}
