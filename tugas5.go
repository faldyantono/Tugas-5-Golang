
package main

import "fmt"
import "math/rand"

type pelajar struct {
 nama string
 umur int
}

func main(){
var n1,n2,n3 pelajar;
n1.nama = "Ahmad"
n2.nama = "Abi"
min := 0
max :=100
n3.nama = "Zaka"
n1.umur, n2.umur, n3.umur = rand.Intn((max-min)+max), rand.Intn((max-min)+max), rand.Intn((max-min)+max)
fmt.Println(n1.nama, n1.umur)
fmt.Println(n2.nama, n2.umur)
fmt.Println(n3.nama, n3.umur)

}
