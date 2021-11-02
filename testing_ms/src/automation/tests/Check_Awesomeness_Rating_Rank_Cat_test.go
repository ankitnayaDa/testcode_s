package tests

import (
  "automation/lib"
  "automation/types"
  "fmt"
  "testing"
  "github.com/stretchr/testify/assert"
)

func Test_Awesomeness_Rating_Rank_Cat_test(t *testing.T){

    Admin_user :=lib.Session{types.LogInAPI,"admin","adminpass",""}

	fmt.Printf("\nLogin using Admin user")
    assert.Nil(t, Admin_user.Login(), "\nFailed to login with Admin User")
    fmt.Printf("\nLogin using Admin user successfully")

    fmt.Printf("\nList all Cats")
    assert.Nil(t, Admin_user.ListCats(), "\nFailed to List all Cats")
    fmt.Printf("\nList all Cats")

    fmt.Printf("\nCalculate Awesomeness Rating and Rank for all the cats\n")
    Admin_user.Calculate_ASCII_from_Charater()
}