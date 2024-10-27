package ptr

import "time"

func FromString(v string) *string {
	return &v
}

func FromInt(v int) *int {
	return &v
}

func FromTime(v time.Time) *time.Time {
	return &v
}

func ToString(p *string) string {
	return *p
}

func ToInt(p *int) int {
	return *p
}

func ToTime(p *time.Time) time.Time {
	return *p
}
