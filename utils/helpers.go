package utils

import "time"

func HandleNil[k any](data any) *k {
	if data == nil {
		return nil
	}
	converted := data.(k)
	return &converted
}

func ToConst[k ~string](data any) *k {
	if data == nil {
		return nil
	}
	stringData := data.(string)
	convertedData := k(stringData)
	return &convertedData
}

func ToTimePtr(t any) *time.Time {
	if t == nil {
		return nil
	}
	timeValue := t.(time.Time)
	return &timeValue
}
