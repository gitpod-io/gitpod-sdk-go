// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"github.com/gitpod-io/flex-sdk-go/internal/apierror"
	"github.com/gitpod-io/flex-sdk-go/shared"
)

type Error = apierror.Error

// An AutomationTrigger represents a trigger for an automation action. The
// `post_environment_start` field indicates that the automation should be triggered
// after the environment has started. The `post_devcontainer_start` field indicates
// that the automation should be triggered after the dev container has started.
//
// This is an alias to an internal type.
type AutomationTrigger = shared.AutomationTrigger

// An AutomationTrigger represents a trigger for an automation action. The
// `post_environment_start` field indicates that the automation should be triggered
// after the environment has started. The `post_devcontainer_start` field indicates
// that the automation should be triggered after the dev container has started.
//
// This is an alias to an internal type.
type AutomationTriggerParam = shared.AutomationTriggerParam

// This is an alias to an internal type.
type EnvironmentClass = shared.EnvironmentClass

// This is an alias to an internal type.
type EnvironmentClassParam = shared.EnvironmentClassParam

// This is an alias to an internal type.
type FieldValue = shared.FieldValue

// This is an alias to an internal type.
type FieldValueParam = shared.FieldValueParam

// This is an alias to an internal type.
type OrganizationRole = shared.OrganizationRole

// This is an alias to an internal value.
const OrganizationRoleUnspecified = shared.OrganizationRoleUnspecified

// This is an alias to an internal value.
const OrganizationRoleAdmin = shared.OrganizationRoleAdmin

// This is an alias to an internal value.
const OrganizationRoleMember = shared.OrganizationRoleMember

// This is an alias to an internal type.
type Principal = shared.Principal

// This is an alias to an internal value.
const PrincipalUnspecified = shared.PrincipalUnspecified

// This is an alias to an internal value.
const PrincipalAccount = shared.PrincipalAccount

// This is an alias to an internal value.
const PrincipalUser = shared.PrincipalUser

// This is an alias to an internal value.
const PrincipalRunner = shared.PrincipalRunner

// This is an alias to an internal value.
const PrincipalEnvironment = shared.PrincipalEnvironment

// This is an alias to an internal value.
const PrincipalServiceAccount = shared.PrincipalServiceAccount

// This is an alias to an internal type.
type RunsOn = shared.RunsOn

// This is an alias to an internal type.
type RunsOnDocker = shared.RunsOnDocker

// This is an alias to an internal type.
type RunsOnParam = shared.RunsOnParam

// This is an alias to an internal type.
type RunsOnDockerParam = shared.RunsOnDockerParam

// This is an alias to an internal type.
type Subject = shared.Subject

// This is an alias to an internal type.
type SubjectParam = shared.SubjectParam

// This is an alias to an internal type.
type TaskExecution = shared.TaskExecution

// This is an alias to an internal type.
type TaskExecutionMetadata = shared.TaskExecutionMetadata

// This is an alias to an internal type.
type TaskExecutionPhase = shared.TaskExecutionPhase

// This is an alias to an internal value.
const TaskExecutionPhaseUnspecified = shared.TaskExecutionPhaseUnspecified

// This is an alias to an internal value.
const TaskExecutionPhasePending = shared.TaskExecutionPhasePending

// This is an alias to an internal value.
const TaskExecutionPhaseRunning = shared.TaskExecutionPhaseRunning

// This is an alias to an internal value.
const TaskExecutionPhaseSucceeded = shared.TaskExecutionPhaseSucceeded

// This is an alias to an internal value.
const TaskExecutionPhaseFailed = shared.TaskExecutionPhaseFailed

// This is an alias to an internal value.
const TaskExecutionPhaseStopped = shared.TaskExecutionPhaseStopped

// This is an alias to an internal type.
type TaskExecutionSpec = shared.TaskExecutionSpec

// This is an alias to an internal type.
type TaskExecutionSpecPlan = shared.TaskExecutionSpecPlan

// This is an alias to an internal type.
type TaskExecutionSpecPlanStep = shared.TaskExecutionSpecPlanStep

// This is an alias to an internal type.
type TaskExecutionSpecPlanStepsTask = shared.TaskExecutionSpecPlanStepsTask

// This is an alias to an internal type.
type TaskExecutionStatus = shared.TaskExecutionStatus

// This is an alias to an internal type.
type TaskExecutionStatusStep = shared.TaskExecutionStatusStep

// This is an alias to an internal type.
type UserStatus = shared.UserStatus

// This is an alias to an internal value.
const UserStatusUnspecified = shared.UserStatusUnspecified

// This is an alias to an internal value.
const UserStatusActive = shared.UserStatusActive

// This is an alias to an internal value.
const UserStatusSuspended = shared.UserStatusSuspended

// This is an alias to an internal value.
const UserStatusLeft = shared.UserStatusLeft
