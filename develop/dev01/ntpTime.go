package ntpTime

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func ntpTime() (time.Time, error) {
	//библиотека возвращает точное время либо ошибку
	ntpTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		//Если есть ошибка то пишем ее в Stderr
		_, writeStringErr := io.WriteString(os.Stderr, err.Error())
		if writeStringErr != nil {
			fmt.Printf("Ошибка записи ошибки в STDERR: %s", writeStringErr)
		}
		return ntpTime, err
	}

	return ntpTime, nil
}
