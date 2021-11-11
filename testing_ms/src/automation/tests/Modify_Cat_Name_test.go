package tests

import (
	"automation/lib"
	"automation/types"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Modify_Cat_Name_test(t *testing.T) {

	Admin_user := lib.Session{types.LogInAPI, "admin", "adminpass", ""}

	fmt.Printf("\nLogin using Admin user")
	assert.Nil(t, Admin_user.Login(), "\nFailed to login with Admin User")
	fmt.Printf("\nLogin using Admin user successfully")

	fmt.Printf("\n Modify Cats Name ")
	assert.Nil(t, Admin_user.ModifyCatsNameAndCheck("Dups", "Loki"), "\nFailed to delete Cat")

	fmt.Printf("\n Modify Cats Name using a duplicate name ")
	assert.Error(t, Admin_user.ModifyCatsNameAndCheck("Harri", "Loki"), "\nFailed to delete Cat")

}
