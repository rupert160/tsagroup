package main

import (
        "fmt"
        "github.com/dongri/phonenumber"
        "regexp"
)

type contact struct {
	Full_Name string `json:"full_name"`
	Email string `json:"title"`
        Phone_Numbers []string `json:"phone_numbers"`
}

var contacts = []contact{
 	{Full_Name: "Alex Bell", Email: "alex@bell-labs.com", Phone_Numbers:[]string{"03 8578 6688", "1800728069"}},
	{Full_Name: "fredrik IDESTAM", Phone_Numbers: []string{"+6139888998"}},
	{Full_Name: "radia perlman", Email: "rperl001@mit.edu", Phone_Numbers: []string{"(03) 9333 7119", "0488445688", "+61488224568"}},
}

func e164(token string) string{
        matched2, err2 := regexp.MatchString(`\+.*`, token)
        matched3, err3 := regexp.MatchString(`\+(61).*`, token)
        if matched2 && !matched3 {
                return "Invalid:International"
        } else if err2 != nil || err3 != nil {return "Error:International"}

        matched, err := regexp.MatchString(`^1800.*`, token)
        if matched {
                reClean := regexp.MustCompile(`([^\d*])`)
                return fmt.Sprintf("61%s", reClean.ReplaceAll([]byte(token), []byte("")))
        }else if err != nil{
                return ""
        }

        matched4, err := regexp.MatchString(`^61.*`, token)
        if matched4 {
                return token
        }else{
                return fmt.Sprintf("%s", phonenumber.ParseWithLandLine(token,"AU"))
        }
}

func reformat_initial_data(){
        var new_contacts []contact
        for _,
        value := range contacts {
                var value3 []string
                for _,
                value2 := range value.Phone_Numbers{
                        value3 = append(value3, e164(value2))
                }
                new_contacts = append(new_contacts, contact{value.Full_Name, value.Email, value3})
        }
        contacts = new_contacts
}
func print_dataset(){
        for _,
        value := range contacts {
                fmt.Printf("%s\n", value)
        }
}

func main() {
	route := initiateRoutes();
        reformat_initial_data()
	route.Run(":3004")
}
