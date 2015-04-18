!=========================================================================================
! {{name}} implements {{basename}} for a {{n}}-d {{type}} parameter
! Paramters:
!   i The {{type}} value to output

  subroutine {{name}}(i)
    implicit none

! Paramters
    {{type}}, value :: i

    write (*,*) 'i=', i

    return
  end subroutine {{name}}
