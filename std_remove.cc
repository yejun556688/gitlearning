#include <iostream>
#include <algorithm>
bool is_odd(int i)
{
    return ((i%2)==1);
}
int main()
{
    int myints[] = {10,20,30,20,3,40,50,37,56};
    int *pbegin = myints;
    int *pend = myints + sizeof(myints)/sizeof(int);
   // pend = std::remove(pbegin,pend,20);
   //pend = std::remove_if(pbegin,pend,is_odd);
   pend = std::remove_if(pbegin,pend,[](int i){return i%2;});
    for(int *p= pbegin ; p!=pend;++p)
        std::cout << ' ' << *p;
    std::cout << std::endl;
 //   std::cout << *pend <<" " << *(pend +1) << " "<< *(pend + 2) <<std::endl;
    return 0;
}
