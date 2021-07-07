#pragma once
#include <iostream>
#include <vector>

class Library {
 public:
  Library(size_t sz);
  ~Library();

  void foo(double x);
  double get() const;

 private:
  std::vector<double> v_;
};