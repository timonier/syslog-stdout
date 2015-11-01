# syslog-stdout

Minimalistic syslog which just prints all messages received from `/dev/log` to standard out.

This is useful in Docker containers where you don't want to install a syslog daemon.

## Installation

Add `dist/syslog-stdout` into an executable folder (like `/usr/local/bin`) inside your container.

This operation can be done in `Dockerfile` via:

```
ADD https://github.com/mauchede/syslog-stdout/raw/master/dist/syslog-stdout /usr/local/bin/syslog-stdout
RUN chmod +x /usr/local/bin/syslog-stdout
```

## Usage

Just run `syslog-stdout` in background.

## Contributing

1. Fork it.
2. Create your branch: `git checkout -b my-new-feature`.
3. Commit your changes: `git commit -am 'Add some feature'`.
4. Push to the branch: `git push origin my-new-feature`.
5. Submit a pull request.

## Credits

The original script (in Python) has been created by [Gryphius](https://github.com/gryphius/syslog-stdout).
