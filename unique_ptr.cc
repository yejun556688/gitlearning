// unique_ptr::operator*
#include <iostream>
#include <memory>

class A{
    public:
    A(){
        std::cout << "A" << '\n';
    }
    ~A(){
        std::cout << "~A" << '\n';
    }
    int* buffer;
};



int main () {

  int* num = new int[100];
   for(int i=0; i< 100;i++){
       num[i] = i;
   }
  { 
    auto deleter = [](A*p){
    //delete p;
    if(p) {
	  for(int i=0; i< 90;i++){
           std::cout << p->buffer[i]<< "\n";
	  }
	  delete[] p->buffer;
	  }
    std::cout << "[deleter called]\n";
   };
   std::unique_ptr<A,decltype(deleter)> foo (new A(),deleter);
   
   foo->buffer = num;
  
  }
  for(int i=0; i < 100; i++)
  std::cout << num[i] << " ";
  return 0;
}
