#include <iostream>
#include <algorithm>
#include <vector>
using namespace std;
int main()
{
    int n1 = 3;
    int n2 = 5;
    vector<int> v = {0,2,3,4,5,7,6,8};
    auto result1 = find(v.begin(),v.end(),n1);
    auto result2 = find(v.begin(),v.end(),n2);
    if(result1 != v.end())
        cout << "find " << n1 <<" in vector position " << result1.base() << endl;
    if(result2 != v.end())
        cout << "find " << n2 <<" in vector position " << result2.base() << endl;

    
}
