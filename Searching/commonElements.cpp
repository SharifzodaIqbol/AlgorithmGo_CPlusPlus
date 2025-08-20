#include<bits/stdc++.h>
using namespace std;
int main(){
    vector<int> arr1{1, 5, 10, 20, 30};
    vector<int> arr2{5, 13, 15, 20};
    vector<int> arr3{5, 20};
    int i{0}, k{0}, j{0};
    int n1 = arr1.size(), n2 = arr2.size(), n3 = arr3.size();
    vector<int> res;
    while(i < n1 && j < n2 && k < n3){
        if(arr1[i] == arr2[j] && arr2[j] == arr3[k]){
            res.push_back(arr1[i]);
            i++;
            k++;
            j++;
            while(i< n1 && arr1[i] == arr1[i-1]){
                i++;
            }
            while(j < n2 && arr2[j] == arr2[j-1]){
                j++;
            }
            while(k < n3 && arr3[k] == arr3[k-1]){
                k++;
            }
        }else if(arr1[i] < arr2[j]){
            i++;
        }else if(arr2[j] < arr3[k]){
            j++;
        }else if(arr3[k] < arr1[i]){
            k++;
        }
    }
    for(auto x: res){
        cout << x << " ";
    }
}