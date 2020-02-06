# sanddorn/checkcapabilities

![make-all](https://github.com/sanddorn/checkcapabilities/workflows/make%20all/badge.svg)
![make-image](https://github.com/sanddorn/checkcapabilities/workflows/make%20image/badge.svg)
[![GoDoc](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=for-the-badge)](https://godoc.org/github.com/genuinetools/amicontained)
[![Github All Releases](https://img.shields.io/github/downloads/sanddorn/checkcapabilities/total.svg?style=for-the-badge)](https://github.com/sanddorn/checkcapabilities/releases)

Tool to check if only certain capabilities (allowed capabilities) are available to a process. 
You can use it to check if the Pod Security Policy (PSP) works on your Kubernetes cluster. 

The tool is forked from [genuinetools/amicontained](https://github.com/genuinetools/amicontained) and adapted to be used as an automated check for linux kernel capabilities.
The features of the original tool are still available, as so, the documentation will only add the new features.

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [Installation](#installation)
    - [Binaries](#binaries)
    - [Via Go](#via-go)
- [Usage](#usage)
- [Examples](#examples)
    - [docker](#docker)
    - [lxc](#lxc)
    - [systemd-nspawn](#systemd-nspawn)
    - [rkt](#rkt)
    - [unshare](#unshare)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Installation

#### Binaries

For installation instructions from binaries please visit the [Releases Page](https://github.com/sanddorn/checkcapabilities/releases).

#### Via Go

```bash
$ go get github.com/sanddorn/checkcapabilities
```

## Usage

```console
$ checkcapabilities -h
checkcapabilities -  Check for capabilities.

Usage: checkcapabilities <command>

Flags:

  -d       enable debug logging (default: false)
  --pcaps  Check for capabilities allowed. (default: All)

Commands:

  version  Show the version information.

```

## Examples

Only examples for docker are available. Others to follow.

#### docker

```console
$ docker run --rm -it sanddorn/checkcapabilities --pcaps mknod
  Container Runtime: docker
  Has Namespaces:
  	pid: true
  	user: false
  AppArmor Profile: docker-default (enforce)
  Capabilities:
  	BOUNDING -> chown dac_override fowner fsetid kill setgid setuid setpcap net_bind_service net_raw sys_chroot mknod audit_write setfcap
  Seccomp: filtering
  Blocked Syscalls (64):
  	MSGRCV SYSLOG SETPGID SETSID USELIB USTAT SYSFS VHANGUP PIVOT_ROOT _SYSCTL ACCT SETTIMEOFDAY MOUNT UMOUNT2 SWAPON SWAPOFF REBOOT SETHOSTNAME SETDOMAINNAME IOPL IOPERM CREATE_MODULE INIT_MODULE DELETE_MODULE GET_KERNEL_SYMS QUERY_MODULE QUOTACTL NFSSERVCTL GETPMSG PUTPMSG AFS_SYSCALL TUXCALL SECURITY LOOKUP_DCOOKIE CLOCK_SETTIME VSERVER MBIND SET_MEMPOLICY GET_MEMPOLICY KEXEC_LOAD ADD_KEY REQUEST_KEY KEYCTL MIGRATE_PAGES UNSHARE MOVE_PAGES PERF_EVENT_OPEN FANOTIFY_INIT NAME_TO_HANDLE_AT OPEN_BY_HANDLE_AT CLOCK_ADJTIME SETNS PROCESS_VM_READV PROCESS_VM_WRITEV KCMP FINIT_MODULE KEXEC_FILE_LOAD BPF USERFAULTFD MEMBARRIER PKEY_MPROTECT PKEY_ALLOC PKEY_FREE RSEQ
  Looking for Docker.sock
  capability chown is not allowed.
  capability dac_override is not allowed.
  capability fowner is not allowed.
  capability fsetid is not allowed.
  capability kill is not allowed.
  capability setgid is not allowed.
  capability setuid is not allowed.
  capability setpcap is not allowed.
  capability net_bind_service is not allowed.
  capability net_raw is not allowed.
  capability sys_chroot is not allowed.
  capability audit_write is not allowed.
  capability setfcap is not allowed.
  
$ echo $?
  1
```

You can see, that the exit code is non-zero if capabilites were granted, which should not be granted.

#### lxc

TODO

#### systemd-nspawn

TODO

#### rkt

TODO

#### unshare

TODO