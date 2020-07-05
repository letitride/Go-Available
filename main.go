package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func exists(domain string) (bool, error) {
	const whoisServer string = "com.whois-servers.net"
	//whois serverへ接続
	conn, err := net.Dial("tcp", whoisServer+":43")
	if err != nil {
		return false, err
	}
	defer conn.Close()
	//接続ソケットへ書き込み(送信) []byte(string) でbyte型に変換できる
	conn.Write([]byte(domain + "\r\n"))
	//net.Conn はインターフェースio.Readerとして適合できる
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		if strings.Contains(strings.ToLower(scanner.Text()), "no match") {
			return false, nil
		}
	}
	return true, nil
}

var marks = map[bool]string{true: "o", false: "x"}

func main() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		domain := s.Text()
		fmt.Print(domain, " ")
		exist, err := exists(domain)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(marks[!exist])
		time.Sleep(1 * time.Second)
	}
}
