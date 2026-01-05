package main 

import (
	"fmt"
)

func main(){
	var ujian   = 90
	var absensi = 90

	var lulusUjian   = ujian > 80
	var lulusAbsensi = absensi > 88

	fmt.Println(lulusUjian)
	fmt.Println(lulusAbsensi)

	var lulus = lulusAbsensi && lulusUjian

	if lulus { 
		fmt.Println("SELAMAT KAMU LULUS")
	}else{
		fmt.Println("JANGAN PATAH SEMANGAT, SILAHKAN MENGULANG TAHUN DEPAN")
	} 


}