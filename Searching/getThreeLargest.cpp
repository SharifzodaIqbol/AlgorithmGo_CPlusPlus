#include<bits/stdc++.h>
using namespace std;
vector<int> getThreeLargest(vector<int> &arr){
    int mx1 = INT_MIN, mx2 = INT_MIN, mx3 = INT_MIN;
    vector<int> res{};
    for (int i = 0; i < arr.size(); ++i){
        if(arr[i] > mx1){
            mx3 = mx2;
            mx2 = mx1;
            mx1 = arr[i];
        }else if(arr[i] > mx2 && arr[i] != mx1){
            mx3 = mx2;
            mx2 = arr[i];
        }else if(arr[i] > mx3 && arr[i]!= mx1 && arr[i]!=mx2){
            mx3 = arr[i];
        }
    }
    if (mx1 == INT_MIN){
        return res;
    }
    res.push_back(mx1);
    if (mx2 == INT_MIN){
        return res;
    }
    res.push_back(mx2);
    if (mx3 == INT_MIN){
        return res;
    }
    res.push_back(mx3);
    return res;
}
int main(){
    vector<int> arr{10, 4, 3, 50, 23, 90, 90};
    auto x = getThreeLargest(arr);
    for(auto c: x){
        cout << c << " ";
    }
    return 0;
}