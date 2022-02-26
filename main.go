package main

import (
	"fmt"
	"os"

	"github.com/laylatheraven/bookstore/helper"
)

func main() {
	args := os.Args
	var sliceOfBooks = []string{"Schindler's List", "Simyacı", "Bülbülü Öldürmek", "Deliduman", "Search Algorithms"}
	lowerCaseArgs := helper.ToLowerSlice(args)

	if len(lowerCaseArgs) == 1 {
		fmt.Printf("	Arama işlemi için girmeniz gereken komut: \"Search\" \n	Listeleme işlemi için girmeniz gereken komut: \"List\"")
		return
	}

	if len(lowerCaseArgs) == 2 {
		if lowerCaseArgs[1] == "search" {
			fmt.Printf("	Arama yapabilmek için bir kitap ismi giriniz!")
			return
		}
		if lowerCaseArgs[1] == "list" {
			fmt.Println("	Kitaplıktaki kitapların listesi:")
			for i, book := range sliceOfBooks {
				fmt.Printf("	%d. kitap : %s\n", i+1, book)
			}
			return
		}

		fmt.Printf("	Bilinmeyen bir komut girdiniz! \n")
		return
	} else {
		if lowerCaseArgs[1] == "search" {
			if bookFound := helper.Search(lowerCaseArgs[2:], sliceOfBooks); len(bookFound) > 0 {
				fmt.Println("	Aramanızla eşleşen kitapların listesi:")
				for _, book := range bookFound {
					fmt.Printf("	%s \n", book)
				}
			} else {
				fmt.Printf("	Aramanızla eşleşen bir kitap bulunmamaktadır!")
			}
			return
		}

		fmt.Printf("	Yapmak istediğiniz işlem gerçekleştirilememektedir!")
		return
	}
}
