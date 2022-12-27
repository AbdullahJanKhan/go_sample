package repository

type Store interface {
	BeginTx() (Store, error)
	Rollback() error
	CommitTx() error
	// mention your database layer function
}
