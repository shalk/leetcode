cpp 模版

# 区间合并

```cpp
#include <iostream>
#include <vector>
#include <algorithm>
using namespace std;

void merge(vector<pair<int,int>> & a) {
  int st = -2e9; int ed = -2e9;
  sort(a.begin(), a.end());
  vector<pair<int, int>> ret;
  for (auto p: a) {
    if (p.first > ed) {
      if (st != -2e9) ret.push_back({st, ed});
      st = p.first; ed = p.second;
    } else {
      ed = max(ed, p.second);
    }
  }
  if (st != -2e9) ret.push_back({st, ed});
  swap(ret, a);
}

int main(){
  int n ; cin >> n;
  vector<pair<int,int>> a;
  for (int i = 0 ;i < n ;i++) {
    int x,y;cin >> x >> y;
    a.push_back({x,y});
  }
  merge(a);
  printf("%d",a.size());
  return 0;
}
```

# 离散化

```cpp
#include <iostream>
#include <vector>
#include <algorithm>
using namespace std;

int find(vector<int>& a, int x) {
  int l = 0; int r= a.size();
  while(l < r) {
    int mid = (l+r)/2;
    if (a[mid] >= x) r = mid;
    else l = mid + 1;
  }
  return l;
}

int main() {
  int n, m ;
  cin >> n >> m;
  vector<int> x(n), c(n);
  for (int i = 0; i < n; i++) {
    cin >> x[i] >> c[i];
  }
  vector<int> l(m), r(m);
  for (int i = 0; i < m ;i++) cin >> l[i] >> r [i];
  vector<int> a;
  for (auto &v: x) a.push_back(v);
  for (auto &v: l) a.push_back(v);
  for (auto &v: r) a.push_back(v);

  sort(a.begin(), a.end());
  a.erase(unique(a.begin(), a.end()), a.end());
  int k = a.size();
  // 前缀和
  vector<int> s(k+1);
  for (int i = 0; i < n; i++) s[find(a,x[i])+1] += c[i];
  for (int i = 1 ; i < k+1 ; i++) s[i] += s[i-1];
  for (int i = 0 ; i < m ; i++) printf("%d\n", s[find(a, r[i])+1] - s[find(a,l[i])]);
  return 0;
}


```

# 位运算

```cpp
#include <iostream>
using namespace std;
int main() {
  int n; cin >> n;
  while(n--) {
    int x; cin >> x;
    int c = 0;
    while(x != 0) {
      x -= (x & -x);
      c++;
    }
    printf("%d ", c);
  }
  return 0;
}
```


# 快速排序



## 单调栈


单调队列模版

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