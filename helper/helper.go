package helper

import (
	"strings"
)

/*
 * Argüman olarak verilen slice'ın tüm elemanlarını lowercase'e
 * çevirerek yeni bir slice döndürür.
 */
func ToLowerSlice(slc []string) []string {
	loweredSlc := []string{}
	for i := 0; i < len(slc); i++ {
		loweredSlc = append(loweredSlc, strings.ToLower(slc[i]))
	}
	return loweredSlc
}

/*
 * Arama argümanları ile eşleşen kitapların listesini (slice)
 * döndürür.
 * Hashmap kullanılarak daha önce bir argüman ile eşleşen bir
 * kitap üzerinden tekrar arama yapılması engellenir.
 */
func Search(slcArg []string, slcBook []string) []string {
	bookLookup := make(map[string]int)
	resultSlc := []string{}

	for _, arg := range slcArg {
		for j, book := range slcBook {
			if _, exists := bookLookup[book]; exists {
				continue
			}

			if strings.Contains(strings.ToLower(book), arg) {
				bookLookup[book] = j
				resultSlc = append(resultSlc, book)
			}
		}
	}
	return resultSlc
}
