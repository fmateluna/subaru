package main

import (
	"crypto/sha1"
	"flag"
	"fmt"
	"subaru/webscraping"
)

func Hashstr(Txt string) string {
	h := sha1.New()
	h.Write([]byte(Txt))
	bs := h.Sum(nil)
	sh := string(fmt.Sprintf("%x\n", bs))
	return sh
}

func main() {
	//"JF1BN9LC2KG024958" https://snaponepc.com/epc/#/
	var vin string
	flag.StringVar(&vin, "vin", "JF1BN9LC2KG024958", "VIN A BUSCAR")
	flag.Parse()
	subaruBot := webscraping.BotSubaru{}
	subaruBot.User = "SBRCHL4090d"
	subaruBot.Pass = "fTAWHJQk!Nyq4V2n3zFPEp"
	subaruBot.Init(vin)

}
