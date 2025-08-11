#include<iostream>
#include<string>
using namespace std;

char find_kth_digit(long long a, long long b, int k){
    long long int res = 1;
    for (int i = 0; i < b;i++){
        res *= a;
    }
    string str = to_string(res);
    return str[str.size() - k];
}
int main(){
    long long a = 5, b = 2;
    int k = 2;
    cout << find_kth_digit(5, 2, 2);
    return 0;
}