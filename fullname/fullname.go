package fullname

import (
	"errors"
	"fmt"
	"regexp"
)

// 構造体をプライベートにすると、他パッケージから参照できない。（変数宣言時の型としても参照できない）
// なのでプライベートにする場合はパッケージ内のみで使われるケースに限るのではないか
// 関数の戻り値を代入した変数がprivateな構造体型と推論される場合はコンパイルは通る
// interfaceを使うことで構造体を隠蔽しつつ他パッケージに公開することができる→Equalメソッドの引数に実装となる構造体を渡す方法がわからない
// Publicな構造体でフィールドがPrivateになっている場合、ゼロ値の構造体インスタンス以外は作成できない。（フィールドは操作できない）

// type FullName interface {
// 	Equal()
// }
type FullName struct {
	first string
	last  string
}

func NewFullName(first, last string) (*FullName, error) {
	if first == "" {
		return nil, errors.New("first name is required")
	}

	if last == "" {
		return nil, errors.New("last name is required")
	}

	if !validateName(first) {
		return nil, errors.New(fmt.Sprintf("first name contains invalid character: %s", first))
	}

	if !validateName(last) {
		return nil, errors.New(fmt.Sprintf("last name contains invalid character: %s", last))
	}

	return &FullName{first, last}, nil
}

func validateName(value string) bool {
	r := regexp.MustCompile("[0-9!-/:-@¥[-`{-~]+")
	return !r.MatchString(value)
}

// 構造体はすべてのフィールドが比較可能である場合、値は比較可能。対応する非空白フィールドが等しい場合、2つの構造体値は等しくなる。
// 今回は明示的に比較メソッドを用意。
func (fn *FullName) Equal(other *FullName) bool {
	return fn.first == other.first && fn.last == other.last
}
