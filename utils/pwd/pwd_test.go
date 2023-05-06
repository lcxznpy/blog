package pwd

import (
	"fmt"
	"testing"
)

func TestHashPwd(t *testing.T) {
	fmt.Println(HashPwd("1234"))
}

func TestCheckPwd(t *testing.T) {
	fmt.Println(CheckPwd("$2a$04$befxr0LhJT1wsuJepiQRFucPD6LzNZq7avyd8NdKBqelHGKMaZgNa", "1234"))
}
