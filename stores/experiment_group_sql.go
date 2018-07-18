package stores

const (
	experimentsGroupsCreateTableSQL = `
		CREATE TABLE IF NOT EXISTS experiments_groups (
			id varchar(20) primary key,
			group_id varchar(20) NOT NULL,
			experiment_id varchar(20) NOT NULL,
			created_at timestamptz DEFAULT NOW(),
			updated_at timestamptz DEFAULT NOW(),
			FOREIGN KEY (group_id) REFERENCES groups (id) ON DELETE CASCADE,
			FOREIGN KEY (experiment_id) REFERENCES experiments (id) ON DELETE CASCADE
		)
	`

	experimentsGroupsGetAllSQL = "SELECT * FROM experiments_groups "

	experimentsGroupsGetByIDSQL = "SELECT * FROM experiments_groups WHERE id=$1"

	experimentsGroupsGetCountSQL = "SELECT COUNT(*) FROM experiments_groups "

	experimentsGroupsInsertSQL = `
		INSERT INTO experiments_groups (id, experiment_id, group_id)
		VALUES ($1, $2, $3)
		RETURNING created_at, updated_at
	`

	experimentsGroupsDeleteSQL = "DELETE FROM experimentsGroups WHERE id=$1"
)
