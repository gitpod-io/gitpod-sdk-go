// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"github.com/gitpod-io/gitpod-sdk-go/internal/apierror"
	"github.com/gitpod-io/gitpod-sdk-go/shared"
)

type Error = apierror.Error
type ErrorCode = apierror.ErrorCode

const ErrorCodeCanceled = apierror.ErrorCodeCanceled
const ErrorCodeUnknown = apierror.ErrorCodeUnknown
const ErrorCodeInvalidArgument = apierror.ErrorCodeInvalidArgument
const ErrorCodeDeadlineExceeded = apierror.ErrorCodeDeadlineExceeded
const ErrorCodeNotFound = apierror.ErrorCodeNotFound
const ErrorCodeAlreadyExists = apierror.ErrorCodeAlreadyExists
const ErrorCodePermissionDenied = apierror.ErrorCodePermissionDenied
const ErrorCodeResourceExhausted = apierror.ErrorCodeResourceExhausted
const ErrorCodeFailedPrecondition = apierror.ErrorCodeFailedPrecondition
const ErrorCodeAborted = apierror.ErrorCodeAborted
const ErrorCodeOutOfRange = apierror.ErrorCodeOutOfRange
const ErrorCodeUnimplemented = apierror.ErrorCodeUnimplemented
const ErrorCodeInternal = apierror.ErrorCodeInternal
const ErrorCodeUnavailable = apierror.ErrorCodeUnavailable
const ErrorCodeDataLoss = apierror.ErrorCodeDataLoss
const ErrorCodeUnauthenticated = apierror.ErrorCodeUnauthenticated

// An AutomationTrigger represents a trigger for an automation action. The `manual`
// field shows a start button in the UI for manually triggering the automation. The
// `post_machine_start` field indicates that the automation should be triggered
// after the machine has started, before the devcontainer is ready. This is used
// for machine-level services like security agents that need to start early. The
// `post_environment_start` field indicates that the automation should be triggered
// after the environment has started (devcontainer ready). The
// `post_devcontainer_start` field indicates that the automation should be
// triggered after the dev container has started. The `prebuild` field starts the
// automation during a prebuild of an environment. This phase does not have user
// secrets available. Note: The prebuild trigger can only be used with tasks, not
// services.
//
// This is an alias to an internal type.
type AutomationTrigger = shared.AutomationTrigger

// An AutomationTrigger represents a trigger for an automation action. The `manual`
// field shows a start button in the UI for manually triggering the automation. The
// `post_machine_start` field indicates that the automation should be triggered
// after the machine has started, before the devcontainer is ready. This is used
// for machine-level services like security agents that need to start early. The
// `post_environment_start` field indicates that the automation should be triggered
// after the environment has started (devcontainer ready). The
// `post_devcontainer_start` field indicates that the automation should be
// triggered after the dev container has started. The `prebuild` field starts the
// automation during a prebuild of an environment. This phase does not have user
// secrets available. Note: The prebuild trigger can only be used with tasks, not
// services.
//
// This is an alias to an internal type.
type AutomationTriggerParam = shared.AutomationTriggerParam

// This is an alias to an internal type.
type EnvironmentClass = shared.EnvironmentClass

// This is an alias to an internal type.
type EnvironmentClassParam = shared.EnvironmentClassParam

// EnvironmentVariableItem represents an environment variable that can be set
// either from a literal value or from a secret reference.
//
// This is an alias to an internal type.
type EnvironmentVariableItem = shared.EnvironmentVariableItem

// EnvironmentVariableItem represents an environment variable that can be set
// either from a literal value or from a secret reference.
//
// This is an alias to an internal type.
type EnvironmentVariableItemParam = shared.EnvironmentVariableItemParam

// EnvironmentVariableSource specifies a source for an environment variable value.
//
// This is an alias to an internal type.
type EnvironmentVariableSource = shared.EnvironmentVariableSource

// EnvironmentVariableSource specifies a source for an environment variable value.
//
// This is an alias to an internal type.
type EnvironmentVariableSourceParam = shared.EnvironmentVariableSourceParam

// This is an alias to an internal type.
type FieldValue = shared.FieldValue

// This is an alias to an internal type.
type FieldValueParam = shared.FieldValueParam

// Gateway represents a system gateway that provides access to services
//
// This is an alias to an internal type.
type Gateway = shared.Gateway

// This is an alias to an internal type.
type OrganizationRole = shared.OrganizationRole

// This is an alias to an internal value.
const OrganizationRoleUnspecified = shared.OrganizationRoleUnspecified

// This is an alias to an internal value.
const OrganizationRoleAdmin = shared.OrganizationRoleAdmin

// This is an alias to an internal value.
const OrganizationRoleMember = shared.OrganizationRoleMember

// This is an alias to an internal type.
type OrganizationTier = shared.OrganizationTier

// This is an alias to an internal value.
const OrganizationTierUnspecified = shared.OrganizationTierUnspecified

// This is an alias to an internal value.
const OrganizationTierFree = shared.OrganizationTierFree

// This is an alias to an internal value.
const OrganizationTierEnterprise = shared.OrganizationTierEnterprise

// This is an alias to an internal value.
const OrganizationTierCore = shared.OrganizationTierCore

// This is an alias to an internal value.
const OrganizationTierFreeOna = shared.OrganizationTierFreeOna

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

// This is an alias to an internal value.
const PrincipalRunnerManager = shared.PrincipalRunnerManager

// This is an alias to an internal value.
const PrincipalAgentExecution = shared.PrincipalAgentExecution

// This is an alias to an internal type.
type ProjectEnvironmentClass = shared.ProjectEnvironmentClass

// This is an alias to an internal type.
type ProjectEnvironmentClassParam = shared.ProjectEnvironmentClassParam

// This is an alias to an internal type.
type ResourceType = shared.ResourceType

// This is an alias to an internal value.
const ResourceTypeUnspecified = shared.ResourceTypeUnspecified

// This is an alias to an internal value.
const ResourceTypeEnvironment = shared.ResourceTypeEnvironment

// This is an alias to an internal value.
const ResourceTypeRunner = shared.ResourceTypeRunner

// This is an alias to an internal value.
const ResourceTypeProject = shared.ResourceTypeProject

// This is an alias to an internal value.
const ResourceTypeTask = shared.ResourceTypeTask

// This is an alias to an internal value.
const ResourceTypeTaskExecution = shared.ResourceTypeTaskExecution

// This is an alias to an internal value.
const ResourceTypeService = shared.ResourceTypeService

// This is an alias to an internal value.
const ResourceTypeOrganization = shared.ResourceTypeOrganization

// This is an alias to an internal value.
const ResourceTypeUser = shared.ResourceTypeUser

// This is an alias to an internal value.
const ResourceTypeEnvironmentClass = shared.ResourceTypeEnvironmentClass

// This is an alias to an internal value.
const ResourceTypeRunnerScmIntegration = shared.ResourceTypeRunnerScmIntegration

// This is an alias to an internal value.
const ResourceTypeHostAuthenticationToken = shared.ResourceTypeHostAuthenticationToken

// This is an alias to an internal value.
const ResourceTypeGroup = shared.ResourceTypeGroup

// This is an alias to an internal value.
const ResourceTypePersonalAccessToken = shared.ResourceTypePersonalAccessToken

// This is an alias to an internal value.
const ResourceTypeUserPreference = shared.ResourceTypeUserPreference

// This is an alias to an internal value.
const ResourceTypeServiceAccount = shared.ResourceTypeServiceAccount

// This is an alias to an internal value.
const ResourceTypeSecret = shared.ResourceTypeSecret

// This is an alias to an internal value.
const ResourceTypeSSOConfig = shared.ResourceTypeSSOConfig

// This is an alias to an internal value.
const ResourceTypeDomainVerification = shared.ResourceTypeDomainVerification

// This is an alias to an internal value.
const ResourceTypeAgentExecution = shared.ResourceTypeAgentExecution

// This is an alias to an internal value.
const ResourceTypeRunnerLlmIntegration = shared.ResourceTypeRunnerLlmIntegration

// This is an alias to an internal value.
const ResourceTypeAgent = shared.ResourceTypeAgent

// This is an alias to an internal value.
const ResourceTypeEnvironmentSession = shared.ResourceTypeEnvironmentSession

// This is an alias to an internal value.
const ResourceTypeUserSecret = shared.ResourceTypeUserSecret

// This is an alias to an internal value.
const ResourceTypeOrganizationPolicy = shared.ResourceTypeOrganizationPolicy

// This is an alias to an internal value.
const ResourceTypeOrganizationSecret = shared.ResourceTypeOrganizationSecret

// This is an alias to an internal value.
const ResourceTypeProjectEnvironmentClass = shared.ResourceTypeProjectEnvironmentClass

// This is an alias to an internal value.
const ResourceTypeBilling = shared.ResourceTypeBilling

// This is an alias to an internal value.
const ResourceTypePrompt = shared.ResourceTypePrompt

// This is an alias to an internal value.
const ResourceTypeCoupon = shared.ResourceTypeCoupon

// This is an alias to an internal value.
const ResourceTypeCouponRedemption = shared.ResourceTypeCouponRedemption

// This is an alias to an internal value.
const ResourceTypeAccount = shared.ResourceTypeAccount

// This is an alias to an internal value.
const ResourceTypeIntegration = shared.ResourceTypeIntegration

// This is an alias to an internal value.
const ResourceTypeWorkflow = shared.ResourceTypeWorkflow

// This is an alias to an internal value.
const ResourceTypeWorkflowExecution = shared.ResourceTypeWorkflowExecution

// This is an alias to an internal value.
const ResourceTypeWorkflowExecutionAction = shared.ResourceTypeWorkflowExecutionAction

// This is an alias to an internal value.
const ResourceTypeSnapshot = shared.ResourceTypeSnapshot

// This is an alias to an internal value.
const ResourceTypePrebuild = shared.ResourceTypePrebuild

// This is an alias to an internal value.
const ResourceTypeOrganizationLlmIntegration = shared.ResourceTypeOrganizationLlmIntegration

// This is an alias to an internal value.
const ResourceTypeCustomDomain = shared.ResourceTypeCustomDomain

// This is an alias to an internal value.
const ResourceTypeRoleAssignmentChanged = shared.ResourceTypeRoleAssignmentChanged

// This is an alias to an internal value.
const ResourceTypeGroupMembershipChanged = shared.ResourceTypeGroupMembershipChanged

// This is an alias to an internal value.
const ResourceTypeWebhook = shared.ResourceTypeWebhook

// This is an alias to an internal value.
const ResourceTypeScimConfiguration = shared.ResourceTypeScimConfiguration

// This is an alias to an internal value.
const ResourceTypeServiceAccountSecret = shared.ResourceTypeServiceAccountSecret

// This is an alias to an internal type.
type RunsOn = shared.RunsOn

// This is an alias to an internal type.
type RunsOnDocker = shared.RunsOnDocker

// This is an alias to an internal type.
type RunsOnParam = shared.RunsOnParam

// This is an alias to an internal type.
type RunsOnDockerParam = shared.RunsOnDockerParam

// SecretRef references a secret by its ID.
//
// This is an alias to an internal type.
type SecretRef = shared.SecretRef

// SecretRef references a secret by its ID.
//
// This is an alias to an internal type.
type SecretRefParam = shared.SecretRefParam

// Current state of the pull request
//
// This is an alias to an internal type.
type State = shared.State

// This is an alias to an internal value.
const StateUnspecified = shared.StateUnspecified

// This is an alias to an internal value.
const StateOpen = shared.StateOpen

// This is an alias to an internal value.
const StateClosed = shared.StateClosed

// This is an alias to an internal value.
const StateMerged = shared.StateMerged

// This is an alias to an internal type.
type Subject = shared.Subject

// This is an alias to an internal type.
type SubjectParam = shared.SubjectParam

// This is an alias to an internal type.
type Task = shared.Task

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
type TaskMetadata = shared.TaskMetadata

// This is an alias to an internal type.
type TaskMetadataParam = shared.TaskMetadataParam

// This is an alias to an internal type.
type TaskSpec = shared.TaskSpec

// This is an alias to an internal type.
type TaskSpecParam = shared.TaskSpecParam

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
