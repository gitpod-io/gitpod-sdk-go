// Code generated by gitpod-log-fields. DO NOT EDIT.

package v1

import (
	logfields "github.com/gitpod-io/flex-go/tools/logfields"
)

func (x *CreateSecretRequest) LogFields() logfields.Collection {
	if x == nil {
		return nil
	}

	var fields logfields.Collection
	fields.Add(&logfields.Entry{Name: "project.id", Value: x.ProjectId})
	return fields
}

func (x *CreateSecretResponse) LogFields() logfields.Collection {
	if x == nil {
		return nil
	}

	var fields logfields.Collection
	fields.Add(logfields.Extract(x.GetSecret())...)
	return fields
}

func (x *UpdateSecretValueRequest) LogFields() logfields.Collection {
	if x == nil {
		return nil
	}

	var fields logfields.Collection
	fields.Add(&logfields.Entry{Name: "secret.id", Value: x.SecretId})
	return fields
}

func (x *UpdateSecretValueResponse) LogFields() logfields.Collection {
	return nil
}

func (x *ListSecretsRequest) LogFields() logfields.Collection {
	if x == nil {
		return nil
	}

	var fields logfields.Collection
	fields.Add(logfields.Extract(x.GetPagination())...)
	fields.Add(logfields.Extract(x.GetFilter())...)
	return fields
}

func (x *ListSecretsResponse) LogFields() logfields.Collection {
	if x == nil {
		return nil
	}

	var fields logfields.Collection
	fields.Add(logfields.Extract(x.GetPagination())...)
	return fields
}

func (x *DeleteSecretRequest) LogFields() logfields.Collection {
	if x == nil {
		return nil
	}

	var fields logfields.Collection
	fields.Add(&logfields.Entry{Name: "secret.id", Value: x.SecretId})
	return fields
}

func (x *DeleteSecretResponse) LogFields() logfields.Collection {
	return nil
}

func (x *GetSecretValueRequest) LogFields() logfields.Collection {
	if x == nil {
		return nil
	}

	var fields logfields.Collection
	fields.Add(&logfields.Entry{Name: "secret.id", Value: x.SecretId})
	return fields
}

func (x *GetSecretValueResponse) LogFields() logfields.Collection {
	return nil
}

func (x *Secret) LogFields() logfields.Collection {
	if x == nil {
		return nil
	}

	var fields logfields.Collection
	fields.Add(&logfields.Entry{Name: "secret.id", Value: x.Id})
	fields.Add(&logfields.Entry{Name: "project.id", Value: x.ProjectId})
	fields.Add(logfields.Extract(x.GetCreatedAt())...)
	fields.Add(logfields.Extract(x.GetUpdatedAt())...)
	fields.Add(logfields.Extract(x.GetCreator())...)
	return fields
}
