!=========================================================================================
! foo_integer implements the subroutine foo with an 0-d integer parameter type

subroutine foo_integer(f, d)
! Constants
  use iso_c_binding

  implicit none

! Parameters
  integer, intent(inout) :: f
  integer(c_size_t) :: d

  write (*,*) 'f=', f, 'd=', d

! Finished
  return

end subroutine foo_integer
