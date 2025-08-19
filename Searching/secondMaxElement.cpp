#include<iostream>
#include<vector>
using namespace std;
int getSecondLargest(vector<int> &arr) {
    int mx1 = INT_MIN, mx2 = INT_MIN;
    for (int i = 0; i < arr.size(); ++i){
        if(arr[i] > mx1){
            mx2 = mx1;
            mx1 = arr[i];
        }else if(arr[i] > mx2 && arr[i] != mx1){
            mx2 = arr[i];
        }
    }
    if (mx2 == INT_MIN){
        return -1;
    }
    return mx2;
}
int main(){
    vector<int> arr{4, 10, 4, 7};
    cout << getSecondLargest(arr);
    return 0;
}