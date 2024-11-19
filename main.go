package main

import (
	"fmt"
)

const(
	n = "119266602355654449829905925063505208968842134526180233175453369559295213121386354196922191643162956460121034863157818925782065747622702315722816311415598427795902403081338292855204992945079164035500729457247904449955009746084103187652651319152343615831982680651688028901550481373382410546825071072631596495071"
	e = "65537"
	d = "103428568263443016946654154829474267154968550187846977620179697049498514465734959708681749848902504958336800515933159601872193184008860085322650445442942456979340537454848143806160768731421324150016528523001318085168488844188953056588564555610694680358100551402163750458463660915289999502611985739257662560673"
)

func main() {
	M1 := 58
	M2 := 384
	S1 := sign(M1)
	S2 := sign(M2)
	S3 := generatesignature(S1,S2)
	fmt.Print(verify(M1 * M2,S3))
}