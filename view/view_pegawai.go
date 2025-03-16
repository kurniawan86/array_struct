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
		fmt.Println("Alamat Pegawai \t: ", emp.Alamat.Jalan)
		fmt.Println("Nomer Rumah \t: ", emp.Alamat.Nomer)
		fmt.Println("Kota Pegawai \t: ", emp.Alamat.Kota)
		fmt.Println("Telp Pegawai \t: ", emp.NoTelp)
		fmt.Println("Email Pegawai \t: ", emp.Email)
		fmt.Println()
	}
}

func Update() {
	var kota, notelp, email string
	var id, nomer int

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("UPDATE DATA PEGAWAI ")

	fmt.Print("Masukkan ID Pegawai yang akan diupdate: ")
	fmt.Scan(&id)

	// Cek apakah ID ada dalam data
	emp, found := model.GetByID(id)
	if !found {
		fmt.Println("Pegawai dengan ID tersebut tidak ditemukan.")
		return
	}

	fmt.Print("Masukkan Alamat Jalan (lama: " + emp.Alamat.Jalan + "): ")
	jalan, _ := reader.ReadString('\n')
	jalan = jalan[:len(jalan)-1]

	fmt.Print("Masukkan Kota (lama: " + emp.Alamat.Kota + "): ")
	fmt.Scan(&kota)

	fmt.Print("Masukkan Nomer Rumah (lama: ", emp.Alamat.Nomer, "): ")
	fmt.Scan(&nomer)

	fmt.Print("Masukkan No Telp (lama: " + emp.NoTelp + "): ")
	fmt.Scan(&notelp)

	fmt.Print("Masukkan Email (lama: " + emp.Email + "): ")
	fmt.Scan(&email)

	updatedEmp := node.Pegawai{
		ID: id,
		Alamat: node.Address{
			Jalan: jalan,
			Kota:  kota,
			Nomer: nomer,
		},
		NoTelp: notelp,
		Email:  email,
	}

	if model.Update(id, updatedEmp) {
		fmt.Println("\nData Pegawai Berhasil Diperbarui!\n")
	} else {
		fmt.Println("\nGagal memperbarui data.\n")
	}
}

func Delete() {
	var id int
	fmt.Println("DELETE DATA PEGAWAI ")

	fmt.Print("Masukkan ID Pegawai yang akan dihapus: ")
	fmt.Scan(&id)

	if model.Delete(id) {
		fmt.Println("\nPegawai Berhasil Dihapus!\n")
	} else {
		fmt.Println("\nPegawai Tidak Ditemukan!\n")
	}
}
