! Paramters:
!   i The {{type}} value to output

  subroutine {{name}}(i)
    implicit none

! Parameters
    {{type}}, value :: i

    write (*,*) 'i=', i

    return
  end subroutine {{name}}
