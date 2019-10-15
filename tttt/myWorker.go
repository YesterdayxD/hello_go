package tttt

import "log"

func init(){
	log.SetPrefix("Trace: ")
	log.SetFlags(log.Ldate|log.Lmicroseconds|log.Llongfile)
}
func main() {
	log.Println("message ")
}

