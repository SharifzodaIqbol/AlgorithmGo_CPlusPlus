#include<iostream>
using namespace std;
bool divBy13(string str){
    int rem{};
    for(auto ch : str){
        rem = (rem * 10 + (ch - '0')) % 13;
    }
    return rem == 0;
}
int main(){
    string str{"2911285"};
    cout << (divBy13(str) ? "Yes" : "No");
}