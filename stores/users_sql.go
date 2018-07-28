package stores

const (
	usersCreateTableSQL = `
		CREATE TABLE IF NOT EXISTS users (
			id varchar(20) primary key,
			name text,
			description text,
			created_at timestamptz DEFAULT NOW(),
			updated_at timestamptz DEFAULT NOW
		)
	`

	usersGetAllSQL = "SELECT * FROM users "

	usersGetByIDSQL = "SELECT * FROM experiments WHERE id=$1"

	usersInsertSQL = `
		INSERT INTO users (id, name, description)
		VALUES ($1, $2, $3) 
		RETURNING created_at, updated_at
	`

	usersUpdateSQL = `
		UPDATE users
		SET name=$2, description=$3, updated_at=NOW()
		WHERE id=$1
		RETURNING updated_at
	`

	usersDeleteSQL = "DELETE FROM users WHERE id=$1"
)
