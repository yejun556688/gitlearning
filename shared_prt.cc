#include <iostream>
#include <memory>

class implementation
{
    public:
        ~implementation(){
            std::cout << "destrorying implementation\n";
        }
        void do_sh(){
            std::cout << "do something\n";
        }
};
std::shared_ptr<implementation> test()
{
    std::shared_ptr<implementation> p1(new implementation());
    std::cout << "now use count :" << p1.use_count() << std::endl;
    std::shared_ptr<implementation> p2 = p1;
    std::cout << "now use count :" << p2.use_count() << std::endl;
    p1.reset();
    std::cout << "p1 after reset, now use count :" << p2.use_count() << std::endl;
    //p2.reset();
    std::cout << "p2 after reset\n";
    return p2;
} 
int main()
{
    std::shared_ptr<implementation> p3 = test();
    std::cout << "in main, now use count :" << p3.use_count() << std::endl;
    std::cout << "main out\n" ;
    return 0;
}
