package repository

var getUserByEmail = `
	SELECT * FROM users WHERE email = @email;
`

var getUserByID = `
	SELECT * FROM users WHERE user_id = @user_id;
`

var createUser = `
	INSERT INTO users (name, country, default_language, email, phones, password)
	VALUES (@name, @country, @default_language, @email, @phones, @password)
	RETURNING *;
`

var deactivateDeleteUserByEmail = `
	UPDATE users SET deleted_at = (NOW() AT TIME ZONE 'UTC') WHERE email = @email AND deleted_at IS NULL RETURNING *;
`

var deactivateDeleteUserByID = `
	UPDATE users SET deleted_at = (NOW() AT TIME ZONE 'UTC') WHERE user_id = @user_id AND deleted_at IS NULL RETURNING *;
`
