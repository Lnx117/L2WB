package ntpTime

import (
	"fmt"
	"testing"
)

func TestNtpTime(t *testing.T) {
	testCount := 5

	for i := 0; i < testCount; i++ {
		time, err := ntpTime()
		if err != nil {
			t.Errorf("Ошибка при вызове функции: %s", err)
			continue
		}
		fmt.Println(time)
	}
}
