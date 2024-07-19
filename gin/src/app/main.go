package main

import (
        "fmt"
        "github.com/dongri/phonenumber"
        "regexp"
	"strconv"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"time"
)

type contact_string struct {
	Full_Name string `json:"full_name"`
	Email string `json:"title"`
        Phone_Numbers []string `json:"phone_numbers"`
}

var contacts_string = []contact_string{
 	{Full_Name: "Alex Bell", Email: "alex@bell-labs.com", Phone_Numbers:[]string{"03 8578 6688", "1800728069"}},
	{Full_Name: "fredrik IDESTAM", Phone_Numbers: []string{"+6139888998"}},
	{Full_Name: "radia perlman", Email: "rperl001@mit.edu", Phone_Numbers: []string{"(03) 9333 7119", "0488445688", "+61488224568"}},
}

type contact struct {
	Full_Name string `json:"full_name"`
	Email string `json:"title"`
        Phone_Numbers []int64 `json:"phone_numbers"`
}

type Contact struct {
    Id     int64
    Name   string
    Email string
    Phone_Numbers []int64
}

func (u Contact) String() string {
    return fmt.Sprintf("User<%d %s %v>", u.Id, u.Name, u.Email, u.Phone_Numbers)
}

var contacts []contact
var Contacts []Contact

func convert_stoi(s string) int64 {
	if s == "" {return 0}
	i, err := strconv.Atoi(s)
	if err != nil {
        	panic(err)
    	}
	return int64(i)
}

func e164(token string) int64{
        matched2, err2 := regexp.MatchString(`\+.*`, token)
        matched3, err3 := regexp.MatchString(`\+(61).*`, token)
        if matched2 && !matched3 { return 0 } 
	if err2 != nil || err3 != nil { return 0 }

        matched, err := regexp.MatchString(`^1800.*`, token)
        if matched {
                reClean := regexp.MustCompile(`([^\d*])`)
                return convert_stoi(fmt.Sprintf("61%s", reClean.ReplaceAll([]byte(token), []byte(""))))
        }else if err != nil{
                return 0
        }

	matched4, err := regexp.MatchString(`^61.*`, token)
        if matched4 { return convert_stoi(token) }

        matched5, err := regexp.MatchString(`^\+61.*`, token)
        if matched5 { return convert_stoi(token[1:]) }

        return convert_stoi(fmt.Sprintf("%s", phonenumber.ParseWithLandLine(token,"AU")))

}

func convert_contact(cs contact_string) contact {
        var numbers []int64
        for _,
        number_string := range cs.Phone_Numbers{
                numbers = append(numbers, e164(number_string))//todo
        }
        return contact{cs.Full_Name, cs.Email, numbers} //todo
}

func reformat_initial_data(db *pg.DB){
        var new_contacts []contact //todo
        for _,
        cs := range contacts_string {
		c := convert_contact(cs)
                new_contacts = append(new_contacts, c)
		_, err := db.Model(&Contact{
    		    Name:  c.Full_Name,
    		    Email: c.Email,
    		    Phone_Numbers: c.Phone_Numbers,
    		}).Insert()
    		if err != nil {
    		    panic(err)
    		}
        }
        contacts = new_contacts //todo
	contacts_string = nil
}

func createSchema(db *pg.DB) error {
    models := []interface{}{
        (*Contact)(nil),
    }

    for _, model := range models {
	//todo risk of infinite loop
	counter := 0
        for {
            err := db.Model(model).CreateTable(&orm.CreateTableOptions{ Temp: false })
            if err != nil {
                time.Sleep(5 * time.Second)
                // Retry the operation
		counter++
		if counter >5 {return err}
                continue
            }
            // If no error, break the loop
            break
        }
    }
    return nil
}

func Model() {
    db := pg.Connect(&pg.Options{
        //User: "postgres",
        //Password: "mypass",
        //Database: "postgres",
	// todo need to fix wait for active db bug
	User: "tsauser",
        Password: "tsapass",
        Database: "tsagroup",
        Addr: "postgres:5432",
    })

    err := createSchema(db)
    if err != nil { panic(err) }

    reformat_initial_data(db)
    db.Close()
}

func main() {
	route := initiateRoutes();
	Model()
	route.Run(":3004")
}
