#include<bits/stdc++.h>
using namespace std;
int main(){
    vector<int> arr{1, 2, 8, 10, 10, 12, 19};
    int left = 0, right = arr.size() - 1;
    int target = 5;
    if (target > arr[right]){
        return -1;
    }
    while(left <= right){
        int mid = left + (right - left) / 2;
        if(arr[mid] >= target){
            right = mid - 1;
        }else{
            left = mid + 1;
        }
    }
    cout << left;
    return 0;
}