# kudo-registry
MVP for KUDO Registry

## Usage

### Prerequisites 

Download for instance the current test repository from https://kudo-test-repo.storage.googleapis.com to `/tmp/kudo-test-repo`

### Start the Registry locally

Start the KUDO Registry with a local Storage Backend:

```.env
go run cmd/kudoregistry/main.go server start --backend=local --root-dir=/tmp/kudo-test-repo
```

* Flag `--backend` defines the StorageBackend, here `local`.
* Flag `--root-dir` defines the directory from which when local it is serving

### Install via the Kudoctl CLI

Make sure you are in the `cli-improvement` branch of the https://github.com/fabianbaier/kudo fork.

Then run: `go run cmd/kubectl-kudo/main.go install kafka --all-dependencies --repo-url=http://127.0.0.1:50051/api/v1/repo`

Or if you want to directly access the Google Cloud Storage:

* run the CLI via: `go run cmd/kubectl-kudo/main.go install kafka --all-dependencies --repo-url=https://kudo-test-repo.storage.googleapis.com`

Or by default this should link to it as well:
* run the CLI via: `go run cmd/kubectl-kudo/main.go install kafka --all-dependencies`

Just be advised that the CLI will automatically if not otherwise specified create `~/.kudo` and caches the data, resulting in the desireable outcome of having the repo cached after the first time of downloading and/or manually creating.
