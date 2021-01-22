#include "./facade/bridge.h"
#include "classes.h"
//https://renenyffenegger.ch/notes/development/languages/C-C-plus-plus/GCC/create-libraries/index
 int Increment(int n)
 {
     return n + 1;
 }

void* New_A(int value)
{
    return new A(value);
}


void Release_A(void* obj)
{
    reinterpret_cast<A*>(obj)->~A();
}

int GetValue_A(void* obj)
{
    auto a = reinterpret_cast<A*>(obj);
    return a->GetValue();
}


A::A(int value): m_value(value){}

int A::GetValue()const{ return m_value;}

void GenerateFibonacci(callback_fcn cb, void* context)
{
    if(!cb(1, 1, context) || !cb(1, 2, context))
    {
        return;
    }
    int a = 1, b = 2, pos = 3;

    while(cb(b, pos++, context))
    {
        int c = a + b;
        a = b; 
        b = c;
    }

}


 /*
    Compilation

    Linux
        dynamic
            mkdir ./cpp/speedUplib/lib/linux/dynamic
            g++ -c ./cpp/speedUplib/lib/bridge.cpp -o ./cpp/speedUplib/lib/linux/dynamic/bridge.o
            g++ -shared ./cpp/speedUplib/lib/linux/dynamic/bridge.o -o ./cpp/speedUplib/lib/linux/dynamic/linuxDynamic.so -fPIC
            ar rcs ./cpp/speedUplib/lib/linux/dynamic/linuxDynamic.a  ./cpp/speedUplib/lib/linux/dynamic/linuxDynamic.so 
        static 
            mkdir ./cpp/speedUplib/lib/linux/static
            g++ -c ./cpp/speedUplib/lib/bridge.cpp -o ./cpp/speedUplib/lib/linux/static/bridge.o
            ar rcs ./cpp/speedUplib/lib/linux/static/linuxStatic.a  ./cpp/speedUplib/lib/linux/static/bridge.o 
    Windows
        dynamic
            mkdir ./cpp/speedUplib/lib/windows/dynamic
            x86_64-w64-mingw32-g++ -c ./cpp/speedUplib/lib/bridge.cpp  -o ./cpp/speedUplib/lib/windows/dynamic/bridge.o -fPIC
            x86_64-w64-mingw32-g++ -shared  ./cpp/speedUplib/lib/windows/dynamic/bridge.o -o ./cpp/speedUplib/lib/windows/dynamic/windowsDynamic.dll
        static
            mkdir ./cpp/speedUplib/lib/windows/static
            x86_64-w64-mingw32-gcc -c ./cpp/speedUplib/lib/bridge.cpp -o ./cpp/speedUplib/lib/windows/static/bridge.o -fPIC
            x86_64-w64-mingw32-gcc-ar rcs ./cpp/speedUplib/lib/windows/static/windowsStatic.lib   ./cpp/speedUplib/lib/windows/static/bridge.o




    g++ ./cpp/speedUplib/example/main.cpp -I./cpp/speedUplib/lib/ -L./cpp/speedUplib/lib/windows/static -lwindowsStatic

    go build -o ./cpp/speedUpLib/lib/windows/dynamic/main.exe .\main\main.go 
 */