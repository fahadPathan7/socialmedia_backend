package accesscontrol

const (
	// Admin role
	Admin = "admin"
	// User role
	User = "user"
	// Guest role
	Guest = "guest"
)

// entityRolePermissionsMap defines all
func entityRolePermissionsMap() map[string]entityPermissions {
	return map[string]entityPermissions{
		Post: {
			Entity:               Post,
			RolesAllowedToCreate: []string{Admin, User},
			RolesAllowedToRead:   []string{Admin, User, Guest},
			RolesAllowedToUpdate: []string{Admin, User},
			RolesAllowedToDelete: []string{Admin, User},
		},
		Comment: {
			Entity:               Comment,
			RolesAllowedToCreate: []string{Admin, User},
			RolesAllowedToRead:   []string{Admin, User, Guest},
			RolesAllowedToUpdate: []string{Admin, User},
			RolesAllowedToDelete: []string{Admin, User},
		},
		React: {
			Entity:               React,
			RolesAllowedToCreate: []string{Admin, User},
			RolesAllowedToRead:   []string{Admin, User, Guest},
			RolesAllowedToUpdate: []string{Admin, User},
			RolesAllowedToDelete: []string{Admin, User},
		},
	}
}
