// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// From FreeBSD's <sys/cpuset.h> <sys/_cpuset.h> and <sys/syscall.h>,
// https://svnweb.freebsd.org/base/head/sys/sys/_cpuset.h?view=log CPU_MAXSIZE(CPU_SETSIZE),
// https://svnweb.freebsd.org/base/head/sys/kern/syscalls.master?view=log syscall entry,
// First SMP awarded FreeBSD/armv6 release 10.1 (https://en.wikipedia.org/wiki/FreeBSD),
// On FreeBSD 10/11/12 _CPU_SETSIZE is 32, on FreeBSD 8/9 _CPU_SETSIZE is 16
const (
	_CPU_LEVEL_WHICH = 3    // Actual mask/id for which.
	_CPU_WHICH_PID   = 2    // Specifies a process ID.
	_CPU_SETSIZE_MAX = 1024 // Max 8192 CPUs supported.
	_CPU_SETSIZE_MIN = 16   // Min CPU setsize.
	_CPU_CURRENT_PID = -1   // Current process ID.
)
