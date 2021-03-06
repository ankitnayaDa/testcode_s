package tests

import (
	"automation/lib"
	"automation/types"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Normal_Login_user_List_Cat_test(t *testing.T) {

	Local_user := lib.Session{types.LogInAPI, "user", "helloworld", ""}

	fmt.Printf("\nLogin using Normal user")
	assert.Nil(t, Local_user.Login(), "\nFailed to login with Normal User")
	fmt.Printf("\nLogin using Normal user successfully")

	fmt.Printf("\nList all Cats")
	assert.Nil(t, Local_user.ListCats(), "\nFailed to List all Cats")
	fmt.Printf("\nList all Cats")
}
