package models

type BloodPressure struct {
	value     string `json:"value"`
	PatientID int    `json:"patient_id"`
}
