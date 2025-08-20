#include<bits/stdc++.h>
using namespace std;
int main(){
    vector<int> arr{1, 1, 1, 0, 0, 0, 0};
    int left = 0, right = arr.size() - 1;
    int target = 1;
    while(left <= right){
        int mid = left + (right - left) / 2;
        if (arr[mid] == target){
            left = mid + 1;
        }else{
            right = mid - 1;
        }
    }
    cout << left;
}