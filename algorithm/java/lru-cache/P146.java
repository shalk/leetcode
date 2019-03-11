class LRUCache {

    private Deque<Integer> queue;
    private HashMap<Integer, Integer> map;
    private int capacity;
    public LRUCache(int capacity) {
        queue = new ArrayDeque<>();
        map = new HashMap<>(); 
        this.capacity = capacity;
    }
    
    public int get(int key) {
        if (map.containsKey(key)) {
            queue.remove(Integer.valueOf(key));
            queue.addLast(Integer.valueOf(key));
            return map.get(key);
        } else {
            return -1;
        }
    }
    
    public void put(int key, int value) {
        if (map.containsKey(key)) {
            queue.remove(Integer.valueOf(key));
            queue.addLast(Integer.valueOf(key));
            map.put(key, value);
        } else {
            if (queue.size() == capacity)  {
                int old = queue.poll();
                map.remove(old);
            }
            queue.addLast(Integer.valueOf(key));
            map.put(key, value);
        }
    }
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * LRUCache obj = new LRUCache(capacity);
 * int param_1 = obj.get(key);
 * obj.put(key,value);
 */
