#include<bits/stdc++.h>
using namespace std;

class Solution {
public:

    int dir[4][2] = {1,0,-1,0,0,1,0,-1};
    int num,n,m;
    bool vis[110][110];

    void dfs(vector<vector<int>>& heights, int x, int y) {
        if(vis[n-1][m-1]) {
            return ;
        }
        for(int i=0;i<4;i++) {
            int nx = x + dir[i][0];
            int ny = y + dir[i][1];
            if(nx >= 0 and nx < n and ny >= 0 and ny < m and !vis[nx][ny]) {
                if (abs(heights[nx][ny] - heights[x][y]) <= num) {
                    vis[nx][ny] = 1;
                    dfs(heights, nx,ny);
                }
            }
        }
    }

    bool check(vector<vector<int>>& heights) {
        for(int i=0;i<n;i++)
            for(int j=0;j<m;j++)
                vis[i][j] = false;
        vis[0][0] = true;
        dfs(heights, 0,0);
        return vis[n-1][m-1];
    }

    int minimumEffortPath(vector<vector<int>>& heights) {
        n = heights.size(),m = heights[0].size();
        int l = 0 ,r = 0;
        for(int i=0;i<n;i++)
            for(int j=0;j<m;j++)
                r = max(heights[i][j], r);
        while(l<r) {
            num = l + r >> 1;
            if (check(heights)) {
                r = num;
            } else {
                l = num + 1;
            }
        }
        return r;
    }
};

int main () {
    Solution a;
    int arr[3][3] = {{1,2,2}, {3,8,2}, {5,3,5}};
    vector<vector<int> > edges(3,vector<int>(3));
    for(int i=0;i<3;i++) {
        edges[i].insert(edges[i].begin(), arr[i], arr[i]+3);
    }
    cout << a.minimumEffortPath(edges) << endl;
    return 0;
}
