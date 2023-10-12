package model

type Role struct {
	Type        string
	Permissions []Permission
}

var (
	AdminRole = Role{
		Type: "admin",
		Permissions: []Permission{
			CreatePredictionModelPermission,
			ReadPredictionModelPermission,
			DeletePredictionModelPermission,
			PredictionPermission,
		},
	}
	UserRole = Role{
		Type: "user",
		Permissions: []Permission{
			ReadPredictionModelPermission,
			PredictionPermission,
		},
	}
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

func (r Role) HasPermission(permission Permission) bool {
	for _, p := range r.Permissions {
		if p == permission {
			return true
		}
	}
	return false
}
