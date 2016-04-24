package models

type HeartRate struct {
	value     string `json:"value"`
	PatientID int    `json:"patient_id"`
}
