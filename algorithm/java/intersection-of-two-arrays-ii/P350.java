class Solution {
    public int[] intersect(int[] nums1, int[] nums2) {
        if (nums1 == null || nums2 == null) {
            return new int[0];
        }
        int l1 = nums1.length;
        int l2 = nums2.length;
        if (l1 == 0 || l2 == 0) {
            return new int[0];
        }
        //ensure l1 <= l2;
        if (l1 > l2) {
            return intersect(nums2, nums1);
        }
        
        HashMap<Integer,Integer> map1 = new HashMap<>();
        // HashMap<Integer,Integer> map2 = new HashMap<>();
        for (int num: nums1) {
            if (map1.containsKey(num)) {
                map1.put(num, map1.get(num)+1);
            } else {
                map1.put(num, 1);
            }
        }
        ArrayList<Integer> ret = new ArrayList<>();
        for (int num: nums2) {
            if (map1.size() == 0) {
                break;
            }
            if (map1.containsKey(num)) {
                ret.add(num);
                int val = map1.get(num);
                val--;
                if (val > 0) {
                    map1.put(num, val);
                } else {
                    map1.remove(num);
                }
            }
        }
        int[] arr = new int[ret.size()];
        int k = 0;
        for (Integer num: ret) {
            arr[k++] = num;
        }
        return arr;
        
    }
}
