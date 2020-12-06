package models

type textTypeCode string

const (
	DescriptionOfEvent              textTypeCode = "Description of Event or Problem"
	ManufacturerEvaluationSummary   textTypeCode = "Manufacturer Evaluation Summary"
	AdditionalManufacturerNarrative textTypeCode = "Additional Manufacturer Narrative"
)

type MdrText struct {
	TextTypeCode string `json:"text_type_code"`
	Text         string `json:"text"`
}
