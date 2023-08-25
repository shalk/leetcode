cpp 模版


单调栈


单调队列模版

单调递减队
```cpp
deque<int> q;
while(!q.empty() && nums[q.back()]<= nums[i]) q.pop_back();
q.push_back(i);
```

```cpp
int a[N]; int hh = 0, int tt = -1;
while(hh <= tt && nums[a[tt]] <= nums[i]) tt--;
a[++tt]=i;
```

```cpp
class Solution {
public:
    vector<int> maxSlidingWindow(vector<int>& nums, int k) {
        deque<int> q;
        vector<int> ret;
        int n = nums.size();
        for (int i = 0; i < n; i++ ) {
            // 队列是否划出窗口
            if (!q.empty() && q.front() < i - (k-1)) q.pop_front();
            // 入队，弹出非单调的队尾
            while(!q.empty() && nums[q.back()] <= nums[i]) q.pop_back();
            q.push_back(i);
            // 窗口完整
            if (i >= k-1) ret.push_back(nums[q.front()]);
        }
        return ret;
    }
};
```