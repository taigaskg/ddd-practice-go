package main

import "ddd-practice-go/username"

func main() {
	// n := new(map[string]string)
	// m := make([]int, 10)
	// fmt.Println("n: ", n, ",", "m: ", m)
	// fmt.Println(&n)
	// // fmt.Println(len(*n), cap(*n))
	// // nn := append(*n, 1)
	// // n = &nn
	// // nn := *n
	// // nn["aaaa"] = "aaaa"
	// fmt.Println(n)
	// h := fullname.FullName{}
	// fmt.Println(h == fullname.FullName{})

	// myMoney, err := money.New(1000, "JPY")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// allowance, err := money.New(3000, "usd")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// res, err := myMoney.Add(allowance)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(res)

	// mn, err := modelnumber.New("ppppp", "bbbbb", "lllll")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(mn.ToString())
}

func CreateUser(name string) *User {
	return &User{Id: name} // フィールド名が間違っているがコンパイルエラーにはならない
}
func CreateUser2(name username.UserName) *User2 {
	// return &User{Id: name} // コンパイルエラー
	return &User2{Name: name} // これで間違った代入を防ぐことができる

}

type User struct {
	Id   string
	Name string
}

type User2 struct {
	Id   string
	Name username.UserName
}
