// C++ program to print all
// three-primes smaller than
// or equal to n without using
// extra space
#include <bits/stdc++.h>
using namespace std;

void numbersWith3Divisors(int);
bool isPrime(int);

void numbersWith3Divisors(int n)
{

    for (int i = 2; i * i <= n; i++) {

        if (isPrime(i)) {
                cout << i * i << " ";
        }
    }
}

bool isPrime(int n)
{
    for (int i = 2; i * i <= n; i++) {
        if (n % i == 0)
            return false;
    }
    return true;
}

int main()
{
    int n = 122;

    numbersWith3Divisors(n);

    return 0;
}