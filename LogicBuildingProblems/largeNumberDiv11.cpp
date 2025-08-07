#include<iostream>
#include<string>
using namespace std;
bool checkDivisible11(string str){
    int n = str.size();
    int sumOdd{}, sumEvent{};
    for (int i = 0; i < n; ++i){
        if(i & 1){
            sumEvent += str[i] - '0';
        }else{
            sumOdd += str[i] - '0';
        }
    }
    return (sumEvent - sumOdd) % 11 == 0;
}
int main(){
    string s = "76945";
    cout << (checkDivisible11(s) ? "true" : "false");

}