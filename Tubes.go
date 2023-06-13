package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

type bankSoal [1000]struct {
	desc           string
	j1, j2, j3, j4 string
	ans            string
	benar, salah   int
	solution       string
}

type player [1000]struct {
	nama string
	skor int
}

func loginAdmin(bS *bankSoal, n *int) {
	var choice int
	var lanjut = true
	fmt.Println("--Selamat Datang--")
	fmt.Println()
	for lanjut {
		fmt.Println("==================")
		fmt.Println("    MENU ADMIN    ")
		fmt.Println("==================")
		fmt.Println("Sebagai Admin anda dapat:")
		fmt.Println("1. Menambahkan Soal \n2. Menghapus Soal \n3. Mengedit Soal\n4. Menampilkan Soal\n5. Keluar")
		fmt.Print("Silahkan pilih: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println()
			addSoal(bS, n)
		} else if choice == 2 {
			fmt.Println()
			deleteSoal(bS, n)
		} else if choice == 3 {
			fmt.Println()
			editSoal(bS, n)
		} else if choice == 4 {
			fmt.Println()
			showSoal(bS, *n)
		} else if choice == 5 {
			fmt.Println()
			fmt.Println("====KEMBALI KE MENU SEBELUMNYA====")
			fmt.Println()
			lanjut = false
		} else {
			fmt.Println()
			fmt.Println("INPUT ERROR. Silahkan masukan lagi")
			fmt.Println()
		}

	}
}

func scanString(s *string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	//scanner.Scan()
	//for scanner.Scan() {
	//	fmt.Println(scanner.Text())
	//}
	*s = scanner.Text()
}

func addSoal(bS *bankSoal, n *int) {
	var lanjut = "Y"
	fmt.Println("===  TAMBAH SOAL  ===")
	for lanjut == "Y" {
		fmt.Println("Jumlah soal pada bank soal saat ini adalah:", *n)
		fmt.Print("Silahkan masukan soal\n")

		fmt.Print("Soal: ")
		scanString(&bS[*n].desc)
		scanString(&bS[*n].desc)

		fmt.Println("Silahkan masukan 4 pilihan jawaban")

		fmt.Print("A. ")
		scanString(&bS[*n].j1)

		fmt.Print("B. ")
		scanString(&bS[*n].j2)

		fmt.Print("C. ")
		scanString(&bS[*n].j3)

		fmt.Print("D. ")
		scanString(&bS[*n].j4)

		fmt.Println("Silahkan masukan opsi jawaban soal tersebut (A/B/C/D)")
		scanString(&bS[*n].ans)

		fmt.Println("Silahkan masukan solusi/penjelasan soal tersebut")
		scanString(&bS[*n].solution)

		*n++

		fmt.Print("Apakah anda masih ingin menambahkan soal? (Y/N)")
		fmt.Scan(&lanjut)
		fmt.Println()
		for lanjut != "Y" && lanjut != "N" {
			fmt.Println("INPUT ERROR. Silahkan masukan Y/N")
			fmt.Scan(&lanjut)
			fmt.Println()
		}

	}
	fmt.Println("====KEMBALI KE MENU SEBELUMNYA====")
	fmt.Println()
}

func deleteSoal(bS *bankSoal, n *int) {
	var choice int
	var lanjut = "Y"
	fmt.Println("===  HAPUS SOAL  ===")
	for lanjut == "Y" {
		if *n == 0 {
			fmt.Println("Saat ini tidak terdapat soal dalam bank Soal. Silahkan masukan soal terlebih dahulu")
			lanjut = "N"
		} else {

			fmt.Println("Silahkan pilih soal yang ingin anda hapus")
			for i := 0; i < *n; i++ {
				fmt.Print(i+1, ": ", bS[i].desc)
				fmt.Println()
			}
			fmt.Print("Masukan nomor soal yg ingin anda hapus: ")
			fmt.Scan(&choice)

			choice--
			if *n == 1 {
				*n--
			} else {
				for i := choice; i < *n-1; i++ {
					bS[i].desc = bS[i+1].desc
					bS[i].j1 = bS[i+1].j1
					bS[i].j2 = bS[i+1].j2
					bS[i].j3 = bS[i+1].j3
					bS[i].j4 = bS[i+1].j4
					bS[i].ans = bS[i+1].ans
				}
				*n--
			}
			fmt.Print("Apakah anda masih ingin menghapus soal? (Y/N) ")
			fmt.Scan(&lanjut)
			for lanjut != "Y" && lanjut != "N" {
				fmt.Println("INPUT ERROR. Silahkan masukan Y/N")
				fmt.Scan(&lanjut)
			}
		}
	}
	fmt.Println()
	fmt.Println("====KEMBALI KE MENU SEBELUMNYA====")
	fmt.Println()
}

func editSoal(bS *bankSoal, n *int) {
	var choice int
	for choice != 3 {
		fmt.Println("====  EDIT SOAL  ====")
		fmt.Println("Silahkan pilih opsi berikut ini")
		fmt.Println("1. Mengedit Soal")
		fmt.Println("2. Mengedit Kunci Jawaban Soal")
		fmt.Println("3. Kembali ke menu sebelumnya")

		fmt.Print("Silahkan masukan pilihan anda: ")
		fmt.Scan(&choice)
		fmt.Println()
		if choice == 1 {
			var nomor int
			showSoal(bS, *n)
			if *n != 0 {
				fmt.Print("Silahkan pilih nomor soal yang ingin anda edit: ")
				fmt.Scan(&nomor)
				nomor--
				fmt.Println("Silahkan masukan Soal Baru")
				scanString(&bS[nomor].desc)
				scanString(&bS[nomor].desc)
			}
		} else if choice == 2 {
			var nomor int
			showSoal(bS, *n)
			if *n != 0 {
				fmt.Print("Silahkan pilih nomor soal yang ingin anda edit: ")
				fmt.Scan(&nomor)
				nomor--
				fmt.Println("Silahkan masukan kunci jawaban yang baru")
				fmt.Print("A. ")
				scanString(&bS[nomor].j1)
				scanString(&bS[nomor].j1)

				fmt.Print("B. ")
				scanString(&bS[nomor].j2)

				fmt.Print("C. ")
				scanString(&bS[nomor].j3)

				fmt.Print("D. ")
				scanString(&bS[nomor].j4)

				fmt.Println("Silahkan masukan solusi soal tersebut")
				scanString(&bS[nomor].ans)
			}
		} else if choice == 3 {
			fmt.Println("====KEMBALI KE MENU SEBELUMNYA====")
		}

	}
}

func showSoal(bS *bankSoal, n int) {
	if n == 0 {
		fmt.Println("Tidak terdapat soal pada bank soal. Silahkan mengisi soal terlebih dahulu")

	} else {
		var choice int
		var bSsorted = bS
		fmt.Println("Silahkan Pilih Opsi Berikut")
		fmt.Println("1. Tampilkan Semua Soal")
		fmt.Println("2. Tampilkan 5 Soal Paling Banyak Benar")
		fmt.Println("3. Tampilkan 5 Soal Paling Banyak Salah")
		fmt.Scan(&choice)
		fmt.Println()

		for choice != 1 && choice != 2 && choice != 3 {
			fmt.Println("Input yang dimasukkan salah. Silahkan masukan kembali (1/2)")
			fmt.Scan(&choice)
		}
		if choice == 1 {

			for i := 0; i < n; i++ {
				fmt.Print(i+1, ". ", bSsorted[i].desc, "\n")
				fmt.Println("A.", bSsorted[i].j1)
				fmt.Println("B.", bSsorted[i].j2)
				fmt.Println("C.", bSsorted[i].j3)
				fmt.Println("D.", bSsorted[i].j4)
				fmt.Println("Jawaban:", bSsorted[i].ans)
				fmt.Println("Benar: ", bSsorted[i].benar, "Salah: ", bSsorted[i].salah)
				fmt.Println()
			}

		} else if choice == 2 {
			showMostBenar(bSsorted, n)
			fmt.Println()
		} else if choice == 3 {
			showMostSalah(bSsorted, n)
			fmt.Println()
		}

	}
}

func showMostBenar(bS *bankSoal, n int) {
	if n < 5 {
		for i := 0; i < n; i++ {
			var ptr = i
			for j := i + 1; j < n; j++ {
				if bS[j].benar > bS[ptr].benar {
					ptr = j
				}
			}
			bS[i], bS[ptr] = bS[ptr], bS[i]

			fmt.Print(i+1, ". ", bS[i].desc, "\n")
			fmt.Println("A.", bS[i].j1)
			fmt.Println("B.", bS[i].j2)
			fmt.Println("C.", bS[i].j3)
			fmt.Println("D.", bS[i].j4)
			fmt.Println("Jawaban:", bS[i].ans)
			fmt.Println("Benar: ", bS[i].benar, "Salah: ", bS[i].salah)

		}
	}
}

func showMostSalah(bS *bankSoal, n int) {
	if n < 5 {
		for i := 0; i < n; i++ {
			var ptr = i
			for j := i + 1; j < n; j++ {
				if bS[j].salah > bS[ptr].salah {
					ptr = j
				}
			}
			bS[i], bS[ptr] = bS[ptr], bS[i]

			fmt.Print(i+1, ". ", bS[i].desc, "\n")
			fmt.Println("A.", bS[i].j1)
			fmt.Println("B.", bS[i].j2)
			fmt.Println("C.", bS[i].j3)
			fmt.Println("D.", bS[i].j4)
			fmt.Println("Jawaban:", bS[i].ans)
			fmt.Println("Benar: ", bS[i].benar, "Salah: ", bS[i].salah)

		}
	} else {
		for i := 0; i < 5; i++ {
			var ptr = i
			for j := i + 1; j < n; j++ {
				if bS[j].salah > bS[ptr].salah {
					ptr = j
				}
			}
			bS[i], bS[ptr] = bS[ptr], bS[i]

			fmt.Print(i+1, ". ", bS[i].desc, "\n")
			fmt.Println("A.", bS[i].j1)
			fmt.Println("B.", bS[i].j2)
			fmt.Println("C.", bS[i].j3)
			fmt.Println("D.", bS[i].j4)
			fmt.Println("Jawaban:", bS[i].ans)
			fmt.Println("Benar: ", bS[i].benar, "Salah: ", bS[i].salah)

		}
	}
}

func loginPlayer(pemain *player, bS *bankSoal, nSoal int, nPlayer *int) {
	var nama string
	var idx int
	var choice int
	fmt.Println("===  SELAMAT DATANG  ====")

	fmt.Println("Silahkan masukan nama anda")
	fmt.Scan(&nama)
	fmt.Println()
	idx = found(*pemain, nama, *nPlayer)

	if idx == -1 {
		fmt.Println("==================================")
		fmt.Println("    SELAMAT DATANG PEMAIN BARU    ")
		fmt.Println("==================================")
		fmt.Println()
		pemain[*nPlayer].nama = nama
		*nPlayer++
	} else {
		fmt.Println("======================")
		fmt.Println("    SELAMAT DATANG    ")
		fmt.Println("    ", pemain[idx].nama, "    ")
		fmt.Println("======================")
		fmt.Println()
	}
	for choice != 3 {
		//fmt.Println(*nPlayer)
		fmt.Println("**** Silahkan Pilih Menu Berikut ****")
		fmt.Println("1. Memulai Quiz\n2. Melihat Skor\n3. Kembali ke menu sebelumnya")
		fmt.Scan(&choice)
		fmt.Println()

		idx = found(*pemain, nama, *nPlayer)
		for choice != 1 && choice != 2 && choice != 3 {
			fmt.Println("Input yang dimasukan salah. Silahkan masukan lagi (1/2)")
			fmt.Scan(&choice)
		}
		if choice == 1 {
			mulaiQuiz(pemain, bS, nSoal, idx)
		} else if choice == 2 {
			printSkor(*pemain, *nPlayer)
		} else if choice == 3 {
			fmt.Println("====  Kembali Ke Menu Sebelumnya  ====")
		}
	}
}

func found(pemain player, nama string, n int) int {
	for i := 0; i < n; i++ {
		if pemain[i].nama == nama {
			return i
		}
	}
	return -1
}

func mulaiQuiz(pemain *player, bS *bankSoal, nSoal, nPlayer int) {
	var jml int
	var lanjut = "Y"
	var skorPemain int

	for lanjut == "Y" {
		fmt.Println("Jumlah soal saat ini: ", nSoal)
		fmt.Println("Silahkan pilih banyak soal yang ingin dijawab")
		fmt.Scan(&jml)

		for jml > nSoal {
			fmt.Println("Maaf anda memilih lebih banyak dari jumlah soal saat ini")
			fmt.Println("Silahkan masukan kembali jumlah soal  ( jumlah soal saat ini =", nSoal, ")")
			fmt.Scan(&jml)
			fmt.Println()
		}

		randNum := rand.Perm(jml)
		fmt.Println()
		fmt.Println("===  MEMULAI PERMAINAN  ===")
		fmt.Println()
		for i := 0; i < jml; i++ {
			var jawaban string
			var ans string

			fmt.Print(i+1, ". ", bS[randNum[i]].desc, "\n")
			fmt.Println("A.", bS[randNum[i]].j1)
			fmt.Println("B.", bS[randNum[i]].j2)
			fmt.Println("C.", bS[randNum[i]].j3)
			fmt.Println("D.", bS[randNum[i]].j4)
			fmt.Print("Jawab (A/B/C/D): ")
			fmt.Scan(&jawaban)
			fmt.Println()

			for jawaban != "A" && jawaban != "B" && jawaban != "C" && jawaban != "D" {
				fmt.Println("Opsi yang anda masukan salah. Silahkan Masukan Kembali (A/B/C/D)")
				fmt.Scan(&jawaban)
				fmt.Println()
			}

			if jawaban == "A" {
				ans = "A"
			} else if jawaban == "B" {
				ans = "B"
			} else if jawaban == "C" {
				ans = "C"
			} else if jawaban == "D" {
				ans = "D"
			}

			fmt.Println("Jawaban:", bS[randNum[i]].ans)
			fmt.Println("Solusi :", bS[randNum[i]].solution)
			fmt.Println()
			if ans == bS[randNum[i]].ans {
				bS[randNum[i]].benar++
				skorPemain += 100
				fmt.Println("Selamat anda benar. Skor Anda =", skorPemain)
			} else {
				bS[randNum[i]].salah++
				fmt.Println("Anda Salah Menjawab. Skor Anda = ", skorPemain)
			}
		}

		fmt.Println()
		fmt.Println("=======================")
		fmt.Println("   PERMAINAN SELESAI   ")
		fmt.Println("=======================")
		fmt.Println("Anda telah menyelesaikan permainan.\nSkor akhir anda:", skorPemain, "Skor anda sebelumnya:", pemain[nPlayer].skor)

		if skorPemain > pemain[nPlayer].skor {
			pemain[nPlayer].skor = skorPemain
		}

		fmt.Println()
		fmt.Print("Apakah Anda masih ingin bermain? (Y/N)")
		fmt.Scan(&lanjut)
		fmt.Println()
		skorPemain = 0
	}
	fmt.Println("===   Kembali Ke Menu Sebelumnya   ===")
	fmt.Println()
}

func sortSkor(pemain *player, nPlayer int) {
	for i := 0; i < nPlayer; i++ {
		var ptr = i
		for j := i + 1; j < nPlayer; j++ {
			if pemain[j].skor > pemain[ptr].skor {
				ptr = j
			}
		}
		pemain[i], pemain[ptr] = pemain[ptr], pemain[i]
	}
}

func printSkor(pemain player, nPlayer int) {
	fmt.Println("==================")
	fmt.Println("    PRINT SKOR    ")
	fmt.Println("==================")
	sortSkor(&pemain, nPlayer)
	for i := 0; i < nPlayer; i++ {
		fmt.Println(i+1, ".", pemain[i].nama, pemain[i].skor)
		fmt.Println()
	}
}

func main() {
	var bS bankSoal
	var pemain player
	var nSoal = 0
	var nPemain = 0
	var choice int

	for choice != 3 {
		fmt.Println("-------------------------------------")
		fmt.Println("     WHO ONE TO BE A MILLIONAIRE     ")
		fmt.Println("-------------------------------------")

		fmt.Println("Pilih Menu Berikut")
		fmt.Println("1. Login as Admin")
		fmt.Println("2. Login as Player")
		fmt.Println("3. Keluar")
		fmt.Print("Silahkan masukan pilihan anda: ")
		fmt.Scanln(&choice)
		fmt.Println()

		for choice != 1 && choice != 2 && choice != 3 {
			fmt.Println("Input yang dimasukan salah. Silahkan masukan kembali (1/2/3)")
			fmt.Scan(&choice)
		}

		if choice == 1 {
			loginAdmin(&bS, &nSoal)
			fmt.Scanln(&choice)
		} else if choice == 2 {
			if nSoal == 0 {
				fmt.Println("Tidak terdapat soal dalam bank Soal. Silahkan masukan Soal terlebih dahulu")
			} else {
				loginPlayer(&pemain, &bS, nSoal, &nPemain)
				fmt.Scanln(&choice)
			}

		} else if choice == 3 {
			fmt.Println("===THANK YOU===")
		}

	}

}
