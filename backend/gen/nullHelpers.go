package gen

import "database/sql"

func makeNullString(str string) sql.NullString {
	if str == "" {
		return sql.NullString{}
	}
	return sql.NullString{
		String: str,
		Valid:  true,
	}
}

func makeNullBool(bl *bool) sql.NullBool {
	if bl == nil {
		return sql.NullBool{}
	}
	return sql.NullBool{
		Bool:  *bl,
		Valid: true,
	}
}

func makeNullFloat(f *float64) sql.NullFloat64 {
	if f == nil {
		return sql.NullFloat64{
			Float64: 0,
			Valid:   false,
		}
	}
	return sql.NullFloat64{
		Float64: *f,
		Valid:   true,
	}
}
