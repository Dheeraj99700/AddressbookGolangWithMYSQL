package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Contact struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
	City      string
	State     string
	Zip       int
	Mobile    int
	Email     string
}

var db *sql.DB
var err error

func main() {
	var option int
	var choice int
	fmt.Println("------------Welcome to Address Book Program----------------")
	connector()
	var exit bool = true
	for exit {
		fmt.Println("Select The option to perform operation given below")
		fmt.Printf("1.View All Contacts\n2.View Contcat By City Or State\n3.Numbers of Contacts BY City OR State\n4.Add Contact\n5.Update Contact\n6.Delete Contact\n")
		fmt.Scanf("%d", &option)
		switch option {
		case 1:
			{
				ViewAllContacts()
				//View All Contact

			}
		case 2:
			{
				ViewCityOrStateContact()
				//ViewCityOrStateContact
			}
		case 3:
			{
				CountOfContacts()
				//CountOfContacts
			}
		case 4:
			{
				AddContact()
				// AddContact
			}
		case 5:
			{
				UpdateContact()
				//UpdateContact
			}
		case 6:
			{
				DeleteContact()
				//DeleteContact()
			}
		default:
			{
				fmt.Println("Invalid Option Entered")
			}
		}
		fmt.Printf("\nDo you wish to continue with the program\nPress 1 to continue\n")
		fmt.Scanf("%d", &choice)
		if choice == 1 {
			continue
		}
		break
	}
	fmt.Println("Thankyou for using Address Book Management System")
}
func connector() {
	db, err = sql.Open("mysql", "root:3421@tcp(127.0.0.1:3306)/Addressbook")
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
}

func ViewAllContacts() {

	var Contacts []Contact

	rows, err := db.Query("SELECT * FROM Contact;")
	if err != nil {
		fmt.Errorf("error in query all Contact: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var data Contact
		if err := rows.Scan(&data.ID, &data.FirstName, &data.LastName, &data.Address, &data.City, &data.State, &data.Zip, &data.Mobile, &data.Email); err != nil {
			fmt.Errorf("error in query all Contact: %v", err)
		}
		Contacts = append(Contacts, data)
	}
	fmt.Println(Contacts)

}

func ViewCityOrStateContact() {
	var op int
	var city string
	var state string
	fmt.Println("Select The option to perform operation given below")
	fmt.Printf("1.City\n2.State\n")
	fmt.Scan(&op)
	switch op {
	case 1:
		fmt.Println("Enter the City Name")
		fmt.Scan(&city)
		rows, err := db.Query("SELECT * FROM Contact WHERE City =?", city)
		if err != nil {
			fmt.Errorf("error in query all Contact: %v", err)
		}
		for rows.Next() {
			var data Contact
			if err := rows.Scan(&data.ID, &data.FirstName, &data.LastName, &data.Address, &data.City, &data.State, &data.Zip, &data.Mobile, &data.Email); err != nil {
				fmt.Errorf("error in query all Contact: %v", err)
			}
			fmt.Println(data)
		}

	case 2:

		fmt.Println("Enter the State Name")
		fmt.Scan(&state)
		rows, err := db.Query("SELECT * FROM Contact WHERE State = ?", state)
		if err != nil {
			fmt.Errorf("error in query all Contact: %v", err)
		}
		defer rows.Close()

		for rows.Next() {
			var data Contact
			if err := rows.Scan(&data.ID, &data.FirstName, &data.LastName, &data.Address, &data.City, &data.State, &data.Zip, &data.Mobile, &data.Email); err != nil {
				fmt.Errorf("error in query all Contact: %v", err)
			}
			fmt.Println(data)
		}

	default:

		fmt.Println("Inavlid Option Enter")

	}
}
func CountOfContacts() {
	var op int
	var city string
	var state string
	var count int
	fmt.Println("Select The option to perform operation given below")
	fmt.Printf("1.City\n2.State\n")
	fmt.Scan(&op)
	switch op {
	case 1:
		fmt.Println("Enter the City Name")
		fmt.Scan(&city)
		rows, err := db.Query("SELECT COUNT(*) FROM Contact WHERE City = ?", city)
		if err != nil {
			fmt.Errorf("error in query all Contact: %v", err)
		}
		defer rows.Close()

		for rows.Next() {
			var data Contact
			if err := rows.Scan(&data.ID, &data.FirstName, &data.LastName, &data.Address, &data.City, &data.State, &data.Zip, &data.Mobile, &data.Email); err != nil {
				fmt.Errorf("error in query all Contact: %v", err)
			}
			if city == data.City {
				count++
			}

		}
		fmt.Printf("Contacts Available in city %s are : %d\n", city, count)

	case 2:

		fmt.Println("Enter the State Name")
		fmt.Scan(&state)
		rows, err := db.Query("SELECT * FROM Contact;")
		if err != nil {
			fmt.Errorf("error in query all Contact: %v", err)
		}
		defer rows.Close()

		for rows.Next() {
			var data Contact
			if err := rows.Scan(&data.ID, &data.FirstName, &data.LastName, &data.Address, &data.City, &data.State, &data.Zip, &data.Mobile, &data.Email); err != nil {
				fmt.Errorf("error in query all Contact: %v", err)
			}
			if state == data.State {
				count++
			}
		}
		fmt.Printf("Contacts Available in State %s are : %d\n", state, count)

	default:

		fmt.Println("Inavlid Option Enter")

	}

}

func AddContact() {
	var Contacts Contact
	fmt.Println("Enter First Name")
	fmt.Scan(&Contacts.FirstName)
	fmt.Println("Enter Last Name")
	fmt.Scan(&Contacts.LastName)
	fmt.Println("Enter Address")
	fmt.Scan(&Contacts.Address)
	fmt.Println("Enter City")
	fmt.Scan(&Contacts.City)
	fmt.Println("Enter State")
	fmt.Scan(&Contacts.State)
	fmt.Println("Enter Zip")
	fmt.Scan(&Contacts.Zip)
	fmt.Println("Enter Mobile")
	fmt.Scan(&Contacts.Mobile)
	fmt.Println("Enter Email")
	fmt.Scan(&Contacts.Email)

	_, err := db.Exec("INSERT INTO Contact (FirstName,LastName,Address,City,State,Zip,Mobile,Email)VALUES (?, ?, ?,?,?, ?, ?,?)", Contacts.FirstName, Contacts.LastName, Contacts.Address, Contacts.City, Contacts.State, Contacts.Zip, Contacts.Mobile, Contacts.Email)
	if err != nil {
		fmt.Errorf("add Contact: %v", err)
	} else {
		fmt.Println("Contact Added Succesfully")
	}

}
func UpdateContact() {
	var Contacts Contact
	var Id int
	var choice int
	fmt.Println("Enter the ID of Contact You Want To Update")
	fmt.Scan(&Id)
	fmt.Println("Enter the Detail You want to update")
	fmt.Printf("1.FirstName\n2.LastName\n3.Address\n4.City\n5.State\n6.Zip\n7.Mobile\n8.Email\n")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		fmt.Println("Enter First Name")
		fmt.Scan(&Contacts.FirstName)
		_, err := db.Exec("UPDATE Contact SET FirstName = ?, WHERE ID= ?", Contacts.FirstName, Id)
		fmt.Println(Contacts.FirstName)
		if err != nil {
			fmt.Errorf("Update Contact: %v", err)
		} else {
			fmt.Println("Contact Updated Succesfully")
		}

	case 2:
		fmt.Println("Enter Last Name")
		fmt.Scan(&Contacts.LastName)
		_, err := db.Exec("UPDATE Contact SET LastName =?, WHERE ID=?", Contacts.LastName, Id)
		if err != nil {
			fmt.Errorf("Update Contact: %v", err)
		} else {
			fmt.Println("Contact Updated Succesfully")
		}
	case 3:
		fmt.Println("Enter Address")
		fmt.Scan(&Contacts.Address)
		_, err := db.Exec("UPDATE Contact SET Address =?, WHERE ID=?", Contacts.Address, Id)
		if err != nil {
			fmt.Errorf("Update Contact: %v", err)
		} else {
			fmt.Println("Contact Updated Succesfully")
		}
	case 4:
		fmt.Println("Enter City")
		fmt.Scan(&Contacts.City)
		_, err := db.Exec("UPDATE Contact SET City =?, WHERE ID=?", Contacts.City, Id)
		if err != nil {
			fmt.Errorf("Update Contact: %v", err)
		} else {
			fmt.Println("Contact Updated Succesfully")
		}
	case 5:
		fmt.Println("Enter State")
		fmt.Scan(&Contacts.State)
		_, err := db.Exec("UPDATE Contact SET State =?, WHERE ID=?", Contacts.State, Id)
		if err != nil {
			fmt.Errorf("Update Contact: %v", err)
		} else {
			fmt.Println("Contact Updated Succesfully")
		}
	case 6:
		fmt.Println("Enter Zip")
		fmt.Scan(&Contacts.Zip)
		_, err := db.Exec("UPDATE Contact SET Zip =?, WHERE ID=?", Contacts.Zip, Id)
		if err != nil {
			fmt.Errorf("Update Contact: %v", err)
		} else {
			fmt.Println("Contact Updated Succesfully")
		}
	case 7:
		fmt.Println("Enter Mobile")
		fmt.Scan(&Contacts.Mobile)
		_, err := db.Exec("UPDATE Contact SET Mobile =?, WHERE ID=?", Contacts.Mobile, Id)
		if err != nil {
			fmt.Errorf("Update Contact: %v", err)
		} else {
			fmt.Println("Contact Updated Succesfully")
		}
	case 8:
		fmt.Println("Enter Email")
		fmt.Scan(&Contacts.Email)
		_, err := db.Exec("UPDATE Contact SET Email =?, WHERE ID=?", Contacts.Email, Id)
		if err != nil {
			fmt.Errorf("Update Contact: %v", err)
		} else {
			fmt.Println("Contact Updated Succesfully")
		}
	default:
		fmt.Println("Invalid Choice")
	}

}

func DeleteContact() {
	var Id int
	fmt.Println("Enter the Id of Contact You want to Delete")
	fmt.Scan(&Id)
	_, err = db.Exec("DELETE from Contact WHERE ID=?", Id)
	if err != nil {
		fmt.Errorf("delete Contact: %v", err)
	} else {
		fmt.Println("Contact Deleted Successfully")
	}

}
