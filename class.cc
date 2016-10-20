#include"class.h"
#include<istream>
#include<stdio.h>
void A::set(int a)
{
   // std::cout << a << std::endl;
   printf("%d\n",a);
}

int main()
{
    A ss;
    ss.set();
}
