package main

func main() {}

type FullName struct {
	first string
	last  string
}

func NewFullName(first, last string) *FullName {
	return &FullName{first, last}
}

// 構造体はすべてのフィールドが比較可能である場合、値は比較可能。対応する非空白フィールドが等しい場合、2つの構造体値は等しくなる。
// 今回は明示的に比較メソッドを用意。
func (fn *FullName) Equal(other *FullName) bool {
	return fn.first == other.first && fn.last == other.last
}
