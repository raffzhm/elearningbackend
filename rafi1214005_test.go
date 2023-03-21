package rafi1214005

import (
	"fmt"
	"testing"
)

func TestRegisterMhs(t *testing.T) {
	nm := "user1"
	em := "user1@gmail.com"
	nps := "12345user"
	hasil := RegisterMhs(nm, em, nps)
	fmt.Println(hasil)

}

func TestEnrolMatakuliah(t *testing.T) {
	mn := "Pemrograman 3"
	mk := "No Code"
	ml := "Gedung Pendidikan"
	hasil := EnrolMatakuliah(mn, mk, ml)
	fmt.Println(hasil)

}