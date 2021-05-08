#include<bits/stdc++.h>
using namespace std;

class Solution {
public:

    int val[20], mid, n;
    vector<int> v;

    bool dfs(int i) {
        if(i == v.size()) {
            return true;
        }
        for(int j=0;j<n;j++) {
            if(val[j] + v[i] <= mid) {
                val[j] += v[i];
                if(!dfs(i+1)) val[j] -= v[i];
                else return true;
            }
            if(!val[j] or val[j] + v[i] == mid) return false;
        }
        return false;
    }

    bool check() {
        // init
        for(int i=0;i<n;i++) val[i] = 0;
        if (dfs(0)) {
            return true;
        }
        return false;
    }

    int minimumTimeRequired(vector<int>& jobs, int k) {
        // 倒序
        sort(jobs.begin(),jobs.end(), greater<int>());
        v = jobs;
        n = k;
        // get l, r
        int l,r = 0;
        for(int i=0;i<jobs.size();i++) r += jobs[i];
        l = r / k;
        while(l<r) {
            mid = l + r >> 1;
            if(check()) r = mid;
            else l = mid + 1;
        }


        return r;
    }
};

int main () {
    Solution a;
    vector<int> v = {250,250,256,251,254,254,251,255,250,252,254,255};
    cout << a.minimumTimeRequired(v, 10) << endl;
    return 0;
}