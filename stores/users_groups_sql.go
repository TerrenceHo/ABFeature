package stores

const (
	usersGroupsCreateTableSQL = `
		CREATE TABLE IF NOT EXISTS users_groups (
			id varchar(20) primary key,
			user_id varchar(20) NOT NULL,
			group_id varchar(20) NOT NULL,
			created_at timestamptz DEFAULT NOW(),
			updated_at timestamptz DEFAULT NOW(),
			FOREIGN KEY (group_id) REFERENCES groups (id) ON DELETE CASCADE,
			FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
		)
	`

	usersGroupsGetAllSQL = "SELECT * FROM users_groups "

	usersGroupsGetAllByUserSQL = `
		SELECT g.* FROM groups g
		INNER JOIN users_groups ug ON g.id = ug.group_id
		WHERE ug.user_id=$1
	`

	usersGroupsGetAllByGroupSQL = `
		SELECT u.* FROM users u
		INNER JOIN users_groups ug ON u.id = ug.user_id
		WHERE ug.group_id=$1
	`

	usersGroupsGetByIDSQL = "SELECT * FROM users_groups WHERE id=$1"

	usersGroupsGetByUserAndGroupSQL = `
		SELECT * FROM users_groups 
		WHERE user_id=$1 AND group_id=$2
	`

	usersGroupsInsertSQL = `
		INSERT INTO users_groups (id, user_id, group_id)
		VALUES ($1, $2, $3)
		RETURNING created_at, updated_at
	`

	usersGroupsDeleteSQL = "DELETE FROM users_groups WHERE user_id=$1 AND group_id=$2"
)
