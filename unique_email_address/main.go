package main

import (
	"fmt"
	"strings"
)

func numUniqueEmail(emails []string) int {
	ans, uniqueAddressMap := 0, make(map[string]struct{})
	for _, email := range emails {
		address := uniqAddress(email)
		if _, ok := uniqueAddressMap[address]; !ok {
			ans++
			uniqueAddressMap[address] = struct{}{}
		}
	}
	return ans
}

func uniqAddress(email string) string {
	parts := strings.Split(email, "@")
	local, domain := parts[0], parts[1]

	//handling rule where '.' is ignored
	local = strings.ReplaceAll(local, ".", "")
	//handling + rule
	plusIndex := strings.Index(local, "+")
	if plusIndex >= 0 {
		local = local [:plusIndex]
	}
	return local + "@" + domain
}

func main () {
	emails := []string {"test.email+alex@leetcode.com","test.e.mail+bob.cathy@leetcode.com","testemail+david@lee.tcode.com"}
	fmt.Println("number of unique emails are: ", numUniqueEmail(emails))
	
}