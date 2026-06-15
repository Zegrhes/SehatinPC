package main

import (
	"fmt"
)

const MAX = 1000

type isiLog struct {
	Suhu       float64
	BebanKerja float64
}

type Komponen struct {
	NomorSeri       string
	Nama            string
	StatusKesehatan string
	Suhu            float64
	BebanKerja      float64
	Log             [1000]isiLog
	NLog            int
}

type TabKomponen [MAX]Komponen

func main() {
	var data TabKomponen
	var n, pilihan int
	var jalan bool
	jalan = true

	inisialisasiData(&data, &n)

	for jalan == true {
		fmt.Println()
		fmt.Println("=====================================")
		fmt.Println(" Sistem Monitoring Kesehatan PC ")
		fmt.Println("=====================================")
		fmt.Println("1. Tambah Komponen")
		fmt.Println("2. Ubah Data Komponen")
		fmt.Println("3. Hapus Komponen")
		fmt.Println("4. Catat Log Status Komponen")
		fmt.Println("5. Urutkan Komponen (Nomor Seri - Insertion Sort - Ascending)")
		fmt.Println("6. Urutkan Komponen (Nama - Selection Sort - Descending)")
		fmt.Println("7. Cari Komponen (Nama - Binary Search)")
		fmt.Println("8. Cari Komponen (Status Kesehatan - Sequential Search)")
		fmt.Println("9. Tampilkan Statistik")
		fmt.Println("10. Tampilkan Semua Komponen")
		fmt.Println("11. Tampilkan Log Status Komponen")
		fmt.Println("12. Statistik Komponen")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)
		fmt.Println()
		if pilihan == 0 {
			fmt.Println("Keluar dari program..............")
			jalan = false
		} else {
			switch pilihan {
			case 1:
				tambahKomponen(&data, &n)
			case 2:
				ubahKomponen(&data, n)
			case 3:
				hapusKomponen(&data, &n)
			case 4:
				catatLog(&data, n)
			case 5:
				AscendingInsertionSorturutNomorSeri(&data, n)
				fmt.Println("Data berhasil diurutkan berdasarkan Nomor Seri (Ascending)")
				tampilSemua(data, n)
			case 6:
				DescendingSelectionSorturutNama(&data, n)
				fmt.Println("Data berhasil diurutkan berdasarkan Nama (Descending)")
				tampilSemua(data, n)
			case 7:
				BinarySearchNama(data, n)
			case 8:
				sequentialSearchStatus(data, n)
			case 9:
				statistikPC(data, n)
			case 10:
				tampilSemua(data, n)
			case 11:
				tampilLog(data, n)
			case 12:
				statistikKomponen(data, n)
			default:
				fmt.Println("Pilihan tidak valid")
			}
		}
	}
}

func tambahKomponen(T *TabKomponen, n *int) {
	if *n >= MAX {
		fmt.Println("Kapasitas penuh!")
	} else {
		fmt.Print("Masukkan Nomor Seri: ")
		fmt.Scan(&T[*n].NomorSeri)
		fmt.Print("Masukkan Nama Komponen: ")
		fmt.Scan(&T[*n].Nama)
		fmt.Print("Masukkan Suhu: ")
		fmt.Scan(&T[*n].Suhu)
		fmt.Print("Masukkan Beban Kerja (%): ")
		fmt.Scan(&T[*n].BebanKerja)
		T[*n].StatusKesehatan = tentukanStatus(T[*n].Suhu, T[*n].BebanKerja)

		*n = *n + 1
		fmt.Println("Komponen berhasil ditambahkan")
	}
}

func ubahKomponen(T *TabKomponen, n int) {
	var ns string
	var i, idx int
	i = 0
	idx = -1

	if n == 0 {
		fmt.Println("Data masih kosong")
	} else {
		tampilSemua(*T, n)

		fmt.Print("Masukkan Nomor Seri komponen yang akan diubah: ")
		fmt.Scan(&ns)

		for i < n && idx == -1 {
			if T[i].NomorSeri == ns {
				idx = i
			}
			i = i + 1
		}
		if idx != -1 {
			fmt.Print("Masukkan Nama baru: ")
			fmt.Scan(&T[idx].Nama)
			fmt.Print("Masukkan Suhu baru: ")
			fmt.Scan(&T[idx].Suhu)
			fmt.Print("Masukkan Beban Kerja baru (%): ")
			fmt.Scan(&T[idx].BebanKerja)

			T[idx].StatusKesehatan = tentukanStatus(T[idx].Suhu, T[idx].BebanKerja)

			fmt.Println("Data berhasil diubah")
		} else {
			fmt.Println("Komponen tidak ditemukan")
		}
	}
}

func hapusKomponen(T *TabKomponen, n *int) {
	var ns string
	var i, idx int
	i = 0
	idx = -1

	if *n == 0 {
		fmt.Println("Data masih kosong")
	} else {
		tampilSemua(*T, *n)

		fmt.Print("Masukkan Nomor Seri komponen yang akan dihapus: ")
		fmt.Scan(&ns)

		for i < *n && idx == -1 {
			if T[i].NomorSeri == ns {
				idx = i
			}
			i = i + 1
		}

		if idx != -1 {
			for i = idx; i < *n-1; i++ {
				T[i] = T[i+1]
			}
			*n = *n - 1
			fmt.Println("Data berhasil dihapus")
		} else {
			fmt.Println("Komponen tidak ditemukan")
		}
	}
}

func AscendingInsertionSorturutNomorSeri(T *TabKomponen, n int) {
	var pass, i int
	var temp Komponen
	if n < 2 {
		fmt.Println("Data masih kosong")
	} else {
		for pass = 1; pass < n; pass++ {
			temp = T[pass]
			i = pass - 1
			for i >= 0 && T[i].NomorSeri > temp.NomorSeri {
				T[i+1] = T[i]
				i = i - 1
			}
			T[i+1] = temp
		}
	}
}

func DescendingSelectionSorturutNama(T *TabKomponen, n int) {
	var pass, i, maxIdx int
	var temp Komponen

	if n < 2 {
		fmt.Println("Data masih kosong")
	} else {
		for pass = 0; pass < n-1; pass++ {
			maxIdx = pass
			for i = pass + 1; i < n; i++ {
				if T[i].Nama > T[maxIdx].Nama {
					maxIdx = i
				}
			}
			temp = T[pass]
			T[pass] = T[maxIdx]
			T[maxIdx] = temp
		}
	}
}

func BinarySearchNama(T TabKomponen, n int) {
	var cari string
	var left, right, foundIdx, mid int
	left = 0
	right = n - 1
	foundIdx = -1

	if n == 0 {
		fmt.Println("Data masih kosong")
	} else {
		fmt.Print("Masukkan Nama komponen yang dicari: ")
		fmt.Scan(&cari)

		for left <= right && foundIdx == -1 {
			mid = (left + right) / 2
			if T[mid].Nama == cari {
				foundIdx = mid
			} else if T[mid].Nama > cari {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}

		if foundIdx != -1 {
			fmt.Println("Komponen ditemukan (metode Binary Search): ")
			fmt.Printf("Nomor Seri: %-8s | Nama: %-12s | Status: %-9s | Suhu: %7.2f | Beban: %6.2f%%\n",
				T[foundIdx].NomorSeri, T[foundIdx].Nama, T[foundIdx].StatusKesehatan, T[foundIdx].Suhu, T[foundIdx].BebanKerja)
		} else {
			fmt.Println("Komponen tidak ditemukan")
		}
	}
}

func sequentialSearchStatus(T TabKomponen, n int) {
	var cari string
	var i int
	var found bool = false

	if n == 0 {
		fmt.Println("Data masih kosong")
	} else {
		fmt.Print("Masukkan Status Kesehatan yang dicari ( sehat / lag / overheat ): ")
		fmt.Scan(&cari)

		fmt.Println("Hasil pencarian (metode Sequential Search):")

		for i = 0; i < n; i++ {
			if T[i].StatusKesehatan == cari {
				fmt.Printf("- Nomor Seri: %-8s | Nama: %-12s | Status: %-9s | Suhu: %7.2f | Beban: %6.2f%%\n",
					T[i].NomorSeri, T[i].Nama, T[i].StatusKesehatan, T[i].Suhu, T[i].BebanKerja)
				found = true
			}
		}

		if found == false {
			fmt.Println("Tidak ada komponen dengan status tersebut")
		}
	}
}

func statistikPC(T TabKomponen, n int) {
	var bermasalah, i int
	var totalSuhu, rataSuhu float64
	var status string

	if n == 0 {
		fmt.Println("Data masih kosong")
	} else {
		for i = 0; i < n; i++ {
			status = T[i].StatusKesehatan
			if status == "lag" || status == "overheat" {
				bermasalah = bermasalah + 1
			}
			totalSuhu = totalSuhu + T[i].Suhu
		}

		rataSuhu = totalSuhu / float64(n)
		fmt.Println("--- Statistik Sistem ---")
		fmt.Printf("Jumlah komponen bermasalah (lag / overheat): %d\n", bermasalah)
		fmt.Printf("Rata-rata suhu komponen: %.2f\n", rataSuhu)
	}
}

func tampilSemua(T TabKomponen, n int) {
	var i int = 0

	if n == 0 {
		fmt.Println("Data masih kosong")
	} else {
		fmt.Println("Daftar Komponen:")
		for i = 0; i < n; i++ {
			fmt.Printf("%d. SN: %-8s | Nama: %-12s | Status: %-9s | Suhu: %7.2f | Beban: %6.2f%%\n",
				i+1, T[i].NomorSeri, T[i].Nama, T[i].StatusKesehatan, T[i].Suhu, T[i].BebanKerja)
		}
	}
}

func tentukanStatus(suhu float64, beban float64) string {
	var status string
	if suhu >= 85.0 {
		status = "overheat"
	} else if suhu >= 75.0 && beban >= 80.0 {
		status = "lag"
	} else if beban >= 95.0 {
		status = "lag"
	} else {
		status = "sehat"
	}
	return status
}

func catatLog(T *TabKomponen, n int) {
	var log isiLog
	var ns string
	var suhuBaru, bebanBaru float64
	var i int
	var found bool
	found = false

	if n == 0 {
		fmt.Println("Data masih kosong")
	} else {
		fmt.Print("Masukkan Nomor Seri komponen: ")
		fmt.Scan(&ns)

		for i < n && found == false {
			if T[i].NomorSeri == ns {
				found = true
				if T[i].NLog < 1000 {
					log.Suhu = T[i].Suhu
					log.BebanKerja = T[i].BebanKerja
					T[i].Log[T[i].NLog] = log
					T[i].NLog = T[i].NLog + 1
				} else {
					fmt.Println("Kapasitas log penuh")
				}

				fmt.Print("Masukkan Suhu baru: ")
				fmt.Scan(&suhuBaru)
				fmt.Print("Masukkan Beban Kerja baru (%): ")
				fmt.Scan(&bebanBaru)

				T[i].Suhu = suhuBaru
				T[i].BebanKerja = bebanBaru
				T[i].StatusKesehatan = tentukanStatus(T[i].Suhu, T[i].BebanKerja)
				fmt.Println("Data komponen berhasil diperbarui")
			}
			i = i + 1
		}
		if found == false {
			fmt.Println("Komponen tidak ditemukan")
		}
	}
}

func tampilLog(T TabKomponen, n int) {
	var ns string
	var i, j int
	var found bool

	if n == 0 {
		fmt.Println("Data masih kosong")
	} else {
		fmt.Print("Masukkan Nomor Seri komponen: ")
		fmt.Scan(&ns)

		for i < n && found == false {
			if T[i].NomorSeri == ns {
				found = true
				fmt.Printf("\nLog Komponen %s (%s):\n", T[i].Nama, T[i].NomorSeri)
				for j = 0; j < T[i].NLog; j++ {
					fmt.Printf("%d. Suhu: %7.2f | Beban: %6.2f%%\n", j+1, T[i].Log[j].Suhu, T[i].Log[j].BebanKerja)
				}
			}
			i = i + 1
		}

		if found == false {
			fmt.Println("Komponen tidak ditemukan")
		}
	}
}

func statistikKomponen(T TabKomponen, n int) {
	var ns string
	var i, j int
	var found bool

	if n == 0 {
		fmt.Println("Data masih kosong")
	} else {
		fmt.Print("Masukkan Nomor Seri: ")
		fmt.Scan(&ns)

		for i < n && found == false {
			if T[i].NomorSeri == ns {
				found = true

				fmt.Println("\n--- Data Komponen ---")
				fmt.Printf("Nomor Seri    : %s\n", T[i].NomorSeri)
				fmt.Printf("Nama          : %s\n", T[i].Nama)
				fmt.Printf("Status        : %s\n", T[i].StatusKesehatan)
				fmt.Printf("Suhu          : %7.2f\n", T[i].Suhu)
				fmt.Printf("Beban Kerja   : %7.2f%%\n", T[i].BebanKerja)

				fmt.Println("\n--- Log Komponen ---")
				for j = 0; j < T[i].NLog; j++ {
					fmt.Printf("%d. Suhu: %7.2f | Beban: %6.2f%%\n", j+1, T[i].Log[j].Suhu, T[i].Log[j].BebanKerja)
				}

				var totalSuhu, totalBeban, maxSuhu, minSuhu, maxBeban, minBeban float64
				minSuhu = T[i].Log[0].Suhu
				minBeban = T[i].Log[0].BebanKerja
				for j = 0; j < T[i].NLog; j++ {
					totalSuhu = totalSuhu + T[i].Log[j].Suhu
					totalBeban = totalBeban + T[i].Log[j].BebanKerja
					if T[i].Log[j].Suhu > maxSuhu {
						maxSuhu = T[i].Log[j].Suhu
					}
					if T[i].Log[j].Suhu < minSuhu {
						minSuhu = T[i].Log[j].Suhu
					}
					if T[i].Log[j].BebanKerja > maxBeban {
						maxBeban = T[i].Log[j].BebanKerja
					}
					if T[i].Log[j].BebanKerja < minBeban {
						minBeban = T[i].Log[j].BebanKerja
					}
				}

				fmt.Println("\n--- Statistik Komponen ---")
				fmt.Printf("Rata-rata Suhu         : %7.2f\n", totalSuhu/float64(T[i].NLog))
				fmt.Printf("Suhu Tertinggi         : %7.2f\n", maxSuhu)
				fmt.Printf("Suhu Terendah          : %7.2f\n", minSuhu)
				fmt.Printf("Rata-rata Beban        : %7.2f%%\n", totalBeban/float64(T[i].NLog))
				fmt.Printf("Beban Tertinggi        : %7.2f%%\n", maxBeban)
				fmt.Printf("Beban Terendah         : %7.2f%%\n", minBeban)
			}
			i = i + 1
		}

		if found == false {
			fmt.Println("Komponen tidak ditemukan")
		}
	}
}

func inisialisasiData(T *TabKomponen, n *int) {
	var i, j int
	type logEntry struct{ suhu, bebanKerja float64 }
	type dataTemp struct {
		nomorSeri, nama, statusKesehatan string
		suhu, bebanKerja                 float64
		logs                             [5]logEntry
	}

	var dataAwal = [8]dataTemp{
		{"PRC-001", "Processor", "lag", 72.5, 65.0, [5]logEntry{{72.5, 65.0}, {73.0, 66.0}, {71.5, 64.0}, {72.0, 65.0}, {73.5, 67.0}}},
		{"GPU-001", "GPU", "overheat", 81.0, 90.0, [5]logEntry{{81.0, 90.0}, {82.0, 91.0}, {80.0, 89.0}, {81.5, 90.0}, {83.0, 92.0}}},
		{"RAM-001", "RAM", "sehat", 45.2, 55.0, [5]logEntry{{45.2, 55.0}, {46.0, 56.0}, {44.5, 54.0}, {45.0, 55.0}, {46.5, 57.0}}},
		{"SSD-001", "SSD", "sehat", 38.0, 30.0, [5]logEntry{{38.0, 30.0}, {39.0, 31.0}, {37.0, 29.0}, {38.5, 30.0}, {39.5, 32.0}}},
		{"MB-001", "Motherboard", "sehat", 52.8, 45.0, [5]logEntry{{52.8, 45.0}, {53.5, 46.0}, {52.0, 44.0}, {53.0, 45.0}, {54.0, 47.0}}},
		{"PSU-001", "PSU", "sehat", 48.3, 70.0, [5]logEntry{{48.3, 70.0}, {49.0, 71.0}, {47.5, 69.0}, {48.0, 70.0}, {49.5, 72.0}}},
		{"CL-001", "Cooler", "sehat", 35.0, 50.0, [5]logEntry{{35.0, 50.0}, {36.0, 51.0}, {34.0, 49.0}, {35.5, 50.0}, {36.5, 52.0}}},
		{"CF-001", "Case Fan", "sehat", 32.5, 40.0, [5]logEntry{{32.5, 40.0}, {33.5, 41.0}, {31.5, 39.0}, {32.0, 40.0}, {34.0, 42.0}}},
	}

	for i = 0; i < 8; i++ {
		T[i].NomorSeri = dataAwal[i].nomorSeri
		T[i].Nama = dataAwal[i].nama
		T[i].StatusKesehatan = dataAwal[i].statusKesehatan
		T[i].Suhu = dataAwal[i].suhu
		T[i].BebanKerja = dataAwal[i].bebanKerja
		T[i].NLog = 5
		for j = 0; j < 5; j++ {
			T[i].Log[j] = isiLog{dataAwal[i].logs[j].suhu, dataAwal[i].logs[j].bebanKerja}
		}
	}

	*n = 8
	fmt.Println("Data dummy berhasil diinisialisasi")
}
