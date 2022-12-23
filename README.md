# aws-env-persist

Persists environment 💾  AWS Credentials 🔐 across different terminals 💻

## Installation

Download the latest binary from the [release page](/releases)
Save it to `/usr/local/bin/aws-env-persist`.
Make it executable `chmod +x /usr/local/bin/aws-env-persist`

## Auto Completion

If you want to have auto-completion, put this into your `.bashrc`

```sh
# Enable auto-completion
complete -C /usr/local/bin/aws-env-persist aws-env-persist
```

## Usage

Add this to your `.bashrc`, if you want to source the aws credentials automatically:

```sh
# Source AWS environment credentials automatically
eval "$(aws-env-persist get-env)"
```

or source it on demand in your current shell via:

```sh
source <(aws-env-persist get-env)
```

Persist current environment:

```sh
aws-env-persist save
```

