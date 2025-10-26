package main

type oeuvre struct {
	Id          int
	Titre       string
	Description string
	Artiste     string
	Image       string
}

func (a *artBox) GetOeuvres() ([]oeuvre, error) {
	oeuvres := []oeuvre{}

	rows, err := a.db.Query("SELECT * FROM oeuvres")
	if err != nil {
		return oeuvres, err
	}
	defer rows.Close()

	for rows.Next() {
		o := oeuvre{}
		rows.Scan(
			&o.Id,
			&o.Titre,
			&o.Description,
			&o.Artiste,
			&o.Image,
		)

		oeuvres = append(oeuvres, o)
	}

	return oeuvres, nil
}

func (a *artBox) GetOeuvreById(id int) (oeuvre, error) {
	o := oeuvre{}

	row := a.db.QueryRow("SELECT * FROM oeuvres WHERE id = :id", id)
	err := row.Scan(
		&o.Id,
		&o.Titre,
		&o.Description,
		&o.Artiste,
		&o.Image,
	)

	return o, err
}

func (a *artBox) addOeuvre(o oeuvre) error {
	_, err := a.db.Exec(
		`INSERT INTO oeuvres (titre, description, artiste, image)
		VALUES (:titre, :description, :artiste, :image)`,
		o.Titre,
		o.Description,
		o.Artiste,
		o.Image,
	)

	return err
}
