# GoLangBasics


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
