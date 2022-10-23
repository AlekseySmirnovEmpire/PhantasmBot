package character

import "fmt"

func PrintList(str *string) string {
	return fmt.Sprintf("---------------------------------\n%s\n---------------------------------\n", *str)
}
