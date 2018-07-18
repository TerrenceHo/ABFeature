package stores

const (
	experimentsCreateTableSQL = `
		CREATE TABLE IF NOT EXISTS experiments (
			id varchar(20) primary key,
			name text,
			description text,
			percentage float,
			enabled bool,
			project_id varchar(20) NOT NULL,
			created_at timestamptz DEFAULT NOW(),
			updated_at timestamptz DEFAULT NOW(),
			FOREIGN KEY (project_id) REFERENCES projects (id) ON DELETE CASCADE
		);
	`

	experimentsGetAllSQL = "SELECT * FROM experiments "

	experimentsGetByIDSQL = "SELECT * FROM experiments WHERE id=$1"

	experimentsGetCountSQL = "SELECT COUNT(*) FROM experiments "

	experimentsInsertSQL = `
		INSERT INTO experiments (id, name, description, 
		percentage, enabled, project_id)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING created_at, updated_at
	`

	experimentsUpdateSQL = `
		UPDATE experiments
		SET name=$2, description=$3, percentage=$4, enabled=$5, updated_at=NOW()
		WHERE id=$1 AND project_id=$6
		RETURNING updated_at
	`

	experimentsDeleteSQL = "DELETE FROM experiments WHERE id=$1"
)
