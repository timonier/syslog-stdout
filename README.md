# README

⚠️ This project is no longer maintained. ⚠️

## Installation

Add `dist/syslog-stdout` into an executable folder (like `/usr/sbin`) inside your container.

This operation can be done in `Dockerfile`:

```
# Get syslog-stdout via ADD
ADD https://github.com/timonier/syslog-stdout/releases/download/v1.1.1/syslog-stdout.tar.gz /tmp/syslog-stdout.tar.gz
RUN tar fxz /tmp/syslog-stdout.tar.gz -C /usr/sbin

# Or

# Get syslog-stdout via RUN (if curl is installed)
RUN curl -ksL "https://github.com/timonier/syslog-stdout/releases/download/v1.1.1/syslog-stdout.tar.gz" | tar fxz - -C /usr/sbin

# Or

# Get syslog-stdout via RUN (if wget is installed)
RUN wget -q O - --no-check-certificate "https://github.com/timonier/syslog-stdout/releases/download/v1.1.1/syslog-stdout.tar.gz" | tar fxz - -C /usr/sbin
```

## Usage

Run `syslog-stdout` in background.

## Credits

The original script (in Python) has been created by [Gryphius](https://github.com/gryphius).

## Links

* [creating statically linked executables in Go](http://blog.xebia.com/2014/07/04/create-the-smallest-possible-docker-container/)
* [gryphius/syslog-stdout](https://github.com/gryphius/syslog-stdout)
