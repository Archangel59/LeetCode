#include<iostream>
#include<vector>
using namespace std;

class Solution {
public:
    int removeElement(vector<int>& nums, int val) {
        if(nums.empty()) {
            return 0;
        }
        int j = 0;
        for(int i=0;i<nums.size();i++) {
            if(nums[i] - val) {
                nums[j++] = nums[i];
            }
        }
        return j++;
    }
};

int main () {
    Solution a;
    vector<int> v = {0, 1, 2, 2, 3, 0, 4, 2};
    cout << a.removeElement(v, 2) << endl;
}