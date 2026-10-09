// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build irix
// +build irix

package syscall

//sysnb raw_ptrace(request int, pid int, addr uintptr, data uintptr) (err Errno)
//ptrace(request _C_int, pid Pid_t, addr *byte, data *byte) _C_long

// IRIX has no getdirentries; read directory entries with getdents.
//sys	getdents(fd int, buf []byte) (n int, err error)
//getdents(fd _C_int, buf *byte, nbytes _C_int) _C_int

func ReadDirent(fd int, buf []byte) (n int, err error) {
	return getdents(fd, buf)
}
