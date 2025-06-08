package main

type SurveyStore interface {
	AddDepartmentSurveyResponse(data requestParameters) error
	UpdateResponseCount() error
	GetTotalResponseCount() (int, error)
	AddDepartmentSurveyResponseWithCount(data requestParameters) error
}
