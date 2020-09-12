package model

// Reason of an error.
type Reason string

// Possible reasons
const (
	ReasonUnknown      Reason = ""
	ReasonExtraLine    Reason = "extra line before"
	ReasonMissingLine  Reason = "missing line after"
	ReasonTooManyLines Reason = "too many lines after"
)

// ReasonFromError extracts reason from the error.
func ReasonFromError(err error) Reason {
	if err == nil {
		return ReasonUnknown
	}

	if ierr, ok := err.(importOrderError); ok {
		return ierr.Reason
	}

	return ReasonUnknown
}
