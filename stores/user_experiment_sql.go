package stores

const (
	usersExperimentsCreateTableSQL = `
		CREATE TABLE IF NOT EXISTS users_experiments (
			id varchar(20) primary key,
			user_id varchar(20) NOT NULL,
			experiment_id varchar(20) NOT NULL,
			created_at timestamptz DEFAULT NOW(),
			updated_at timestamptz DEFAULT NOW(),
			FOREIGN KEY (experiment_id) REFERENCES experiments (id) ON DELETE CASCADE,
			FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
		)
	`

	usersExperimentsGetAllSQL = "SELECT * FROM users_experiments "

	usersExperimentsGetAllByUserSQL = `
		SELECT e.* FROM experiments e
		INNER JOIN users_experiments ue ON e.id = ue.experiment_id
		WHERE ue.user_id=$1
	`

	usersExperimentsGetAllByExperimentSQL = `
		SELECT u.* FROM users u
		INNER JOIN users_experiments ue on u.id = ue.user_id
		WHERE ue.experiment_id=$1
	`

	usersExperimentsGetByIDSQL = "SELECT * FROM users_experiments WHERE id=$1"

	usersExperimentsGetByUserAndExperimentSQL = `
		SELECT * FROM users_experiments
		WHERE user_id=$1 AND experiment_id=$2
	`

	usersExperimentsInsertSQL = `
		INSERT INTO users_experiments (id, user_id, experiment_id)
		VALUES ($1, $2, $3)
		RETURNING created_at, updated_at
	`

	usersExperimentsDeleteSQL = `
		DELETE FROM users_experiments
		WHERE user_id=$1 AND experiment_id=$2
	`
)
