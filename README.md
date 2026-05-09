# madbench-examples

Downloadable fixtures for the [madbench](https://github.com/MadAppGang/madbench) tutorials. Mirrors the fixture trees that the tutorials build by hand, so you can skip the typing and curl them in one command.

## Layout

```
codenav/sample-go-repo/      # Tutorial 02 — first live LLM eval (codenav)
codegen-fizzbuzz/repo/       # Tutorial 03 — codegen with exec
bugfix-add/repo/             # Tutorial 04 — semantic judging with rubrics
```

Each subdirectory is a complete fixture — drop it under your tutorial project's `fixtures/` directory and you're ready to run.

## Quick download

For example, to pull the codenav fixture into your tutorial project:

```bash
mkdir -p fixtures
curl -sL https://github.com/MadAppGang/madbench-examples/tarball/main \
  | tar -xz --strip-components=2 -C fixtures \
      'MadAppGang-madbench-examples-*/codenav/sample-go-repo'
```

The strip count and path glob differ per fixture — see each tutorial's "Step 1 — Build the fixture" section for the exact command.

## Why a separate repo?

`madbench` itself is private; this repo holds only the tutorial fixture trees so they're publicly downloadable without exposing the harness source. Fixture content here is generated from `examples/<tutorial>/fixtures/` in the main repo and is intended to stay in sync.

## License

Same as `madbench`. Fixtures are sample code only — not for production use.
