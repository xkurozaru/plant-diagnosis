package model

type Role struct {
	Type string
}

var (
	AdminRole = Role{Type: "admin"}
	UserRole  = Role{Type: "user"}
)

func NewRole(roleType string) Role {
	switch roleType {
	case AdminRole.Type:
		return AdminRole
	case UserRole.Type:
		return UserRole
	default:
		return Role{}
	}
}

func (r Role) IsAdmin() bool {
	return r.Type == AdminRole.Type
}

func (r Role) IsUser() bool {
	return r.Type == UserRole.Type
}
