package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Sdat struct {
	Email    string
	Password string
}

type data struct {
	lis []Sdat
}

func Getdata(filename string) data {
	var dSdat Sdat
	var d data
	f, err := os.Open(filename)
	if err != nil {
		fmt.Print(err)
	}
	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		//var temp1 [2]string
		//fileLines = append(fileLines, fileScanner.Text())
		temp1 := strings.Split(fileScanner.Text(), " ")
		dSdat.Email = temp1[0]
		dSdat.Password = temp1[1]
		d.lis = append(d.lis, dSdat)
	}

	return d

}

func (d data) CheckPass(email, password string) (bool, string) {
	for _, r := range d.lis {
		if r.Email == email {
			if r.Password == password {
				return true, "found"
			}

		}
		return false, "incorrect password"

	}
	return false, "user not found"
}
