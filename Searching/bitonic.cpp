#include<bits/stdc++.h>
using namespace std;
int main(){
    vector<int> arr{120, 100, 80, 20, 0};
    int left = 0, right = arr.size() - 1;
    while(left <= right){
        int mid = left + (right - left) / 2;
        if(arr[mid] < arr[mid + 1] && mid + 1 < arr.size()){
            left = mid + 1;
        }else{
            right = mid - 1;
        }
    }
    cout << arr[left];
}