package rules

import (
	"errors"
	"reflect"

	"ensi-cloud-integration/internal/domain"
)

func ActionTypeRule(v interface{}, _ string) error {
	st := reflect.ValueOf(v)
	if st.Kind() != reflect.String {
		return errors.New("action must be a string")
	}

	for _, value := range domain.GetActionTypes() {
		if string(value) == st.String() {
			return nil
		}
	}

	return errors.New("invalid action type")
}

func SortTypeRule(v interface{}, _ string) error {
	st := reflect.ValueOf(v)
	if st.Kind() != reflect.String {
		return errors.New("sort must be a string")
	}

	for _, value := range domain.GetSortTypes() {
		if string(value) == st.String() {
			return nil
		}
	}

	return errors.New("invalid sort type")
}
