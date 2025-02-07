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
const OrganizationRoleOrganizationRoleUnspecified = shared.OrganizationRoleOrganizationRoleUnspecified

// This is an alias to an internal value.
const OrganizationRoleOrganizationRoleAdmin = shared.OrganizationRoleOrganizationRoleAdmin

// This is an alias to an internal value.
const OrganizationRoleOrganizationRoleMember = shared.OrganizationRoleOrganizationRoleMember

// This is an alias to an internal type.
type Principal = shared.Principal

// This is an alias to an internal value.
const PrincipalPrincipalUnspecified = shared.PrincipalPrincipalUnspecified

// This is an alias to an internal value.
const PrincipalPrincipalAccount = shared.PrincipalPrincipalAccount

// This is an alias to an internal value.
const PrincipalPrincipalUser = shared.PrincipalPrincipalUser

// This is an alias to an internal value.
const PrincipalPrincipalRunner = shared.PrincipalPrincipalRunner

// This is an alias to an internal value.
const PrincipalPrincipalEnvironment = shared.PrincipalPrincipalEnvironment

// This is an alias to an internal value.
const PrincipalPrincipalServiceAccount = shared.PrincipalPrincipalServiceAccount

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
const TaskExecutionPhaseTaskExecutionPhaseUnspecified = shared.TaskExecutionPhaseTaskExecutionPhaseUnspecified

// This is an alias to an internal value.
const TaskExecutionPhaseTaskExecutionPhasePending = shared.TaskExecutionPhaseTaskExecutionPhasePending

// This is an alias to an internal value.
const TaskExecutionPhaseTaskExecutionPhaseRunning = shared.TaskExecutionPhaseTaskExecutionPhaseRunning

// This is an alias to an internal value.
const TaskExecutionPhaseTaskExecutionPhaseSucceeded = shared.TaskExecutionPhaseTaskExecutionPhaseSucceeded

// This is an alias to an internal value.
const TaskExecutionPhaseTaskExecutionPhaseFailed = shared.TaskExecutionPhaseTaskExecutionPhaseFailed

// This is an alias to an internal value.
const TaskExecutionPhaseTaskExecutionPhaseStopped = shared.TaskExecutionPhaseTaskExecutionPhaseStopped

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
const UserStatusUserStatusUnspecified = shared.UserStatusUserStatusUnspecified

// This is an alias to an internal value.
const UserStatusUserStatusActive = shared.UserStatusUserStatusActive

// This is an alias to an internal value.
const UserStatusUserStatusSuspended = shared.UserStatusUserStatusSuspended

// This is an alias to an internal value.
const UserStatusUserStatusLeft = shared.UserStatusUserStatusLeft
