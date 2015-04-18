! Parameters:
!   f_p The pointer to allocate
!   {{for;0;{{n}};1;d%d;, }} The dimension of the array to allocate.

  function {{name}}(f_p, {{for;0;{{n}};1;d%d;, }}) result(rc)
    use iso_c_binding

    implicit none

! Parmeters
    {{type}}, dimension({{for;0;{{n}};1;:;,}}), allocatable :: f_p
    {{for;0;{{n}};1;integer(c_size_t), value :: d%d;\n    }}
! Return value
    integer :: rc

    rc = 0
    allocate(f_p({{for;0;{{n}};1;d%d;, }}), iostat=rc)

    return
  end function {{name}}
