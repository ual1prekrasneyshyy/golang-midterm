package main

func GetRoleByName(name string) Role {
	var role Role
	GetConnection().Find(&role, "name = ?", name)
	return role
}
