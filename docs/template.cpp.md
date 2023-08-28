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
  int n, m ;cin >> n >> m;
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

# 差分
0 a[1] .... a[n] 有差分数组 B[1] ... B[n] 

B[i] = A[i] - A[i-1]; 
当[l,r] 加C; B[r+1]-=c; B[l] += c;
A[i] = 求和B[i] (0..i)

```cpp
#include<iostream>
#include<vector>
int main() {
  int n,m;cin >> n >> m;
  // 数组前面补0，方便计算查分
  vector<int> a(n+1);
  for (int i =1; i <= n; i++) cin >> a[i];
  // 查分数组后面多一位
  vector<int> b(n+2);
  for (int i = 1; i <= n ;i++) b[i] = a[i] - a[i-1];
  while(m--) {
    int l,r,c; cin >> l >> r>> c;
    // [l,r] +c 公式
    b[r+1] -=c; b[l]+= c;
  }
  for (int i = 1; i <= n ;i++) {
    a[i] = a[i-1] + b[i];
    printf("%d ", a[i]);
  }
  return 0;
}
```

# 二维差分

```cpp
#include<iostream>
#include<vector>
using namespace std;
//公式
void insert(vector<vector<int> >& a, int x1, int y1, int x2, int y2, int c) {
  a[x1][y1] += c;
  a[x1][y2+1] -=c;
  a[x2+1][y1] -=c;
  a[x2+1][y2+1] +=c;
}

int main() {
  // 输入，注意下标
  int n,m,q;scanf("%d %d %d",&n,&m,&q);
  vector<vector<int>> a(n+1, vector<int>(m+1));
  for (int i = 1; i <= n ;i++)
    for (int j = 1; j <=m; j++)
      scanf("%d",&a[i][j]);  
  // 初始化差分矩阵, 注意下标
  vector<vector<int>> b(n+2, vector<int>(m+2));
  for (int i = 1; i <= n; i++)
    for (int j=1; j <=m; j++)
      insert(b, i,j,i,j,a[i][j]);
  
  // 操作
  while(q--) {
    int x1,y1,x2,y2,c;
    scanf("%d %d %d %d %d\n",&x1,&y1,&x2,&y2,&c);
    insert(b, x1,y1,x2,y2,c);
  }
  // 反求前缀和
  for (int i =1; i <=n ;i++)
    for (int j=1; j <=m; j++)
      a[i][j] = a[i-1][j]+a[i][j-1]-a[i-1][j-1]+b[i][j];
  // 输出
  for (int i = 1; i <=n ; i++) {
    for (int j = 1;j <= m; j++) printf("%d ", a[i][j]);
    printf("\n");
  }
  return 0;
}
```

# 二维前缀和
a[i][j] 前缀和 b[i][j] ,常用于计算子矩阵的和

b[i][j] = a[i][j]+b[i-1][j] + b[i][j-1] - b[i-1][j-1]
求 x1,y1, x2,y2 的和  = b[x2][y2] - b[x2][y1-1] - b[x1-1][y2] + b[x1-1][y1-1];

```
#include<iostream>
#include<vector>
using namespace std;

int main() {
  int n,m,q;scanf("%d%d%d",&n,&m,&q);
  vector<vector<int> > a(n+1,vector<int>(m+1,0));
  vector<vector<int>> b(n+1, vector<int>(m+1,0));
  for (int i = 1; i <= n; i++) {
    for (int j = 1; j <=m;j++) {
      cin >> a[i][j];
      b[i][j] = b[i-1][j] + b[i][j-1] - b[i-1][j-1] + a[i][j];
    }
  }
  while(q--) {
    int x1,y1,x2,y2; scanf("%d%d%d%d\n",&x1,&y1,&x2,&y2);
    printf("%d\n", b[x2][y2] - b[x2][y1-1] - b[x1-1][y2] + b[x1-1][y1-1]);
  }
  return 0;
}
```


# 前缀和
a[0]...a[n] 前缀数组 B[0]....B[n] ,适合求区间和

B[i] = B[i-1] + a[i];
[ l , r ] 求和， B[L] - B[R-1];

```
#include<iostream>
#include<vector>
using namespace std;
int main() {
  int n,m; scanf("%d%d",&n,&m);
  vector<int> a(n+1,0),b(n+1,0);
  for (int i = 1; i <= n ;i++) {
    scanf("%d", &a[i]);
    b[i] = b[i-1] + a[i];
  }
  while(m--) {
    int l,r; scanf("%d %d\n",&l,&r);
    printf("%d\n",b[r]-b[l-1]);
  }
  return 0;
}
```

# 快速排序

## 二分查找

```cpp
//https://www.acwing.com/problem/content/791/
#include<iostream>
#include<vector>
using namespace std;
int n,q;
int main(){
  cin >> n >> q;
  vector<int> a(n);
  for (int i = 0 ;i < n;i++) cin >> a[i];
  while(q--) {
    int x; cin >> x;
    int l = 0, r = n-1;
    while(l < r) {
      int mid = (l+r)/2;
      if (a[mid] >= x) r = mid;
      else l = mid + 1;
    }
    int first = l;
    l =0; r = n-1;
    while(l < r) {
      int mid = (l+r+1)/2;
      if (a[mid] <= x) l = mid;
      else r = mid -1;
    }
    int second = l;
    if (a[first] != x) printf("-1 -1\n");
    else printf("%d %d\n",first, second);
  }
  return 0;
}
```


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
