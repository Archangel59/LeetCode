#include<bits/stdc++.h>
using namespace std;

class Solution {
public:
    int fa[100010],fb[100010];
    int findFa(int x) {
        if(x != fa[x]) fa[x] = findFa(fa[x]);
        return fa[x];
    }
    int findFb(int x) {
        if(x != fb[x]) fb[x] = findFb(fb[x]);
        return fb[x];
    }
    int maxNumEdgesToRemove(int n, vector<vector<int> >& edges) {
        int ans = 0;
        for(int i=1;i<=n;i++) fa[i] = i;
        // 先处理公共边
        for(int i=0;i<edges.size();i++)
            if(edges[i][0] == 3) {
                if(findFa(edges[i][1]) != findFa(edges[i][2]))
                    fa[findFa(edges[i][1])] = fa[findFa(edges[i][2])];
                else
                    ans++;
            }
        // 复制一份给B使用
        for(int i=1;i<=n;i++) fb[i] = fa[i];
        // 接着处理A
        for(int i=0;i<edges.size();i++)
            if(edges[i][0] == 1) {
                if(findFa(edges[i][1]) != findFa(edges[i][2]))
                    fa[findFa(edges[i][1])] = fa[findFa(edges[i][2])];
                else
                    ans++;
            }
        // 最后处理B
        for(int i=0;i<edges.size();i++)
            if(edges[i][0] == 2) {
                if(findFb(edges[i][1]) != findFb(edges[i][2]))
                    fb[findFb(edges[i][1])] = fb[findFb(edges[i][2])];
                else
                    ans++;
            }
        // 判断fa和fb是否是完整图
        for(int i=2;i<=n;i++) {
            if (findFa(1) != findFa(i)) return -1;
            if (findFb(1) != findFb(i)) return -1;
        }
        return ans;
    }
};

int main () {
    Solution a;
    int arr[6][3] = {{3,1,2}, {3,2,3}, {1,1,3},{1,2,4},{1,1,2},{2,3,4}};
    vector<vector<int> > edges(6,vector<int>(3));
    for(int i=0;i<6;i++) {
        edges[i].insert(edges[i].begin(), arr[i], arr[i]+3);
    }
    cout << a.maxNumEdgesToRemove(4, edges) << endl;
    return 0;
}
