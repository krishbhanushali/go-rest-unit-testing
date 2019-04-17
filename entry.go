package main

type entry struct {
	ID           int    `json:"id,omitempty"`
	FirstName    string `json:"first_name,omitempty"`
	LastName     string `json:"last_name,omitempty"`
	EmailAddress string `json:"email_address,omitempty"`
	PhoneNumber  string `json:"phone_number,omitempty"`
}
