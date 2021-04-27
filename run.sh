#!/bin/bash
# PID 1 doesn't generate coredumps unless you create a handler: https://medium.com/hackernoon/my-process-became-pid-1-and-now-signals-behave-strangely-b05c52cc551c
GOTRACEBACK=crash ./coredump
