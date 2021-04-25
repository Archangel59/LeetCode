#include<iostream>
#include<algorithm>
#include<vector>
using namespace std;
class Solution {
public:
    struct node {
        int l,r,value;
    }tree[400010];

    void build(int l,int r,int k) {
        tree[k].l = l;
        tree[k].r = r;
        if(l == r) {
            tree[k].value = 0;
            return ;
        }
        int mid = l + r >> 1;
        build(l,mid, k<<1);
        build(mid+1, r, (k<<1)+1);
        tree[k].value = tree[k<<1].value + tree[(k<<1)+1].value;
    }

    void updateOne(int idx, int k) {
        if(tree[k].l == tree[k].r) {
            tree[k].value++;
            return ;
        }
        int mid = tree[k].l + tree[k].r >> 1;
        if(idx <= mid) updateOne(idx, k<<1);
        else updateOne(idx, (k<<1)+1);
        tree[k].value = tree[k<<1].value + tree[(k<<1)+1].value;
    }

    int getRange(int l, int r, int k) {
        if(tree[k].l >= l and tree[k].r <= r) {
            return tree[k].value;
        }
        int mid = tree[k].l + tree[k].r >> 1;
        int ans = 0;
        if(l <= mid) ans += getRange(l, r, k<<1);
        if(mid < r) ans += getRange(l,r, (k<<1)+1);
        return ans;
    }

    struct arrNode {
        int value,idx;
    };

    static bool cmp(arrNode a, arrNode b) {
        if(a.value == b.value) return a.idx < b.idx;
        return a.value < b.value;
    }

    static bool cmp1(arrNode a, arrNode b) {
        return a.idx < b.idx;
    }

    vector<int> getHashArr(vector<int> &nums) {
        arrNode a[100010];
        for(int i=0;i<nums.size();i++) {
            a[i].idx = i;
            a[i].value = nums[i];
        }
        sort(a,a+nums.size(),cmp);
        int head = a[0].value, idx = 1;
        for(int i=0;i<nums.size();i++) {
            if(!i) {
                a[i].value = idx;
                continue;
            }
            if(a[i].value != head) idx++;
            a[i].value = idx;
        }
        sort(a,a+nums.size(),cmp1);
        vector<int> v;
        for(int i=0;i<nums.size();i++) {
            v.push_back(a[i].value);
        }
        return v;
    }

    vector<int> countSmaller(vector<int>& nums) {
        vector<int> v = getHashArr(nums),res;
        build(1, nums.size(), 1);
        for(int i=nums.size()-1;i>=0;i--) {
            updateOne(v[i], 1);
            if(v[i] == 1) {
                res.push_back(0);
                continue;
            }
            res.push_back(getRange(1, v[i]-1, 1));
        }
        reverse(res.begin(),res.end());

        return res;
    }
};

int main () {
    Solution a;
    vector<int> v = {1,9,7,8,5}, v1;
    v1 = a.countSmaller(v);
    for(int i=0;i<v1.size();i++)
        cout << v1[i] << " ";
    cout << endl;

    return 0;
}
