// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build freebsd

package runtime_test

import (
	"os/exec"
	"testing"
)

func TestBSDNumCPU(t *testing.T) {
	_, err := exec.LookPath("cpuset")
	if err != nil {
		// can not test without cpuset command
		t.SkipNow()
	}
	got := runTestProg(t, "testprog", "BSDNumCPU")
	want := "OK\n"
	if got != want {
		t.Fatalf("expected %q, but got:\n%s", want, got)
	}
}
