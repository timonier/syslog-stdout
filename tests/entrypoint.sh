#!/bin/sh
set -e

tar fxz /dist/syslog-stdout.tar.gz -C /usr/sbin

/usr/sbin/syslog-stdout &
sleep 1

logger "Test 1 syslog-stdout"
logger "Test 2 syslog-stdout"
sleep 1
