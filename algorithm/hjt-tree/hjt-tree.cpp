#include<iostream>
#include<vector>
#include<algorithm>
#include<stdio.h>
using namespace std;

inline void read(int &x){x=0;int f=1;char c=getchar();while(c>'9'||c<'0') {if(c=='-') f=-1;c=getchar();}
    while(c>='0'&&c<='9') {x=x*10+c-'0'; c=getchar();}x*=f;}

const int maxn = 2e5+5;
int a[maxn];
vector<int> v;
int getId(int x) {
    return lower_bound(v.begin(), v.end(), x) - v.begin() + 1;
}

struct node {
    int l,r,sum;
}Tree[maxn * 40];

int cnt,root[maxn];

void insert(int l,int r,int pre, int &now, int p) {
    Tree[++cnt] = Tree[pre];
    now = cnt;
    Tree[now].sum++;
    if(l == r) return ;
    int m = l + r >> 1;
    if(p <= m) insert(l, m, Tree[pre].l, Tree[now].l, p);
    else insert(m+1, r, Tree[pre].r, Tree[now].r, p);
}

int query(int l,int r,int L,int R,int k) {
    if(l == r) return l;
    int m = l + r >> 1;
    int tmp = Tree[Tree[R].l].sum - Tree[Tree[L].l].sum;
    if(k <= tmp) return query(l, m, Tree[L].l, Tree[R].l, k);
    else return query(m+1, r, Tree[L].r, Tree[R].r, k-tmp);
}

int main () {
    int n,m;cin>>n>>m;
    for(int i=1;i<=n;i++) {
        read(a[i]);
        v.push_back(a[i]);
    }
    sort(v.begin(),v.end());
    v.erase(unique(v.begin(), v.end()), v.end());
    for(int i=1;i<=n;i++) {
        insert(1,n,root[i-1],root[i],getId(a[i]));
    }
    while(m--) {
        int l,r,k;read(l),read(r),read(k);
        printf("%d\n",v[query(1, n, root[l-1], root[r], k)-1]);
    }
    return 0;
}