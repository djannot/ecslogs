package main

import (
  "fmt"
  "golang.org/x/crypto/ssh"
  "log"
  "os"
  "strings"
)

var f *os.File

func main() {
  if len(os.Args) != 7 && len(os.Args) != 8 {
                log.Fatalf("Usage: %s <user> <host:port> <pattern> <input log> <# days> <ip1,ip2,ip3|dynamic> <pipe|file> <output file>", os.Args[0])
        }

  var err error
  if os.Args[7] == "file" {
    f, err = os.Create(os.Args[8])
    if err != nil {
      panic(err)
    }

    defer f.Close()
  }

  client, session, err := connectToHost(os.Args[1], os.Args[2])
  if err != nil {
          panic(err)
  }

  ipAddresses := []string{}
  if os.Args[6] == "dynamic" {
    outIp, err := session.CombinedOutput(`sudo getrackinfo -a | egrep "MA|SA" | awk '{ print $5 }'`)
    if err != nil {
            panic(err)
    }
    ipAddresses = strings.Split(string(outIp),"\n")
  } else {
    ipAddresses = strings.Split(os.Args[6],",")
  }

  for _, ip := range ipAddresses {
    if ip != "" {
      session, err := client.NewSession()
      if err != nil {
                panic(err)
        }
      outFiles, err := session.CombinedOutput(`ssh ` + ip + ` 'sudo find /opt/emc/caspian/fabric/agent/services/object/main/log/ -mtime -` + os.Args[5] + ` | grep ` + os.Args[4] + `'`)
        if err != nil {
                panic(err)
        }
      for _, log := range strings.Split(string(outFiles),"\n") {
        if log != "" {
          session, err := client.NewSession()
          if err != nil {
                        panic(err)
                }
          outGrep, _ := session.CombinedOutput(`ssh ` + ip + ` 'sudo zgrep "` + os.Args[3] + `" ` + log + `' | awk '{ print "` + ip + ` ` + log + ` "$0 }'`)
          if string(outGrep) != "" {
            if os.Args[7] == "pipe" {
              fmt.Println(string(outGrep))
            } else {
              if _, err = f.Write(outGrep); err != nil {
                fmt.Println(err)
              }
            }
          }
        }
      }
    }
  }
  client.Close()
}

func getPass(user, instruction string, questions []string, echos []bool) (answers []string, err error) {
  if len(questions) > 0 {
    fmt.Println(questions[0])
        var pass string
        fmt.Scanf("%s\n", &pass)
    return []string{pass}, nil
  } else {
    return []string{}, nil
  }
}

func connectToHost(user, host string) (*ssh.Client, *ssh.Session, error) {

        sshConfig := &ssh.ClientConfig{
                User: user,
                //Auth: []ssh.AuthMethod{ssh.Password(pass)},
      Auth: []ssh.AuthMethod{
        ssh.KeyboardInteractive(getPass),
      },
        }

        client, err := ssh.Dial("tcp", host, sshConfig)
        if err != nil {
                return nil, nil, err
        }

        session, err := client.NewSession()
        if err != nil {
                client.Close()
                return nil, nil, err
        }

        return client, session, nil
}
