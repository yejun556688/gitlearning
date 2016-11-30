#include <iostream>
#include <math.h>
int To2Power(int val)
{
    if(val < 0)
        return 0;
    --val;
    val |= val >> 1;
    val |= val >> 2;
    val |= val >> 4;
    val |= val >> 8;
    val |= val >> 16;
    return val+1;
}
int main()
{
    //std::cout << To2Power(165) << " " << std::endl;
    int stride = 172;
    int width = stride - 7;
    std::cout << "width " << width << std::endl;
    std::cout << "  " << ceil(width/8.0)*8 << std::endl;
}
