package leedcode

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
)

func jamOne() {
	var t, n, m int
	var tmp string

	fmt.Scanf("%d", &t)
	for i := 1; i <= t; i++ {
		fmt.Scanf("%d", &n)
		fmt.Scanf("%d", &m)
		fmt.Printf("Case #%d:\n", i)

		for j := 1; j <= 2*n+1; j++ {
			tmp = ""
			for k := 1; k <= 2*m+1; k++ {
				r := j % 2
				l := k % 2

				if (j == 1 || j == 2) && (k == 1 || k == 2) {
					tmp += "."
					continue
				}

				if r == 0 && l == 0 {
					tmp += "."
				} else if r == 1 && l == 1 {
					tmp += "+"
				} else if r == 1 && l == 0 {
					tmp += "-"
				} else {
					tmp += "|"
				}

			}
			fmt.Println(tmp)
		}
	}

}

func jamTwo() {
	var t int
	var arr = [3][4]int{}
	fmt.Scanf("%d", &t)
	for i := 1; i <= t; i++ {
		var m = [4]int{1000000, 1000000, 1000000, 1000000}
		ret := make([][]int, 0)
		for j := 0; j < 3; j++ {
			for k := 0; k < 4; k++ {
				fmt.Scanf("%d", &arr[j][k])

			}
		}

		for k := 0; k < 4; k++ {
			for j := 0; j < 3; j++ {
				if m[k] > arr[j][k] {
					m[k] = arr[j][k]
				}
			}
		}

		backtrackJam(&ret, []int{0, 0, 0, 0}, 0, m, 0, m[0]+m[1]+m[2]+m[3])
		fmt.Printf("Case #%d: ", i)
		if len(ret) < 1 {
			fmt.Printf("IMPOSSIBLE\n")
		} else {
			fmt.Printf("%d %d %d %d\n", ret[0][0], ret[0][1], ret[0][2], ret[0][3])
		}

	}

}

func backtrackJam(ret *[][]int, tmp []int, start int, max [4]int, sum int, rem int) {
	if len(*ret) > 0 {
		return
	}
	if start < 0 {
		return
	}

	if sum == 1000000 {
		out := make([]int, 4)
		copy(out, tmp)
		*ret = append(*ret, out)
		return

	}
	if start == 4 {
		return
	}

	for i := max[start]; i >= 0; i-- {
		tmp[start] = i

		if sum+tmp[start] > 1000000 {
			i = 1000000 - sum + 1
			continue
		}

		if sum+i+rem-max[start] < 1000000 {
			return
		}

		sum += tmp[start]
		rem = rem - max[start]
		start++
		backtrackJam(ret, tmp, start, max, sum, rem)
		start--
		sum -= tmp[start]
		rem = rem + max[start]
	}
}

//use hash to store to prevent TLE
func jamFour() {
	scanner := bufio.NewScanner(os.Stdin) //from standand input for keyboard or terminal
	infn := ""
	if infn == "" && len(os.Args) > 1 { //from file path input
		infn = os.Args[1]
	}
	if infn != "" {
		f, e := os.Open(infn)
		if e != nil {
			panic(e)
		}

		scanner = bufio.NewScanner(f)
	}
	scanner.Split(bufio.ScanWords)
	scanner.Buffer(make([]byte, 1024), 1000000000)

	var t, n int
	t = getInt(scanner)
	var fun []int
	var tmp int
	for i := 1; i <= t; i++ {
		n = getInt(scanner)
		fun = make([]int, n)
		hash := make(map[int][]int)
		res := 0
		index := 0
		for j := 0; j < n; j++ {
			fun[j] = getInt(scanner)
		}
		for j := 0; j < n; j++ {
			tmp = getInt(scanner)
			hash[tmp] = append(hash[tmp], j)

			if tmp > index {
				index = tmp
			}
		}

		for index >= 0 {
			min := 1000000001
			//_, j := range hash[index] is right, direct j:=range hash[index] is wrong
			// range will return key and value, unlike for only value returned
			for j := 0; j < len(hash[index]); j++ {
				res += fun[hash[index][j]]
				if min > fun[hash[index][j]] {
					min = fun[hash[index][j]]
				}
			}

			if index != 0 && len(hash[index]) > 0 {
				res -= min
				if min > fun[index-1] {
					fun[index-1] = min
				}
			}

			index--
		}

		fmt.Printf("Case #%d: %d\n", i, res)
	}
}

// No.5 include:the mind of sample and calculate when yield a poor estimate. a small subset has
// a degree much higher or much lower than the median degree

func TwistyLittlePassages() {
	var t, n, k, r, p, res int
	fmt.Scanf("%d", &t)

	for i := 1; i <= t; i++ {
		sum := 0 //point num sum
		cnt := 0 //point num
		unknow := 0
		res = 0
		fmt.Scanf("%d", &n)
		fmt.Scanf("%d", &k)
		hash := make(map[int]int)

		fmt.Scanf("%d", &r)
		fmt.Scanf("%d", &p)
		hash[r] = p
		for j := 0; j < k/2; j++ {
			fmt.Printf("T %d\n", rand.Intn(n)+1)

			fmt.Scanf("%d", &r)
			if r == -1 {
				return
			}
			fmt.Scanf("%d", &p)

			hash[r] = p
			sum += p
			cnt++

			fmt.Printf("W\n")
			fmt.Scanf("%d", &r)
			if r == -1 {
				return
			}
			fmt.Scanf("%d", &p)
			hash[r] = p
		}

		for j := 1; j <= n; j++ {
			if _, ok := hash[j]; ok {
				res += hash[j]
			} else {
				unknow++
			}
		}
		res += int(float64(unknow) * float64(sum) / float64(cnt))

		fmt.Printf("E %d\n", res/2)

	}

}

//the third and forth question refer something about input efficiently for big data stream
//use bufio package to prevent TLE
func diceStraight() {
	scanner := bufio.NewScanner(os.Stdin) //from standand input for keyboard or terminal
	infn := ""
	if infn == "" && len(os.Args) > 1 { //from file path input
		infn = os.Args[1]
	}
	if infn != "" {
		f, e := os.Open(infn)
		if e != nil {
			panic(e)
		}

		scanner = bufio.NewScanner(f)
	}
	scanner.Split(bufio.ScanWords)
	scanner.Buffer(make([]byte, 1024), 1000000000)

	var t, n int
	t = getInt(scanner)

	for i := 1; i <= t; i++ {
		res := 0
		n = getInt(scanner)

		arr := make([]int, n)
		for j := 0; j < n; j++ {
			arr[j] = getInt(scanner)
		}

		sort.Ints(arr)

		for j := 0; j < n; j++ {
			if res+1 <= arr[j] {
				res++
			}
		}

		fmt.Printf("Case #%d: %d\n", i, res)
	}

}
func getInt(scanner *bufio.Scanner) int {
	num, error := strconv.Atoi(getStr(scanner))
	if error != nil {
		panic(error)
	}
	return num

}

func getStr(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

// NO.round-A 1 greed
func doubleOrOneThing() {
	scanner := bufio.NewScanner(os.Stdin)
	infn := ""
	if infn == "" && len(os.Args) > 1 {
		infn = os.Args[1]
	}
	if infn != "" {
		f, e := os.Open(infn)
		if e != nil {
			panic(e)
		}

		scanner = bufio.NewScanner(f)
	}
	scanner.Split(bufio.ScanWords)
	scanner.Buffer(make([]byte, 1024), 1000000000)

	var t int
	var str string
	t = getInt(scanner)

	for i := 1; i <= t; i++ {
		str = getStr(scanner)
		res := ""
		if len(str) == 1 {
			res = str
		}
		for i := 1; i < len(str); i++ {
			if str[i-1] < str[i] {
				res += string(str[i-1]) + string(str[i-1])
			} else if str[i-1] == str[i] {
				mul := 1
				var j int
				for j = i; j < len(str) && str[j] == str[j-1]; j++ {
					mul++
				}

				if j >= len(str) {
					res += str[i-1:]
					i = j - 1
					continue
				}
				if str[j-1] < str[j] {
					res += str[i-1:j] + str[i-1:j]
				} else {
					res += str[i-1 : j]
				}

				i = j

			} else {
				res += string(str[i-1])
			}
		}
		if len(str) > 1 && str[len(str)-1] != str[len(str)-2] {
			res += string(str[len(str)-1])
		}
		fmt.Printf("Case #%d: %s\n", i, res)
	}
}

//No round A-2
//as asked the question needs to be solved in polynomial time,
// so don't use brute force(the k power of n)
// any number can be expressed by binary,
func equalSum() {
	scanner := bufio.NewScanner(os.Stdin)
	infn := ""
	if infn == "" && len(os.Args) > 1 {
		infn = os.Args[1]
	}
	if infn != "" {
		f, e := os.Open(infn)
		if e != nil {
			panic(e)
		}

		scanner = bufio.NewScanner(f)
	}
	scanner.Split(bufio.ScanWords)
	scanner.Buffer(make([]byte, 1024), 1000000000)

	var t, n int
	t = getInt(scanner)

	for i := 1; i <= t; i++ {
		n = getInt(scanner)
		if n != 100 {
			return
		}
		sum := 0
		arr := make([]int, 2*n)
		for j := 0; j < 30; j++ {
			arr[j] = 1 << uint(j)
			sum += arr[j]
		}
		for j := 30; j < n; j++ {
			arr[j] = int(1e9) - (n - 1 - j)
			sum += arr[j]
		}

		for j := 0; j < n; j++ {
			fmt.Printf("%d ", arr[j])
		}
		fmt.Printf("\n")

		for j := n; j < 2*n; j++ {
			arr[j] = getInt(scanner)
			sum += arr[j]
		}
		if sum%2 != 0 {
			return
		}

		sum = sum / 2
		res := make([]int, 0)
		bound := 1 << uint(30)
		for j := 2*n - 1; sum >= bound; j-- {
			if sum > arr[j] {
				res = append(res, arr[j])
				sum = sum - arr[j]
			}
		}
		for j := 0; j < 30; j++ {
			if sum>>uint(j)&1 == 1 { //&&for bool type. 10&10=10
				res = append(res, arr[j])
			}
		}

		for j := 0; j < len(res); j++ {
			fmt.Printf("%d ", res[j])
		}
		fmt.Printf("\n")
	}
}
