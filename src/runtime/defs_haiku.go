// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

/*
Input to cgo.

GOARCH=amd64 go tool cgo -cdefs defs_dragonfly.go >defs_dragonfly_amd64.h
*/

package runtime

/*
#include <Errors.h>
#include <sys/time.h>
#include <sys/mman.h>
#include <unistd.h>
#include <errno.h>
#include <signal.h>
#include <time.h>
*/
import "C"

const (
	EINTR  = C.EINTR
	EFAULT = C.EFAULT
	EBUSY  = C.EBUSY
	EAGAIN = C.EAGAIN

	PROT_NONE  = C.PROT_NONE
	PROT_READ  = C.PROT_READ
	PROT_WRITE = C.PROT_WRITE
	PROT_EXEC  = C.PROT_EXEC

	MAP_ANON    = C.MAP_ANON
	MAP_PRIVATE = C.MAP_PRIVATE
	MAP_FIXED   = C.MAP_FIXED

	MADV_FREE = C.POSIX_MADV_DONTNEED

	SA_SIGINFO = C.SA_SIGINFO
	SA_RESTART = C.SA_RESTART
	SA_ONSTACK = C.SA_ONSTACK

	SI_USER = C.SI_USER

	NSIG = C.NSIG

	SIGHUP       = C.SIGHUP
	SIGINT       = C.SIGINT
	SIGQUIT      = C.SIGQUIT
	SIGILL       = C.SIGILL
	SIGCHLD      = C.SIGCHLD
	SIGABRT      = C.SIGABRT
	SIGPIPE      = C.SIGPIPE
	SIGFPE       = C.SIGFPE
	SIGKILL      = C.SIGKILL
	SIGSTOP      = C.SIGSTOP
	SIGSEGV      = C.SIGSEGV
	SIGCONT      = C.SIGCONT
	SIGTSTP      = C.SIGTSTP
	SIGALRM      = C.SIGALRM
	SIGTERM      = C.SIGTERM
	SIGTTIN      = C.SIGTTIN
	SIGTTOU      = C.SIGTTOU
	SIGUSR1      = C.SIGUSR1
	SIGUSR2      = C.SIGUSR2
	SIGWINCH     = C.SIGWINCH
	SIGKILLTHR   = C.SIGKILLTHR
	SIGTRAP      = C.SIGTRAP
	SIGPOLL      = C.SIGPOLL
	SIGPROF      = C.SIGPROF
	SIGSYS       = C.SIGSYS
	SIGURG       = C.SIGURG
	SIGVTALRM    = C.SIGVTALRM
	SIGXCPU      = C.SIGXCPU
	SIGXFSZ      = C.SIGXFSZ
	SIGBUS       = C.SIGBUS
	SIGRESERVED1 = C.SIGRESERVED1
	SIGRESERVED2 = C.SIGRESERVED2

	FPE_INTDIV = C.FPE_INTDIV
	FPE_INTOVF = C.FPE_INTOVF
	FPE_FLTDIV = C.FPE_FLTDIV
	FPE_FLTOVF = C.FPE_FLTOVF
	FPE_FLTUND = C.FPE_FLTUND
	FPE_FLTRES = C.FPE_FLTRES
	FPE_FLTINV = C.FPE_FLTINV
	FPE_FLTSUB = C.FPE_FLTSUB

	BUS_ADRALN = C.BUS_ADRALN
	BUS_ADRERR = C.BUS_ADRERR
	BUS_OBJERR = C.BUS_OBJERR

	SEGV_MAPERR = C.SEGV_MAPERR
	SEGV_ACCERR = C.SEGV_ACCERR

	ITIMER_REAL    = C.ITIMER_REAL
	ITIMER_VIRTUAL = C.ITIMER_VIRTUAL
	ITIMER_PROF    = C.ITIMER_PROF
)

type Sigset C.sigset_t
type StackT C.stack_t

type Siginfo C.siginfo_t

type Mcontext C.mcontext_t
type Ucontext C.ucontext_t

type Timespec C.struct_timespec
type Timeval C.struct_timeval
type Itimerspec C.struct_itimerspec
type Itimerval C.struct_itimerval
