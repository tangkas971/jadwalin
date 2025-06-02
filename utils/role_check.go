package utils

import "fmt"

func RoleCheck(inputRoleUser string, allowedRole string) error {
	if inputRoleUser != allowedRole {
		return fmt.Errorf("akses ditolak! role %s tidak diizinkan", inputRoleUser)
	}
	return nil 
}