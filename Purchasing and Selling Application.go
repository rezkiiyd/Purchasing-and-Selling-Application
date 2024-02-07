package main

import "fmt"

type barang struct {
	ID              int
	Nama            string
	Kategori        string
	Harga           int
	Stok            int
	TotalPendapatan int
}

type transaksi struct {
	ID       int
	BarangID int
	Jumlah   int
	Tanggal  string
}

type arrBarang [1000]barang
type arrTransaksi [1000]transaksi

var totalModal, totalInputanModal int

func main() {
	var pilihan int
	var dataBarang arrBarang
	var dataTransaksi arrTransaksi
	var jumlahBarang, jumlahTransaksi, modal int

	for {
		menu(&pilihan)
		if pilihan == 0 {
			fmt.Println("Terima kasih! Program telah berakhir.")
			return
		} else if pilihan == 1 {
			tambahModal(&modal)
		} else if pilihan == 2 {
			tambahBarang(&dataBarang, &jumlahBarang)
		} else if pilihan == 3 {
			fmt.Print("Masukkan ID Barang yang akan diubah: ")
			var idBarang int
			fmt.Scanln(&idBarang)
			editBarang(&dataBarang, jumlahBarang, idBarang)
		} else if pilihan == 4 {
			fmt.Print("Masukkan ID Barang yang akan dihapus: ")
			var idBarang int
			fmt.Scanln(&idBarang)
			hapusBarang(&dataBarang, &jumlahBarang, idBarang)
		} else if pilihan == 5 {
			tambahTransaksi(&dataTransaksi, &jumlahTransaksi, &dataBarang, jumlahBarang)
		} else if pilihan == 6 {
			fmt.Print("Masukkan ID Transaksi yang akan diubah: ")
			var idTransaksi int
			fmt.Scanln(&idTransaksi)
			editTransaksi(&dataTransaksi, jumlahTransaksi, idTransaksi)
		} else if pilihan == 7 {
			fmt.Print("Masukkan ID Transaksi yang akan dihapus: ")
			var idTransaksi int
			fmt.Scanln(&idTransaksi)
			hapusTransaksi(&dataTransaksi, &jumlahTransaksi, idTransaksi)
		} else if pilihan == 8 {
			fmt.Println("=== Daftar Barang Terurut ===")
			fmt.Print("Masukkan Daftar Barang Berdasarkan Harga (Tertinggi/Terendah): ")
			var kategori string
			fmt.Scanln(&kategori)
			if kategori == "Tertinggi" {
				sortBarangByHargaTertinggi(&dataBarang, jumlahBarang, kategori)
				tampilkanBarang(dataBarang, jumlahBarang)
			} else if kategori == "Terendah" {
				sortBarangByHargaTerendah(&dataBarang, jumlahBarang, kategori)
				tampilkanBarang(dataBarang, jumlahBarang)
			} else {
				fmt.Println("Pilihan Tidak Valid")
			}
		} else if pilihan == 9 {
			fmt.Println("=== Cari Data Barang ===")
			fmt.Print("Masukkan kata kunci pencarian (Nama Barang): ")
			var keyword string
			fmt.Scanln(&keyword)
			cariDataBarang(dataBarang, jumlahBarang, keyword)
		} else if pilihan == 10 {
			tampilkanModal(totalInputanModal, totalModal)
		} else if pilihan == 11 {
			tampilkanDataPendapatan(dataBarang, jumlahBarang)
		} else if pilihan == 12 {
			tampilkanDataBarangTerjualTerbanyak(dataBarang, dataTransaksi, jumlahBarang, jumlahTransaksi)
		} else if pilihan == 13 {
			tampilkanDataBarangTerkini(dataBarang, jumlahBarang)
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func menu(pilihan *int) {
	fmt.Println("=== APLIKASI PEGAWAI TOKO ===")
	fmt.Println("    Halo, Selamat Datang!    ")
	fmt.Println("1. Tambah Modal")
	fmt.Println("2. Tambah Data Barang")
	fmt.Println("3. Edit Data Barang")
	fmt.Println("4. Hapus Data Barang")
	fmt.Println("5. Tambah Data Transaksi")
	fmt.Println("6. Edit Data Transaksi")
	fmt.Println("7. Hapus Data Transaksi")
	fmt.Println("8. Lihat Data Barang Terurut")
	fmt.Println("9. Cari Data Barang")
	fmt.Println("10. Tampilkan Modal")
	fmt.Println("11. Tampilkan Pendapatan")
	fmt.Println("12. Tampilkan Barang Terjual Terbanyak")
	fmt.Println("13. Tampilkan Data Barang Terkini")
	fmt.Println("0. Keluar")
	fmt.Print("Masukkan pilihan Anda: ")
	fmt.Scanln(pilihan)
}
func tambahModal(modal *int) {
	fmt.Print("Masukkan modal baru: ")
	fmt.Scanln(*&modal)
	totalModal += *modal
	totalInputanModal += *modal
	fmt.Println("Modal berhasil ditambahkan.")
}

func tampilkanModal(totalInputanModal, totalModal int) {
	fmt.Println("Total modal awal:", totalInputanModal)
	fmt.Println("Total modal saat ini:", totalModal)
}

func tambahBarang(t *arrBarang, n *int) {
	var brg barang
	fmt.Print("Masukkan ID Barang: ")
	fmt.Scanln(&brg.ID)
	for brg.ID != 0 && *n < 1000 {
		fmt.Print("Masukkan Nama Barang: ")
		fmt.Scanln(&brg.Nama)
		fmt.Print("Masukkan Kategori Barang: ")
		fmt.Scanln(&brg.Kategori)
		fmt.Print("Masukkan Harga Barang: ")
		fmt.Scanln(&brg.Harga)
		fmt.Print("Masukkan Stok Barang: ")
		fmt.Scanln(&brg.Stok)

		modalRequired := brg.Stok * brg.Harga
		if modalRequired > totalModal {
			fmt.Println("Modal tidak cukup. Barang tidak dapat diproses.")
		} else {
			t[*n] = brg
			totalModal -= modalRequired
			*n++
			fmt.Println("Barang berhasil ditambahkan.")
		}

		fmt.Print("Masukkan ID Barang (0 untuk berhenti): ")
		fmt.Scanln(&brg.ID)
	}
}

func cariBarang(t arrBarang, n int, ID int) int {
	var ketemu int = -1
	var i int = 0
	for i < n && ketemu == -1 {
		if t[i].ID == ID {
			ketemu = i
		}
		i++
	}
	return ketemu
}

func editBarang(t *arrBarang, n int, ID int) {
	found := cariBarang(*t, n, ID)
	if found == -1 {
		fmt.Println("ID Barang tidak ditemukan")
	} else {
		var nama, kategori string
		var harga, stok int

		fmt.Print("Masukkan nama baru Barang: ")
		fmt.Scanln(&nama)
		fmt.Print("Masukkan kategori baru Barang: ")
		fmt.Scanln(&kategori)
		fmt.Print("Masukkan harga baru Barang: ")
		fmt.Scanln(&harga)
		fmt.Print("Masukkan stok baru Barang: ")
		fmt.Scanln(&stok)

		t[found].Nama = nama
		t[found].Kategori = kategori
		t[found].Harga = harga
		t[found].Stok = stok

		fmt.Println("Data barang berhasil diubah.")
	}
}

func hapusBarang(t *arrBarang, n *int, ID int) {
	found := cariBarang(*t, *n, ID)
	if found == -1 {
		fmt.Println("Barang tidak ditemukan")
	} else {
		for i := found; i < *n-1; i++ {
			t[i] = t[i+1]
		}
		*n--
		fmt.Println("Data barang berhasil dihapus.")
	}
}

func cariTransaksi(t arrTransaksi, n int, ID int) int {
	var ketemu int = -1
	var i int = 0
	for i < n && ketemu == -1 {
		if t[i].ID == ID {
			ketemu = i
		}
		i++
	}
	return ketemu
}

func tambahTransaksi(t *arrTransaksi, n *int, b *arrBarang, nb int) {
	var trx transaksi
	fmt.Print("Masukkan ID Transaksi: ")
	fmt.Scanln(&trx.ID)
	for trx.ID != 0 && *n < 1000 {
		var validasiStok bool = true

		fmt.Print("Masukkan ID Barang: ")
		fmt.Scanln(&trx.BarangID)
		fmt.Print("Masukkan Jumlah Barang: ")
		fmt.Scanln(&trx.Jumlah)
		fmt.Print("Masukkan Tanggal Transaksi: ")
		fmt.Scanln(&trx.Tanggal)

		found := cariBarang(*b, nb, trx.BarangID)
		if found != -1 {
			if trx.Jumlah > (*b)[found].Stok {
				fmt.Println("Jumlah barang yang dimasukkan melebihi stok yang tersedia.")
				validasiStok = false
			} else {
				(*b)[found].Stok -= trx.Jumlah
				(*b)[found].TotalPendapatan += trx.Jumlah * (*b)[found].Harga
			}
		} else {
			validasiStok = false
		}

		if validasiStok {
			t[*n] = trx
			*n++
			fmt.Println("Transaksi berhasil diproses.")
		} else {
			fmt.Println("Transaksi tidak dapat diproses.")
		}

		fmt.Print("Masukkan ID Transaksi (0 untuk berhenti): ")
		fmt.Scanln(&trx.ID)
	}
}

func editTransaksi(t *arrTransaksi, n int, ID int) {
	found := cariTransaksi(*t, n, ID)
	if found == -1 {
		fmt.Println("ID Transaksi tidak ditemukan")
	} else {
		var barangID, jumlah int
		var tanggal string

		fmt.Print("Masukkan ID Barang Baru: ")
		fmt.Scanln(&barangID)
		fmt.Print("Masukkan Jumlah Barang Baru: ")
		fmt.Scanln(&jumlah)
		fmt.Print("Masukkan Tanggal Transaksi Baru: ")
		fmt.Scanln(&tanggal)

		t[found].BarangID = barangID
		t[found].Jumlah = jumlah
		t[found].Tanggal = tanggal

		fmt.Println("Data transaksi berhasil diubah.")
	}
}

func hapusTransaksi(t *arrTransaksi, n *int, ID int) {
	found := cariTransaksi(*t, *n, ID)
	if found == -1 {
		fmt.Println("ID Transaksi tidak ditemukan")
	} else {
		for i := found; i < *n-1; i++ {
			t[i] = t[i+1]
		}
		*n--
		fmt.Println("Data transaksi berhasil dihapus.")
	}
}

func sortBarangByHargaTerendah(t *arrBarang, n int, kategori string) {
	var pass, i int
	var temp barang
	pass = 1
	for pass <= n-1 {
		i = pass
		temp = t[pass]
		for i > 0 && temp.Harga < t[i-1].Harga {
			t[i] = t[i-1]
			i--
		}
		t[i] = temp
		pass++
	}
}

func sortBarangByHargaTertinggi(t *arrBarang, n int, kategori string) {
	var pass, i int
	var temp barang
	pass = 1
	for pass <= n-1 {
		i = pass
		temp = t[pass]
		for i > 0 && temp.Harga > t[i-1].Harga {
			t[i] = t[i-1]
			i--
		}
		t[i] = temp
		pass++
	}
}

func cariDataBarang(t arrBarang, n int, tipeData string) {
	fmt.Println("Hasil Pencarian:")
	found := false
	for i := 0; i < n; i++ {
		if t[i].Nama == tipeData {
			fmt.Println("ID:", t[i].ID)
			fmt.Println("Nama:", t[i].Nama)
			fmt.Println("Kategori:", t[i].Kategori)
			fmt.Println("Harga:", t[i].Harga)
			fmt.Println("Stok:", t[i].Stok)
			fmt.Println("====================")
			found = true
		}
	}
	if !found {
		fmt.Println("Data tidak ditemukan.")
	}
}

func tampilkanDataPendapatan(b arrBarang, nb int) {
	fmt.Println("=== Data Pendapatan Barang ===")
	for i := 0; i < nb; i++ {
		fmt.Println("Nama Barang:", b[i].Nama)
		fmt.Println("Total Pendapatan:", b[i].TotalPendapatan)
		fmt.Println("=============================")
	}
	totalPendapatan := 0
	for i := 0; i < nb; i++ {
		totalPendapatan += b[i].TotalPendapatan
	}
	fmt.Println("Total Pendapatan Keseluruhan:", totalPendapatan)
}

func tampilkanBarang(t arrBarang, n int) {
	fmt.Println("=== Daftar Barang ===")
	for i := 0; i < n; i++ {
		fmt.Println("ID:", t[i].ID)
		fmt.Println("Nama:", t[i].Nama)
		fmt.Println("Kategori:", t[i].Kategori)
		fmt.Println("Harga:", t[i].Harga)
		fmt.Println("Stok:", t[i].Stok)
		fmt.Println("====================")
	}
}

func tampilkanDataBarangTerjualTerbanyak(b arrBarang, t arrTransaksi, nb, nt int) {
	var penjualan [1000]int
	maxBarang := min(5, nb)
	for i := 0; i < nb; i++ {
		for j := 0; j < nt; j++ {
			if b[i].ID == t[j].BarangID {
				penjualan[i] += t[j].Jumlah
			}
		}
	}
	fmt.Println("=== Data Barang Terjual Terbanyak ===")
	for i := 0; i < maxBarang; i++ {
		maxIndex := 0
		maxTerjual := penjualan[maxIndex]

		for j := 1; j < nb; j++ {
			if penjualan[j] > maxTerjual {
				maxIndex = j
				maxTerjual = penjualan[j]
			}
		}

		fmt.Println("Nama Barang: ", b[maxIndex].Nama)
		fmt.Println("Jumlah Terjual: ", penjualan[maxIndex])
		fmt.Println("=============================")

		penjualan[maxIndex] = -1
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func tampilkanDataBarangTerkini(t arrBarang, n int) {
	fmt.Println("=== Data Barang Terkini ===")
	for i := 0; i < n; i++ {
		fmt.Println("ID:", t[i].ID)
		fmt.Println("Nama:", t[i].Nama)
		fmt.Println("Kategori:", t[i].Kategori)
		fmt.Println("Harga:", t[i].Harga)
		fmt.Println("Stok:", t[i].Stok)
		fmt.Println("====================")
	}
}
