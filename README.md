# madbench-examples

Downloadable fixtures for the [madbench](https://github.com/MadAppGang/madbench) tutorials. Mirrors the fixture trees that the tutorials build by hand, so you can skip the typing and curl them in one command.

## Layout

```
tutorial/
├── 02-first-live-llm-eval/     # for docs/tutorials/02-first-live-llm-eval.md
│   ├── go.mod
│   ├── auth/login.go
│   ├── routes/router.go
│   ├── cmd/server/main.go
│   └── internal/session/store.go
├── 03-codegen-with-exec/       # for docs/tutorials/03-codegen-with-exec.md
│   ├── go.mod
│   └── fizzbuzz_test.go
└── 04-semantic-judging/        # for docs/tutorials/04-semantic-judging-with-rubrics.md
    ├── go.mod
    ├── add.go
    └── add_test.go
```

Folder names mirror the tutorial filename slugs in `docs/tutorials/` so the mapping is trivially traceable in both directions.

## Quick download

To pull the codenav fixture (Tutorial 02) into your project:

```bash
mkdir -p fixtures
curl -sL https://github.com/MadAppGang/madbench-examples/tarball/main \
  | tar -xz --strip-components=2 -C fixtures \
      'MadAppGang-madbench-examples-*/tutorial/02-first-live-llm-eval'
```

The path argument at the end picks which tutorial to extract — replace `02-first-live-llm-eval` with `03-codegen-with-exec` or `04-semantic-judging` for the other tutorials. See each tutorial's "Step 1 — Build the fixture" section for the exact command.

## Why a separate repo?

`madbench` itself is private; this repo holds only the tutorial fixture trees so they're publicly downloadable without exposing the harness source. Fixture content is generated from `examples/<tutorial>/fixtures/` in the main repo.

## License

Same as `madbench`. Fixtures are sample code only — not for production use.
