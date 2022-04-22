package model

import (
	"errors"
	"fmt"
)

type ModelNumber struct {
	productCode string
	branch      string
	lot         string
}

func NewModelNumber(productCode, branch, lot string) (*ModelNumber, error) {
	if productCode == "" {
		return nil, errors.New("productCode must not be empty")
	}
	if branch == "" {
		return nil, errors.New("branch must not be empty")
	}
	if lot == "" {
		return nil, errors.New("lot must not be empty")
	}
	return &ModelNumber{productCode, branch, lot}, nil
}

func (mn *ModelNumber) ToString() string {
	return fmt.Sprintf("%s-%s-%s", mn.productCode, mn.branch, mn.lot)
}
