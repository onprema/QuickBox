#!/usr/bin/env bash

# Start sshd as daemon and log to stderr (-e)
exec /usr/sbin/sshd -e
