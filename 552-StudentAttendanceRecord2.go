package main

import (
	"fmt"
	"math"
)

func checkRecord2(n int) int {
	total := int(math.Pow(3, float64(n)))
	ls := getLs(n) + getLs(n-1)*2
	as := getAs(n-1) + getAs(n-2)*2 + getAs(n-3)*4
	fmt.Println("L: ", ls, "A: ", as)
	return total - ls - as
}

func getLs(n int) int {
	if n > 2 {
		return int(math.Pow(3, float64(n-3)))
	}
	return 0
}

func getAs(n int) int {
	if n < 2 {
		return int(math.Max(float64(n), 0))
	} else if n == 2 {
		return 5
	}
	return int(math.Pow(3, float64(n-1)))
}

/*

n   recs  count   LL+s   A+s     2*+x
1    3      3      0      0       0
2    9      8      0      1       2
3    27     19     1      7       3
4    81     43     5      33      5
5    243    94     15     134     8
6    729    200    45     484     12

---

## LL logic

func getLs(n int) {
	if	n > 2 {
		return math.Pow(3, n-3)
	}
	return 0
}

Ls := getLs(n) + getLs(n-1)*2


## A logic



# A steps:
0    1    2    3    4     5
0 -> 1 -> 5 -> 9 -> 27 -> 81

# nonA:
0 -> 1 -> 7
# A:
1 -> 5 -> 19


#1
A1
#2
5
#3


n + (n-1)*4 + (n-2)*4


#2

LL
LA
LP
AL
- AA
AP
PL
PA
PP


#3

= LLL
LLP
LLA
LAL
LAP
- LAA
LPL
LPP
LPA

ALL
- ALA
ALP
- AAL
- AAA
- AAP
APL
- APA
APP

PLL
PLA
PLP
PPL
PPP
PPA
PAL
PAP
- PAA


#4

# LLLL
# LLLA
# LLLP
LLAL
- LLAA
LLAP
LLPL
LLPA
LLPP
LALL
- LALA
LALP
- LAAL
- LAAA
- LAAP
LAPL
- LAPA
LAPP
LPLL
LPLA
LPLP
LPAL
- LPAA
LPAP
LPPL
LPPA
LPPP

# ALLL
- ALLA
ALLP
- ALAL
- ALAA
- ALAP
ALPL
- ALPA
ALPP
- AALA
- AALP
- AAAL
- AAAA
- AAAP
- AAPL
- AAPA
- AAPP
- AALL
APLL
- APLA
APLP
- APAL
- APAA
- APAP
APPL
- APPA
APPP

5+9+5

P
P
P ... same as L


#5

LLLLL
LLLL
LLLL
LLL
LLL
LLL
LLL
LLL
LLL
LL
LL

ALLLL
ALLL
ALLL
ALL
ALL
ALL
ALL
ALL
ALL
ALA
ALA
ALA


*/
