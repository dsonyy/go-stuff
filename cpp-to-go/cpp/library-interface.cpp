#include "library-interface.h"
#include "library.h"

CLibrary CLibrary_ctor(size_t sz) {
  Library* obj = new Library(sz);
  return (void*)obj;
}

void CLibrary_dctor(CLibrary cobj) {
  Library* obj = (Library*)cobj;
  delete obj;
}

void CLibrary_foo(CLibrary cobj, double x) {
  Library* obj = (Library*)cobj;
  obj->foo(x);
}

double CLibrary_get(CLibrary cobj) {
  Library* obj = (Library*)cobj;
  return obj->get();
}
