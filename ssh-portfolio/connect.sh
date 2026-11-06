#!/bin/sh
# Open a fresh connection instead of reusing a client ControlMaster socket.
exec ssh -tt -o ControlMaster=no -o ControlPath=none -o ControlPersist=no -p 2222 visitor@127.0.0.1
