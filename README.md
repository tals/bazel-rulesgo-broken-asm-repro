Demonstration project for bazelbuild/rules_go that fails picking up on the .S files

Run

```bash
bazel run --sandbox_debug //:hello-go
```

# Workaround
Push the c++ code to a cc_library and add it as a cdep. See [workaround branch](https://github.com/tals/bazel-rulesgo-broken-asm-repro/tree/workaround)