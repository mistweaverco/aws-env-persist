# aws-env-persist

Persists environment 💾  AWS Credentials 🔐 across different terminals 💻

## Usage

Add this to your `.bashrc`:

```sh
source <(aws-env-persist get-env)
```

Persist current environment:

```sh
aws-env-persist save
```

