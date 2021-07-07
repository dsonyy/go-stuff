#ifdef __cplusplus
extern "C" {
#endif

typedef void* CLibrary;
CLibrary CLibrary_ctor(size_t);
void CLibrary_dctor(CLibrary);
void CLibrary_foo(CLibrary, double);
double CLibrary_get(CLibrary);

#ifdef __cplusplus
}
#endif