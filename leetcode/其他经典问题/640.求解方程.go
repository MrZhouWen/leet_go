package 其他经典问题

import (
	"strconv"
	"unicode"
)

/**
输入: equation = "x+5-3+x=6+x-2"
输出: "x=2"
如果方程没有解，请返回"No solution"。如果方程有无限解，则返回 “Infinite solutions” 。
题目保证，如果方程中只有一个解，则 'x' 的值是一个整数。

经典的移项 解析问题：难点在于 细节的处理
T = O(n)
S = O(1)
*/
func solveEquation(equation string) string {
	factor, val := 0, 0
	i, n, sign := 0, len(equation), 1 //等式的左边系数默认为正
	for i < n {
		//1、等号分界线
		if equation[i] == '=' {
			sign = -1
			i++
			continue
		}
		//2、最后的符号：由 = 和 +，-共同判断
		s := sign
		if equation[i] == '+' {
			i++ //跳过+号
		} else if equation[i] == '-' {
			s = -s //符号取反
			i++
		}
		//3、获取数字项目 或者 x前的系数
		num, valid := 0, false
		for i < n && unicode.IsDigit(rune(equation[i])) {
			valid = true
			num = num*10 + int(equation[i]-'0')
			i++
		}
		//4、一串数字读取完，读取到x或者+，-号
		if i < n && equation[i] == 'x' { //变量
			if valid { //因为可能x的系数=1,所以要做判断系数的有效性
				s *= num
			}
			factor += s
			i++
		} else { //读完一个数值
			val += s * num
		}
	}
	//到此形成了 factor*X + val = 0的等式
	if factor == 0 {
		if val == 0 {
			return "Infinite solutions"
		} else {
			return "No solution"
		}
	}
	return "x=" + strconv.Itoa(-val/factor)
}
