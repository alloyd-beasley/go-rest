package models

type Patient struct {
	SequenceNumberTreatment []string `json:"sequence_number_treatment"`
	PatientSequenceNumber   string   `json:"patient_sequence_number"`
	SequenceNumberOutcome   []string `json:"sequence_number_outcome"`
}
