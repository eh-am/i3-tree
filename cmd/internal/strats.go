package internal

// BadStratError represents an error when a strategy could not be
// parsed into a existing strategy correctly
type BadStratError struct {
	StratName string
}

func (e BadStratError) Error() string {
	return "invalid strat: " + e.StratName
}
