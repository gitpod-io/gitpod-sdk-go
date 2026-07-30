# Changelog

## 0.14.0 (2026-07-30)

Full Changelog: [v0.13.0...v0.14.0](https://github.com/gitpod-io/gitpod-sdk-go/compare/v0.13.0...v0.14.0)

### ⚠ BREAKING CHANGES

* **api:** narrow public SDK surface
* **agent:** remove loop condition mechanism

### Features

* add base snapshot resource types ([896bc63](https://github.com/gitpod-io/gitpod-sdk-go/commit/896bc6335a2f5b845756174348d49e82d54026cd))
* **agent:** add devcontainer rebuild wait schema ([3182ece](https://github.com/gitpod-io/gitpod-sdk-go/commit/3182ece8209938e7ac9ac1a668a39b18d138019c))
* **agents:** add native goal set control ([1a6d658](https://github.com/gitpod-io/gitpod-sdk-go/commit/1a6d6586181d43513cb37ee4d2b6c68589b384d8))
* **api:** add agent turn controls ([570d9de](https://github.com/gitpod-io/gitpod-sdk-go/commit/570d9ded37b8b7ca4b8e05f7a8a938bdacb6420b))
* **api:** add agent_execution_cnf value to RunnerCapability enum ([a2b9e42](https://github.com/gitpod-io/gitpod-sdk-go/commit/a2b9e4207c41110bf2bd7b4122cf9266babe1204))
* **api:** add agent_execution_ids filter to agent list executions ([f5698dd](https://github.com/gitpod-io/gitpod-sdk-go/commit/f5698dd36af6abc0c7a71c3e427d7e53c182d35a))
* **api:** add allow_unverified_email_scim_fallback_match to scim_configurations ([23c9014](https://github.com/gitpod-io/gitpod-sdk-go/commit/23c90146a14ee8ab1c12135caf0f09c37473d3e6))
* **api:** add base snapshot build environment role ([aadee26](https://github.com/gitpod-io/gitpod-sdk-go/commit/aadee262c4c70167da3ad1feb44a92f08a686a61))
* **api:** add base snapshot runner capability ([1cb9b25](https://github.com/gitpod-io/gitpod-sdk-go/commit/1cb9b25ec59d98cc0170d8c70c4891864efae7bc))
* **api:** add Claude Sonnet 5 model enum ([29243cf](https://github.com/gitpod-io/gitpod-sdk-go/commit/29243cf9ff545b0bffbdc3bcaa7d021225199393))
* **api:** add Codex agent settings shape ([4ed1e07](https://github.com/gitpod-io/gitpod-sdk-go/commit/4ed1e074c73c1444b752b9fcbbd92c085b6e4a4c))
* **api:** add Codex fast mode setting ([16756cf](https://github.com/gitpod-io/gitpod-sdk-go/commit/16756cfd0cc71e06e113a4fbfcdf8be2b61ba18c))
* **api:** add codex model policy states ([d5b77e6](https://github.com/gitpod-io/gitpod-sdk-go/commit/d5b77e64815e2059913bd13e71174f737325b551))
* **api:** add Codex policy allowlists ([38ad2bb](https://github.com/gitpod-io/gitpod-sdk-go/commit/38ad2bbb80fc7531dd6abb258bc7311ca46e0aab))
* **api:** add credential_proxy to secrets, remove format from environment secrets ([4cb045e](https://github.com/gitpod-io/gitpod-sdk-go/commit/4cb045ee2a236cdba1f43cf3e28d9abafe434801))
* **api:** add default environment class IDs to project creation defaults ([ca59a8b](https://github.com/gitpod-io/gitpod-sdk-go/commit/ca59a8bfab949948ee14c51ec14c31711c3a10e9))
* **api:** add deprecated accessToken field to runner response types ([0a1c411](https://github.com/gitpod-io/gitpod-sdk-go/commit/0a1c4110fbda059d061363db6ef2002b8cf67024))
* **api:** add disable_from_scratch field to organization policy ([2cfc5d6](https://github.com/gitpod-io/gitpod-sdk-go/commit/2cfc5d60cbe13cae3a36f98cc0ee612d66f80437))
* **api:** add dynamic LLM header contracts ([8add73b](https://github.com/gitpod-io/gitpod-sdk-go/commit/8add73b06c30395101d4fdbacdf6b7f74c34ac44))
* **api:** add goal field to agent execution status ([9675b1f](https://github.com/gitpod-io/gitpod-sdk-go/commit/9675b1fe7c27d8780fd798bb2f6300aaef4e57d7))
* **api:** add incident trigger support to workflow triggers and executions ([2d74b1b](https://github.com/gitpod-io/gitpod-sdk-go/commit/2d74b1b1e2557b899477c2e16de80d3dc200297a))
* **api:** add integration_id, update webhook_id in pull request trigger ([71d164a](https://github.com/gitpod-io/gitpod-sdk-go/commit/71d164aa1e7e29b16ee8531437e1e9969933aa08))
* **api:** add MaxPortAdmissionLevel field to organization policy ([395d6af](https://github.com/gitpod-io/gitpod-sdk-go/commit/395d6af87f23155027dc42519e5b7b7245af84e2))
* **api:** add old_path field to environment git changed file model ([9fe1925](https://github.com/gitpod-io/gitpod-sdk-go/commit/9fe19258fc36ca036af7788f5f6a9999734c2563))
* **api:** add OpenAI Auto value to AgentExecutionStatusSupportedModel enum ([ebeb95e](https://github.com/gitpod-io/gitpod-sdk-go/commit/ebeb95e3b757fac267685faa725d38a277901fff))
* **api:** add Opus 4.8 supported model ([0b2c1fb](https://github.com/gitpod-io/gitpod-sdk-go/commit/0b2c1fb482f4daaa582929312fe00918947720d2))
* **api:** add OPUS_4_7 to agent supported models enum ([d2a744b](https://github.com/gitpod-io/gitpod-sdk-go/commit/d2a744be9876078da2a967eea5b71e6d10709fca))
* **api:** add pagination support and query parameter to runners.ListScmOrganizations ([c5ba3fd](https://github.com/gitpod-io/gitpod-sdk-go/commit/c5ba3fdd26a149c7cbec69a8e938eabbc4afd85d))
* **api:** add port_authentication capability to RunnerCapability ([6e3178e](https://github.com/gitpod-io/gitpod-sdk-go/commit/6e3178eff7946662bdd2aede907d7d86b4d003f1))
* **api:** add prebuild_requires_success field to TaskSpec ([7acaa8f](https://github.com/gitpod-io/gitpod-sdk-go/commit/7acaa8f0604a85a56588a14016ef91b1e3c8fe7b))
* **api:** add project_creation_defaults field to organization policy ([8c14eb0](https://github.com/gitpod-io/gitpod-sdk-go/commit/8c14eb081e6ee7cd91126d4e0e5a7e6eafc8cf91))
* **api:** add readinessTimeout parameter to automation services ([71ef113](https://github.com/gitpod-io/gitpod-sdk-go/commit/71ef113396de91acb089a6a04593b780921f877a))
* **api:** add redis_stream capability to runner ([0ff3d08](https://github.com/gitpod-io/gitpod-sdk-go/commit/0ff3d08535c6a90cae7be6729f1711a4f6a6da2b))
* **api:** add ResourceRoleOrgEnvironmentsReader to ResourceRole enum ([9693bfa](https://github.com/gitpod-io/gitpod-sdk-go/commit/9693bfa0c4843da9a589b7a46d4680a946631c92))
* **api:** add review_requested event to workflow trigger pull request events ([ac58c85](https://github.com/gitpod-io/gitpod-sdk-go/commit/ac58c8578b8564674f5204cadbee1a3ce89e5b94))
* **api:** add security policy CRUD ([22b4985](https://github.com/gitpod-io/gitpod-sdk-go/commit/22b49851741dc41ead2ac8c0f1bfc33d5bcabe29))
* **api:** add web_browser_disabled field to organization policy ([380a20e](https://github.com/gitpod-io/gitpod-sdk-go/commit/380a20e93b7c60a3fdbb8c46fbb6f1a645e10238))
* **api:** expose audit log entry kind ([049c177](https://github.com/gitpod-io/gitpod-sdk-go/commit/049c1779cf7fbd4dd1dc2aa6de42fcbe1afc119d))
* **api:** expose default security policy assignment ([72a8ea5](https://github.com/gitpod-io/gitpod-sdk-go/commit/72a8ea5a8232d2654a6ca4b8427a2e67355977be))
* **api:** expose usage insights API ([eb4cf35](https://github.com/gitpod-io/gitpod-sdk-go/commit/eb4cf35646fd6f20d9ac72b54520e7ad5cb0baff))
* **api:** expose Veto Exec security policies ([7c48ec7](https://github.com/gitpod-io/gitpod-sdk-go/commit/7c48ec7282a1ddd976ef196fe6bad0226d4d142f))
* **api:** remove deprecated AccessToken field from runner responses ([79d8fcf](https://github.com/gitpod-io/gitpod-sdk-go/commit/79d8fcfe9d3591bf13fff2d0f79336242f36e733))
* **api:** remove Terminal field from RunsOn type ([5b02fe5](https://github.com/gitpod-io/gitpod-sdk-go/commit/5b02fe5111f66ac18a52efb079dcc0a86d82181a))
* **audit:** curate audit events at their producers ([e86e47b](https://github.com/gitpod-io/gitpod-sdk-go/commit/e86e47b4dbbb9cc0eb5a8e914e4b6ab2ad9927b0))
* **audit:** store and retrieve Veto Exec details ([b268002](https://github.com/gitpod-io/gitpod-sdk-go/commit/b26800200e449b63b61d18a27612f0f15fd8cae1))
* **automations:** pin agent and Codex model per automation ([625dcf5](https://github.com/gitpod-io/gitpod-sdk-go/commit/625dcf5893d8b04bb2af8bcc6059b0ad25930776))
* **codex:** add GPT-5.6 model selection ([ace7d76](https://github.com/gitpod-io/gitpod-sdk-go/commit/ace7d7690d04989e74d396b952f5bccdbed20d38))
* **go:** add default http client with timeout ([5aa4fca](https://github.com/gitpod-io/gitpod-sdk-go/commit/5aa4fca6ee96b6fc19569aa1f419cbe53517d61e))
* **rbac:** add billing viewer role ([29562e3](https://github.com/gitpod-io/gitpod-sdk-go/commit/29562e3be0ea996cbe05b201c62e5e430b8ccf38))
* **rbac:** add insights viewer role ([9d7bc14](https://github.com/gitpod-io/gitpod-sdk-go/commit/9d7bc14ad0192d122ec4e24526d69314b8a6b37d))
* **secrets:** support JFrog OIDC secret sources ([bd41640](https://github.com/gitpod-io/gitpod-sdk-go/commit/bd416406c95e41ddbfdee43b2ecdd1037ccf9792))
* **stlc:** configurable CI runner and private-production-repo support in workflow templates ([852cade](https://github.com/gitpod-io/gitpod-sdk-go/commit/852cade5831dcf39436267266556ea1a71b696ae))
* support setting headers via env ([a05eb6b](https://github.com/gitpod-io/gitpod-sdk-go/commit/a05eb6b7268ba0e12c80549506c4a8a01b6d30b5))
* **types:** add prebuild value to AutomationsFileServicesTriggeredBy enum ([efe0872](https://github.com/gitpod-io/gitpod-sdk-go/commit/efe08724860d1d93c1782725f688a588b755be40))
* **types:** extract Report types to top-level from WorkflowStep ([d58de25](https://github.com/gitpod-io/gitpod-sdk-go/commit/d58de258a0f8f2f3e1400778b07b91d4ff5ee91e))


### Bug Fixes

* **api:** narrow public SDK surface ([e2a4ec0](https://github.com/gitpod-io/gitpod-sdk-go/commit/e2a4ec0d63fc3262945dee9776f573a3422af66e))
* **api:** wrap prebuild default updates ([7d0e27a](https://github.com/gitpod-io/gitpod-sdk-go/commit/7d0e27ace5c69cb51ff7c0fa943a3019678b1ca1))
* **go:** avoid panic when http.DefaultTransport is wrapped ([b5215b9](https://github.com/gitpod-io/gitpod-sdk-go/commit/b5215b91002d11266b5c3eb948c59bc03e115aee))


### Chores

* avoid embedding reflect.Type for dead code elimination ([c03525f](https://github.com/gitpod-io/gitpod-sdk-go/commit/c03525f5f6ef24ce53d1af2b4f3995df1d670e14))
* **internal:** more robust bootstrap script ([7d8a76e](https://github.com/gitpod-io/gitpod-sdk-go/commit/7d8a76ef4e36d2cfba13489347ec03ec5ba83565))
* **internal:** regenerate SDK with no functional changes ([32e924f](https://github.com/gitpod-io/gitpod-sdk-go/commit/32e924f8ceecdfcf58597b59c671b1489e47cf80))
* **internal:** regenerate SDK with no functional changes ([f8b376e](https://github.com/gitpod-io/gitpod-sdk-go/commit/f8b376ef7d1eebbdf66252e7ec29893d2fc1541f))
* redact api-key headers in debug logs ([3ba5280](https://github.com/gitpod-io/gitpod-sdk-go/commit/3ba5280f40993a3b29e630881257c844aca5c06b))


### Documentation

* **api:** add customer-facing descriptions to runner enums and fix CreateRunner examples ([a71730b](https://github.com/gitpod-io/gitpod-sdk-go/commit/a71730b1161024a42873d2bd06728dddda2cb8a3))
* **api:** deprecate isAdmin field in runner list_scm_organizations response ([ed8323d](https://github.com/gitpod-io/gitpod-sdk-go/commit/ed8323d73d95edceb3e445698adde781498a6a58))


### Refactors

* **agent:** remove loop condition mechanism ([0c5e903](https://github.com/gitpod-io/gitpod-sdk-go/commit/0c5e9037f7f6c1c414e46b287463ca2157a048c9))
* **api:** default StartAgent agent ID ([ebc1c49](https://github.com/gitpod-io/gitpod-sdk-go/commit/ebc1c493843892f2dc5db9c972357168175cdc63))
* **api:** remove project defaults API surface ([37623df](https://github.com/gitpod-io/gitpod-sdk-go/commit/37623df41e541fccc175b00c475b7635fbef1825))

## 0.13.0 (2026-04-02)

Full Changelog: [v0.12.0...v0.13.0](https://github.com/gitpod-io/gitpod-sdk-go/compare/v0.12.0...v0.13.0)

### Features

* [backend/api] add backend search for group members ([49c3d91](https://github.com/gitpod-io/gitpod-sdk-go/commit/49c3d916444caae5dcc07f121455482195133bd8))
* **api:** add acknowledge_token parameter to environment start method ([ac70010](https://github.com/gitpod-io/gitpod-sdk-go/commit/ac7001022051285de9b32ffef87c709f5af1ebcc))
* **api:** add agent_message parameter to agents.send_to_execution ([3621415](https://github.com/gitpod-io/gitpod-sdk-go/commit/362141594f1753f4a593358ff70beeac558d7469))
* **api:** add AgentExecutionViewer to ResourceRole enum ([710fa8b](https://github.com/gitpod-io/gitpod-sdk-go/commit/710fa8bb14ec0285d4199b04dc5d8878017c4bae))
* **api:** add ASG warm pool capability to runner ([eb8feed](https://github.com/gitpod-io/gitpod-sdk-go/commit/eb8feede61b4578197b0e9267251d973b73a0f96))
* **api:** add automations API, move sort types to shared ([f504dd5](https://github.com/gitpod-io/gitpod-sdk-go/commit/f504dd5b650a32739711f156e01ed02ccb9e8672))
* **api:** add claims_expression field to organization SSO configuration ([1667cf8](https://github.com/gitpod-io/gitpod-sdk-go/commit/1667cf8cd216030769e53589b94adc3e16d2f173))
* **api:** add conversation_sharing_policy to agent policy ([15724f3](https://github.com/gitpod-io/gitpod-sdk-go/commit/15724f3f6623c8e9e7ede8b4f5fa7059c4f7c184))
* **api:** add credential_proxy field to environment secrets ([825707d](https://github.com/gitpod-io/gitpod-sdk-go/commit/825707d29e2c22e1e10ef8735613639760aa270b))
* **api:** add desired_size field to WarmPoolStatus ([5bda4f8](https://github.com/gitpod-io/gitpod-sdk-go/commit/5bda4f88ca1cc834df41be493805dbf8b9c80df7))
* **api:** add disabled parameter to automation update method ([dce6505](https://github.com/gitpod-io/gitpod-sdk-go/commit/dce6505a3cde12a211f222ad619e821461e4d158))
* **api:** add environment field to agent wake events ([2f676cf](https://github.com/gitpod-io/gitpod-sdk-go/commit/2f676cf71ddf36498f388fb6fe5e445afe603ee5))
* **api:** add exclude_prompt_content parameter to agent list prompts ([f2cf0ac](https://github.com/gitpod-io/gitpod-sdk-go/commit/f2cf0ac9ac9cc407c4dde90991558fdbd776d110))
* **api:** add exclude_team_ids parameter to organization list members ([b5e305e](https://github.com/gitpod-io/gitpod-sdk-go/commit/b5e305e1ace58ab6586efaae8d438077d9a848a6))
* **api:** add filter parameter to groups list method ([14b0c2a](https://github.com/gitpod-io/gitpod-sdk-go/commit/14b0c2adf6c9d3c5c184d14170a1ceac5b8d8c6d))
* **api:** add has_failed_execution_since filter to automation list ([5a78576](https://github.com/gitpod-io/gitpod-sdk-go/commit/5a785768a421f90e5205ec4303c08822f40c0b78))
* **api:** add lockdown_at field to environment status ([34081af](https://github.com/gitpod-io/gitpod-sdk-go/commit/34081afd3b244d3f9a18f2aab6c3f200297c4039))
* **api:** add loop conditions/retrigger to agent spec and wake event ([adcf2de](https://github.com/gitpod-io/gitpod-sdk-go/commit/adcf2de763ac97f072cf24825996ca75f8f62520))
* **api:** add ManagedMetricsEnabled field to MetricsConfiguration ([1821753](https://github.com/gitpod-io/gitpod-sdk-go/commit/1821753daf0e46d09be2b87deb8bf83a6f4878fc))
* **api:** add max_subagents_per_environment field to agent policy ([f418eb8](https://github.com/gitpod-io/gitpod-sdk-go/commit/f418eb89a71943db1a404ba8cf1d01e055837ede))
* **api:** add min_size and max_size fields to prebuild warm pools ([dead742](https://github.com/gitpod-io/gitpod-sdk-go/commit/dead7420ef3f284d284df800c03ae330d018da06))
* **api:** add notification resource type ([00bd054](https://github.com/gitpod-io/gitpod-sdk-go/commit/00bd054040c7ca7ccd2318f9971545e82a731fff))
* **api:** add organization_tier field to identity response ([501a8eb](https://github.com/gitpod-io/gitpod-sdk-go/commit/501a8eb6bf3023fa92687168b55b34920447394d))
* **api:** add read_only field to PersonalAccessToken proto messages ([b9e5956](https://github.com/gitpod-io/gitpod-sdk-go/commit/b9e5956f058355e6336bb736a8c20bbe8a3b95c9))
* **api:** add resource type filtering to events watch ([32e133c](https://github.com/gitpod-io/gitpod-sdk-go/commit/32e133cd77876be3982bcaa33b2bf86aecb91637))
* **api:** add resource_ids to group filters, remove action from veto/policy types ([dca7e89](https://github.com/gitpod-io/gitpod-sdk-go/commit/dca7e8917089556865c9da48f43ee22e3354bdad))
* **api:** add search parameter to agent list prompts method ([d8da94a](https://github.com/gitpod-io/gitpod-sdk-go/commit/d8da94ac4e4b0ee253d24e288be6f29583581022))
* **api:** add session_id parameter to agent/environment methods ([9e1bb27](https://github.com/gitpod-io/gitpod-sdk-go/commit/9e1bb27c4ab0e24f54f27b888b37b02da240787f))
* **api:** add session_ids filter to agent/environment list methods ([e2a1abf](https://github.com/gitpod-io/gitpod-sdk-go/commit/e2a1abfc2b17924eb2f3c046b28ad2af0c9387c7))
* **api:** add SessionAdmin and SessionUser to ResourceRole enum ([247ec3a](https://github.com/gitpod-io/gitpod-sdk-go/commit/247ec3acd13e42723d05297ee6bfd026ee6f5ce1))
* **api:** add sessionId parameter to environment.NewFromProject method ([cff411c](https://github.com/gitpod-io/gitpod-sdk-go/commit/cff411ccd0ecd9ad400dbfff9038b832725221ab))
* **api:** add snapshotSizeBytes field to PrebuildStatus ([0501ee3](https://github.com/gitpod-io/gitpod-sdk-go/commit/0501ee30c65161292edc7b58fc2381a2ac25d4d1))
* **api:** add Sonnet 4.6 model variants to agent supported models ([2d13268](https://github.com/gitpod-io/gitpod-sdk-go/commit/2d1326867ed33b91226cdd03c17cc4966aa5a0ae))
* **api:** add sort option to ListAuditLogs ([98822e2](https://github.com/gitpod-io/gitpod-sdk-go/commit/98822e27f8b515666eed82d949732202d93e17fb))
* **api:** add sort parameter to projects list, refactor organization sort types ([08c907a](https://github.com/gitpod-io/gitpod-sdk-go/commit/08c907aeb8c94bcd8ea343a58e4665cffa0cb6be))
* **api:** add SUPPORTED_MODEL_HAIKU_4_5 to AgentExecutionStatusSupportedModel ([14ca472](https://github.com/gitpod-io/gitpod-sdk-go/commit/14ca472a83834020cb6c93cd6b3830be40ed3b59))
* **api:** add terminal field to RunsOn model ([b5134ce](https://github.com/gitpod-io/gitpod-sdk-go/commit/b5134ce63fa0424cd5e38fb023f34d932498027a))
* **api:** add time range filters to event list method ([7407251](https://github.com/gitpod-io/gitpod-sdk-go/commit/7407251a46ca25ede1377d8486fc02077ced4110))
* **api:** add update_window field and UpdateWindow type to runner ([829d303](https://github.com/gitpod-io/gitpod-sdk-go/commit/829d3034d52de1706055d5d55af2886af3bd4720))
* **api:** add WakeEvent parameter to agent send_to_execution method ([209477e](https://github.com/gitpod-io/gitpod-sdk-go/commit/209477e5a08eaf472a37c0bef47747a9f091ee4b))
* **api:** add warm pool methods to prebuilds ([baab482](https://github.com/gitpod-io/gitpod-sdk-go/commit/baab482ec21147cb070437c847b470d4dd712882))
* **backend/api,dashboard:** add server-side `recentlyCompleted` sort for `ListWorkflows` ([77e6c2e](https://github.com/gitpod-io/gitpod-sdk-go/commit/77e6c2ed11a0812f106a96d350cbb6c076bb1c8a))
* **internal:** support comma format in multipart form encoding ([038baa4](https://github.com/gitpod-io/gitpod-sdk-go/commit/038baa4aea542a7a0ebbe18ae8f2939f15a8d413))
* **types:** add ResourceRoleOrgAuditLogReader to ResourceRole enum ([c171119](https://github.com/gitpod-io/gitpod-sdk-go/commit/c171119f93f7e47a2243d7ecbf1bfd740674520d))
* **types:** add team_admin and team_viewer values to ResourceRole ([52c94f0](https://github.com/gitpod-io/gitpod-sdk-go/commit/52c94f027f4e60e96eedd22517832273028bc459))
* **types:** add warm pool capability to RunnerCapability enum ([ddd43e4](https://github.com/gitpod-io/gitpod-sdk-go/commit/ddd43e41e5c26f00b24274157ce6d7d9e6b3a416))


### Bug Fixes

* allow canceling a request while it is waiting to retry ([ec8099b](https://github.com/gitpod-io/gitpod-sdk-go/commit/ec8099b7f44da7197cbdaca0396d2c0b21f01eed))
* **api:** remove acknowledgeToken parameter from environment start method ([9f314d6](https://github.com/gitpod-io/gitpod-sdk-go/commit/9f314d65a530a01666021f2efa1041e1660b1123))
* **api:** rename ExecutableDenyList to VetoExecPolicy in organization policies ([59d7ce5](https://github.com/gitpod-io/gitpod-sdk-go/commit/59d7ce52b1626387770b7eab7a5fce73bf5f7d85))
* fix request delays for retrying to be more respectful of high requested delays ([86e683a](https://github.com/gitpod-io/gitpod-sdk-go/commit/86e683a1957f85a2a2bd7737b832f6b98f1d1ae4))
* **internal:** skip tests that depend on mock server ([0af24f8](https://github.com/gitpod-io/gitpod-sdk-go/commit/0af24f8c3d984471ca831d518773d88f794a2903))
* prevent duplicate ? in query params ([1e72155](https://github.com/gitpod-io/gitpod-sdk-go/commit/1e721558d9ad62ffd94abc3d123d2536b45a551e))
* **types:** generate shared enum types that are not referenced by other schemas ([1288811](https://github.com/gitpod-io/gitpod-sdk-go/commit/12888119896f45faf465d8c091af89e57910d19c))
* **types:** remove AgentExecution enum values from Principal and ResourceRole ([46a315a](https://github.com/gitpod-io/gitpod-sdk-go/commit/46a315abab9e7e0688a56efb71623634dd8cd4a4))


### Chores

* **ci:** add build step ([b3fc37a](https://github.com/gitpod-io/gitpod-sdk-go/commit/b3fc37a3b502d6aff16ea10a20aaad6f3c0fb41a))
* **ci:** skip lint on metadata-only changes ([f288dca](https://github.com/gitpod-io/gitpod-sdk-go/commit/f288dcae2df429304920d4edd0b7a0bdde8eece8))
* **ci:** skip uploading artifacts on stainless-internal branches ([b938324](https://github.com/gitpod-io/gitpod-sdk-go/commit/b9383240765de55c0a1f9963b9ffc59507e3e1d6))
* **ci:** support opting out of skipping builds on metadata-only commits ([e86d9b7](https://github.com/gitpod-io/gitpod-sdk-go/commit/e86d9b73a4c93ffcddeb60c2faae1ecba8f9809f))
* **docs:** add missing descriptions ([aff93a2](https://github.com/gitpod-io/gitpod-sdk-go/commit/aff93a262a316c089dbc4eb53bbf2f1295e67f50))
* **internal:** minor cleanup ([457e4e5](https://github.com/gitpod-io/gitpod-sdk-go/commit/457e4e54ec3d59111a073afc835bbe767a8f5cec))
* **internal:** move custom custom `json` tags to `api` ([4b6a3a9](https://github.com/gitpod-io/gitpod-sdk-go/commit/4b6a3a9aea35be9d96fdeced6be5f69c889d71c9))
* **internal:** regenerate SDK with no functional changes ([a878cc7](https://github.com/gitpod-io/gitpod-sdk-go/commit/a878cc7776983afdf1085933205f4599a79a04d4))
* **internal:** regenerate SDK with no functional changes ([4c85fa6](https://github.com/gitpod-io/gitpod-sdk-go/commit/4c85fa6bbbf090329053bce27f83a795be67f329))
* **internal:** regenerate SDK with no functional changes ([b6ed717](https://github.com/gitpod-io/gitpod-sdk-go/commit/b6ed71708bbf64a2f33ed98dbc91c4e5fb70621b))
* **internal:** regenerate SDK with no functional changes ([c67ca16](https://github.com/gitpod-io/gitpod-sdk-go/commit/c67ca160abb9e3485c9cb36a9b663234d5996cbb))
* **internal:** regenerate SDK with no functional changes ([8b0c202](https://github.com/gitpod-io/gitpod-sdk-go/commit/8b0c2026d434956eadbf6bebf17602ad2467c465))
* **internal:** remove mock server code ([48be3e9](https://github.com/gitpod-io/gitpod-sdk-go/commit/48be3e9a15291c7588193ed32cacfc7ffc98278c))
* **internal:** tweak CI branches ([2793802](https://github.com/gitpod-io/gitpod-sdk-go/commit/2793802d39acaf0e2acbf0eed0811a7af56d2999))
* **internal:** update gitignore ([cf1eb98](https://github.com/gitpod-io/gitpod-sdk-go/commit/cf1eb98d304dc41745060ec689dfdc579e545ec2))
* **internal:** use explicit returns ([2ac60b9](https://github.com/gitpod-io/gitpod-sdk-go/commit/2ac60b9befe3522599a86a586de76d8eaafd2d53))
* **internal:** use explicit returns in more places ([607cf23](https://github.com/gitpod-io/gitpod-sdk-go/commit/607cf23d9de758d28cabfdc9b1f92823adc0d6dc))
* remove unnecessary error check for url parsing ([ffe759b](https://github.com/gitpod-io/gitpod-sdk-go/commit/ffe759b8839dde00b2d396d93d80cd392916afb1))
* update mock server docs ([8e5bad9](https://github.com/gitpod-io/gitpod-sdk-go/commit/8e5bad965bd13abad2082bc22d1b32bdedcc3816))


### Documentation

* **api:** add validation constraints to timeout fields in environment and policy ([ba57b2a](https://github.com/gitpod-io/gitpod-sdk-go/commit/ba57b2a80c5f456bb5f5ab842412d8aa5acc8d7a))
* **api:** update annotations examples in agent start execution ([886895c](https://github.com/gitpod-io/gitpod-sdk-go/commit/886895c0c74cca2427955c7ee07e73bd4b1f092f))
* **api:** update MinSize validation constraints in prebuild warm pool ([ae8c301](https://github.com/gitpod-io/gitpod-sdk-go/commit/ae8c301910deff4b3a9efa4e170c79d5fab917a1))
* **api:** update resource filter descriptions in group role assignment list ([75c864c](https://github.com/gitpod-io/gitpod-sdk-go/commit/75c864c48bea726d89af5bf369c1f0e098b72967))

## 0.12.0 (2026-02-18)

Full Changelog: [v0.11.0...v0.12.0](https://github.com/gitpod-io/gitpod-sdk-go/compare/v0.11.0...v0.12.0)

### Features

* **api:** add audit_only to VetoExec and ExecutableDenyList ([3957348](https://github.com/gitpod-io/gitpod-sdk-go/commit/39573483e5b4013a283ce4075a6b71e79b1d1ce4))
* **api:** add runner_side_agent capability to RunnerCapability ([926e67e](https://github.com/gitpod-io/gitpod-sdk-go/commit/926e67e5bd00498dc91b7735414c796b5db3784a))
* **api:** add WarmpoolAdmin and WarmpoolViewer roles to ResourceRole ([e00b441](https://github.com/gitpod-io/gitpod-sdk-go/commit/e00b441d8ed0c1a0eb7d373415e373000e5b9439))


### Bug Fixes

* **types:** rename ExecutableDenyList to VetoExecPolicy in organization policies ([8353604](https://github.com/gitpod-io/gitpod-sdk-go/commit/8353604971c6a4d35ed29ddd046120e424fefc8b))

## 0.11.0 (2026-02-11)

Full Changelog: [v0.10.0...v0.11.0](https://github.com/gitpod-io/gitpod-sdk-go/compare/v0.10.0...v0.11.0)

### Features

* **api:** add bulk create/delete/update methods to projects ([f71fbc0](https://github.com/gitpod-io/gitpod-sdk-go/commit/f71fbc0090470ac5e6495af47fcb4af32891fa85))

## 0.10.0 (2026-02-11)

Full Changelog: [v0.9.0...v0.10.0](https://github.com/gitpod-io/gitpod-sdk-go/compare/v0.9.0...v0.10.0)

### Features

* [backend] Resource admin should be able to see all resource shares ([ca240e4](https://github.com/gitpod-io/gitpod-sdk-go/commit/ca240e4338b7af995b6dbc0d74f48903cb73423d))
* [SCIM] Configurable token expiration duration ([576e470](https://github.com/gitpod-io/gitpod-sdk-go/commit/576e4705dff1c62cacb2d726ad96c13c362cf51f))
* **api:** add additional_scopes to organization SSO configuration ([6e90868](https://github.com/gitpod-io/gitpod-sdk-go/commit/6e90868eab91485fa2236f1e9b393e4f02eb5d63))
* **api:** add annotations to agent execution metadata, filtering, and start ([c8dc7c3](https://github.com/gitpod-io/gitpod-sdk-go/commit/c8dc7c33805c7945022f8c65b7c8772909a63e2f))
* **api:** add announcement banner fields to organization proto ([a053bfa](https://github.com/gitpod-io/gitpod-sdk-go/commit/a053bfabb106a47fbe2080daba81b3e865015613))
* **api:** add derivedFromOrgRole field to RoleAssignment ([aca4f5d](https://github.com/gitpod-io/gitpod-sdk-go/commit/aca4f5dc7c5a19df88ad3d54fdababc341de6785))
* **api:** add dev runner parameters and fields to agent/runner ([85b1d51](https://github.com/gitpod-io/gitpod-sdk-go/commit/85b1d51656d561841510c1b101840fe351d777d0))
* **api:** add exclude_group_id filter to organization list members ([dca8e3d](https://github.com/gitpod-io/gitpod-sdk-go/commit/dca8e3d680f052df99c5684e06e459c3451a6564))
* **api:** add ExecutableDenyList to environment and organization policy ([a9780be](https://github.com/gitpod-io/gitpod-sdk-go/commit/a9780be351c6ecf9f1fa39fac34637448ce28e57))
* **api:** add filters to ListPrebuilds for inventory page ([5ffa984](https://github.com/gitpod-io/gitpod-sdk-go/commit/5ffa984912e737984d66525a92ad6987743f8c6e))
* **api:** add mcp_integrations field to agent execution status ([46ae3a5](https://github.com/gitpod-io/gitpod-sdk-go/commit/46ae3a560ea956f41cf8a104513206cd77ff234d))
* **api:** add opus 4.6 model variants to agent supported models ([4ba4075](https://github.com/gitpod-io/gitpod-sdk-go/commit/4ba4075bd7d52d4887e65bc4a04f00f9b2e2e838))
* **api:** add port access methods and organization admission level to environments ([0f7e97f](https://github.com/gitpod-io/gitpod-sdk-go/commit/0f7e97fd2b092c211805023224bdf103bb78554f))
* **api:** add ResourceRoleOrgAutomationsAdmin to ResourceRole enum ([d61b656](https://github.com/gitpod-io/gitpod-sdk-go/commit/d61b656054e841882b80c526d263932930f31a76))
* **api:** add ResourceRoleOrgGroupsAdmin to ResourceRole enum ([e17fccd](https://github.com/gitpod-io/gitpod-sdk-go/commit/e17fccd0d326d696c953f9a2010fb113399a2101))
* **api:** add restrict_account_creation_to_scim field to organization policy ([b6e23d0](https://github.com/gitpod-io/gitpod-sdk-go/commit/b6e23d05cf0c0ce7ca003be7510da42e7e04accd))
* **api:** add scope field and enum to environment secrets ([5c71212](https://github.com/gitpod-io/gitpod-sdk-go/commit/5c712120c5349317969792da6073639b63e5f290))
* **api:** add user_ids filter to organization list members method ([8483ca5](https://github.com/gitpod-io/gitpod-sdk-go/commit/8483ca578e298f2eff0042f4772b1ee5a0c0311e))
* **api:** add warm pools resource to prebuilds ([629edbe](https://github.com/gitpod-io/gitpod-sdk-go/commit/629edbe2a216446086e02e075854d1e63b5a3a88))
* **api:** implement GetAnnouncementBanner and UpdateAnnouncementBanner handlers ([025165d](https://github.com/gitpod-io/gitpod-sdk-go/commit/025165d03294ac90bf9672f131a31876ec48842a))
* **api:** remove GetChatIdentityToken method from accounts ([ee4a7aa](https://github.com/gitpod-io/gitpod-sdk-go/commit/ee4a7aa2789c34cb8fb63ce439eae2e951ded9f0))
* **chat:** add Pylon identity verification for chat widget ([cba45c3](https://github.com/gitpod-io/gitpod-sdk-go/commit/cba45c3d2fb196a48626481ffdb46577653f1e1b))
* **cli:** add --name flag to environment create command ([f406773](https://github.com/gitpod-io/gitpod-sdk-go/commit/f406773b1c7de04ff8915c20b700e73bd21ceb42))
* **db:** add service_account_tokens table ([17dd72a](https://github.com/gitpod-io/gitpod-sdk-go/commit/17dd72af458736138f41f8a80757daf4a68b56b0))
* Introduce projects admin org role ([860789d](https://github.com/gitpod-io/gitpod-sdk-go/commit/860789d1f286aa8a96da2d7858c2ae50d4af307e))
* **runner:** add capability check for ListSCMOrganizations ([465f090](https://github.com/gitpod-io/gitpod-sdk-go/commit/465f0900c63c943aba22d53d8ccae17d62221a90))
* **types:** add RoleAssignment value to ResourceType enum ([5a52b0e](https://github.com/gitpod-io/gitpod-sdk-go/commit/5a52b0eb845d19d7d674895632d93af5618d0a98))


### Documentation

* **api:** clarify port_sharing_disabled behavior in organization policy ([f2676c3](https://github.com/gitpod-io/gitpod-sdk-go/commit/f2676c35ec87e2e9ffbaa9de70f0da5cfdf1916a))
* **api:** update FilePath parameter description in secret ([5975445](https://github.com/gitpod-io/gitpod-sdk-go/commit/5975445f2fe59c311e1c43600ffc7b7e7e17ec86))

## 0.9.0 (2026-01-21)

Full Changelog: [v0.8.0...v0.9.0](https://github.com/gitpod-io/gitpod-sdk-go/compare/v0.8.0...v0.9.0)

### Features

* [api] Introduce RPCs to share resources with individual users ([d7ffceb](https://github.com/gitpod-io/gitpod-sdk-go/commit/d7ffceb2eaa72981a62c72f283d60744ca62892a))
* [api] sorting for `ListMembers` ([cabae1a](https://github.com/gitpod-io/gitpod-sdk-go/commit/cabae1a3485c18f2bc4906b3f93ae3f407ef0109))
* [backend] Adding direct_share field to groups ([143f723](https://github.com/gitpod-io/gitpod-sdk-go/commit/143f7236d939a962c38a40cf3504d09a8a02d791))
* [backend] Introduce org:runners_admin organization role ([7993dc4](https://github.com/gitpod-io/gitpod-sdk-go/commit/7993dc4480f7f676181f7837e41caa7f337febaa))
* [backend] Introduce role and member status filtering for `ListMembers` ([fda937f](https://github.com/gitpod-io/gitpod-sdk-go/commit/fda937f8cbdec42a802895bf6c23e0cb74317c10))
* **agent:** add spec mode for planning before interactive implementation ([47daf98](https://github.com/gitpod-io/gitpod-sdk-go/commit/47daf9841f0ba302ea5c442bd578cd6912b18ee9))
* API for SCIM configuration management ([a0f4eed](https://github.com/gitpod-io/gitpod-sdk-go/commit/a0f4eed6d124adafdd7f4fb5aa5685ebe409ff5e))
* **api:** add CheckRepositoryAccess API for repository access validation ([1879407](https://github.com/gitpod-io/gitpod-sdk-go/commit/1879407adceddbbce473b79bbfd8ed91bfe9f1c1))
* **api:** add draft and state fields to PullRequest proto ([b9a0d47](https://github.com/gitpod-io/gitpod-sdk-go/commit/b9a0d47bdd27ff86502afbba720b3d718cdd5f99))
* **api:** add inputs array to UserInputBlock proto ([a1e5a8a](https://github.com/gitpod-io/gitpod-sdk-go/commit/a1e5a8ab944d912f2bbc10f16ad676dba3376bf7))
* **api:** add ListSCMOrganizations endpoint ([6ee682f](https://github.com/gitpod-io/gitpod-sdk-go/commit/6ee682f9bd97ec6d9af9b7e9368cdd25e16c8adf))
* **api:** improve SearchRepositories pagination with next_page and total_count ([5616440](https://github.com/gitpod-io/gitpod-sdk-go/commit/56164403356b880c01ec2494cedb2f4862fabaf7))
* **automations:** add before_snapshot trigger type ([7e9a241](https://github.com/gitpod-io/gitpod-sdk-go/commit/7e9a24185ce20eeaf8221d0333a7c16ee37ca972))
* **dashboard:** show tier badge in org selector ([3c6cbfa](https://github.com/gitpod-io/gitpod-sdk-go/commit/3c6cbfae5c174df931ae0e365f6279f61a28ca21))
* Define SCIMConfiguration database schema ([92e9eaf](https://github.com/gitpod-io/gitpod-sdk-go/commit/92e9eafacdd3573ba3805625c21ea476d459aeb8))
* move agent mode from Spec to Status, add AgentModeChange signals ([32e7ca6](https://github.com/gitpod-io/gitpod-sdk-go/commit/32e7ca6a4a091dfe2d8e5e624ccfd78546e69465))
* **secrets:** add ServiceAccountSecret entity with full support ([ec680a1](https://github.com/gitpod-io/gitpod-sdk-go/commit/ec680a173f2724fd28eb9c58e16e936cadbd1eba))


### Bug Fixes

* **docs:** add missing pointer prefix to api.md return types ([82efe0a](https://github.com/gitpod-io/gitpod-sdk-go/commit/82efe0afe041ce2ffc4ee6452bc1f202c6ca169a))


### Chores

* **internal:** update `actions/checkout` version ([ad9a294](https://github.com/gitpod-io/gitpod-sdk-go/commit/ad9a2944946a952fa632be07ef188a32ae46d8ab))

## 0.8.0 (2026-01-09)

Full Changelog: [v0.7.0...v0.8.0](https://github.com/gitpod-io/gitpod-sdk-go/compare/v0.7.0...v0.8.0)

### Features

* **agent:** add group-based SCM tools access control ([02a3ac4](https://github.com/gitpod-io/gitpod-sdk-go/commit/02a3ac402e5b07ea0efffaf6a35ea3740b800211))
* **api:** add ImageInput to UserInputBlock proto ([c48d8ec](https://github.com/gitpod-io/gitpod-sdk-go/commit/c48d8ec1113b38bdc7cf11495001b6c29cb112a4))
* **api:** add recommended editors configuration to project settings ([673574f](https://github.com/gitpod-io/gitpod-sdk-go/commit/673574ffc184dc7b7f5e66599e173899d9dd4e4e))
* **db:** add webhooks table with trigger reference ([41d8762](https://github.com/gitpod-io/gitpod-sdk-go/commit/41d8762b09efce9967748d6c4e57308d8ee2dc7c))
* **prebuild:** expose snapshot completion percentage in prebuild status ([5e6837c](https://github.com/gitpod-io/gitpod-sdk-go/commit/5e6837cb70420eeba8a19f2b10f965e6c4d4478f))
* **skills:** add organization-level skills support ([dbd94be](https://github.com/gitpod-io/gitpod-sdk-go/commit/dbd94bebb9f959939d0bb2b882181316b771d0b9))


### Bug Fixes

* skip usage tests that don't work with Prism ([2e9de50](https://github.com/gitpod-io/gitpod-sdk-go/commit/2e9de50431b0d2ab64db8b7bbba3c83aedc7d74b))
* streaming endpoints should pass through errors correctly ([0897659](https://github.com/gitpod-io/gitpod-sdk-go/commit/08976594ce2bdb431fc4a662ec0fa21d6989a8c3))


### Chores

* **internal:** codegen related update ([7118d8d](https://github.com/gitpod-io/gitpod-sdk-go/commit/7118d8d561c9b7bacfa89e78e7fa1731c8c9f37b))
* pin GitHub Actions to SHA ([af1d69c](https://github.com/gitpod-io/gitpod-sdk-go/commit/af1d69c695071e89c7ede111953b510914b6dff0))

## 0.7.0 (2025-12-15)

Full Changelog: [v0.6.1...v0.7.0](https://github.com/gitpod-io/gitpod-sdk-go/compare/v0.6.1...v0.7.0)

### Features

* **api:** RBAC APIs ([dfee19e](https://github.com/gitpod-io/gitpod-sdk-go/commit/dfee19e63c7697876819a9bf516ec4466a3179ee))

## 0.6.1 (2025-12-06)

Full Changelog: [v0.6.0...v0.6.1](https://github.com/gitpod-io/gitpod-sdk-go/compare/v0.6.0...v0.6.1)

### Bug Fixes

* **mcp:** correct code tool API endpoint ([e94dac8](https://github.com/gitpod-io/gitpod-sdk-go/commit/e94dac8aeb3a3704d65044e7a9c7d7aa865c02a0))
* rename param to avoid collision ([230f5ad](https://github.com/gitpod-io/gitpod-sdk-go/commit/230f5adf57fa28e012fe769b611f3ad81ba283ba))


### Chores

* elide duplicate aliases ([515de72](https://github.com/gitpod-io/gitpod-sdk-go/commit/515de72284fc4276378d364665da670a92d6e756))
* **internal:** codegen related update ([7e15eec](https://github.com/gitpod-io/gitpod-sdk-go/commit/7e15eec605c407c7fe6db51625a3b2ab5da9447a))

## 0.6.0 (2025-09-25)

Full Changelog: [v0.5.0...v0.6.0](https://github.com/gitpod-io/gitpod-sdk-go/compare/v0.5.0...v0.6.0)

### Features

* **api:** gitpod -&gt; ona ([552d759](https://github.com/gitpod-io/gitpod-sdk-go/commit/552d759cc37edb71984b3a8d8dc740e999fa42a4))
* **client:** add debug log helper ([80b9e2b](https://github.com/gitpod-io/gitpod-sdk-go/commit/80b9e2b14e06850bda8d96cb4bdbcc611d9684d8))
* **client:** support optional json html escaping ([1b5fe8d](https://github.com/gitpod-io/gitpod-sdk-go/commit/1b5fe8d71198d7c27d346ced65b1a5916c1ca5f8))


### Bug Fixes

* bugfix for setting JSON keys with special characters ([8cb7fa5](https://github.com/gitpod-io/gitpod-sdk-go/commit/8cb7fa5e315ee8407b030c7203e95b66003d9473))
* **client:** process custom base url ahead of time ([aa5fbb8](https://github.com/gitpod-io/gitpod-sdk-go/commit/aa5fbb87c9b247fbe9ca6d74f72ccbd775d295b6))
* don't try to deserialize as json when ResponseBodyInto is []byte ([52bb01f](https://github.com/gitpod-io/gitpod-sdk-go/commit/52bb01f2bc2aae5736c292faf09d98a515be2a08))
* **pagination:** check if page data is empty in GetNextPage ([8d9d6fb](https://github.com/gitpod-io/gitpod-sdk-go/commit/8d9d6fb9ca5a3913d0400620dd18686290f8e57d))
* use slices.Concat instead of sometimes modifying r.Options ([083789f](https://github.com/gitpod-io/gitpod-sdk-go/commit/083789f37af9f36fe39d0e45cb9e0ed087cf54cf))


### Chores

* bump minimum go version to 1.22 ([92bbcb3](https://github.com/gitpod-io/gitpod-sdk-go/commit/92bbcb34c44340cfc9eb20416ac7f9c565462299))
* **ci:** enable for pull requests ([2c731e3](https://github.com/gitpod-io/gitpod-sdk-go/commit/2c731e32d19595c013749cea2a7a8c50603aa0a4))
* **ci:** only run for pushes and fork pull requests ([5214c86](https://github.com/gitpod-io/gitpod-sdk-go/commit/5214c8625e015da0f06ab2f817c96a31346475ec))
* do not install brew dependencies in ./scripts/bootstrap by default ([99dff1a](https://github.com/gitpod-io/gitpod-sdk-go/commit/99dff1a321df7284bf84be3b4639ddbbbba3f588))
* **internal:** fix lint script for tests ([076dd1d](https://github.com/gitpod-io/gitpod-sdk-go/commit/076dd1d2cef96adf0c5b21e4413387307f52971d))
* **internal:** update comment in script ([b55d442](https://github.com/gitpod-io/gitpod-sdk-go/commit/b55d44229d14d7bc61b3bf18ed181689b280dfc8))
* lint tests ([1bba107](https://github.com/gitpod-io/gitpod-sdk-go/commit/1bba1075657651b1ac7b261bbd9c2fcd3fc08fdd))
* lint tests in subpackages ([7c466c2](https://github.com/gitpod-io/gitpod-sdk-go/commit/7c466c272c807ecda01497bd3565408e973cc177))
* update @stainless-api/prism-cli to v5.15.0 ([b8acf38](https://github.com/gitpod-io/gitpod-sdk-go/commit/b8acf385df0d9b2ca8d2cf65635b6a2ab437e9c7))
* update more docs for 1.22 ([d1761ab](https://github.com/gitpod-io/gitpod-sdk-go/commit/d1761aba3bc28977b3227bcdf8161099feeee04d))

## 0.5.0 (2025-06-06)

Full Changelog: [v0.4.0...v0.5.0](https://github.com/gitpod-io/gitpod-sdk-go/compare/v0.4.0...v0.5.0)

### Features

* **api:** manual updates ([04616ff](https://github.com/gitpod-io/gitpod-sdk-go/commit/04616fff8d30f184d182f2f446c9d91ee82b896b))
* **api:** manual updates ([ccca7e5](https://github.com/gitpod-io/gitpod-sdk-go/commit/ccca7e5f582b1cb79fa8f18111ac2d59f3c8f42e))
* **api:** manual updates ([60b3e5e](https://github.com/gitpod-io/gitpod-sdk-go/commit/60b3e5e20df48bbb2e6d544ccbd266bf418eea7c))
* **api:** manual updates ([3ecd0c4](https://github.com/gitpod-io/gitpod-sdk-go/commit/3ecd0c4894b471a99b52da94d98ee4b1c12cde30))
* **api:** manual updates ([77f5109](https://github.com/gitpod-io/gitpod-sdk-go/commit/77f510928da90a197e8e556203677aa1a5838b9b))
* **client:** add support for reading base URL from environment variable ([e074cd4](https://github.com/gitpod-io/gitpod-sdk-go/commit/e074cd4d1977e7d9c861c9a08a5e8fd0b182d944))
* **client:** allow custom baseurls without trailing slash ([#55](https://github.com/gitpod-io/gitpod-sdk-go/issues/55)) ([20b4808](https://github.com/gitpod-io/gitpod-sdk-go/commit/20b480853025ee9b5402ae5a361010c7de9eabdc))
* **client:** improve default client options support ([#57](https://github.com/gitpod-io/gitpod-sdk-go/issues/57)) ([0cb0d29](https://github.com/gitpod-io/gitpod-sdk-go/commit/0cb0d293a64defd394978cd2b4da6cb006753a69))
* **client:** support custom http clients ([#65](https://github.com/gitpod-io/gitpod-sdk-go/issues/65)) ([ba2780f](https://github.com/gitpod-io/gitpod-sdk-go/commit/ba2780f5173f3e569527a24e9151e258e90df509))


### Bug Fixes

* **client:** return error on bad custom url instead of panic ([#64](https://github.com/gitpod-io/gitpod-sdk-go/issues/64)) ([656a0d1](https://github.com/gitpod-io/gitpod-sdk-go/commit/656a0d137cfa42afa6d540d80d306d254f4841aa))
* **client:** unmarshal stream events into fresh memory ([#63](https://github.com/gitpod-io/gitpod-sdk-go/issues/63)) ([9cf0811](https://github.com/gitpod-io/gitpod-sdk-go/commit/9cf08112fe7b5e702311882b79f167956b09753f))
* handle empty bodies in WithJSONSet ([6155e91](https://github.com/gitpod-io/gitpod-sdk-go/commit/6155e915016da00b885ed17f2452e0b48e1e4fd6))
* **pagination:** handle errors when applying options ([8853449](https://github.com/gitpod-io/gitpod-sdk-go/commit/8853449c7377b7d2976647e4f84598513394b4ab))
* **test:** return early after test failure ([#61](https://github.com/gitpod-io/gitpod-sdk-go/issues/61)) ([0295d21](https://github.com/gitpod-io/gitpod-sdk-go/commit/0295d214d2934e8d4789883255885d3e5ea474c9))


### Chores

* add request options to client tests ([#60](https://github.com/gitpod-io/gitpod-sdk-go/issues/60)) ([b575336](https://github.com/gitpod-io/gitpod-sdk-go/commit/b5753364f5bc6c5617205b2145ade577bc43ac72))
* **ci:** add timeout thresholds for CI jobs ([41d7ceb](https://github.com/gitpod-io/gitpod-sdk-go/commit/41d7cebf060b9a6d6f432d2caa6a82032ca3408c))
* **ci:** only use depot for staging repos ([9fbbeb1](https://github.com/gitpod-io/gitpod-sdk-go/commit/9fbbeb1fd88a3dc49943ef661c1d1c39da956fa2))
* **docs:** document pre-request options ([f2e2e2e](https://github.com/gitpod-io/gitpod-sdk-go/commit/f2e2e2e6286c2641387a0fc11b7ada23d8afc9fc))
* **docs:** improve security documentation ([#59](https://github.com/gitpod-io/gitpod-sdk-go/issues/59)) ([2d3d9e7](https://github.com/gitpod-io/gitpod-sdk-go/commit/2d3d9e7bdb95db8579085db1a23fa7f558b3ab2b))
* fix typos ([#62](https://github.com/gitpod-io/gitpod-sdk-go/issues/62)) ([2ab745f](https://github.com/gitpod-io/gitpod-sdk-go/commit/2ab745f9136f7c6f4528f05e26c1c22a4b15b959))
* **internal:** codegen related update ([a14dae0](https://github.com/gitpod-io/gitpod-sdk-go/commit/a14dae09ffdfec9ef7e1cb3c4b0c3ead27b14e14))
* **internal:** codegen related update ([#56](https://github.com/gitpod-io/gitpod-sdk-go/issues/56)) ([dc521f6](https://github.com/gitpod-io/gitpod-sdk-go/commit/dc521f608d7475d2d98f64fe118433450ab5b2af))
* **internal:** expand CI branch coverage ([7d00669](https://github.com/gitpod-io/gitpod-sdk-go/commit/7d00669f67ba5d7007feceb6f25863de34a237df))
* **internal:** reduce CI branch coverage ([3b674ac](https://github.com/gitpod-io/gitpod-sdk-go/commit/3b674ac888efc38b3713547841770b6c77af2ac8))
* **internal:** remove extra empty newlines ([#58](https://github.com/gitpod-io/gitpod-sdk-go/issues/58)) ([80ff63b](https://github.com/gitpod-io/gitpod-sdk-go/commit/80ff63b57f6df81750f4acaee91003902d2ddae6))
* **tests:** improve enum examples ([#66](https://github.com/gitpod-io/gitpod-sdk-go/issues/66)) ([359070f](https://github.com/gitpod-io/gitpod-sdk-go/commit/359070f05052c648ae75a806ddd037c8eb28dbe7))


### Documentation

* update documentation links to be more uniform ([df86410](https://github.com/gitpod-io/gitpod-sdk-go/commit/df86410ba6bf256e7dc7ebd8b41fc5e15d762fa0))
* update URLs from stainlessapi.com to stainless.com ([#53](https://github.com/gitpod-io/gitpod-sdk-go/issues/53)) ([a5f0af7](https://github.com/gitpod-io/gitpod-sdk-go/commit/a5f0af754b7fb4ef5d967aa53ce3ad22648826be))

## 0.4.0 (2025-02-21)

Full Changelog: [v0.3.2...v0.4.0](https://github.com/gitpod-io/gitpod-sdk-go/compare/v0.3.2...v0.4.0)

### Features

* **api:** manual updates ([#48](https://github.com/gitpod-io/gitpod-sdk-go/issues/48)) ([78b2504](https://github.com/gitpod-io/gitpod-sdk-go/commit/78b2504189e98dcfaef1427444378752ccede6cc))
* **api:** manual updates ([#50](https://github.com/gitpod-io/gitpod-sdk-go/issues/50)) ([19015d3](https://github.com/gitpod-io/gitpod-sdk-go/commit/19015d3d486ef7569342f1f08bcadebc9e5f89e3))

## 0.3.2 (2025-02-18)

Full Changelog: [v0.3.1...v0.3.2](https://github.com/gitpod-io/gitpod-sdk-go/compare/v0.3.1...v0.3.2)

### Features

* feat(shared): add standard error codes for API errors ([#45](https://github.com/gitpod-io/gitpod-sdk-go/issues/45)) ([bab81e0](https://github.com/gitpod-io/gitpod-sdk-go/commit/bab81e0dc68c2499dc57faf6ae25f7162da06427))

## 0.3.1 (2025-02-18)

Full Changelog: [v0.3.0...v0.3.1](https://github.com/gitpod-io/gitpod-sdk-go/compare/v0.3.0...v0.3.1)

### Features

* **api:** manual updates ([#42](https://github.com/gitpod-io/gitpod-sdk-go/issues/42)) ([faa83bf](https://github.com/gitpod-io/gitpod-sdk-go/commit/faa83bff9c5d16b8a514fef030b79b9b4de37da8))

## 0.3.0 (2025-02-18)

Full Changelog: [v0.2.0...v0.3.0](https://github.com/gitpod-io/gitpod-sdk-go/compare/v0.2.0...v0.3.0)

### Features

* **api:** manual updates ([#38](https://github.com/gitpod-io/gitpod-sdk-go/issues/38)) ([853ae5c](https://github.com/gitpod-io/gitpod-sdk-go/commit/853ae5cf56e0d0fef0d2371166bd167b33e4ffa5))
* **api:** manual updates ([#40](https://github.com/gitpod-io/gitpod-sdk-go/issues/40)) ([a71a350](https://github.com/gitpod-io/gitpod-sdk-go/commit/a71a350d93dc72bb856668893583195f2d22b5c5))

## 0.2.0 (2025-02-18)

Full Changelog: [v0.1.1...v0.2.0](https://github.com/gitpod-io/gitpod-sdk-go/compare/v0.1.1...v0.2.0)

### Features

* **api:** manual updates ([#33](https://github.com/gitpod-io/gitpod-sdk-go/issues/33)) ([8b18e3a](https://github.com/gitpod-io/gitpod-sdk-go/commit/8b18e3af0ace255d515d1f53e9d8573a57ba8617))
* **api:** manual updates ([#36](https://github.com/gitpod-io/gitpod-sdk-go/issues/36)) ([9612bdb](https://github.com/gitpod-io/gitpod-sdk-go/commit/9612bdbb4a43b3fdc23e72b446b0ac0b295fb326))
* **api:** Organizations Open API docs ([#31](https://github.com/gitpod-io/gitpod-sdk-go/issues/31)) ([8c99874](https://github.com/gitpod-io/gitpod-sdk-go/commit/8c99874c9c17e63becdbd0fb30a45d77c51bdf2f))


### Bug Fixes

* **client:** don't truncate manually specified filenames ([#34](https://github.com/gitpod-io/gitpod-sdk-go/issues/34)) ([bb87acd](https://github.com/gitpod-io/gitpod-sdk-go/commit/bb87acddf4cf002b37c63e8a74bb2f044deac221))

## 0.1.1 (2025-02-14)

Full Changelog: [v0.1.0-alpha.3...v0.1.1](https://github.com/gitpod-io/gitpod-sdk-go/compare/v0.1.0-alpha.3...v0.1.1)

### Features

* **api:** manual updates ([#28](https://github.com/gitpod-io/gitpod-sdk-go/issues/28)) ([7fd8f6c](https://github.com/gitpod-io/gitpod-sdk-go/commit/7fd8f6c7ef83ba84bbca6e5e9cf052ccf0c6cc4c))

## 0.1.0-alpha.3 (2025-02-13)

Full Changelog: [v0.1.0-alpha.2...v0.1.0-alpha.3](https://github.com/gitpod-io/gitpod-sdk-go/compare/v0.1.0-alpha.2...v0.1.0-alpha.3)

### Features

* **api:** manual updates ([#17](https://github.com/gitpod-io/gitpod-sdk-go/issues/17)) ([762e8f7](https://github.com/gitpod-io/gitpod-sdk-go/commit/762e8f70ad9507a80a2cddefe31e79d128919c06))
* **api:** manual updates ([#18](https://github.com/gitpod-io/gitpod-sdk-go/issues/18)) ([3d70885](https://github.com/gitpod-io/gitpod-sdk-go/commit/3d708857be80530a1b8cc5169b5b03dbc38bebd3))
* **api:** manual updates ([#20](https://github.com/gitpod-io/gitpod-sdk-go/issues/20)) ([47bc0b1](https://github.com/gitpod-io/gitpod-sdk-go/commit/47bc0b103590ceace9c56bec72a8c6da133f5f17))
* **api:** manual updates ([#22](https://github.com/gitpod-io/gitpod-sdk-go/issues/22)) ([fd57922](https://github.com/gitpod-io/gitpod-sdk-go/commit/fd57922f815b9c876056ea9a878ac5c352af1ca9))
* **api:** manual updates ([#23](https://github.com/gitpod-io/gitpod-sdk-go/issues/23)) ([33e7aed](https://github.com/gitpod-io/gitpod-sdk-go/commit/33e7aed13c50f6458e8f0c4a7fba7fcb9e775f24))
* **api:** manual updates ([#25](https://github.com/gitpod-io/gitpod-sdk-go/issues/25)) ([d8140c6](https://github.com/gitpod-io/gitpod-sdk-go/commit/d8140c673b66d53bfe58ff315ad4e9062b7cc5ac))
* **api:** Overview page updates ([#26](https://github.com/gitpod-io/gitpod-sdk-go/issues/26)) ([a85f0c2](https://github.com/gitpod-io/gitpod-sdk-go/commit/a85f0c2a5d99d9616bb3013fb98b88b23967e1df))
* **api:** update examples ([#21](https://github.com/gitpod-io/gitpod-sdk-go/issues/21)) ([48bf713](https://github.com/gitpod-io/gitpod-sdk-go/commit/48bf7138ab06ae118aaefc46c9ff2a9809a25d77))
* **api:** update with latest API spec ([#24](https://github.com/gitpod-io/gitpod-sdk-go/issues/24)) ([7232ba6](https://github.com/gitpod-io/gitpod-sdk-go/commit/7232ba6189894f53d16d94218c7b0e2d361c4ea7))

## 0.1.0-alpha.2 (2025-02-11)

Full Changelog: [v0.1.0-alpha.1...v0.1.0-alpha.2](https://github.com/gitpod-io/gitpod-sdk-go/compare/v0.1.0-alpha.1...v0.1.0-alpha.2)

### Bug Fixes

* do not call path.Base on ContentType ([#13](https://github.com/gitpod-io/gitpod-sdk-go/issues/13)) ([8901c2e](https://github.com/gitpod-io/gitpod-sdk-go/commit/8901c2e2c7b1e01a388713b7c87a8bae642b4826))


### Chores

* go live ([#15](https://github.com/gitpod-io/gitpod-sdk-go/issues/15)) ([090fe43](https://github.com/gitpod-io/gitpod-sdk-go/commit/090fe43b7c871c25fcb2c4fd33eb21843cc8e646))

## 0.1.0-alpha.1 (2025-02-07)

Full Changelog: [v0.0.1-alpha.0...v0.1.0-alpha.1](https://github.com/gitpod-io/flex-sdk-go/compare/v0.0.1-alpha.0...v0.1.0-alpha.1)

### Features

* **api:** add events streaming ([7b06b6d](https://github.com/gitpod-io/flex-sdk-go/commit/7b06b6d80655159603946e90bdd2a7f290aae1a9))
* **api:** dedupe paginations ([844ce0c](https://github.com/gitpod-io/flex-sdk-go/commit/844ce0c0a6346e2bc09c1ca0774218853a42ed7a))
* **api:** fix pagination field names ([5d9b63a](https://github.com/gitpod-io/flex-sdk-go/commit/5d9b63a349e96e0d7d6c40907aa6dd5a16e5f1b9))
* **api:** manual updates ([0ce1c80](https://github.com/gitpod-io/flex-sdk-go/commit/0ce1c80121b3c57f64e090f6e8f6d38511f0e6fc))
* **api:** manual updates ([7f97b68](https://github.com/gitpod-io/flex-sdk-go/commit/7f97b68a42fb25dbd4814de7dae16a856df491b9))
* **api:** pagination config ([2fba4d4](https://github.com/gitpod-io/flex-sdk-go/commit/2fba4d49efb2157186289a03bd937f391867470e))
* **api:** properly produce empty request bodies ([#4](https://github.com/gitpod-io/flex-sdk-go/issues/4)) ([99f62a8](https://github.com/gitpod-io/flex-sdk-go/commit/99f62a8c92c6591b8bea5c17e0360f62a1f9254e))
* **api:** strip enum prefix ([#11](https://github.com/gitpod-io/flex-sdk-go/issues/11)) ([14372b8](https://github.com/gitpod-io/flex-sdk-go/commit/14372b8a15aeac0128174c55be895a936c6620b4))
* **api:** try to fix updateenvironmentrequest ([#8](https://github.com/gitpod-io/flex-sdk-go/issues/8)) ([6a8ac0f](https://github.com/gitpod-io/flex-sdk-go/commit/6a8ac0f0c043c86f0d39fa1a38075f7770123d36))
* **api:** update to latest changes ([#10](https://github.com/gitpod-io/flex-sdk-go/issues/10)) ([9ea005e](https://github.com/gitpod-io/flex-sdk-go/commit/9ea005e0c22967d6f1d22859b39b070e70a6db54))
* **api:** update via SDK Studio ([95e28e4](https://github.com/gitpod-io/flex-sdk-go/commit/95e28e4046adbdac2368d83afb7293183c4feb05))
* **api:** update via SDK Studio ([c6ac608](https://github.com/gitpod-io/flex-sdk-go/commit/c6ac60843e1d9ac02d5d0641352eb942f7d6b3a9))
* **api:** update via SDK Studio ([23d2c90](https://github.com/gitpod-io/flex-sdk-go/commit/23d2c90c4890948732515b629d5a3e85c032cd3a))
* **api:** update via SDK Studio ([01b0d82](https://github.com/gitpod-io/flex-sdk-go/commit/01b0d82e4f08d5488610192804219ce17575efec))
* **api:** update via SDK Studio ([6ef2201](https://github.com/gitpod-io/flex-sdk-go/commit/6ef220173096a1ff89c4ee833e9cfd8562e4a126))
* **api:** update via SDK Studio ([f8874b6](https://github.com/gitpod-io/flex-sdk-go/commit/f8874b60e449349dac91dcce5776df76e34402f5))
* **api:** update via SDK Studio ([1e08279](https://github.com/gitpod-io/flex-sdk-go/commit/1e0827902b909fa3897368d6a7108c1064d588ad))
* **api:** update via SDK Studio ([00bf9ad](https://github.com/gitpod-io/flex-sdk-go/commit/00bf9ad9b89dc7eaa39f6a61215c7047498339da))
* **api:** update via SDK Studio ([3f0d5cd](https://github.com/gitpod-io/flex-sdk-go/commit/3f0d5cd57af9f17286f2e8a0b0104b4254cb6204))
* **api:** update via SDK Studio ([1a2a1f6](https://github.com/gitpod-io/flex-sdk-go/commit/1a2a1f66c8491d894fd5ebe0d2b8652448937ae6))
* **api:** update via SDK Studio ([e085c74](https://github.com/gitpod-io/flex-sdk-go/commit/e085c7455cb169e7433100b070f570331cdc29bf))
* **api:** update via SDK Studio ([4dbfaf1](https://github.com/gitpod-io/flex-sdk-go/commit/4dbfaf1c5982af8e68599a6cb42fcd7625dd4027))
* **api:** update via SDK Studio ([fae2a51](https://github.com/gitpod-io/flex-sdk-go/commit/fae2a514edf4816e2cc927e236e8febe5392e47c))
* **api:** update via SDK Studio ([5b939b4](https://github.com/gitpod-io/flex-sdk-go/commit/5b939b47cb7e63ea96192c1dba7b68ccb66dc330))
* **api:** update via SDK Studio ([0475403](https://github.com/gitpod-io/flex-sdk-go/commit/04754035486a77cd93411f113a69f69cf208f128))
* **api:** update via SDK Studio ([2e48633](https://github.com/gitpod-io/flex-sdk-go/commit/2e4863370ac6829882755e5d2c7a91c818136a42))
* **api:** update via SDK Studio ([daada19](https://github.com/gitpod-io/flex-sdk-go/commit/daada19224feef8c8198d46982c5cad59cb135a4))
* **api:** update via SDK Studio ([d41d24a](https://github.com/gitpod-io/flex-sdk-go/commit/d41d24ad6b1f5081cd12faf48be9e25b90bdb71c))
* **api:** update via SDK Studio ([0c29b09](https://github.com/gitpod-io/flex-sdk-go/commit/0c29b09e2219a5dec7b468d387fa1479adb2682f))
* **api:** update via SDK Studio ([94304f9](https://github.com/gitpod-io/flex-sdk-go/commit/94304f9e90884f325779b6592f9b58b9e5dd6584))
* **api:** update via SDK Studio ([7df303b](https://github.com/gitpod-io/flex-sdk-go/commit/7df303b41eb2bee5cdeab7155f0d3291ce6c6869))
* **api:** update via SDK Studio ([f38f5bf](https://github.com/gitpod-io/flex-sdk-go/commit/f38f5bfe96b6c0d760f35c24fe31faf20010ea71))
* **api:** update via SDK Studio ([d91ff3e](https://github.com/gitpod-io/flex-sdk-go/commit/d91ff3e8b5d2003abc1032fbe1b7ca4460507fc0))
* **api:** update via SDK Studio ([587e248](https://github.com/gitpod-io/flex-sdk-go/commit/587e248f96871423b4385a75a98d11d5bbafaa81))
* **api:** update via SDK Studio ([121a64f](https://github.com/gitpod-io/flex-sdk-go/commit/121a64f139fc0fc3e29827fac48856ac6744d85b))
* **client:** send `X-Stainless-Timeout` header ([2c7e88f](https://github.com/gitpod-io/flex-sdk-go/commit/2c7e88f0c39b3958401849e2d6c0e1dc753f6877))
* pagination responses ([8f4569a](https://github.com/gitpod-io/flex-sdk-go/commit/8f4569a96b08e814164884475ea69864ecc4e610))


### Bug Fixes

* fix apijson.Port for embedded structs ([efeae7c](https://github.com/gitpod-io/flex-sdk-go/commit/efeae7c812d04cc957c3305a7447b58306119c7c))
* fix apijson.Port for embedded structs ([60e9a40](https://github.com/gitpod-io/flex-sdk-go/commit/60e9a40244a5f77f899e693c646b59f030f00557))
* fix early cancel when RequestTimeout is provided for streaming requests ([#9](https://github.com/gitpod-io/flex-sdk-go/issues/9)) ([85c3877](https://github.com/gitpod-io/flex-sdk-go/commit/85c3877513d216815d84fd69e8ad77de7497d1a0))
* fix unicode encoding for json ([da625f4](https://github.com/gitpod-io/flex-sdk-go/commit/da625f44e263be3dd03385efae9ff3392b346cc4))
* pagination example ([c2514e2](https://github.com/gitpod-io/flex-sdk-go/commit/c2514e25269e4662709dfb997b3b66077569c8fe))
* pagination response ([2a82e77](https://github.com/gitpod-io/flex-sdk-go/commit/2a82e775066fd42201e8481a8ad46e40f77212ac))
* **tests:** disable mock tests ([#5](https://github.com/gitpod-io/flex-sdk-go/issues/5)) ([d05b1fd](https://github.com/gitpod-io/flex-sdk-go/commit/d05b1fd4a3df4b467f9b170006d28c64e4ac9c75))
* **tests:** disable test mocks ([#3](https://github.com/gitpod-io/flex-sdk-go/issues/3)) ([ba10691](https://github.com/gitpod-io/flex-sdk-go/commit/ba10691f2b65fcbdc7407065a13493c6b6dbe4e2))


### Chores

* add UnionUnmarshaler for responses that are interfaces ([#6](https://github.com/gitpod-io/flex-sdk-go/issues/6)) ([d6fff80](https://github.com/gitpod-io/flex-sdk-go/commit/d6fff8045ed918f66fbb7889a654fb332c46a77a))
* bump license year ([44e7174](https://github.com/gitpod-io/flex-sdk-go/commit/44e717413ceb1d122ae94aff35077a710fc0f020))
* configure new SDK language ([a58f282](https://github.com/gitpod-io/flex-sdk-go/commit/a58f2822fc2d2f91517bd78b4c73e4d8bfef6adc))
* go live ([#1](https://github.com/gitpod-io/flex-sdk-go/issues/1)) ([966a7b3](https://github.com/gitpod-io/flex-sdk-go/commit/966a7b3d20e4ff7c98df72443859991a843aa8c1))
* **internal:** codegen related update ([e2267e2](https://github.com/gitpod-io/flex-sdk-go/commit/e2267e2a1e4508201c11d3141e6a4f9a54bc791a))
* **internal:** codegen related update ([b1ae6b0](https://github.com/gitpod-io/flex-sdk-go/commit/b1ae6b0b7dfb7ec14e3cb599728b081576790777))
* **internal:** codegen related update ([0c1b414](https://github.com/gitpod-io/flex-sdk-go/commit/0c1b4143dc7788efec18e3d35f4bbc65114aac5e))
* **internal:** codegen related update ([7aaf5e0](https://github.com/gitpod-io/flex-sdk-go/commit/7aaf5e0f9f94e64a35f0af3fdd20f34262f3e37c))
* **internal:** update examples ([4778169](https://github.com/gitpod-io/flex-sdk-go/commit/4778169765fb1b912e66de33ddd1db2cd85e4018))
* rebuild project due to codegen change ([bc60ea7](https://github.com/gitpod-io/flex-sdk-go/commit/bc60ea747051741bde334b8caa4c6c2e5f66c734))
* rebuild project due to codegen change ([e75888d](https://github.com/gitpod-io/flex-sdk-go/commit/e75888d3ffbbf647c42bf34c62853ec50088d2f7))
* rebuild project due to codegen change ([e6aaa0b](https://github.com/gitpod-io/flex-sdk-go/commit/e6aaa0bf1f9ecbd7ead59a8fe881ad31398117b2))


### Documentation

* document raw responses ([bbb2a23](https://github.com/gitpod-io/flex-sdk-go/commit/bbb2a2347f1ad8bbcd219ea10390bb0d1eb70087))
