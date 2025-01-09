// Code generated by gitpod-log-fields. DO NOT EDIT.

package v1

import (
	logfields "github.com/gitpod-io/flex-api-go/tools/logfields"
)

func (x *CreateRunnerRequest) LogFields() logfields.Collection {
	if x == nil {
		return nil
	}

	var fields logfields.Collection
	fields.Add(logfields.Extract(x.GetSpec())...)
	return fields
}

func (x *CreateRunnerResponse) LogFields() logfields.Collection {
	if x == nil {
		return nil
	}

	var fields logfields.Collection
	fields.Add(logfields.Extract(x.GetRunner())...)
	return fields
}

func (x *GetRunnerRequest) LogFields() logfields.Collection {
	if x == nil {
		return nil
	}

	var fields logfields.Collection
	fields.Add(&logfields.Entry{Name: "runner.id", Value: x.RunnerId})
	return fields
}

func (x *GetRunnerResponse) LogFields() logfields.Collection {
	if x == nil {
		return nil
	}

	var fields logfields.Collection
	fields.Add(logfields.Extract(x.GetRunner())...)
	return fields
}

func (x *ListRunnersRequest) LogFields() logfields.Collection {
	if x == nil {
		return nil
	}

	var fields logfields.Collection
	fields.Add(logfields.Extract(x.GetPagination())...)
	fields.Add(logfields.Extract(x.GetFilter())...)
	return fields
}

func (x *ListRunnersResponse) LogFields() logfields.Collection {
	if x == nil {
		return nil
	}

	var fields logfields.Collection
	fields.Add(logfields.Extract(x.GetPagination())...)
	return fields
}

func (x *UpdateRunnerRequest) LogFields() logfields.Collection {
	if x == nil {
		return nil
	}

	var fields logfields.Collection
	fields.Add(&logfields.Entry{Name: "runner.id", Value: x.RunnerId})
	fields.Add(logfields.Extract(x.GetSpec())...)
	return fields
}

func (x *UpdateRunnerResponse) LogFields() logfields.Collection {
	return nil
}

func (x *DeleteRunnerRequest) LogFields() logfields.Collection {
	if x == nil {
		return nil
	}

	var fields logfields.Collection
	fields.Add(&logfields.Entry{Name: "runner.id", Value: x.RunnerId})
	return fields
}

func (x *DeleteRunnerResponse) LogFields() logfields.Collection {
	return nil
}

func (x *CreateRunnerTokenRequest) LogFields() logfields.Collection {
	if x == nil {
		return nil
	}

	var fields logfields.Collection
	fields.Add(&logfields.Entry{Name: "runner.id", Value: x.RunnerId})
	return fields
}

func (x *CreateRunnerTokenResponse) LogFields() logfields.Collection {
	return nil
}

func (x *ParseContextURLRequest) LogFields() logfields.Collection {
	if x == nil {
		return nil
	}

	var fields logfields.Collection
	fields.Add(&logfields.Entry{Name: "runner.id", Value: x.RunnerId})
	return fields
}

func (x *ParseContextURLResponse) LogFields() logfields.Collection {
	if x == nil {
		return nil
	}

	var fields logfields.Collection
	fields.Add(logfields.Extract(x.GetGit())...)
	return fields
}

func (x *ParseContextURLPreconditionFailureDetails) LogFields() logfields.Collection {
	return nil
}

func (x *CheckAuthenticationForHostRequest) LogFields() logfields.Collection {
	if x == nil {
		return nil
	}

	var fields logfields.Collection
	fields.Add(&logfields.Entry{Name: "runner.id", Value: x.RunnerId})
	return fields
}

func (x *CheckAuthenticationForHostResponse) LogFields() logfields.Collection {
	return nil
}

func (x *ListRunnerPoliciesRequest) LogFields() logfields.Collection {
	if x == nil {
		return nil
	}

	var fields logfields.Collection
	fields.Add(logfields.Extract(x.GetPagination())...)
	fields.Add(&logfields.Entry{Name: "runner.id", Value: x.RunnerId})
	return fields
}

func (x *ListRunnerPoliciesResponse) LogFields() logfields.Collection {
	if x == nil {
		return nil
	}

	var fields logfields.Collection
	fields.Add(logfields.Extract(x.GetPagination())...)
	return fields
}

func (x *CreateRunnerPolicyRequest) LogFields() logfields.Collection {
	if x == nil {
		return nil
	}

	var fields logfields.Collection
	fields.Add(&logfields.Entry{Name: "runner.id", Value: x.RunnerId})
	fields.Add(&logfields.Entry{Name: "group.id", Value: x.GroupId})
	return fields
}

func (x *CreateRunnerPolicyResponse) LogFields() logfields.Collection {
	if x == nil {
		return nil
	}

	var fields logfields.Collection
	fields.Add(logfields.Extract(x.GetPolicy())...)
	return fields
}

func (x *UpdateRunnerPolicyRequest) LogFields() logfields.Collection {
	if x == nil {
		return nil
	}

	var fields logfields.Collection
	fields.Add(&logfields.Entry{Name: "runner.id", Value: x.RunnerId})
	fields.Add(&logfields.Entry{Name: "group.id", Value: x.GroupId})
	return fields
}

func (x *UpdateRunnerPolicyResponse) LogFields() logfields.Collection {
	if x == nil {
		return nil
	}

	var fields logfields.Collection
	fields.Add(logfields.Extract(x.GetPolicy())...)
	return fields
}

func (x *DeleteRunnerPolicyRequest) LogFields() logfields.Collection {
	if x == nil {
		return nil
	}

	var fields logfields.Collection
	fields.Add(&logfields.Entry{Name: "runner.id", Value: x.RunnerId})
	fields.Add(&logfields.Entry{Name: "group.id", Value: x.GroupId})
	return fields
}

func (x *DeleteRunnerPolicyResponse) LogFields() logfields.Collection {
	return nil
}

func (x *Runner) LogFields() logfields.Collection {
	if x == nil {
		return nil
	}

	var fields logfields.Collection
	fields.Add(&logfields.Entry{Name: "runner.id", Value: x.RunnerId})
	fields.Add(logfields.Extract(x.GetCreatedAt())...)
	fields.Add(logfields.Extract(x.GetUpdatedAt())...)
	fields.Add(logfields.Extract(x.GetSpec())...)
	fields.Add(logfields.Extract(x.GetStatus())...)
	fields.Add(logfields.Extract(x.GetCreator())...)
	return fields
}

func (x *RunnerSpec) LogFields() logfields.Collection {
	if x == nil {
		return nil
	}

	var fields logfields.Collection
	fields.Add(logfields.Extract(x.GetConfiguration())...)
	return fields
}

func (x *RunnerConfiguration) LogFields() logfields.Collection {
	return nil
}

func (x *RunnerStatus) LogFields() logfields.Collection {
	if x == nil {
		return nil
	}

	var fields logfields.Collection
	fields.Add(logfields.Extract(x.GetUpdatedAt())...)
	return fields
}

func (x *RunnerPolicy) LogFields() logfields.Collection {
	if x == nil {
		return nil
	}

	var fields logfields.Collection
	fields.Add(&logfields.Entry{Name: "group.id", Value: x.GroupId})
	return fields
}
