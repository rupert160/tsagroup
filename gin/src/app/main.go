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

//type contact struct {
//	Full_Name string `json:"full_name"`
//	Email string `json:"title"`
//        Phone_Numbers []int `json:"phone_numbers"`
//}
type Contact struct {
    Id     int64
    Name   string
    Email string
    Phone_Numbers []int64
}

func (u Contact) String() string {
    return fmt.Sprintf("User<%d %s %v>", u.Id, u.Name, u.Email, u.Phone_Numbers)
}

//var contacts []Contact

func convert_stoi(s string) int {
	if s == "" {return 0}
	i, err := strconv.Atoi(s)
	if err != nil {
        	panic(err)
    	}
	return i
}

func e164(token string) int{
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

//func convert_contact(cs contact_string) Contact {
//        var numbers []int
//        //for _,
//        //number_string := range cs.Phone_Numbers{
//        //        numbers = append(numbers, e164(number_string))//todo
//        //}
//        return Contact{cs.Full_Name, cs.Email, numbers} //todo
//}
//
//func reformat_initial_data(){
//        //var new_contacts []Contact //todo
//        //for _,
//        //cs := range contacts_string {
//        //        new_contacts = append(new_contacts, convert_contact(cs))//todo
//        //}
//        //contacts = new_contacts //todo
//	contacts_string = nil
//}

func createSchema(db *pg.DB) error {
    models := []interface{}{
        (*Contact)(nil),
    }

    for _, model := range models {
        err := db.Model(model).CreateTable(&orm.CreateTableOptions{
            Temp: false,
        })
        if err != nil {
            return err
        }
    }
    return nil
}

func Model() {
    db := pg.Connect(&pg.Options{
        User: "tsauser",
        Password: "tsapass",
        Addr: "postgres:5432",
        Database: "tsagroup",
    })
    defer db.Close()

    err := createSchema(db)
    if err != nil {
        panic(err)
    }

    user1 := &Contact{
        Name:   "Rupert Bailey",
        Email: "admin1@admin.com",
        Phone_Numbers: []int64{611800111222, 61470436111},
    }
    _, err = db.Model(user1).Insert()
    if err != nil {
        panic(err)
    }

    _, err = db.Model(&Contact{
        Name:   "Bob Bailey",
        Email: "bob@admin.com",
        Phone_Numbers: []int64{611800111333, 61470436222},
    }).Insert()
    if err != nil {
        panic(err)
    }

    // Select user by primary key.
    user := &Contact{Id: user1.Id}
    err = db.Model(user).WherePK().Select()
    if err != nil {
        panic(err)
    }

    // Select all users.
    var users []Contact
    err = db.Model(&users).Select()
    if err != nil {
        panic(err)
    }

    fmt.Println(user)
    fmt.Println(users)
    // Output: Contact<1 admin [admin1@admin admin2@admin]>
    // [Contact<1 admin [admin1@admin admin2@admin]> Contact<2 root [root1@root root2@root]>]
}

func main() {
	route := initiateRoutes();
	time.Sleep(15 * time.Second)
	Model()
	route.Run(":3004")
}
