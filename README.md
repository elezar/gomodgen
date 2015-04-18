# gomodgen
A Fortran module generator written in Go.

# Description
I often have to write generic interface implementations that follow the following pattern:
```
interface to_string
  module procedure to_string_integer
  module procedure to_string_integer_1d
  module procedure to_string_real4
  module procedure to_string_real4_1d
end interface
```
where each of the module procedures is an implementation of the generic interface ```to_string``` for a given type-signature.

For example (for the scalar integer variant):
```
function to_string_integer(i) result(s)
  implicit none
  integer, value :: i
  character(len=:), allocatable :: s

  integer :: required_len

  required_len = get_required_len(i)

  allocate(s(len=required_len))

  write (s, '(I0)') i

  return
end function to_string_integer
```

I also want to learn some Go.


# Copyright

Copyright (c) 2015 Evan Lezar (http://www.evanlezar.com)
