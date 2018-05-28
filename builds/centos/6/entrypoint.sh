#!/usr/bin/env bash

# Generate host keys
ssh-keygen -t dsa -f /etc/ssh/ssh_host_dsa_key -N ''
ssh-keygen -t rsa -f /etc/ssh/ssh_host_rsa_key -N ''

# Start sshd as daemon and log to stderr (-e)
exec /usr/sbin/sshd -e
