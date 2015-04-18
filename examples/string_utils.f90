module string_utils

  implicit none

  private

  public :: output

  ! Interface description
  interface output
    module procedure output_integer
    module procedure output_real4
  end interface

contains

!=========================================================================================
! Description for output_integer
  subroutine output_integer(i)
    implicit none

! Paramters
    integer, value :: i

    write (*,*) 'i=', i

    return
  end subroutine output_integer


!=========================================================================================
! Description for output_real4
  subroutine output_real4(i)
    implicit none

! Paramters
    real(4), value:: i

    write (*,*) 'i=', i

    return
  end subroutine output_real4

end module string_utils
