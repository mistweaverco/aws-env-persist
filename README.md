# aws-env-persist

Persists environment 💾  AWS Credentials 🔐 across different terminals 💻

## Usage

Add this to your `.bashrc`:

```sh
# Source AWS Environment Credentials
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

