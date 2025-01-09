// Code generated by gitpod-log-fields. DO NOT EDIT.

package v1

import (
	logfields "github.com/gitpod-io/flex-sdk-go/tools/logfields"
)

func (x *GetIDTokenRequest) LogFields() logfields.Collection {
	return nil
}

func (x *GetIDTokenResponse) LogFields() logfields.Collection {
	return nil
}

func (x *Subject) LogFields() logfields.Collection {
	if x == nil {
		return nil
	}

	var fields logfields.Collection
	fields.Add(&logfields.Entry{Name: "subject.id", Value: x.Id})
	return fields
}

func (x *GetAuthenticatedIdentityRequest) LogFields() logfields.Collection {
	return nil
}

func (x *GetAuthenticatedIdentityResponse) LogFields() logfields.Collection {
	if x == nil {
		return nil
	}

	var fields logfields.Collection
	fields.Add(logfields.Extract(x.GetSubject())...)
	return fields
}

func (x *ExchangeTokenRequest) LogFields() logfields.Collection {
	return nil
}

func (x *ExchangeTokenResponse) LogFields() logfields.Collection {
	return nil
}
