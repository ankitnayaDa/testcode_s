Supermetrics Test Automation Assessment Automation Framework

#Test cases:
1. Check for Log-in/log-out as a Admin user
2. Check for Log-in/log-out as a Normal user
3. List Cats when login as Admin user
4. List Cats when login as Normal user
5. Check for awesomeness rating and rank by calculating the sum of the ASCII character codes for the letters of the cat's name
6. Check if the cat's name is exactly "James", the awesomeness is infinite
7. Check if The cats are presented in descending order of awesomeness
8. Check if only admin user is able to delete cats
9. Check if Admin user and Normal user is able to modify cat names
10. Check if users are not able to add duplicate cat names
11. Check if modified cat names are same after logout operation


#Procedure To Run Test Scripts:
1. Please use linux based os
2. go should be installed (https://golang.org/doc/install)
3. GOPATH should be set
	Ex: export GOPATH=<path>/testing_ms/
4. Scripts can be run using below commands
	Ex: go test -v Login_Admin_user_List_Cat_test.go


