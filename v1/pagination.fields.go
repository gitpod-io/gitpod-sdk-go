// Code generated by gitpod-log-fields. DO NOT EDIT.

package v1

import (
	logfields "github.com/gitpod-io/flex-sdk-go/tools/logfields"
)

func (x *PaginationRequest) LogFields() logfields.Collection {
	return nil
}

func (x *PaginationResponse) LogFields() logfields.Collection {
	return nil
}

func (x *Sort) LogFields() logfields.Collection {
	if x == nil {
		return nil
	}

	var fields logfields.Collection
	fields.Add(&logfields.Entry{Name: "sort.field", Value: x.Field})
	return fields
}
