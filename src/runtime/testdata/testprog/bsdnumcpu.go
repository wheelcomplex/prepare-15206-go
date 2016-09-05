// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build freebsd

package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"syscall"
)

func init() {
	register("BSDNumCPU", BSDNumCPU)
}

func BSDNumCPU() {
	const CHILDENV = "BSD_NUMCPU_CHILD"

	if os.Getenv(CHILDENV) == "YES" {
		print(runtime.NumCPU())
		os.Exit(0)
	}

	// return OK when only one cpu avaible
	cmdncpu := getcpus(-1)
	if cmdncpu <= 1 {
		fmt.Println("OK")
		os.Exit(0)
	}

	// launch limited proc with env
	err := os.Setenv(CHILDENV, "YES")
	if err != nil {
		fmt.Printf("Setenv %s failed: %s\n", CHILDENV, err.Error())
		os.Exit(1)
	}

	// check n-1 cpus
	list := ""
	cmdncpu--
	for n := 0; n < cmdncpu; n++ {
		if list == "" {
			list += strconv.Itoa(n)
		} else {
			list += "," + strconv.Itoa(n)
		}
	}
	checkncpu(list, cmdncpu)

	fmt.Println("OK")
}

// testprog should print runtime.NumCPU() as execpted.
func checkncpu(list string, n int) {
	args := []string{"-l", list}
	args = append(args, os.Args...)
	cmd := exec.Command("cpuset", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("launch cpuset failed: %s\n", err.Error())
		os.Exit(1)
	}
	ret, err := strconv.Atoi(string(output))
	if err != nil {
		fmt.Printf("except NumCPU == %d, got error %s\n", n, err.Error())
		os.Exit(1)
	}
	if ret != n {
		fmt.Printf("NumCPU() test failed, except %d got %d.\n", n, ret)
		os.Exit(1)
	}
}

// get number of cpus avaible for this pid
func getcpus(pid int) int {
	if pid == -1 {
		pid = syscall.Getpid()
	}
	// cpuset -g -p <pid>
	cmd := exec.Command("cpuset", "-g", "-p", strconv.Itoa(pid))
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("launch cpuset failed: %s\n", err.Error())
		os.Exit(1)
	}
	// pid <pid> mask: 0, 1
	pos := bytes.Index(output, []byte(":"))
	if pos == -1 {
		fmt.Printf("unknow cpuset output: %s\n", output)
		os.Exit(1)
	}
	if len(output) <= pos+1 {
		fmt.Printf("invalid cpuset output: %s\n", output)
		os.Exit(1)
	}
	ncpu := len(bytes.Split(output[pos+1:], []byte(",")))
	if ncpu == 0 {
		fmt.Printf("error: cpuset return zero form %s\n", output)
		os.Exit(1)
	}
	return ncpu
}
