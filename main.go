package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	conso   = []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "r", "s", "t", "v", "w", "x", "y", "z"}
	vocal   = []string{"a", "e", "i", "o", "u"}
	spchars = []string{"!", "@", "#", "$", "%", "^", "*", "&", "*", "-", "+", "?"}
)

func pwgen(pwlen int, spchar bool) string {
	var password string

	if pwlen < 12 {
		pwlen = 12
	}

	if pwlen%2 == 1 {
		pwlen += 1
	}

	length := pwlen - 2
	max := length / 2

	for i := 1; i <= max; i++ {
		password += conso[rand.Intn(len(conso))]
		password += vocal[rand.Intn(len(vocal))]
	}

	if spchar {
		rnd_char := rand.Intn(len(spchars))
		rnd_place := rand.Intn(len(password))
		password = password[:rnd_place] + spchars[rnd_char] + password[rnd_place+1:]
	}

	rnd_place := rand.Intn(len(password))
	password = password[:rnd_place] + strings.ToUpper(password[rnd_place:rnd_place+1]) + password[rnd_place+1:]
	password += fmt.Sprint(rand.Intn(89) + 10)

	return password
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	length := flag.Int("l", 12, "Password length")
	num := flag.Int("n", 1, "Number of generated password")
	special := flag.Bool("s", false, "Special character")
	flag.Parse()

	if *length < 12 {
		fmt.Fprintln(os.Stderr, "Minimal password length is 12.")
		os.Exit(1)
	}

	if *length%2 == 1 {
		*length += 1
	}

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < *num; i++ {
		fmt.Println(pwgen(*length, *special))
	}
}
