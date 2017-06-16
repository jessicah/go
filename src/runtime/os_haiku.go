// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import (
	"runtime/internal/atomic"
	"unsafe"
)

type mOS struct {
	waitsemacount uint32
}

const (
	//_ESRCH       = 3
	//_EAGAIN      = 35
	//_EWOULDBLOCK = _EAGAIN
	//_ENOTSUP     = 91

	// From OpenBSD's sys/time.h
	//_CLOCK_REALTIME  = 0
	//_CLOCK_VIRTUAL   = 1
	//_CLOCK_PROF      = 2
	//_CLOCK_MONOTONIC = 3

	_SS_ONSTACK  = 1
	_SS_DISABLE  = 2

	_SIG_BLOCK   = 1
	_SIG_UNBLOCK = 2
	_SIG_SETMASK = 3
)

//go:noescape
func setitimer(mode int32, new, old *itimerval)

//go:noescape
func sigaction(sig uint32, new, old *sigactiont)

//go:noescape
func sigaltstack(new, old *stackt)

//go:noescape
func sigprocmask(how int32, new, old *sigset)

var sigset_all = ^sigset(0)

//go:noescape
func getrlimit(kind int32, limit unsafe.Pointer) int32
func raise(sig uint32)
func raiseproc(sig uint32)

func osyield()

//go:nosplit
func semacreate(mp *m) {
}

//go:nosplit
func semasleep(ns int64) int32 {
	_g_ := getg()

	// Compute sleep deadline.
	/*
	var tsp *timespec
	if ns >= 0 {
		var ts timespec
		var nsec int32
		ns += nanotime()
		//ts.set_sec(timediv(ns, 1000000000, &nsec))
		//ts.set_nsec(nsec)
		tsp = &ts
		throw("semasleep: need set_sec/set_nsec")
	}
	*/

	for {
		v := atomic.Load(&_g_.m.waitsemacount)
		if v > 0 {
			if atomic.Cas(&_g_.m.waitsemacount, v, v-1) {
				return 0 // semaphore acquired
			}
			continue
		}

		// sleep(tsp, 0, unsafe.Pointer(&_g_.m.waitsemacount), nil)
		throw("semasleep: not implemented")
	}
}

//go:nosplit
func semawakeup(mp *m) {
	atomic.Xadd(&mp.waitsemacount, 1)
	// wake(int32(mp.procid), unsafe.Pointer(&mp.waitsemacount))
	throw("semawakeup: not implemented")
}

// May run with m.p==nil, so write barriers are not allowed.
//go:nowritebarrier
func newosproc(mp *m, stk unsafe.Pointer) {
	if false {
		print("newosproc stk=", stk, " m=", mp, " g=", mp.g0, " id=", mp.id, " ostk=", &mp, "\n")
	}

	throw("newosproc: not implemented")
}

func osinit() {
	ncpu = 1
	physPageSize = 4096
}

var urandom_dev = []byte("/dev/urandom\x00")

//go:nosplit
func getRandomData(r []byte) {
	fd := open(&urandom_dev[0], 0 /* O_RDONLY */, 0)
	n := read(fd, unsafe.Pointer(&r[0]), int32(len(r)))
	closefd(fd)
	extendRandom(r, int(n))
}

func goenvs() {
	goenvs_unix()
}

// Called to initialize a new m (including the bootstrap m).
// Called on the parent thread (main thread in case of bootstrap), can allocate memory.
func mpreinit(mp *m) {
	mp.gsignal = malg(32 * 1024)
	mp.gsignal.m = mp
}

// Called to initialize a new m (including the bootstrap m).
// Called on the new thread, cannot allocate memory.
func minit() {
	// m.procid is a uint64, but lwp_start writes an int32. Fix it up.
	_g_ := getg()
	_g_.m.procid = uint64(*(*int32)(unsafe.Pointer(&_g_.m.procid)))

	minitSignals()
}

// Called from dropm to undo the effect of an minit.
//go:nosplit
func unminit() {
	unminitSignals()
}

func memlimit() uintptr {
	/*
		                TODO: Convert to Go when something actually uses the result.

				Rlimit rl;
				extern byte runtime·text[], runtime·end[];
				uintptr used;

				if(runtime·getrlimit(RLIMIT_AS, &rl) != 0)
					return 0;
				if(rl.rlim_cur >= 0x7fffffff)
					return 0;

				// Estimate our VM footprint excluding the heap.
				// Not an exact science: use size of binary plus
				// some room for thread stacks.
				used = runtime·end - runtime·text + (64<<20);
				if(used >= rl.rlim_cur)
					return 0;

				// If there's not at least 16 MB left, we're probably
				// not going to be able to do much. Treat as no limit.
				rl.rlim_cur -= used;
				if(rl.rlim_cur < (16<<20))
					return 0;

				return rl.rlim_cur - used;
	*/
	return 0
}

func sigtramp()

type sigactiont struct {
	sa_sigaction uintptr
	sa_mask      sigset
	sa_flags     int32
	sa_userdata  uintptr
}

//go:nosplit
//go:nowritebarrierrec
func setsig(i uint32, fn uintptr) {
	var sa sigactiont
	sa.sa_flags = _SA_SIGINFO | _SA_ONSTACK | _SA_RESTART
	//sa.sa_mask = sigset_all
	if fn == funcPC(sighandler) {
		fn = funcPC(sigtramp)
	}
	sa.sa_sigaction = fn
	sigaction(i, &sa, nil)
}

//go:nosplit
//go:nowritebarrierrec
func setsigstack(i uint32) {
	throw("setsigstack")
}

//go:nosplit
//go:nowritebarrierrec
func getsig(i uint32) uintptr {
	var sa sigactiont
	sigaction(i, nil, &sa)
	return sa.sa_sigaction
}

// setSignaltstackSP sets the ss_sp field of a stackt.
//go:nosplit
func setSignalstackSP(s *stackt, sp uintptr) {
	s.ss_sp = sp
}

//go:nosplit
//go:nowritebarrierrec
func sigaddset(mask *sigset, i int) {
	//mask.__bits[(i-1)/32] |= 1 << ((uint32(i) - 1) & 31)
}

func sigdelset(mask *sigset, i int) {
	//mask.__bits[(i-1)/32] &^= 1 << ((uint32(i) - 1) & 31)
}

func (c *sigctxt) fixsigcode(sig uint32) {
}
