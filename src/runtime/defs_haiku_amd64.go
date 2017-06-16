// created by cgo -cdefs and then converted to Go
// cgo -cdefs defs_dragonfly.go

package runtime

//import "unsafe"

const (
	_EINTR  = 0x800000000000000F
	_EFAULT = 0x8000000000001300
	_EBUSY  = 0x8000000000000013
	_EAGAIN = 0x8000000000000010

	_PROT_NONE  = 0x0
	_PROT_READ  = 0x1
	_PROT_WRITE = 0x2
	_PROT_EXEC  = 0x4

	_MAP_ANON    = 0x8
	_MAP_PRIVATE = 0x2
	_MAP_FIXED   = 0x4

	_MADV_FREE = 0x5

	_SA_SIGINFO = 0x40
	_SA_RESTART = 0x10
	_SA_ONSTACK = 0x20

	_SI_USER = 0

	_NSIG = 65

	_SIGHUP       = 1
	_SIGINT       = 2
	_SIGQUIT      = 3
	_SIGILL       = 4
	_SIGCHLD      = 5
	_SIGABRT      = 6
	_SIGPIPE      = 7
	_SIGFPE       = 8
	_SIGKILL      = 9
	_SIGSTOP      = 10
	_SIGSEGV      = 11
	_SIGCONT      = 12
	_SIGTSTP      = 13
	_SIGALRM      = 14
	_SIGTERM      = 15
	_SIGTTIN      = 16
	_SIGTTOU      = 17
	_SIGUSR1      = 18
	_SIGUSR2      = 19
	_SIGWINCH     = 20
	_SIGKILLTHR   = 21
	_SIGTRAP      = 22
	_SIGPOLL      = 23
	_SIGPROF      = 24
	_SIGSYS       = 25
	_SIGURG       = 26
	_SIGVTALRM    = 27
	_SIGXCPU      = 28
	_SIGXFSZ      = 29
	_SIGBUS       = 30
	_SIGRESERVED1 = 31
	_SIGRESERVED2 = 32

	_FPE_INTDIV = 20
	_FPE_INTOVF = 21
	_FPE_FLTDIV = 22
	_FPE_FLTOVF = 23
	_FPE_FLTUND = 24
	_FPE_FLTRES = 25
	_FPE_FLTINV = 26
	_FPE_FLTSUB = 27

	_BUS_ADRALN = 40
	_BUS_ADRERR = 41
	_BUS_OBJERR = 42

	_SEGV_MAPERR = 30
	_SEGV_ACCERR = 31

	_ITIMER_REAL    = 1
	_ITIMER_VIRTUAL = 2
	_ITIMER_PROF    = 3
)

type sigset uint64

type stackt struct {
	ss_sp     uintptr
	ss_size   uintptr
	ss_flags  int32
}

type siginfo struct {
	si_signo  int32
	si_code   int32
	si_errno  int32
	si_pid    int32
	si_uid    int32
	si_addr   uintptr
	si_status int32
	si_band   int64
	si_value  uintptr
}

type mcontext struct {
	rax uint64
	rbx uint64
	rcx uint64
	rdx uint64
	rdi uint64
	rsi uint64
	rbp uint64
	r8  uint64
	r9  uint64
	r10 uint64
	r11 uint64
	r12 uint64
	r13 uint64
	r14 uint64
	r15 uint64
	rsp uint64
	rip uint64
	rflags uint64
	fpregs [512]int32
}

type ucontext struct {
	uc_link     *ucontext
	uc_sigmask  sigset
	uc_stack    stackt
	uc_mcontext mcontext
}

type timespec struct {
	tv_sec  int32
	tv_nsec int64
}

//func (ts *timespec) set_sec(x int32) {
//	ts.tv_sec = x
//}

type timeval struct {
	tv_sec  int32
	tv_usec int32
}

func (tv *timeval) set_usec(x int32) {
	tv.tv_usec = x
}

type itimerspec struct {
	it_interval timespec
	it_value    timespec
}

type itimerval struct {
	it_interval timeval
	it_value    timeval
}
