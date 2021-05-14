package main

import (
	"fmt"
	"net"
)

func main() {
	art :=
		`

  _________                      ________   _________ 
 /   _____/ ____ _____    ____   \_____  \  \______  \
 \_____  \_/ ___\\__  \  /    \   /  ____/      /    /
 /        \  \___ / __ \|   |  \ /       \     /    / 
/_______  /\___  >____  /___|  / \_______ \   /____/  
        \/     \/     \/     \/          \/           by Abhinivesh

`
	fmt.Println(art)
	var input string
	fmt.Print("Enter IP address or Host name: ")
	fmt.Scanln(&input)
	fmt.Print("\033[H\033[2J")
	fmt.Printf("Open Ports in %s are listed below\n", input)
	for i := 1; 1 <= 1000; i++ {
		go func(j int) {
			address := fmt.Sprintf("%s:%d", input, j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("%d open\n", j)
		}(i)
	}
}
