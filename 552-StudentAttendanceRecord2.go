package main

import (
	"fmt"
	"math"
)

func checkRecord2(n int) int {
	prevAdd := 0
	fmt.Println(getAs(n))
	return checkHelper(n, &prevAdd)
}

func checkHelper(n int, prevAdd *int) int {
	if n == 1 {
		return 3
	}
	val := checkHelper(n-1, prevAdd) * 2
	if n == 2 {
		*prevAdd = 2
	}
	*prevAdd += n - 2
	return val + *prevAdd
}

func getLs(n int) int {
	if n > 2 {
		return int(math.Pow(3, float64(n-3)))
	}
	return 0
}

func getAs(n int) int {
	if n > 1 {
		prev := getAs(n - 1)
		prevCenter := int(math.Min(float64(prev+4), 9))
		if n == 2 {
			prevCenter = 1
		}
		fmt.Println("p: ", prev, prevCenter)
		return prev*2 + prevCenter*2 + prevCenter
	}
	return 0
}

/*

.                                               count 2*+x
n   recs   count   LL+s   A+s     2*+x   diff   diff
1    3       3      0      0       0      0     0
2    9       8      0      1       2      2     2
3    27      19     1      7       3      1     3
4    81      43     8      33      5      2     5
5    243     94     18     131     8      3     10
6    729     200    45     484     12     4     12
7    2187    418    135    1634    18     6     18
8    6561    861    405    5295    25     7     25
9    19683   1753   1215   16715   31     6     31
10   59049   3536   3645   51868   30     -1    30

---

getAs(3) -> getAs(2) -> getAs(1)
.prev*2+prev+4  1			 0

---


.       1
.       5
.       1
.       5
.       9
.       5
.       1
.       5
.       1
.     1 5
.     5 9
.     1 5
. . 1 5 9
. 1 5 9 9
. . 1 5 9
.     1 5
.     5 9
.     1 5
.       1
.       5
.       1
.       5
.       9
.       5
.       1
.       5
.       1



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
# LLAL
- LLAA
LLAP
# LLPL
LLPA
LLPP
# LALL
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

# LLLLL
# LLLLA
# LLLLP
# LLLAL
#- LLLAA
# LLLAP
# LLLPL
# LLLPA
# LLLPP
# LLALL
#- LLALA
# LLALP
#- LLAAL
- LLAAA
- LLAAP
# LLAPL
- LLAPA
LLAPP
# LLPLL
# LLPLA
# LLPLP
# LLPAL
- LLPAA
LLPAP
# LLPPL
LLPPA
LLPPP

LALLL
- LALLA
LALLP
- LALAL
- LALAA
- LALAP
LALPL
- LALPA
LALPP
- LAALL
- LAALA
- LAALP
- LAAAL
- LAAAA
- LAAAP
- LAAPL
- LAAPA
- LAAPP
LAPLL
- LAPLA
LAPLP
- LAPAL
- LAPAA
- LAPAP
LAPPL
- LAPPA
LAPPP


LLP
LLP
LLP
LLP
LLP
LLP
LLP
LLP
LLP
LLP
LLP
LLP
LLP
LLP
LLP
LLP
LLP
LLP
LLP
LLP
LLP
LLP
LLP
LLP
LLP
LLP
LLP

ALLLL
- ALLLA
ALLLP
- ALLAL
- ALLAA
- ALLAP
ALLPL
- ALLPA
ALLPP
- ALALL
- ALALA
- ALALP
- ALAAL
- ALAAA
- ALAAP
- ALAPL
- ALAPA
- ALAPP
ALPLL
- ALPLA
ALPLP
- ALPAL
- ALPAA
- ALPAP
ALPPL
- ALPPA
ALPPP

- AALLL
- AALLA
- AALLP
- AALAL
- AALAA
- AALAP
- AALPL
- AALPA
- AALPP
- AAALL
- AAALA
- AAALP
- AAAAL
- AAAAA
- AAAAP
- AAAPL
- AAAPA
- AAAPP
- AAPLL
- AAPLA
- AAPLP
- AAPAL
- AAPAA
- AAPAP
- AAPPL
- AAPPA
- AAPPP

*/
