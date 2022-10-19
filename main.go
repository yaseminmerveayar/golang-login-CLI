package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var username string = "yasemin"
	var password string = "123"

	var usernameEntry string = ""
	var passwordEntry string = ""

	fmt.Printf("\nGiriş yapmak için kullanıcı adı ve şifrenizi giriniz \n --Beş adet deneme hakkınız bulunmaktadır-- \n\n")

prog:
	for i := 4; i >= 0; i-- {

		fmt.Print("Kullanıcı adını giriniz : ")
		fmt.Scanln(&usernameEntry)

		fmt.Print("Şifreyi giriniz : ")
		fmt.Scanln(&passwordEntry)

		date := time.Now()

		if username == usernameEntry && password == passwordEntry {
			text := []string{"Kullanıcı adı : ", usernameEntry, "Şifre : ", passwordEntry, "Tarih : ", date.Format("2006-01-02 15:04:05"), "Durum : ", "Başarılı"}
			fileWrite(text)
			fmt.Println("İşlem Başarılı")
			break prog
		} else {
			fmt.Println("--Hatalı giriş yaptınız--")
			fmt.Printf("-- %v hakkınız kaldı-- \n\n", i)

			text := []string{"Kullanıcı adı : ", usernameEntry, "Şifre : ", passwordEntry, "Tarih : ", date.Format("2006-01-02 15:04:05"), "Durum : ", "Başarısız"}
			fileWrite(text)
		}

	}

}

func fileWrite(text []string) {
	file, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)

	if err != nil {
		fmt.Println("Hata : ", err)
	} else {
		for index, t := range text {
			_, writeErr := file.WriteString(t)

			if index%2 != 0 {
				file.WriteString("\n")
			}

			if writeErr != nil {
				fmt.Println(writeErr)
			}
		}
		file.WriteString("\n")
	}
	defer file.Close()
}
