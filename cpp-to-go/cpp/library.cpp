#include "library.h"

Library::Library(size_t sz) : v_(sz, 0) {
  std::cout << "Hello " << v_.size() << std::endl;
}

Library::~Library() {
  std::cout << "Goodbye " << v_.size() << std::endl;
}

void Library::foo(double x) {
  for (auto& e : v_) {
    e += x;
  }
}

double Library::get() const {
  if (v_.empty())
    return -1;

  return v_[0];
}
