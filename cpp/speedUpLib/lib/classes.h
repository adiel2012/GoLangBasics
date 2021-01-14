
#pragma once

class A{
public:
    A(int value);
    ~A(){};
    int GetValue()const;
private:
    int m_value;
};