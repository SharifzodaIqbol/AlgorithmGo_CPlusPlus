#include <iostream>
#include<vector>
using namespace std;
int main(){
    vector<int> nums{10, 20, 30, 40, 50};
    for (int i = 0; i < nums.size();i+=2){
        cout << nums[i] << " ";
    }
}
