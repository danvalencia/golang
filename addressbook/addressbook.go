package addressbook

import (
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/proto"
)

// NewAddressBook returns a new AddressBookInstance
func NewAddressBook() *AddressBook {
	return &AddressBook{}
}

// AddToBook will add the given person to the addressBook
func (addressBook *AddressBook) AddToBook(person *Person) {
	people := addressBook.GetPeople()
	addressBook.People = append(people, person)
}

// Write will write the addressBook to the given file
func (addressBook *AddressBook) Write(filename string) error {
	out, err := proto.Marshal(addressBook)
	if err != nil {
		log.Println("There was an issue marhalling the address book")
		return err
	}

	if err := ioutil.WriteFile(filename, out, 0644); err != nil {
		log.Println("There was an issue writing address book to disk")
		return err
	}

	return nil

}

// Read will read the contents of the file pointed by filename and
// will unmarshall the content into an AddressBook structure.
// If the file does not exist, an error is returned.
func Read(filename string) (*AddressBook, error) {
	in, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("There was an error reading file %v\n", filename)
		return nil, err
	}

	addressBook := &AddressBook{}

	proto.Unmarshal(in, addressBook)

	return addressBook, nil
}
