# Gitpod MCP Server

A Model Control Protocol (MCP) server that provides access to Gitpod resources and operations through Claude.

## Features

- List Gitpod projects
- List Gitpod environments
- Create new environments
- Stop environments
- Execute commands in environments

## Claude Desktop Configuration

1. Create a file containing your Gitpod personal access token:
```bash
echo "your_api_key_here" > /tmp/gitpod-personal-access-token.txt
```

2. Build the server:
```bash
go build -o /tmp/gitpod-mcp
```

3. Add the following to your Claude Desktop configuration to enable Gitpod integration:

```json
{
  "mcpServers": {
    "gitpod": {
      "command": "/tmp/gitpod-mcp",
      "args": ["/tmp/gitpod-personal-access-token.txt"]
    }
  }
}
```

>**Note:** Don't forget to delete the token file after you're done.

## Available Resources

- `gitpod://projects` - List all available Gitpod projects
- `gitpod://environments` - List current Gitpod environments

## Available Tools

### create_environment
Creates a new Gitpod environment for a project.
- Parameters:
  - `project_id` (string, required): The ID of the project to create the environment in

### stop_environment
Stops a running Gitpod environment.
- Parameters:
  - `environment_id` (string, required): The ID of the environment to stop

### execute_command
Executes a command in a Gitpod environment.
- Parameters:
  - `environment_id` (string, required): The ID of the environment to execute the command in
  - `command` (string, required): The command to execute (runs as a bash script in project root)
  - `description` (string, required): A short description of the command (max 200 characters)
```
