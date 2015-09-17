# syslog-stdout

Minimalistic syslog which just prints all messages received from `/dev/log` to standard out.

This is useful in Docker containers where you don't want to install a syslog daemon.

## How to use it?

* Add `dist/syslog-stdout` into an executable folder (like `/usr/local/bin`) inside your container.
* Run `syslog-stdout` in background.

## Credits

The original script (in Python) has been created by [Gryphius](https://github.com/gryphius/syslog-stdout).
