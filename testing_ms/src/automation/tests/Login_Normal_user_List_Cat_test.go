package tests

import (
	"automation/lib"
	"automation/types"
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_Normal_Login_user_List_Cat_test(t *testing.T){

	Local_user :=lib.Session{types.LogInAPI,"user","helloworld"}

	fmt.Printf("\nLogin using Normal user")
	assert.Nil(t, Local_user.Login(), "\nFailed to login with Normal User")
	fmt.Printf("\nLogin using Normal user successfully")

}