package domain

import (
	"fmt"
)

func ToCheck(b Book, reqTitle string) error {
	if b.Title == reqTitle {
		return fmt.Errorf("This book already exists in the collection.")
	}
	return nil
}
