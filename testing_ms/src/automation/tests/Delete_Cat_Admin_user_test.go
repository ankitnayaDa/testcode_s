package tests

import (
	"automation/lib"
	"automation/types"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Delete_Cat_Admin_user_test(t *testing.T) {

	Admin_user := lib.Session{types.LogInAPI, "admin", "adminpass", ""}

	fmt.Printf("\nLogin using Admin user")
	assert.Nil(t, Admin_user.Login(), "\nFailed to login with Admin User")
	fmt.Printf("\nLogin using Admin user successfully")

	fmt.Printf("\n Delete Cats from inventory ")
	assert.Nil(t, Admin_user.DeleteCatsAndCheck("Dups"), "\nFailed to delete Cat")

}
