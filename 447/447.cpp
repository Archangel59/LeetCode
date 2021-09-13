#include<bits/stdc++.h>
using namespace std;

class Solution {
public:
    int numberOfBoomerangs(vector<vector<int>>& points) {
        int ans = 0;
        for(int i=0;i<points.size();i++) {
            map<int,int> mp;
            map<int,int> ::iterator it;
            for(int j=0;j<points.size();j++) {
                int dis = (points[i][0] - points[j][0]) * (points[i][0] - points[j][0])  + (points[i][1] - points[j][1]) * (points[i][1] - points[j][1]);
                mp[dis]++;
            }
            for(it=mp.begin();it!=mp.end();it++) {
                ans += it->second * (it->second - 1);
            }
        }
        return ans;
    }
};