package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	nums := make([]*big.Int, 0)
	for i := 0; i < 4; i++ {
		scanner.Scan()
		count, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, big.NewInt(int64(count)))
	}
	n_n := big.NewInt(int64(n)).Mul(big.NewInt(int64(n)), big.NewInt(int64(n)))
	a_b_c_d := new(big.Int)
	a_b_c_d.Mul(jiecheng(nums[0]), jiecheng(nums[1])).Mul(a_b_c_d, jiecheng(nums[2])).Mul(a_b_c_d, jiecheng(nums[3]))
	ans := n_n.Div(jiecheng(n_n), a_b_c_d)
	ans = ans.Mod(ans, big.NewInt(998244353))
	fmt.Println(ans)
}

func jiecheng(n *big.Int) *big.Int {
	if n.Int64() == 1 || n.Int64() == 0 {
		return big.NewInt(1)
	}
	return n.Mul(n, jiecheng(big.NewInt(n.Int64()-1)))
}
