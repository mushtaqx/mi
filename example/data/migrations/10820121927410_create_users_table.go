package migrations

type CreateUsersTable struct{}

func (m CreateUsersTable) Up() bool {
	// setup db
	return true
}

func (m CreateUsersTable) Down() bool {
	// take down db
	return true
}
