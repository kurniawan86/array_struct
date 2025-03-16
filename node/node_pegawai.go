package node

type Address struct {
	Jalan, Kota string
	Nomer       int
}

type Pegawai struct {
	ID     int
	Alamat Address
	NoTelp string
	Email  string
}
