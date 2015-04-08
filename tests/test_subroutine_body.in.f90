!=========================================================================================
! {{name}} implements the subroutine for an {{n}}-d {{type}} type

subroutine {{name}}(f, d)
! Constants
  use iso_c_binding

  implicit none

! Parameters
  {{type}}, intent(inout) :: f
  integer(c_size_t) :: d

  write (*,*) 'f=', f, 'd=', d

! Finished
  return

end subroutine {{name}}
