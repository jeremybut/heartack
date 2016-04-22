
package main

type User struct {
    Firstname             string            `json:"firstname"`
    Lastname              string            `json:"lastname"`
    Email                 string            `json:"email"`
    Password              string            `json:"password"`
    // Patients              []Patient         `json:"patients"`
}

type Patient struct {
    Firstname             string            `json:"firstname"`
    Lastname              string            `json:"lastname"`
    Email                 string            `json:"email"`
    Password              string            `json:"password"`
    SocialSecurityNumber  string            `json:"social_security_number"`
    // Birthdate             date
    Gender                string            `json:"gender"`
    Address               string            `json:"address"`
    Zipcode               string            `json:"zipcode"`
    City                  string            `json:"city"`
    Country               string            `json:"country"`
    HeartRateMin          string            `json:"heart_rate_min"`
    HeartRateMax          string            `json:"heart_rate_max"`
    Weight                string            `json:"weight"`
    UserID                uint              `json:"doctor_id"`
    // BloodPressures        []BloodPressure   `json:"blood_pressures"`
    // HeartRates            []HeartRate       `json:"heart_rates"`
}

type BloodPressure struct {
    value                 string            `json:"value"`
    PatientID             int               `json:"patient_id"`
}

type HeartRate struct {
    value                 string            `json:"value"`
    PatientID             int               `json:"patient_id"`
}
