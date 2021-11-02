Supermetrics Test Automation Assessment Automation Framework

#Test cases:

    Check for Log-in/log-out as a Admin user
    Check for Log-in/log-out as a Normal user
    List Cats when login as Admin user
    List Cats when login as Normal user
    Check for awesomeness rating and rank by calculating the sum of the ASCII character codes for the letters of the cat's name
    Check if the cat's name is exactly "James", the awesomeness is infinite
    Check if The cats are presented in descending order of awesomeness
    Check if only admin user is able to delete cats
    Check if Admin user and Normal user is able to modify cat names
    Check if users are not able to add duplicate cat names
    Check if modified cat names are same after logout operation

#Procedure To Run Test Scripts:

    go should be installed (https://golang.org/doc/install)
    GOPATH should be set Ex: export GOPATH=/<path>/testing_ms/
    Scripts can be run using below commands Ex: go test -v Login_Admin_user_List_Cat_test.go
