package main

import "fmt"

type dataPolusi struct {
	lokasi          string
	Date            string
	AQI             int
	Sumber          string
	tingkatKeamanan string
}

const maxData = 10

var data [maxData]dataPolusi
var hitungData int

func dataMasuk(in dataPolusi) {
	if hitungData >= maxData {
		fmt.Println("Data penuh, tidak bisa menambahkan data baru.")
		return
	}
	data[hitungData] = in
	hitungData++
	checkAlert(in)
	fmt.Println("Data berhasil ditambahkan.")
}

func checkAlert(d dataPolusi) {
	if d.AQI > 150 {
		fmt.Printf("PERINGATAN: Polusi tinggi di %s pada %s\n", d.lokasi, d.Date)
	}
}

func dataTambahan(lokasi, Date string, dataBaru dataPolusi) {
	var i int

	for i = 0; i < hitungData; i++ {
		if data[i].lokasi == lokasi && data[i].Date == Date {
			data[i] = dataBaru
			fmt.Println("Data berhasil diperbarui.")
			checkAlert(dataBaru)
			return
		}

	}

	fmt.Println("Data tidak ditemukan.")
}

func hapusData(lokasi, Date string) {
	var i, j int

	for i = 0; i < hitungData; i++ {
		if data[i].lokasi == lokasi && data[i].Date == Date {
			for j = i; j < hitungData-1; j++ {
				data[j] = data[j+1]
			}
			hitungData--
			fmt.Println("Data berhasil dihapus.")
			return
		}
	}
	fmt.Println("Data tidak ditemukan.")
}

func tampilkanData() {
	if hitungData == 0 {
		fmt.Println("Tidak ada data.")
		return
	}
	fmt.Println("== Semua Data Polusi ==")

	var i int

	for i = 0; i < hitungData; i++ {
		d := data[i]
		fmt.Printf("[%d] %s - %s | AQI: %d | Sumber: %s | Level: %s\n",
			i+1, d.lokasi, d.Date, d.AQI, d.Sumber, d.tingkatKeamanan)
	}
}

func searchSequential(lokasi string) {
	found := false
	for i := 0; i < hitungData; i++ {
		if data[i].lokasi == lokasi {
			fmt.Println(data[i])
			found = true
		}
	}
	if !found {
		fmt.Println("Tidak ditemukan.")
	}
}

func sortByLocation() {
	var i, j int
	for i = 0; i < hitungData-1; i++ {
		for j = i + 1; j < hitungData; j++ {
			if data[i].lokasi > data[j].lokasi {
				data[i], data[j] = data[j], data[i]
			}
		}
	}
}

func searchBinary(lokasi string) {
	sortByLocation()
	low := 0
	high := hitungData - 1
	for low <= high {
		mid := (low + high) / 2
		if data[mid].lokasi == lokasi {
			fmt.Println("Ditemukan:", data[mid])
			return
		} else if data[mid].lokasi < lokasi {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	fmt.Println("Tidak ditemukan.")
}

func sortByAQISelection() {
	for i := 0; i < hitungData-1; i++ {
		maxIdx := i
		for j := i + 1; j < hitungData; j++ {
			if data[j].AQI > data[maxIdx].AQI {
				maxIdx = j
			}
		}
		data[i], data[maxIdx] = data[maxIdx], data[i]
	}
	fmt.Println("Diurutkan berdasarkan AQI (tinggi ke rendah).")
}

func sortByDateInsertion() {
	for i := 1; i < hitungData; i++ {
		key := data[i]
		j := i - 1
		for j >= 0 && data[j].Date < key.Date {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
	fmt.Println("Diurutkan berdasarkan Tanggal (terbaru ke lama).")
}


func main() {
	var onGoing bool
	
	onGoing = true

	for onGoing {
		fmt.Println("\n== AIR QUALITY INDEX ==")
		fmt.Println("1. Tambah Data")
		fmt.Println("2. Perbarui Data")
		fmt.Println("3. Hapus Data")
		fmt.Println("4. Tampilkan Semua Data")
		fmt.Println("5. Cari Data (Sequential)")
		fmt.Println("6. Cari Data (Binary Search)")
		fmt.Println("7. Urutkan berdasarkan AQI (Selection Sort)")
		fmt.Println("8. Urutkan berdasarkan Tanggal (Insertion Sort)")
		fmt.Println("9. Keluar")
		fmt.Print("Pilih menu: ")


		var pilihan int
		fmt.Scanln(&pilihan)

		if pilihan == 1 {
			var d dataPolusi
			fmt.Print("Lokasi: ")
			fmt.Scanln(&d.lokasi)
			fmt.Print("Tanggal (YYYY-MM-DD): ")
			fmt.Scanln(&d.Date)
			fmt.Print("AQI: ")
			fmt.Scanln(&d.AQI)
			fmt.Print("Sumber: ")
			fmt.Scanln(&d.Sumber)
			fmt.Print("Tingkat Bahaya: ")
			fmt.Scanln(&d.tingkatKeamanan)
			dataMasuk(d)

		} else if pilihan == 2 {
			var loc, date string
			var d dataPolusi
			fmt.Print("Masukkan lokasi yang ingin diubah: ")
			fmt.Scanln(&loc)
			fmt.Print("Masukkan tanggal (YYYY-MM-DD): ")
			fmt.Scanln(&date)
			fmt.Print("Lokasi baru: ")
			fmt.Scanln(&d.lokasi)
			fmt.Print("Tanggal baru (YYYY-MM-DD): ")
			fmt.Scanln(&d.Date)
			fmt.Print("AQI baru: ")
			fmt.Scanln(&d.AQI)
			fmt.Print("Sumber baru: ")
			fmt.Scanln(&d.Sumber)
			fmt.Print("Tingkat bahaya baru: ")
			fmt.Scanln(&d.tingkatKeamanan)
			dataTambahan(loc, date, d)

		} else if pilihan == 3 {
			var loc, date string
			fmt.Print("Masukkan lokasi yang ingin dihapus: ")
			fmt.Scanln(&loc)
			fmt.Print("Masukkan tanggal (YYYY-MM-DD): ")
			fmt.Scanln(&date)
			hapusData(loc, date)

		} else if pilihan == 4 {
			tampilkanData()

		} else if pilihan == 5 {
			var loc string
			fmt.Print("Masukkan lokasi: ")
			fmt.Scanln(&loc)
			searchSequential(loc)

		} else if pilihan == 6 {
			var loc string
			fmt.Print("Masukkan lokasi: ")
			fmt.Scanln(&loc)
			searchBinary(loc)

		} else if pilihan == 7 {
			sortByAQISelection()
			tampilkanData()
		}  else if pilihan == 8 {
			sortByDateInsertion()
			tampilkanData()

		} else if pilihan == 9 {
			fmt.Println("Keluar dari program.")
			onGoing = false

		} else {
			fmt.Println("Menu tidak valid.")
		}
	}
}
