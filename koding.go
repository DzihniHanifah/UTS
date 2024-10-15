package main

import "fmt"

var (
	Username = "Dzihni"
	Password = "2406421516"
	History  []string
)

type Pengguna struct {
	Nama         string
	NPM          string
	Umur         int
	JenisKelamin string
	Alamat       string
	BukuFavorit  string
}

type Buku struct {
	Nama   string
	Jumlah int
}

var daftarBuku = []Buku{
	{Nama: "Pemrograman", Jumlah: 10},
	{Nama: "Film", Jumlah: 5},
	{Nama: "Printing", Jumlah: 20},
}

var pengguna = Pengguna{
	Nama:         "Dzihni Hanifah",
	NPM:          "2406421516",
	Umur:         18,
	JenisKelamin: "Perempuan",
	Alamat:       "Bogor",
	BukuFavorit:  "Tentang Kamu - Tere Liye",
}

func main() {
	var inputUsername, inputPassword string

	fmt.Println("========================================")
	fmt.Println("= Selamat Datang di Perpustakan Vokasi =")
	fmt.Println("========================================")

	fmt.Print("Silahkan input username: ")
	fmt.Scanf("%s\n", &inputUsername)

	fmt.Print("Silahkan input password: ")
	fmt.Scanf("%s\n", &inputPassword)

	fmt.Println("========================================")

	if inputUsername == Username && inputPassword == Password {
		fmt.Println("Login Berhasil!")

		for {
			fmt.Println("============== Menu Program ============")
			fmt.Println("1. Lihat Informasi Pengguna Program")
			fmt.Println("2. Lihat Daftar Buku")
			fmt.Println("3. Tambah Daftar Buku")
			fmt.Println("4. Tambah Peminjaman Buku")
			fmt.Println("5. Histori Peminjaman Buku")
			fmt.Println("6. Keluar dari Program")

			var pilihan int
			fmt.Print("Pilih Menu: ")
			fmt.Scanf("%d\n", &pilihan)

			fmt.Println("========================================")

			switch pilihan {
			case 1:
				LihatInformasiPenggunaProgram()
			case 2:
				LihatDaftarBuku()
			case 3:
				TambahDaftarBuku()
			case 4:
				TambahPeminjamanBuku()
			case 5:
				LihatHistoriPeminjaman()
			case 6:
				fmt.Println("== Terima Kasih! Sampai Jumpa Kembali ==")
				fmt.Println("========================================")
				return
			default:
				fmt.Println("Pilihan tidak valid.")
				fmt.Println("========================================")

			}
		}
	} else {
		fmt.Println("Username atau Password salah.")
	}
}

func LihatInformasiPenggunaProgram() {
	fmt.Println("========================================")
	fmt.Println("INFORMASI PENGGUNA PROGRAM")
	fmt.Printf("Nama: %s\n", pengguna.Nama)
	fmt.Printf("NPM: %s\n", pengguna.NPM)
	fmt.Printf("Umur: %d\n", pengguna.Umur)
	fmt.Printf("Jenis Kelamin: %s\n", pengguna.JenisKelamin)
	fmt.Printf("Alamat: %s\n", pengguna.Alamat)
	fmt.Printf("Buku Favorit: %s\n", pengguna.BukuFavorit)
	fmt.Println("========================================")
}

func LihatDaftarBuku() {
	fmt.Println("========================================")
	fmt.Println("DAFTAR BUKU")
	for i, buku := range daftarBuku {
		fmt.Printf("%d. Nama: %s - Jumlah: %d\n", i+1, buku.Nama, buku.Jumlah)
	}
	fmt.Println("========================================")
}

func TambahDaftarBuku() {
	var bukuBaru Buku
	fmt.Println("========================================")
	fmt.Println("TAMBAH DAFTAR BUKU")
	fmt.Print("Masukkan nama buku: ")
	fmt.Scanln(&bukuBaru.Nama)
	fmt.Print("Masukkan jumlah buku: ")
	fmt.Scanln(&bukuBaru.Jumlah)

	if bukuBaru.Jumlah <= 0 {
		fmt.Println("Jumlah harus lebih besar dari 0.")
		fmt.Println("========================================")
		return
	}

	daftarBuku = append(daftarBuku, bukuBaru)
	History = append(History, fmt.Sprintf("Menambahkan buku: %s - Jumlah: %d", bukuBaru.Nama, bukuBaru.Jumlah))
	fmt.Println("Buku berhasil ditambahkan!")
	fmt.Println("========================================")
}

func TambahPeminjamanBuku() {
	var namaBuku string
	var jumlahPinjam int
	fmt.Println("========================================")
	fmt.Println("TAMBAH PEMINJAMAN BUKU")
	fmt.Print("Masukkan nama buku: ")
	fmt.Scanln(&namaBuku)
	fmt.Print("Masukkan jumlah yang ingin dipinjam: ")
	fmt.Scanln(&jumlahPinjam)

	if jumlahPinjam <= 0 {
		fmt.Println("Jumlah harus lebih besar dari 0.")
		fmt.Println("========================================")
		return
	}

	for i, buku := range daftarBuku {
		if buku.Nama == namaBuku {
			if buku.Jumlah >= jumlahPinjam {
				daftarBuku[i].Jumlah -= jumlahPinjam
				History = append(History, fmt.Sprintf("Meminjam buku: %s - Jumlah: %d", namaBuku, jumlahPinjam))
				fmt.Println("Peminjaman berhasil.")
				fmt.Println("========================================")
				return
			} else {
				fmt.Println("Jumlah buku tidak mencukupi.")
				fmt.Println("========================================")
				return
			}
		}
	}
	fmt.Println("Buku tidak ditemukan.")
	fmt.Println("========================================")
}

func LihatHistoriPeminjaman() {
	if len(History) == 0 {
		fmt.Println("Tidak ada histori peminjaman.")
		fmt.Println("========================================")
		return
	}

	fmt.Println("========================================")
	fmt.Println("HISTORI PEMINJAMAN BUKU")
	for _, entry := range History {
		fmt.Println(entry)
	}
	fmt.Println("========================================")
}
