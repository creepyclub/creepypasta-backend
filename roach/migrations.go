package roach

const migrationTable = "creepypasta_migrations"

func (r *Roach) applyMigrations() error {
	lastMigrationId, err := r.lastMigrationId()
	if err != nil {
		return err
	}
	trx, err := r.Db.Begin()
	if err != nil {
		return err
	}
	switch lastMigrationId {
	case 0:
		_, err = trx.Exec("CREATE TABLE topics (" +
			"topic_id BIGSERIAL PRIMARY KEY NOT NULL," +
			"topic_title VARCHAR(250) NOT NULL," +
			"topic_slug VARCHAR(500) NOT NULL DEFAULT ''," +
			"topic_text TEXT NOT NULL," +
			"topic_text_source TEXT NOT NULL," +
			"topic_text_hash VARCHAR(32) NOT NULL)")
		if err != nil {
			return err
		}
		err = r.updateMigrationId(1)
		if err != nil {
			return err
		}
		fallthrough
	case 1:
		//r.updateMigrationId(2)
		//fallthrough
	}
	return trx.Commit()
}

func (r *Roach) updateMigrationId(id int) error {
	_, err := r.Db.Exec("UPDATE "+migrationTable+" SET last_migration = $1", id)
	return err
}

func (r *Roach) lastMigrationId() (int, error) {
	_, err := r.Db.Exec("CREATE TABLE IF NOT EXISTS " + migrationTable + " (last_migration integer NOT NULL DEFAULT 0)")
	if err != nil {
		return 0, err
	}
	var lastMigration int
	row := r.Db.QueryRow("SELECT last_migration FROM " + migrationTable)
	err = row.Scan(&lastMigration)
	if err != nil {
		_, err = r.Db.Exec("INSERT INTO " + migrationTable + " (last_migration) VALUES (0)")
		return 0, err
	}
	return lastMigration, nil
}
