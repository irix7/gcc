// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build irix
// +build irix

package runtime

import (
	"unsafe"
)

// IRIX provides POSIX semaphores and sched_yield, but no futex and no
// sem_timedwait.  IRIX has no auxv; see auxv_none.go.  This file is
// modelled on os_aix.go and os_solaris.go, with the timing primitives
// reduced to what IRIX 6.5 actually exports.

//extern sysconf
func sysconf(int32) _C_long

//extern sysmp
func libc_sysmp(int32) int32

// IRIX <sys/sysmp.h>.  sysconf has no _SC_NPROCESSORS_ONLN on IRIX 6.5,
// so the online processor count comes from sysmp(MP_NPROCS).
const _MP_NPROCS = 1

// IRIX <sys/signal.h> defines NSIG (65) but -fdump-go-spec does not
// emit it, so the signal count is derived from SIGRTMAX.
const _NSIG = _SIGRTMAX + 1

type mOS struct {
	waitsema uintptr // semaphore for parking on locks
}

func getProcID() uint64 {
	return uint64(getpid())
}

//extern malloc
func libc_malloc(uintptr) unsafe.Pointer

//go:noescape
//extern sem_init
func sem_init(sem *semt, pshared int32, value uint32) int32

//go:noescape
//extern sem_wait
func sem_wait(sem *semt) int32

//go:noescape
//extern sem_trywait
func sem_trywait(sem *semt) int32

//go:noescape
//extern sem_post
func sem_post(sem *semt) int32

//go:noescape
//extern usleep
func c_usleep(uint32) int32

//go:nosplit
func semacreate(mp *m) {
	if mp.waitsema != 0 {
		return
	}

	var sem *semt

	// Call libc's malloc rather than malloc. This will
	// allocate space on the C heap. We can't call malloc
	// here because it could cause a deadlock.
	sem = (*semt)(libc_malloc(unsafe.Sizeof(*sem)))
	if sem_init(sem, 0, 0) != 0 {
		throw("sem_init")
	}
	mp.waitsema = uintptr(unsafe.Pointer(sem))
}

//go:nosplit
func semasleep(ns int64) int32 {
	mp := getg().m
	if ns >= 0 {
		// IRIX has no sem_timedwait, so poll sem_trywait against a
		// deadline, sleeping a little between attempts.
		deadline := nanotime() + ns
		for {
			if sem_trywait((*semt)(unsafe.Pointer(mp.waitsema))) == 0 {
				return 0
			}
			err := errno()
			if err != _EAGAIN && err != _EINTR {
				throw("sem_trywait")
			}
			if nanotime() >= deadline {
				return -1
			}
			c_usleep(1000)
		}
	}
	for {
		r1 := sem_wait((*semt)(unsafe.Pointer(mp.waitsema)))
		if r1 == 0 {
			break
		}
		if errno() == _EINTR {
			continue
		}
		throw("sem_wait")
	}
	return 0
}

//go:nosplit
func semawakeup(mp *m) {
	if sem_post((*semt)(unsafe.Pointer(mp.waitsema))) != 0 {
		throw("sem_post")
	}
}

func osinit() {
	ncpu = int32(libc_sysmp(_MP_NPROCS))
	if ncpu < 1 {
		ncpu = 1
	}
	physPageSize = uintptr(sysconf(__SC_PAGE_SIZE))
}

func setProcessCPUProfiler(hz int32) {
	setProcessCPUProfilerTimer(hz)
}

func setThreadCPUProfiler(hz int32) {
	setThreadCPUProfilerHz(hz)
}

//go:nosplit
func validSIGPROF(mp *m, c *sigctxt) bool {
	return true
}
