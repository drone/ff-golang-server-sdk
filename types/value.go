package types

// ValueType interface used for different clause types
type ValueType interface {
	StartsWith(value interface{}) bool
	EndsWith(value interface{}) bool
	Match(value interface{}) bool
	Contains(value interface{}) bool
	EqualSensitive(value interface{}) bool
	Equal(value interface{}) bool
	GreaterThan(value interface{}) bool
	GreaterThanEqual(value interface{}) bool
	LessThan(value interface{}) bool
	LessThanEqual(value interface{}) bool
	In(value interface{}) bool
}
