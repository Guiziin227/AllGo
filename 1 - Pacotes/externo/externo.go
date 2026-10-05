package externo

import "github.com/badoux/checkmail"

func ValidandoEmail(email string) bool {
	err := checkmail.ValidateFormat(email)
	return err == nil
}
