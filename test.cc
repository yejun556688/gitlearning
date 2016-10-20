#include <iostream>

const int get()
{
    int a =5;
    const int aa = a;
    return aa;
}
int main()
{
    int a ;
    a= 5;
    const int  K = a;
    int mun[get()];
    std::cout << K << " " << mun[0] << std::endl;
}
