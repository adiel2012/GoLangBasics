
#pragma once

#ifdef __cplusplus
extern "C" {
#endif
int Increment(int);
void* New_A(int value);
void Release_A(void* obj);
int GetValue_A(void* obj);
typedef int (*callback_fcn)(int);
void GenerateFibonacci(callback_fcn);

#ifdef __cplusplus
}  // extern "C"
#endif