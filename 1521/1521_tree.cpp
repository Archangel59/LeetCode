#include<iostream>
#include<vector>
using namespace std;

class Solution {
public:
    struct node {
        int l,r,value;
    }tree[400010];

    void build(vector<int>& arr, int l, int r, int k) {
        tree[k].l = l, tree[k].r = r;
        if(l == r) {
            tree[k].value = arr[l-1];
            cout << arr[l-1] << ' ' << l << ' ' << k << endl;
            return ;
        }
        int mid = l + r >> 1;
        build(arr, l, mid, k<<1);
        build(arr, mid+1, r, (k<<1)+1);
        tree[k].value = tree[k<<1].value & tree[(k<<1)+1].value;
    }
    int getRange(int l,int r, int k) {
        if(tree[k].l >= l and tree[k].r <= r) {
            return tree[k].value;
        }
        int ans = (1 << 31) - 1;
        int mid = tree[k].l + tree[k].r >> 1;
        if(l <= mid) ans &= getRange(l, r, k<<1);
        if(mid < r) ans &= getRange(l, r, (k<<1)+1);

        return ans;
    }
    int closestToTarget(vector<int>& arr, int target) {
        build(arr, 1, arr.size(), 1);
        cout << "build success" << endl;
        while(1) {
            int l,r;cin>>l>>r;
            cout << getRange(l, r, 1) << endl;
        }

        return 0;
    }
};



int main () {

    Solution a;
    vector<int> v = {9,12,3,7,15};
    cout << a.closestToTarget(v, 5) << endl;

    return 0;
}
