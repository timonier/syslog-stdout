#!/bin/sh
set -e

/usr/local/bin/syslog-stdout &
sleep 1

logger "Test 1 syslog-stdout"
logger "Test 2 syslog-stdout"
sleep 1
