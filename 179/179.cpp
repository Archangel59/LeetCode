#include<bits/stdc++.h>
using namespace std;

class Solution {
public:
    string largestNumber(vector<int>& nums) {
        string strList[110],res;
        for(int i=0;i<nums.size();i++) strList[i] = to_string(nums[i]);
        sort(strList, strList + nums.size(), cmp);
        bool flag = false;
        for(int i=0;i<nums.size();i++) {
            if (strList[i] == "0" and !flag) {
                continue;
            }
            flag = true;
            res += strList[i];
        }
        if(res.size() == 0) res = "0";
        return res;
    }

    static bool cmp(string a,string b) {
            return a+b > b+a;
    }
};

int main () {
    Solution a;
    vector<int> nums;
    nums.push_back(3);
    nums.push_back(30);
    nums.push_back(34);
    nums.push_back(5);
    nums.push_back(9);
    cout << a.largestNumber(nums) << endl;
}
