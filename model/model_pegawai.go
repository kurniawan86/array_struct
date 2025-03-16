package model

import "array_struct/node"

var DaftarPegawai []node.Pegawai

func Create(emp node.Pegawai) {
	DaftarPegawai = append(DaftarPegawai, emp)
}

func Read() []node.Pegawai {
	return DaftarPegawai
}

func Update(index int, updated node.Pegawai) bool {
	for i, emp := range DaftarPegawai {
		if emp.ID == index {
			DaftarPegawai[i] = updated
			return true
		}
	}
	return false
}

func Delete(index int) bool {
	for i, emp := range DaftarPegawai {
		if emp.ID == index {
			DaftarPegawai = append(DaftarPegawai[:i], DaftarPegawai[i+1:]...)
			return true
		}
	}
	return false
}
