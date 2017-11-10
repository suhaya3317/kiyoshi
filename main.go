package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

func main() {
	z := "ズン"
	d := "ドコ"

	zd := []string{z, d}
	zundoko := []string{z, z, z, z, d}
	var zzdoko []string

	rand.Seed(time.Now().UnixNano())
	for {
		zzdoko = append(zzdoko, zd[rand.Intn(2)])
		if len(zzdoko) < 5 {
			continue
		}
		if reflect.DeepEqual(zundoko, zzdoko[len(zzdoko)-5:]) {
			fmt.Println(zzdoko)
			break
		}
	}
	fmt.Println("キ・ヨ・シ！")
}
