#include<iostream>
#include<vector>

using namespace std;
int main(){
    vector<int> arr{1,2,2,1};
    // 67 41 
    for (int i = 1; i < arr.size(); i++){
        if ((i + 1) % 2!=0){
            if (arr[i]>arr[i-1]){
                swap(arr[i], arr[i - 1]);
            }
        }else{
            if (arr[i]<arr[i-1]){
                swap(arr[i], arr[i - 1]);
            }
        }
        
    }
    for(auto x: arr){
        cout << x << " ";
    }
    return 0;
}