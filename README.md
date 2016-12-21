# README

Minimalistic syslog for Docker containers

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

Run, in background, the binary corresponding to your architecture:

```sh
# For linux i386
syslog-stdout-linux-386 &

# For linux amd64
syslog-stdout-linux-amd64 &

# For linux armv5
syslog-stdout-linux-armv5 &

# For linux armv6
syslog-stdout-linux-armv6 &

# For linux armv7
syslog-stdout-linux-armv7 &

# For linux armv8
syslog-stdout-linux-armv8 &
```

## Contributing

1. Fork it.
2. Create your branch: `git checkout -b my-new-feature`.
3. Commit your changes: `git commit -am 'Add some feature'`.
4. Push to the branch: `git push origin my-new-feature`.
5. Submit a pull request.

__Note__: Use the script `bin/build` to test your modifications locally.

## Credits

The original script (in Python) has been created by [Gryphius](https://github.com/gryphius).

## Links

* [creating statically linked executables in Go](http://blog.xebia.com/2014/07/04/create-the-smallest-possible-docker-container/)
* [cross compiling go](http://stackoverflow.com/questions/27412601/cross-compiling-go)
* [gryphius/syslog-stdout](https://github.com/gryphius/syslog-stdout)
