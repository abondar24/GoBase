package main
import (
	"net"
	"log"
	"os"
	"io"
)

func main() {
	conn,err :=net.Dial("tcp","localhost:8000")
	if err !=nil{
		log.Fatal(err)
	}
	done :=make(chan struct{})
	go func(){
		io.Copy(os.Stdout,conn)
		done<-struct{}{}
	}()
	mustCop(conn,os.Stdin)
	conn.Close()
	<-done
}

func mustCop(dst io.Writer,src io.Reader){
	if _,err:=io.Copy(dst,src); err !=nil{
		log.Fatal(err)
	}
}