// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// From FreeBSD's <sys/cpuset.h> <sys/_cpuset.h> and <sys/syscall.h>,
// https://svnweb.freebsd.org/base/head/sys/sys/_cpuset.h?view=log for CPU_MAXSIZE,
// https://svnweb.freebsd.org/base/head/sys/kern/syscalls.master?view=log for syscall entry,
// First SMP awarded FreeBSD/armv6 release 10.1 (https://en.wikipedia.org/wiki/FreeBSD),
// Befor Revision 270222(FreeBSD 8/9), _CPU_SETSIZE = 16 and after that(FreeBSD 10/11/12) _CPU_SETSIZE = 32
const (
	_CPU_LEVEL_WHICH = 3    // Actual mask/id for which.
	_CPU_WHICH_PID   = 2    // Specifies a process id.
	_CPU_SETSIZE_MAX = 1024 // Max cpu setsize, 8192 cpus supported in the future
	_CPU_SETSIZE_MIN = 16   // Min cpu setsize
)

// size of pid_t is 32 bit in amd64/386/arm
//go:noescape
func getpid() int32
