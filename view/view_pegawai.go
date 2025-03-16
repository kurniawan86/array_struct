package view

import (
	"array_struct/model"
	"array_struct/node"
	"bufio"
	"fmt"
	"os"
)

func Insert() {
	var kota, notelp, email string
	var id, nomer int

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("INSERT DATA PEGAWAI BARU ")

	fmt.Print("masukkan ID Pegawai : ")
	fmt.Scan(&id)

	fmt.Print("maukkan Alamat Jalan : ")
	jalan, _ := reader.ReadString('\n')
	jalan = jalan[:len(jalan)-1]

	fmt.Print("masukkan kota : ")
	fmt.Scan(&kota)

	fmt.Print("masukkan nomer rumah : ")
	fmt.Scan(&nomer)

	fmt.Print("masukkan notelp : ")
	fmt.Scan(&notelp)

	fmt.Print("masukkan email : ")
	fmt.Scan(&email)

	address := node.Address{
		Jalan: jalan,
		Kota:  kota,
		Nomer: nomer,
	}

	emp := node.Pegawai{
		ID:     id,
		Alamat: address,
		NoTelp: notelp,
		Email:  email,
	}

	model.Create(emp)
	fmt.Println("")
}

func Display() {
	fmt.Println("daftar pegawai perusahaan XXX")
	for i, emp := range model.Read() {
		fmt.Println("pegawai ke-", i+1)
		fmt.Println("ID Pegawai \t: ", emp.ID)

	}
}
