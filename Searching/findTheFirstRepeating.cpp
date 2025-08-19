#include<bits/stdc++.h>
using namespace std;
int main(){
    set<int> st;
    vector<int> arr{10, 5, 3, 4, 3, 5, 6};
    int n = arr.size();
    int minElement = INT_MAX;
    for (int i = 0; i < n;++i){
        if(st.find(arr[n-i]) != st.end()){
            minElement = min(minElement, i);
        }
        st.insert(arr[n-i]);
    }
    cout << minElement;
}