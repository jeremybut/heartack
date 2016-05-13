package models

type Patient struct {
	ID                   int    `json:"id"`
	Firstname            string `json:"firstname"`
	Lastname             string `json:"lastname"`
	Email                string `json:"email"`
	Password             string `json:"password"`
	SocialSecurityNumber string `json:"social_security_number"`
	// Birthdate         date
	Gender       string `json:"gender"`
	Address      string `json:"address"`
	Zipcode      string `json:"zipcode"`
	City         string `json:"city"`
	Country      string `json:"country"`
	HeartRateMin string `json:"heart_rate_min"`
	HeartRateMax string `json:"heart_rate_max"`
	Weight       string `json:"weight"`
	UserID       uint   `json:"doctor_id"`
}

func AllPatients() ([]*Patient, error) {
	rows, err := db.Query("SELECT id, firstname, lastname, email FROM patients")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	pts := make([]*Patient, 0)
	for rows.Next() {
		pt := new(Patient)
		err := rows.Scan(&pt.ID, &pt.Firstname, &pt.Lastname, &pt.Email)
		if err != nil {
			return nil, err
		}
		pts = append(pts, pt)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return pts, nil
}
