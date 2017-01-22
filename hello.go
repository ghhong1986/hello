package hello

import "fmt"
import "github.com/ghhong1986/stringutil"

func main() {
	pharse := "Note that you can run this command from anywhere"
	fmt.Println(pharse)
	fmt.Println(stringutil.Reverse(pharse))
}
