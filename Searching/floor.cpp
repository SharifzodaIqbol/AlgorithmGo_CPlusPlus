#include<bits/stdc++.h>
using namespace std;
int main(){
    vector<int> arr{1, 2, 8, 10, 10, 12, 19};
    int left = 0, right = arr.size() - 1;
    int res = -1;
    int target = 0;
    while(left <= right){
        int mid = left + (right - left) / 2;
        if(arr[mid] < target){
            left = mid + 1;
            res = left - 1;
        }else{
            right = mid - 1;
        }
    }
    cout << res;
    return 0;
}