/* Definitions for MIPS running Haiku
   Copyright (C) 1998, 1999, 2000, 2001, 2002, 2004
   Free Software Foundation, Inc.

This file is part of GCC.

GCC is free software; you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation; either version 2, or (at your option)
any later version.

GCC is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with GCC; see the file COPYING.  If not, write to
the Free Software Foundation, 59 Temple Place - Suite 330,
Boston, MA 02111-1307, USA.  */


#define TARGET_OS_CPP_BUILTINS()					\
  do									\
    {									\
	builtin_define ("__HAIKU__");					\
	builtin_define ("__MIPS__");					\
	builtin_define ("__STDC_ISO_10646__=201103L");			\
	builtin_assert ("system=haiku");				\
	if (flag_pic)							\
	  {								\
	    builtin_define ("__PIC__");					\
	    builtin_define ("__pic__");					\
	  }								\
    }									\
  while (0)

/* The byte order macros (__MIPSEB__ or __MIPSEL__) come from the MIPS
   back end.  The port for the SGI Indy is big-endian and uses the n32
   ABI.  */

#undef	LINK_SPEC
#define LINK_SPEC "-m elf_mips_haiku %{!r:-shared} %{nostart:-e 0} " \
  "%{shared:-e 0} %{!shared: %{!nostart: -no-undefined}}"

/* Give data symbols a type and a size, as elfos.h does.  mips.h declares
   them without either, and Haiku's runtime_loader looks symbols up by
   type.  */
#undef	ASM_DECLARE_OBJECT_NAME
#define ASM_DECLARE_OBJECT_NAME mips_declare_object_name

/* long double is the same as double, as on Haiku's 32 bit ports.  The
   n32 default is 128 bits.  */
#undef	MIPS_LONG_DOUBLE_TYPE_SIZE
#define MIPS_LONG_DOUBLE_TYPE_SIZE 64

/* Code for Haiku is abicalls code, which cannot use the small data
   section.  */
#undef	MIPS_DEFAULT_GVALUE
#define MIPS_DEFAULT_GVALUE 0
