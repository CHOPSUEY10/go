package gcd

// Golden rule 1 : nama package selalu nama relative folder

type greatestCommonDivisor func(a, b int) int

// Golden rule 2 : Huruf kapital pada fungsi/struct/type/variabel artinya bisa diekspor
func HitungGCD(a, b int, gcd greatestCommonDivisor) int {

	return gcd(a, b)
}

// Golden rule 4 : tidak perlu import file jika ingin menggunakan source code dengan package yang sama (tinggal gunakan)
