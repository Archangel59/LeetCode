#include<bits/stdc++.h>
using namespace std;

class Solution {
public:
    vector<int> v,ans;
    int idx = 0;
    struct node {
        int l,r,value;
    }tree[30010 * 4];

    void build(int l,int r,int k) {
        tree[k].l = l;
        tree[k].r = r;
        if(l == r) {
            tree[k].value = v[idx++];
            return ;
        }
        int mid = l + r >> 1;
        build(l,mid, k<<1);
        build(mid+1, r, (k<<1)+1);
        tree[k].value = tree[k<<1].value ^ tree[(k<<1)+1].value;
    }
    int getRange(int l, int r, int k) {
        if(tree[k].l >= l and tree[k].r <= r) {
            return tree[k].value;
        }
        int mid = tree[k].l + tree[k].r >> 1;
        int ans = 0;
        if(l <= mid) ans ^= getRange(l, r, k<<1);
        if(mid < r) ans ^= getRange(l,r, (k<<1)+1);
        return ans;
    }
    vector<int> xorQueries(vector<int>& arr, vector<vector<int>>& queries) {
        v = arr;
        build(1, arr.size(), 1);
        for(int i=0;i<queries.size();i++) {
            ans.push_back(getRange(queries[i][0]+1, queries[i][1]+1, 1));
        }
        return ans;
    }
};
int main () {
    Solution a;
    vector<int> v = {4,8,2,10},ans;
    vector<vector<int> > q = {{2,3},{1,3},{0,0},{0,3}};
    ans = a.xorQueries(v, q);
    for(int i=0;i<ans.size();i++) cout << ans[i] << ' ';

    return 0;
}