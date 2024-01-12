
#include <bits/stdc++.h>

using namespace std;

vector<long long> Calculate(vector<long long> v)
{
    vector<long long> prefixArray(v.size() + 1);

    for (int i = 0; i < v.size(); i++)
    {
        prefixArray[i + 1] += prefixArray[i] + v[i];
    }

    return prefixArray;
}

int main()
{
    long long int n, x;
    cin >> n;
    vector<long long> v(n);
    for (int i = 0; i < n; i++)
    {
        cin >> x;
        v[i] = x;
    }
    vector<long long> arraySum = Calculate(v);
    
    for (int i = 0; i < arraySum.size(); i++)
    {
        cout << arraySum[i] << ' ';
    }
}

