package addressbook

import (
	"fmt"
	"testing"
)

func TestShouldSerializeObject(t *testing.T) {
	addressBookFile := "/Users/dvalencia/.addressbook.txt"
	var addressBook *AddressBook
	addressBook, err := Read(addressBookFile)

	if err != nil {
		addressBook = NewAddressBook()
	}

	person := &Person{
		Id:    1234,
		Name:  "Daniel Valencia",
		Email: "danvalencia@gmail.com",
		Phones: []*Person_PhoneNumber{
			{Number: "331-347-3493", Type: Person_HOME},
		},
	}

	addressBook.AddToBook(person)

	addressBook.Write(addressBookFile)

	fmt.Printf("AddressBook: %v", addressBook)
}
