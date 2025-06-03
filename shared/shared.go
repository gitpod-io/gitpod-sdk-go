// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"time"

	"github.com/gitpod-io/gitpod-sdk-go/internal/apijson"
	"github.com/gitpod-io/gitpod-sdk-go/internal/param"
)

// An AutomationTrigger represents a trigger for an automation action. The
// `post_environment_start` field indicates that the automation should be triggered
// after the environment has started. The `post_devcontainer_start` field indicates
// that the automation should be triggered after the dev container has started.
type AutomationTrigger struct {
	Manual                bool                  `json:"manual"`
	PostDevcontainerStart bool                  `json:"postDevcontainerStart"`
	PostEnvironmentStart  bool                  `json:"postEnvironmentStart"`
	JSON                  automationTriggerJSON `json:"-"`
}

// automationTriggerJSON contains the JSON metadata for the struct
// [AutomationTrigger]
type automationTriggerJSON struct {
	Manual                apijson.Field
	PostDevcontainerStart apijson.Field
	PostEnvironmentStart  apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AutomationTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r automationTriggerJSON) RawJSON() string {
	return r.raw
}

// An AutomationTrigger represents a trigger for an automation action. The
// `post_environment_start` field indicates that the automation should be triggered
// after the environment has started. The `post_devcontainer_start` field indicates
// that the automation should be triggered after the dev container has started.
type AutomationTriggerParam struct {
	Manual                param.Field[bool] `json:"manual"`
	PostDevcontainerStart param.Field[bool] `json:"postDevcontainerStart"`
	PostEnvironmentStart  param.Field[bool] `json:"postEnvironmentStart"`
}

func (r AutomationTriggerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentClass struct {
	// id is the unique identifier of the environment class
	ID string `json:"id,required"`
	// runner_id is the unique identifier of the runner the environment class belongs
	// to
	RunnerID string `json:"runnerId,required"`
	// configuration describes the configuration of the environment class
	Configuration []FieldValue `json:"configuration"`
	// description is a human readable description of the environment class
	Description string `json:"description"`
	// display_name is the human readable name of the environment class
	DisplayName string `json:"displayName"`
	// enabled indicates whether the environment class can be used to create new
	// environments.
	Enabled bool                 `json:"enabled"`
	JSON    environmentClassJSON `json:"-"`
}

// environmentClassJSON contains the JSON metadata for the struct
// [EnvironmentClass]
type environmentClassJSON struct {
	ID            apijson.Field
	RunnerID      apijson.Field
	Configuration apijson.Field
	Description   apijson.Field
	DisplayName   apijson.Field
	Enabled       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *EnvironmentClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentClassJSON) RawJSON() string {
	return r.raw
}

type EnvironmentClassParam struct {
	// id is the unique identifier of the environment class
	ID param.Field[string] `json:"id,required"`
	// runner_id is the unique identifier of the runner the environment class belongs
	// to
	RunnerID param.Field[string] `json:"runnerId,required"`
	// configuration describes the configuration of the environment class
	Configuration param.Field[[]FieldValueParam] `json:"configuration"`
	// description is a human readable description of the environment class
	Description param.Field[string] `json:"description"`
	// display_name is the human readable name of the environment class
	DisplayName param.Field[string] `json:"displayName"`
	// enabled indicates whether the environment class can be used to create new
	// environments.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r EnvironmentClassParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FieldValue struct {
	Key   string         `json:"key"`
	Value string         `json:"value"`
	JSON  fieldValueJSON `json:"-"`
}

// fieldValueJSON contains the JSON metadata for the struct [FieldValue]
type fieldValueJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FieldValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fieldValueJSON) RawJSON() string {
	return r.raw
}

type FieldValueParam struct {
	Key   param.Field[string] `json:"key"`
	Value param.Field[string] `json:"value"`
}

func (r FieldValueParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Gateway represents a system gateway that provides access to services
type Gateway struct {
	// name is the human-readable name of the gateway. name is unique across all
	// gateways.
	Name string `json:"name,required"`
	// url of the gateway
	URL string `json:"url,required"`
	// region is the geographical region where the gateway is located
	Region string      `json:"region"`
	JSON   gatewayJSON `json:"-"`
}

// gatewayJSON contains the JSON metadata for the struct [Gateway]
type gatewayJSON struct {
	Name        apijson.Field
	URL         apijson.Field
	Region      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Gateway) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayJSON) RawJSON() string {
	return r.raw
}

type OrganizationRole string

const (
	OrganizationRoleUnspecified OrganizationRole = "ORGANIZATION_ROLE_UNSPECIFIED"
	OrganizationRoleAdmin       OrganizationRole = "ORGANIZATION_ROLE_ADMIN"
	OrganizationRoleMember      OrganizationRole = "ORGANIZATION_ROLE_MEMBER"
)

func (r OrganizationRole) IsKnown() bool {
	switch r {
	case OrganizationRoleUnspecified, OrganizationRoleAdmin, OrganizationRoleMember:
		return true
	}
	return false
}

type Principal string

const (
	PrincipalUnspecified    Principal = "PRINCIPAL_UNSPECIFIED"
	PrincipalAccount        Principal = "PRINCIPAL_ACCOUNT"
	PrincipalUser           Principal = "PRINCIPAL_USER"
	PrincipalRunner         Principal = "PRINCIPAL_RUNNER"
	PrincipalEnvironment    Principal = "PRINCIPAL_ENVIRONMENT"
	PrincipalServiceAccount Principal = "PRINCIPAL_SERVICE_ACCOUNT"
	PrincipalRunnerManager  Principal = "PRINCIPAL_RUNNER_MANAGER"
)

func (r Principal) IsKnown() bool {
	switch r {
	case PrincipalUnspecified, PrincipalAccount, PrincipalUser, PrincipalRunner, PrincipalEnvironment, PrincipalServiceAccount, PrincipalRunnerManager:
		return true
	}
	return false
}

type RunsOn struct {
	Docker RunsOnDocker `json:"docker,required"`
	JSON   runsOnJSON   `json:"-"`
}

// runsOnJSON contains the JSON metadata for the struct [RunsOn]
type runsOnJSON struct {
	Docker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunsOn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runsOnJSON) RawJSON() string {
	return r.raw
}

type RunsOnDocker struct {
	Environment []string         `json:"environment"`
	Image       string           `json:"image"`
	JSON        runsOnDockerJSON `json:"-"`
}

// runsOnDockerJSON contains the JSON metadata for the struct [RunsOnDocker]
type runsOnDockerJSON struct {
	Environment apijson.Field
	Image       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunsOnDocker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runsOnDockerJSON) RawJSON() string {
	return r.raw
}

type RunsOnParam struct {
	Docker param.Field[RunsOnDockerParam] `json:"docker,required"`
}

func (r RunsOnParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunsOnDockerParam struct {
	Environment param.Field[[]string] `json:"environment"`
	Image       param.Field[string]   `json:"image"`
}

func (r RunsOnDockerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Subject struct {
	// id is the UUID of the subject
	ID string `json:"id"`
	// Principal is the principal of the subject
	Principal Principal   `json:"principal"`
	JSON      subjectJSON `json:"-"`
}

// subjectJSON contains the JSON metadata for the struct [Subject]
type subjectJSON struct {
	ID          apijson.Field
	Principal   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Subject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subjectJSON) RawJSON() string {
	return r.raw
}

type SubjectParam struct {
	// id is the UUID of the subject
	ID param.Field[string] `json:"id"`
	// Principal is the principal of the subject
	Principal param.Field[Principal] `json:"principal"`
}

func (r SubjectParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Task struct {
	ID string `json:"id,required" format:"uuid"`
	// dependencies specifies the IDs of the automations this task depends on.
	DependsOn     []string     `json:"dependsOn" format:"uuid"`
	EnvironmentID string       `json:"environmentId" format:"uuid"`
	Metadata      TaskMetadata `json:"metadata"`
	Spec          TaskSpec     `json:"spec"`
	JSON          taskJSON     `json:"-"`
}

// taskJSON contains the JSON metadata for the struct [Task]
type taskJSON struct {
	ID            apijson.Field
	DependsOn     apijson.Field
	EnvironmentID apijson.Field
	Metadata      apijson.Field
	Spec          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *Task) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskJSON) RawJSON() string {
	return r.raw
}

type TaskExecution struct {
	ID       string                `json:"id,required" format:"uuid"`
	Metadata TaskExecutionMetadata `json:"metadata"`
	Spec     TaskExecutionSpec     `json:"spec"`
	Status   TaskExecutionStatus   `json:"status"`
	JSON     taskExecutionJSON     `json:"-"`
}

// taskExecutionJSON contains the JSON metadata for the struct [TaskExecution]
type taskExecutionJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Spec        apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TaskExecution) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskExecutionJSON) RawJSON() string {
	return r.raw
}

type TaskExecutionMetadata struct {
	// completed_at is the time the task execution was done.
	CompletedAt time.Time `json:"completedAt" format:"date-time"`
	// created_at is the time the task was created.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// creator describes the principal who created/started the task run.
	Creator Subject `json:"creator"`
	// environment_id is the ID of the environment in which the task run is executed.
	EnvironmentID string `json:"environmentId" format:"uuid"`
	// started_at is the time the task execution actually started to run.
	StartedAt time.Time `json:"startedAt" format:"date-time"`
	// started_by describes the trigger that started the task execution.
	StartedBy string `json:"startedBy"`
	// task_id is the ID of the main task being executed.
	TaskID string                    `json:"taskId" format:"uuid"`
	JSON   taskExecutionMetadataJSON `json:"-"`
}

// taskExecutionMetadataJSON contains the JSON metadata for the struct
// [TaskExecutionMetadata]
type taskExecutionMetadataJSON struct {
	CompletedAt   apijson.Field
	CreatedAt     apijson.Field
	Creator       apijson.Field
	EnvironmentID apijson.Field
	StartedAt     apijson.Field
	StartedBy     apijson.Field
	TaskID        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TaskExecutionMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskExecutionMetadataJSON) RawJSON() string {
	return r.raw
}

type TaskExecutionPhase string

const (
	TaskExecutionPhaseUnspecified TaskExecutionPhase = "TASK_EXECUTION_PHASE_UNSPECIFIED"
	TaskExecutionPhasePending     TaskExecutionPhase = "TASK_EXECUTION_PHASE_PENDING"
	TaskExecutionPhaseRunning     TaskExecutionPhase = "TASK_EXECUTION_PHASE_RUNNING"
	TaskExecutionPhaseSucceeded   TaskExecutionPhase = "TASK_EXECUTION_PHASE_SUCCEEDED"
	TaskExecutionPhaseFailed      TaskExecutionPhase = "TASK_EXECUTION_PHASE_FAILED"
	TaskExecutionPhaseStopped     TaskExecutionPhase = "TASK_EXECUTION_PHASE_STOPPED"
)

func (r TaskExecutionPhase) IsKnown() bool {
	switch r {
	case TaskExecutionPhaseUnspecified, TaskExecutionPhasePending, TaskExecutionPhaseRunning, TaskExecutionPhaseSucceeded, TaskExecutionPhaseFailed, TaskExecutionPhaseStopped:
		return true
	}
	return false
}

type TaskExecutionSpec struct {
	// desired_phase is the phase the task execution should be in. Used to stop a
	// running task execution early.
	DesiredPhase TaskExecutionPhase `json:"desiredPhase"`
	// plan is a list of groups of steps. The steps in a group are executed
	// concurrently, while the groups are executed sequentially. The order of the
	// groups is the order in which they are executed.
	Plan []TaskExecutionSpecPlan `json:"plan"`
	JSON taskExecutionSpecJSON   `json:"-"`
}

// taskExecutionSpecJSON contains the JSON metadata for the struct
// [TaskExecutionSpec]
type taskExecutionSpecJSON struct {
	DesiredPhase apijson.Field
	Plan         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *TaskExecutionSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskExecutionSpecJSON) RawJSON() string {
	return r.raw
}

type TaskExecutionSpecPlan struct {
	Steps []TaskExecutionSpecPlanStep `json:"steps"`
	JSON  taskExecutionSpecPlanJSON   `json:"-"`
}

// taskExecutionSpecPlanJSON contains the JSON metadata for the struct
// [TaskExecutionSpecPlan]
type taskExecutionSpecPlanJSON struct {
	Steps       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TaskExecutionSpecPlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskExecutionSpecPlanJSON) RawJSON() string {
	return r.raw
}

type TaskExecutionSpecPlanStep struct {
	// ID is the ID of the execution step
	ID        string                         `json:"id" format:"uuid"`
	DependsOn []string                       `json:"dependsOn"`
	Label     string                         `json:"label"`
	ServiceID string                         `json:"serviceId" format:"uuid"`
	Task      TaskExecutionSpecPlanStepsTask `json:"task"`
	JSON      taskExecutionSpecPlanStepJSON  `json:"-"`
}

// taskExecutionSpecPlanStepJSON contains the JSON metadata for the struct
// [TaskExecutionSpecPlanStep]
type taskExecutionSpecPlanStepJSON struct {
	ID          apijson.Field
	DependsOn   apijson.Field
	Label       apijson.Field
	ServiceID   apijson.Field
	Task        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TaskExecutionSpecPlanStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskExecutionSpecPlanStepJSON) RawJSON() string {
	return r.raw
}

type TaskExecutionSpecPlanStepsTask struct {
	ID   string                             `json:"id" format:"uuid"`
	Spec TaskSpec                           `json:"spec"`
	JSON taskExecutionSpecPlanStepsTaskJSON `json:"-"`
}

// taskExecutionSpecPlanStepsTaskJSON contains the JSON metadata for the struct
// [TaskExecutionSpecPlanStepsTask]
type taskExecutionSpecPlanStepsTaskJSON struct {
	ID          apijson.Field
	Spec        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TaskExecutionSpecPlanStepsTask) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskExecutionSpecPlanStepsTaskJSON) RawJSON() string {
	return r.raw
}

type TaskExecutionStatus struct {
	// failure_message summarises why the task execution failed to operate. If this is
	// non-empty the task execution has failed to operate and will likely transition to
	// a failed state.
	FailureMessage string `json:"failureMessage"`
	// log_url is the URL to the logs of the task's steps. If this is empty, the task
	// either has no logs or has not yet started.
	LogURL string `json:"logUrl"`
	// the phase of a task execution represents the aggregated phase of all steps.
	Phase TaskExecutionPhase `json:"phase"`
	// version of the status update. Task executions themselves are unversioned, but
	// their status has different versions. The value of this field has no semantic
	// meaning (e.g. don't interpret it as as a timestamp), but it can be used to
	// impose a partial order. If a.status_version < b.status_version then a was the
	// status before b.
	StatusVersion string `json:"statusVersion"`
	// steps provides the status for each individual step of the task execution. If a
	// step is missing it has not yet started.
	Steps []TaskExecutionStatusStep `json:"steps"`
	JSON  taskExecutionStatusJSON   `json:"-"`
}

// taskExecutionStatusJSON contains the JSON metadata for the struct
// [TaskExecutionStatus]
type taskExecutionStatusJSON struct {
	FailureMessage apijson.Field
	LogURL         apijson.Field
	Phase          apijson.Field
	StatusVersion  apijson.Field
	Steps          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TaskExecutionStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskExecutionStatusJSON) RawJSON() string {
	return r.raw
}

type TaskExecutionStatusStep struct {
	// ID is the ID of the execution step
	ID string `json:"id" format:"uuid"`
	// failure_message summarises why the step failed to operate. If this is non-empty
	// the step has failed to operate and will likely transition to a failed state.
	FailureMessage string `json:"failureMessage"`
	// output contains the output of the task execution. setting an output field to
	// empty string will unset it.
	Output map[string]string `json:"output"`
	// phase is the current phase of the execution step
	Phase TaskExecutionPhase          `json:"phase"`
	JSON  taskExecutionStatusStepJSON `json:"-"`
}

// taskExecutionStatusStepJSON contains the JSON metadata for the struct
// [TaskExecutionStatusStep]
type taskExecutionStatusStepJSON struct {
	ID             apijson.Field
	FailureMessage apijson.Field
	Output         apijson.Field
	Phase          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TaskExecutionStatusStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskExecutionStatusStepJSON) RawJSON() string {
	return r.raw
}

type TaskMetadata struct {
	// created_at is the time the task was created.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// creator describes the principal who created the task.
	Creator Subject `json:"creator"`
	// description is a user-facing description for the task. It can be used to provide
	// context and documentation for the task.
	Description string `json:"description"`
	// name is a user-facing name for the task. Unlike the reference, this field is not
	// unique, and not referenced by the system. This is a short descriptive name for
	// the task.
	Name string `json:"name"`
	// reference is a user-facing identifier for the task which must be unique on the
	// environment. It is used to express dependencies between tasks, and to identify
	// the task in user interactions (e.g. the CLI).
	Reference string `json:"reference"`
	// triggered_by is a list of trigger that start the task.
	TriggeredBy []AutomationTrigger `json:"triggeredBy"`
	JSON        taskMetadataJSON    `json:"-"`
}

// taskMetadataJSON contains the JSON metadata for the struct [TaskMetadata]
type taskMetadataJSON struct {
	CreatedAt   apijson.Field
	Creator     apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Reference   apijson.Field
	TriggeredBy apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TaskMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskMetadataJSON) RawJSON() string {
	return r.raw
}

type TaskMetadataParam struct {
	// created_at is the time the task was created.
	CreatedAt param.Field[time.Time] `json:"createdAt" format:"date-time"`
	// creator describes the principal who created the task.
	Creator param.Field[SubjectParam] `json:"creator"`
	// description is a user-facing description for the task. It can be used to provide
	// context and documentation for the task.
	Description param.Field[string] `json:"description"`
	// name is a user-facing name for the task. Unlike the reference, this field is not
	// unique, and not referenced by the system. This is a short descriptive name for
	// the task.
	Name param.Field[string] `json:"name"`
	// reference is a user-facing identifier for the task which must be unique on the
	// environment. It is used to express dependencies between tasks, and to identify
	// the task in user interactions (e.g. the CLI).
	Reference param.Field[string] `json:"reference"`
	// triggered_by is a list of trigger that start the task.
	TriggeredBy param.Field[[]AutomationTriggerParam] `json:"triggeredBy"`
}

func (r TaskMetadataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TaskSpec struct {
	// command contains the command the task should execute
	Command string `json:"command"`
	// runs_on specifies the environment the task should run on.
	RunsOn RunsOn       `json:"runsOn"`
	JSON   taskSpecJSON `json:"-"`
}

// taskSpecJSON contains the JSON metadata for the struct [TaskSpec]
type taskSpecJSON struct {
	Command     apijson.Field
	RunsOn      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TaskSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskSpecJSON) RawJSON() string {
	return r.raw
}

type TaskSpecParam struct {
	// command contains the command the task should execute
	Command param.Field[string] `json:"command"`
	// runs_on specifies the environment the task should run on.
	RunsOn param.Field[RunsOnParam] `json:"runsOn"`
}

func (r TaskSpecParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserStatus string

const (
	UserStatusUnspecified UserStatus = "USER_STATUS_UNSPECIFIED"
	UserStatusActive      UserStatus = "USER_STATUS_ACTIVE"
	UserStatusSuspended   UserStatus = "USER_STATUS_SUSPENDED"
	UserStatusLeft        UserStatus = "USER_STATUS_LEFT"
)

func (r UserStatus) IsKnown() bool {
	switch r {
	case UserStatusUnspecified, UserStatusActive, UserStatusSuspended, UserStatusLeft:
		return true
	}
	return false
}
