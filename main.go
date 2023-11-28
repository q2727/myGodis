package main

import "github.com/qc2727/myGodis/tcp"

func main() {
	tcp.ListenAndServe(":8000")
}
