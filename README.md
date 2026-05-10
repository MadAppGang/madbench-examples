# madbench-examples

Downloadable fixtures for the [madbench](https://github.com/MadAppGang/madbench) tutorials. Each tutorial has its own top-level directory; the `code/` subdirectory inside holds the fixture files that the tutorial expects under your project's `code/` folder.

## Layout

```
02-first-live-llm-eval/        # for docs/tutorials/02-first-live-llm-eval.md
└── code/
    ├── go.mod
    ├── auth/login.go
    ├── routes/router.go
    ├── cmd/server/main.go
    └── internal/session/store.go

03-codegen-with-exec/           # for docs/tutorials/03-codegen-with-exec.md
└── code/
    ├── go.mod
    └── fizzbuzz_test.go

04-semantic-judging/            # for docs/tutorials/04-semantic-judging-with-rubrics.md
└── code/
    ├── go.mod
    ├── add.go
    └── add_test.go
```

Top-level directory names mirror the tutorial filename slugs in `docs/tutorials/` so the mapping is trivially traceable. The inner `code/` folder lands directly inside your project root after extraction.

## Quick download

From inside your tutorial project directory, pull the codenav fixture (Tutorial 02):

```bash
curl -sL https://github.com/MadAppGang/madbench-examples/tarball/main \
  | tar -xz --strip-components=2 \
      'MadAppGang-madbench-examples-*/02-first-live-llm-eval/code'
```

Result on disk:

```
your-project/
├── madbench.yaml         # you create this in Step 2
└── code/                 # extracted by the curl command
    ├── go.mod
    ├── auth/login.go
    └── ...
```

For other tutorials, replace `02-first-live-llm-eval` with `03-codegen-with-exec` or `04-semantic-judging`. See each tutorial's "Step 1 — Build the fixture" for the exact command.

## Why a separate repo?

`madbench` itself is private; this repo holds only the tutorial fixture trees so they're publicly downloadable without exposing the harness source. Fixture content is mirrored from `examples/<tutorial>/fixtures/` in the main repo.

## License

Same as `madbench`. Fixtures are sample code only — not for production use.
