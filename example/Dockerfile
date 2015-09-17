FROM debian:jessie

ADD https://raw.githubusercontent.com/mauchede/syslog-stdout/master/dist/syslog-stdout /usr/local/bin/syslog-stdout
RUN chmod +x /usr/local/bin/syslog-stdout

COPY ./entrypoint.sh /entrypoint.sh

ENTRYPOINT [ "/entrypoint.sh" ]
