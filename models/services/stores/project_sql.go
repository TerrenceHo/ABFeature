package stores

const (
	projectsCreateTableSQL = `
		CREATE TABLE IF NOT EXISTS projects (
			id varchar(20) primary key,
			name text,
			description text,
			created_at timestamptz DEFAULT NOW(),
			updated_at timestamptz DEFAULT NOW()
		);
	`
	projectsGetAllSQL = "SELECT * FROM projects;"

	projectsGetCountSQL = "SELECT COUNT(*) FROM projects "

	projectsGetByIDSQL = "SELECT * FROM projects WHERE id=$1;"

	projectsInsertSQL = `
		INSERT INTO projects (id, name, description)
		VALUES ($1, $2, $3)
		RETURNING created_at, updated_at
	`

	projectsUpdateSQL = `
		UPDATE projects 
		SET name=$2, description=$3, updated_at=NOW()
		WHERE id=$1
		RETURNING updated_at
	`
	projectsDeleteSQL = "DELETE FROM projects WHERE id=$1;"
)
