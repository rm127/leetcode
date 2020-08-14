package main

import (
	"fmt"
	"math"
)

func checkRecord2(n int) int {
	fmt.Println("gA: ", getAs(n))
	if n == 0 {
		return 0
	}
	return int(math.Pow(3, float64(n))) - getLs(n) - getAs(n)
}

func getLs(n int) int {
	if n > 2 {
		return int(math.Pow(3, float64(n-3)))
	}
	return 0
}

func getAs(n int) int {
	if n == 1 {
		return 1
	}
	prev := getAs(n - 1)
	count := int(math.Pow(3, float64(n-2)))
	prevCenter := int(math.Min(float64(prev+count*4), float64(count*9)))
	fmt.Println("p: ", prev, count)
	return prev*2 + prevCenter
}

/*

.                                               count 2*+x
n   recs   count   LL+s   A+s     2*+x   diff   diff
1    3       3      0      0       0      0     0
2    9       8      0      1       2      2     2
3    27      19     1      7       3      1     3
4    81      43     9      33      5      2     5
5    243     94     51     131     8      3     10
6    729     200    ?      484     12     4     12
7    2187    418    ?      1634    18     6     18
8    6561    861    ?      5295    25     7     25
9    19683   1753   ?      16715   31     6     31
10   59049   3536   ?      51868   30     -1    30

---

27/3 = 9

9 = L, 3 = L, 1 = L
       3 = A, 1 = A

9 = A, 3 = L, 1 = A
	   3 = A
	   3 = P, 1 = A
9 = P, 3 = A, 1 = A

---


A:
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
.   1 5 9
. 1 5 9 9
.   1 5 9
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


L:
.
.
.       33
.       9
.       9
.   1 7 7
. . . 1 1
.   . 1 1
.       7
.       1
.       1
.
.
.



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

# LLL
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
# LPLL
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

# LALLL
#- LALLA
#LALLP
#- LALAL
- LALAA
- LALAP
# LALPL
- LALPA
LALPP
#- LAALL
- LAALA
- LAALP
- LAAAL
- LAAAA
- LAAAP
- LAAPL
- LAAPA
- LAAPP
# LAPLL
- LAPLA
LAPLP
- LAPAL
- LAPAA
- LAPAP
LAPPL
- LAPPA
LAPPP

# LPLLL
# LPLLA
# LPLLP
# LPLAL
- LPLAA
LPLAP
# LPLPL
LPLPA
LPLPP
# LPALL
- LPALA
LPALP
- LPAAL
- LPAAA
- LPAAP
LPAPL
- LPAPA
LPAPP
# LPPLL
LPPLA
LPPLP
LPPAL
- LPPAA
LPPAP
LPPPL
LPPPA
LPPPP

# ALLLL
#- ALLLA
# ALLLP
#- ALLAL
- ALLAA
- ALLAP
# ALLPL
- ALLPA
ALLPP
#- ALALL
- ALALA
- ALALP
- ALAAL
- ALAAA
- ALAAP
- ALAPL
- ALAPA
- ALAPP
# ALPLL
- ALPLA
ALPLP
- ALPAL
- ALPAA
- ALPAP
ALPPL
- ALPPA
ALPPP

#- AALLL
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

APLLL
- APLLA
ALPLP
- APALL
- APLPA
- APLAP
APLPL
- APLPA
APLPP
- APALL
- APALA
APALP
- APAAL
- APAAA
- APAAP
- APAPL
- APAPA
- APAPP
APPLL
- APPLA
APPLP
- APPAL
- APPAA
- APPAP
APPPL
- APPPA
APPPP

# PLLLL
# PLLLA
# PLLLP
# PLLAL
- PLLAA
PLLAP
# PLLPL
PLLPA
PLLPP
# PLALL
- PLALA
PLALP
- PLAAL
- PLAAA
- PLAAP
PLAPL
- PLAPA
PLAPP
# PLPLL
PLPLA
PLPLP
PLPAL
- PLPAA
PLPAP
PLPPL
PLPPA
PLPPP

# PALLL
- PALLA
PALLP
- PALAL
- PALAA
- PALAP
PALPL
- PALPA
PALPP
- PAALL
- PAALA
- PAALP
- PAAAL
- PAAAA
- PAAAP
- PAAPL
- PAAPA
- PAAPP
PAPLL
- PAPLA
PAPLP
- PAPAL
- PAPAA
- PAPAP
PAPPL
- PAPPA
PAPPP

PPLLL
PPLLA
PPLLP
PPLAL
- PPLAA
PPLAP
PPLPL
PPLPA
PPLPP
PPALL
- PPALA
PPALP
- PPAAL
- PPAAA
- PPAAP
PPAPL
- PPAPA
PPAPP
PPPLL
PPPLA
PPPLP
PPPAL
- PPPAA
PPPAP
PPPPL
PPPPA
PPPPP

A = 131

L=33
A=9
P=9
*/
