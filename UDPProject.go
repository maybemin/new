package main

import (
	"fmt"
	"net"
	"time"
	"os"
)

func main() {

	portNum := os.Args[1]
	port := ":"+portNum
	protocol :="udp"

	udpAddr, err := net.ResolveUDPAddr(protocol, port)
	if err != nil {
		return
	}
	fmt.Println("protocol : "+protocol + "\nfrom : "+udpAddr.String())


	udpConn, err := net.ListenUDP(protocol, udpAddr)
	udpConn.SetReadBuffer(81920000)
	//프로그램 실행 시간
	pNow := time.Now()
	pTime := pNow.Format("20060102150405")
	afterTime := pNow.Add(60*time.Second).Format("20060102150405")

	//파일경로 받기
	filePath := os.Args[2]
	file, err := os.OpenFile(filePath+pTime+".txt", os.O_CREATE|os.O_WRONLY| os.O_APPEND | os.O_RDWR, 0644)

	buffer := make([]byte, 1024)

	for {
		fNow := time.Now()
		currTime := fNow.Format("20060102150405")

		if afterTime  == currTime {
			defer file.Close()
			file, _ = os.OpenFile(currTime+".txt", os.O_WRONLY | os.O_CREATE | os.O_RDWR, 0644)
			afterTime = fNow.Add(60*time.Second).Format("20060102150405")
		}
		n, addr, err := udpConn.ReadFromUDP(buffer)
		if err != nil {
			return
		}
		str := "[" + addr.IP.String() + "]" + " | " + "[" + pTime + "]" + " | " + "[" + string(buffer[0:n]) + "]" + "\r\n"
		_, err = file.WriteString(str)
	}
}

