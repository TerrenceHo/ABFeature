package stores

const (
	experimentsGroupsCreateTableSQL = `
		CREATE TABLE IF NOT EXISTS experiments_groups (
			id varchar(20) primary key,
			experiment_id varchar(20) NOT NULL,
			group_id varchar(20) NOT NULL,
			created_at timestamptz DEFAULT NOW(),
			updated_at timestamptz DEFAULT NOW(),
			FOREIGN KEY (group_id) REFERENCES groups (id) ON DELETE CASCADE,
			FOREIGN KEY (experiment_id) REFERENCES experiments (id) ON DELETE CASCADE
		)
	`
)
