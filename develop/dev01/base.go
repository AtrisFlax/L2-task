package base

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
)

func PrintExactTime() error {
	const address = "0.beevik-ntp.pool.ntp.org"
	time, err := ntp.Time(address)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Can't get time: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(time)
	return nil
}
