package stores

const (
	groupsCreateTableSQL = `
		CREATE TABLE IF NOT EXISTS groups (
			id varchar(20) primary key,
			name text,
			description text, 
			created_at timestamptz DEFAULT NOW(),
			updated_at timestamptz DEFAULT NOW()
		)
	`

	groupsGetAllSQL = "SELECT * FROM groups"

	groupsGetCountSQL = "SELECT COUNT(*) FROM groups "

	groupsGetByIDSQL = "SELECT * FROM groups WHERE id=$1"

	groupsInsertSQL = `
		INSERT INTO groups (id, name, description)
		VALUES ($1, $2, $3)
		RETURNING created_at, updated_at
	`

	groupsUpdateSQL = `
		UPDATE groups
		SET name=$2, description=$3, updated_at=NOW()
		WHERE id=$1
		RETURNING updated_at
	`

	groupsDeleteSQL = "DELETE FROM groups WHERE id=$1"
)
