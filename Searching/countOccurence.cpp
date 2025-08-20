#include<bits/stdc++.h>
using namespace std;
int main(){
    vector<int> arr{3, 1, 2, 2, 1, 2, 3, 3};
    unordered_map<int, int> myMap;
    vector<int> res;
    int n = arr.size();
    int k = 3;
    k = n / k;
    for (int i = 0; i < n; ++i){
        myMap[arr[i]]++;
    }
    for(auto x: myMap){
        if(x.second > k){
            cout << x.first;
        }
    }
}