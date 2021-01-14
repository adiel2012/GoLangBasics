#include <iostream>
#include <../lib/facade/bridge.h>

int main() {
    std::cout << Increment(4) << std::endl;
    auto a = New_A(44);    
    std::cout << GetValue_A(a) << std::endl;
    Release_A(a);

    return 0;
}