#include<iostream>
#include<vector>
#include<unordered_map>
using namespace std;

int main(){
    unordered_map<int, int> unMap;
    vector<int> vt{1, 2, 3, 4, 5};
    int k = 3;
    for (int i = 0; i < vt.size(); ++i){
        if(unMap.count(vt[i]) && (i - unMap[vt[i]])<=k){
            cout << "Yes\n";
            return 0;
        }
        unMap[vt[i]] = i;
    }
    cout << "No\n";


    /*for (int i = 0; i < vt.size(); i++){
        unMap[vt[i]].push_back(i);
    }
    for(auto x: unMap){
        if (x.second.size() > 1){
            for (int i = 1; i < x.second.size(); i++){
                if(abs(x.second[i-1] - x.second[i]) <= k){
                    cout << "Yes\n";
                    return 0;
                }
            }
        }
    }
    cout << "No\n";*/
    return 0;
}