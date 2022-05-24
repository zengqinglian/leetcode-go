package alertUsingSameKeyCardThreeorMoreTimesinaOneHourPeriod

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	names := []string{"daniel", "daniel", "daniel", "luis", "luis", "luis", "luis"}
	times := []string{"10:00", "10:40", "11:00", "09:00", "11:00", "13:00", "15:00"}
	actual := alertNames(names, times)
	fmt.Println(actual)
}
