package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	arr := []string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}
	a := subdomainVisits(arr)
	fmt.Println("a=", a)
}

func subdomainVisits(cpdomains []string) []string {
	ret := []string{}
	mc := map[string]int{}

	for n := range cpdomains {
		count, _ := strconv.Atoi(strings.Split(cpdomains[n], " ")[0])
		domains := strings.Split(strings.Split(cpdomains[n], " ")[1], ".")
		domain := strings.Split(cpdomains[n], " ")[1]
		if len(domains) == 2 {
			tdomain := domains[1]
			_, ok := mc[tdomain]
			if !ok {
				mc[tdomain] = count
			} else {
				mc[tdomain] += count
			}
			_, ok = mc[domain]
			if !ok {
				mc[domain] = count
			} else {
				mc[domain] += count
			}
		} else if len(domains) == 3 {
			tdomain := domains[2]
			sdomain := domains[1] + "." + domains[2]
			_, ok := mc[tdomain]
			if !ok {
				mc[tdomain] = count
			} else {
				mc[tdomain] += count
			}
			_, ok = mc[sdomain]
			if !ok {
				mc[sdomain] = count
			} else {
				mc[sdomain] += count
			}
			_, ok = mc[domain]
			if !ok {
				mc[domain] = count
			} else {
				mc[domain] += count
			}
		}
	}
	for k, v := range mc {
		ret = append(ret, strconv.Itoa(v)+" "+k)
	}
	return ret
}
