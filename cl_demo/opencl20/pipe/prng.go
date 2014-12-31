package main

import (
	"gocl/cl"
	"math"
)

/***
 * BEGIN - CLASS PM_PRNG
 ***/

type PM_PRNG struct {
	iv   [MAX_NTAB]cl.CL_int
	seed cl.CL_int
}

func (this *PM_PRNG) rngInit(lSeed cl.CL_int) {
	var j cl.CL_int
	var k cl.CL_int

	//xor seed with the mask to avoid its value being zero.
	lSeed ^= MASK

	//fill the shuffling buffer.
	this.iv[0] = lSeed

	for j = 1; j < MAX_NTAB; j++ {
		k = lSeed / IQ
		lSeed = IA*(lSeed-k*IQ) - IR*k

		if lSeed < 0 {
			lSeed += IM
		}

		this.iv[j] = lSeed
	}
}

func (this *PM_PRNG) rngPM(prn, ch cl.CL_int) cl.CL_int {
	var j cl.CL_int
	var k cl.CL_int
	var nrn cl.CL_int

	j = prn / NDIV
	nrn = this.iv[j+ch*NTAB]

	k = prn / IQ
	prn = IA*(prn-k*IQ) - IR*k

	if prn < 0 {
		prn += IM
	}

	this.iv[j+ch*NTAB] = prn

	return nrn
}

/***
 * END -  CLASS PM_PRNG
 ***/

/* uniform to gaussian */
func boxMuller(u [2]cl.CL_float) [2]cl.CL_float {
	var g [2]cl.CL_float

	r := cl.CL_float(math.Sqrt(-2 * math.Log(float64(u[0]))))
	theta := cl.CL_float(2.0 * PI * u[1])

	g[0] = r * cl.CL_float(math.Sin(float64(theta)))
	g[1] = r * cl.CL_float(math.Cos(float64(theta)))

	return g
}
