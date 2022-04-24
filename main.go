package main

import (
	"ddd-practice-go/domain/repository"
	"ddd-practice-go/hogehoge.go"
	"ddd-practice-go/infrastructure/mysql"
	"ddd-practice-go/usecase"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	userName := "user name"

	db, err := mysql.OpenDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var userRepository repository.UserRepository = mysql.NewUserRepository(db)
	uu := usecase.NewUserUsecase(&userRepository)
	err = uu.CreateUser(userName)
	if err != nil {
		// TODO: error
	}
}

// func CreateUser(name string) *user.User {
// 	// return &user.User{id: name} // フィールド名が間違っているがコンパイルエラーにはならない
// }

// func CreateUser2(name username.UserName) *User2 {
// 	// return &User{Id: name} // コンパイルエラー
// 	return &User2{Name: name} // これで間違った代入を防ぐことができる

// }

// type User struct {
// 	Id   string
// 	Name string
// }

// type User2 struct {
// 	Id   string
// 	Name username.UserName
// }

func HogeHoge() {
	// var i hogehoge.I = s{} // ?????TODO:
}

// TODO: 別件のテストコード。
type s struct {
	hogehoge.UnimplementedI
}

func (s *s) M1() (string, error) {
	return "aaa", nil
}

func (s *s) M2() (int, error) {
	return 100, nil
}
