FROM debian:jessie

COPY syslog-stdout /usr/local/bin/syslog-stdout
RUN chmod +x /usr/local/bin/syslog-stdout

COPY ./entrypoint.sh /entrypoint.sh

ENTRYPOINT [ "/entrypoint.sh" ]
