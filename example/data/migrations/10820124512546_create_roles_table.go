package migrations

type CreateRolesTable struct{}

func (m CreateRolesTable) Up() bool {
	// setup db
	return true
}

func (m CreateRolesTable) Down() bool {
	// take down db
	return true
}
