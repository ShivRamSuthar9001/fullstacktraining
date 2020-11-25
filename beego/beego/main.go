package main

import (
	"fmt"

	"github.com/astaxie/beego/client/orm"

	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/go-sql-driver/mysql" // import your required driver
)

//Model Struct
type User struct {
	Id   int
	Name string `orm:"size(100)"`
}

func init() {
	// register model
	orm.RegisterModel(new(User))

	// set default database
	orm.RegisterDataBase("default", "mysql", "shiva:shivrfd1!@tcp(10.3.250.47:3306)/mydb")
}

func handler(e User) (string, error) {
	return fmt.Sprintf("<h1>Hello %v from lambda Go</h1>", e.Name), nil
}
func main() {
	o := orm.NewOrm()

	user := User{Id: 3, Name: "shivram"}

	// insert
	Id, err := o.Insert(&user)
	fmt.Printf("Id: %d, ERR: %v\n", Id, err)

	// update
	user.Name = "astaxie"
	num, err := o.Update(&user)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	// read one
	u := User{Id: 2}
	err = o.Read(&u)
	fmt.Printf("ERR: %v\n", err)

	lambda.Start(handler)

	// delete
	num, err = o.Delete(&u)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

}
