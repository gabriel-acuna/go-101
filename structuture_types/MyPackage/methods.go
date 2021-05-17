package MyPackage

import "fmt"

func (self *Person) GetFullName() string {
	return fmt.Sprintf("%s %s", self.FirstName, self.LastName)
}
