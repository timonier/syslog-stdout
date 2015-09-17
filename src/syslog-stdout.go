package main

import (
    "fmt"
    "net"
    "log"
    "os"
    "strconv"
)

type Syslog struct {}

func (syslog Syslog) getFacility(code int) string {
    switch code >> 3 {
        case 0: return "kern"
        case 1: return "user"
        case 2: return "mail"
        case 3: return "daemon"
        case 4: return "auth"
        case 5: return "syslog"
        case 6: return "lpr"
        case 7: return "news"
        case 8: return "uucp"
        case 9: return "cron"
        case 10: return "authpriv"
        case 11: return "ftp"
        case 12: return "ntp"
        case 13: return "security"
        case 14: return "console"
        case 15: return "mark"
        case 16: return "local0"
        case 17: return "local1"
        case 18: return "local2"
        case 19: return "local3"
        case 20: return "local4"
        case 21: return "local5"
        case 22: return "local6"
        case 23: return "local7"
        default: return "unknown"
    }
}

func (syslog Syslog) getSeverity(code int) string {
    switch code & 0x07 {
        case 0: return "emerg"
        case 1: return "alert"
        case 2: return "crit"
        case 3: return "err"
        case 4: return "warning"
        case 5: return "notice"
        case 6: return "info"
        case 7: return "debug"
        default: return "unknown"
    }
}

func (syslog Syslog) listen(connection net.Conn) {
    for {
        buffer := make([]byte, 2048)
        nr, error := connection.Read(buffer)
        if error != nil {
            return
        }

        header,message := syslog.readData(buffer[0:nr])
        if "" != header {
            fmt.Printf("%s: %s\n", header, message)
        } else {
            fmt.Println(message)
        }
    }
}

func (syslog Syslog) readData(data []byte) (string, string) {
    header := ""
    message := string(data)

    if ">" == string(data[2]) {
        code, error := strconv.Atoi(string(data[1]))

        header = "unknown:unknown"
        if nil == error {
            header = fmt.Sprintf("%s:%s", syslog.getFacility(code), syslog.getSeverity(code))
        }

        message = string(data[3:])
    } else if ">" == string(data[3]) {
        code, error := strconv.Atoi(string(data[1:3]))

        header = "unknown:unknown"
        if nil == error {
            header = fmt.Sprintf("%s:%s", syslog.getFacility(code), syslog.getSeverity(code))
        }

        message = string(data[4:])
    }

    return header, message
}

func (syslog Syslog) run() {
    if _, err := os.Stat("/dev/log"); nil == err {
        os.Remove("/dev/log")
    }

    listen, error := net.Listen("unix", "/dev/log")
    if error != nil {
        log.Fatal("Listen error: ", error)
        return
    }

    for {
        connection, error := listen.Accept()
        if error != nil {
            log.Fatal("Accept error: ", error)
            return
        }

        go syslog.listen(connection)
    }
}

func main() {
    syslog := Syslog{}
    syslog.run()
}
