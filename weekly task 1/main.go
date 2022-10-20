package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type Books struct {
	ID     		uuid.UUID
	judul  		string
	kategori 	string
	harga 		uint
}

func CreateBook(list_book *[]Books, judul, kategori string, harga uint) {
	*list_book = append(*list_book, Books{
		ID: uuid.New(),
		judul: judul,
		harga: uint(harga),
		kategori: kategori,
	})
	fmt.Println("Book Created")
}

func UpdateBook(list_book []Books, id uuid.UUID, judul, kategori string, harga uint) {
	for i, book := range list_book {
		if book.ID == id {
			list_book[i].judul = judul
			list_book[i].harga = harga
			list_book[i].kategori = kategori
			break
		}
	}
	fmt.Println("Book Updated")
}

func DeleteBook(list_book *[]Books, id uuid.UUID) {
	for i, book := range *list_book {
		if book.ID == id {
			*list_book = append((*list_book)[:i], (*list_book)[i+1:]...)
			break
		}
	}
	fmt.Println("Book Deleted")
}

func GetBook(list_book []Books, id uuid.UUID) Books {
	for _, book := range list_book {
		if book.ID == id {
			return book
		}
	}
	return Books{}
}

func PrintMenu() {
	fmt.Println("===BOOK MANAGEMENT===")
	fmt.Println("1. Get All Book")
	fmt.Println("2. Create a book")
	fmt.Println("3. Update a Book")
	fmt.Println("4. Delete a Book")
	fmt.Println("5. Exit")
	fmt.Println("Choose your menu : ")
}

func ReadInput(params string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(params)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func ValidateInt(input string) bool {
	val, err := strconv.Atoi(input)
	return val > 0 && err == nil
}

func main() {
	list_book := []Books{}
	loop := true
	for loop {
		PrintMenu()
		var menu int
		fmt.Scanln(&menu)
		switch menu {
			case 1:
				fmt.Println("===GET ALL BOOK===")
				fmt.Println("Total book : ", len(list_book))
				sort.Slice(list_book, func(i, j int) bool {
					return list_book[i].judul < list_book[j].judul
				})

				for idx, book := range list_book {
					fmt.Println("[#] => ", idx+1)
					fmt.Println("ID\t\t: ", book.ID)
					fmt.Println("Judul\t\t: ", book.judul)
					fmt.Println("Kategori\t: ", book.kategori)
					fmt.Println("Harga\t\t: ", book.harga)
					fmt.Println("===================")
				}
			case 2:
				fmt.Println("===CREATE A BOOK===")
				judul := ReadInput("Judul\t\t: ")
				kategori := ReadInput("Kategori\t: ")
				harga := ReadInput("Harga\t\t: ")
				for !ValidateInt(harga) {
					fmt.Println("Harga harus berupa angka")
					harga = ReadInput("Harga\t\t: ")
				}
				harga_int, _ := strconv.Atoi(harga)
				harga_uint := uint(harga_int)

				CreateBook(&list_book, judul, kategori, harga_uint)
			case 3:
				fmt.Println("===UPDATE A BOOK===")
				id := uuid.MustParse(ReadInput("ID\t\t: "))
				for GetBook(list_book, id).ID == uuid.Nil {
					fmt.Println("ID tidak ditemukan")
					id = uuid.MustParse(ReadInput("ID\t\t: "))
				}

				judul := ReadInput("Judul\t\t: ")
				kategori := ReadInput("Kategori\t: ")
				harga := ReadInput("Harga\t\t: ")
				for !ValidateInt(harga) {
					fmt.Println("Harga harus berupa angka")
					harga = ReadInput("Harga\t\t: ")
				}
				harga_int, _ := strconv.Atoi(harga)
				harga_uint := uint(harga_int)

				UpdateBook(list_book, id, judul, kategori, harga_uint)
			case 4:
				fmt.Println("===DELETE A BOOK===")
				id := uuid.MustParse(ReadInput("ID\t\t: "))
				fmt.Scanln(&id)
				DeleteBook(&list_book, id)
			case 5:
				fmt.Println("===EXIT===")
				loop = false
			default:
				fmt.Println("Menu tidak tersedia")
		}
	}
}