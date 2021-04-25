#include<iostream>
using namespace std;

int maxn = 100010;

struct node {
    int l,r,value;
}tree[maxn * 4];

void build(int l,int r,int k) {
    tree[k].l = l;
    tree[k].r = r;
    if(l == r) {
        cin>>tree[k].value;
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

int main () {
    int n;cin>>n;
    build(1, n, 1);

    return 0;
}
