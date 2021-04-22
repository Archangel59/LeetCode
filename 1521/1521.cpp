#include<iostream>
#include<set>
#include<vector>
using namespace std;

class Solution {
public:
    int closestToTarget(vector<int>& arr, int target) {
        int ans = abs(arr[0] - target);
        set<int> s;
        s.insert(arr[0]);
        set<int> ::iterator it;
        for(int i=0;i<arr.size();i++) {
            set<int> newS;
            newS.insert(arr[i]);
            ans = min(ans, abs(arr[i] - target));
            for(it=s.begin();it!=s.end();it++) {
                newS.insert(arr[i] & *it);
                ans = min(ans, abs((arr[i] & *it) - target));
            }
            s = newS;
        }
        return ans;
    }
};

int main () {

    Solution a;
    vector<int> v = {1000,1000,1000};
    cout << a.closestToTarget(v, 1) << endl;

    return 0;
}