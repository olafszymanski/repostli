package user

type UserRepository interface {
	Create(email, password string) error
}

type postgresRepository struct {
}

func NewPostgresRepository() UserRepository {
	return &postgresRepository{}
}

func (r *postgresRepository) Create(email, password string) error {
	// Connect to database
	// Create user
	return nil
}
