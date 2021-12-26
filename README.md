# go-generics-tests-bug

## Steps to reproduce

Assuming you have go1.18beta1 installed, you can reproduce the bug by running:

```sh
git clone https://github.com/akshaybharambe14/go-generics-tests-bug
cd go-generics-tests-bug
go1.18beta1 test
```

This produces following output:

```text
# github.com/akshaybharambe14/go-generics-tests-bug [github.com/akshaybharambe14/go-generics-tests-bug.test]
panic: runtime error: invalid memory address or nil pointer dereference
[signal 0xc0000005 code=0x0 addr=0x18 pc=0x845ae5]

goroutine 1 [running]:
cmd/compile/internal/noder.transformCompLit(0xc000384400)
        c:/go/src/cmd/compile/internal/noder/transform.go:1019 +0x8a5
cmd/compile/internal/noder.(*irgen).compLit(0xc000142240, {0xad5150, 0xc000541680}, 0xc0003ebcc0)
        c:/go/src/cmd/compile/internal/noder/expr.go:372 +0x276
cmd/compile/internal/noder.(*irgen).expr0(0xc000142240, {0xad5150, 0xc000541680}, {0xad60a8?, 0xc0003ebcc0?})
        c:/go/src/cmd/compile/internal/noder/expr.go:107 +0x585
cmd/compile/internal/noder.(*irgen).expr(0xc000142240, {0xad60a8?, 0xc0003ebcc0?})
        c:/go/src/cmd/compile/internal/noder/expr.go:81 +0x5e8
cmd/compile/internal/noder.(*irgen).compLit(0xc000142240, {0xad51f0, 0xc00058bad0}, 0xc0003ebc70)
        c:/go/src/cmd/compile/internal/noder/expr.go:353 +0x57b
cmd/compile/internal/noder.(*irgen).expr0(0xc000142240, {0xad51f0, 0xc00058bad0}, {0xad60a8?, 0xc0003ebc70?})
        c:/go/src/cmd/compile/internal/noder/expr.go:107 +0x585
cmd/compile/internal/noder.(*irgen).expr(0xc000142240, {0xad60a8?, 0xc0003ebc70?})
        c:/go/src/cmd/compile/internal/noder/expr.go:81 +0x5e8
cmd/compile/internal/noder.(*irgen).compLit(0xc000142240, {0xad51c8, 0xc000415790}, 0xc0003ebc20)
        c:/go/src/cmd/compile/internal/noder/expr.go:360 +0x747
cmd/compile/internal/noder.(*irgen).expr0(0xc000142240, {0xad51c8, 0xc000415790}, {0xad60a8?, 0xc0003ebc20?})
        c:/go/src/cmd/compile/internal/noder/expr.go:107 +0x585
cmd/compile/internal/noder.(*irgen).expr(0xc000142240, {0xad60a8?, 0xc0003ebc20?})
        c:/go/src/cmd/compile/internal/noder/expr.go:81 +0x5e8
cmd/compile/internal/noder.(*irgen).exprs(0xc00007f360?, {0xc00007f350, 0x1, 0x200c00004cc20?})
        c:/go/src/cmd/compile/internal/noder/expr.go:329 +0x8e
cmd/compile/internal/noder.(*irgen).exprList(0x8?, {0xad60a8?, 0xc0003ebc20?})
        c:/go/src/cmd/compile/internal/noder/expr.go:312 +0x85
cmd/compile/internal/noder.(*irgen).stmt(0xc000142240, {0xad5f28?, 0xc000057ac0?})
        c:/go/src/cmd/compile/internal/noder/stmt.go:79 +0x418
cmd/compile/internal/noder.(*irgen).stmts(0xc00036d570?, {0xc000057b00?, 0x4, 0xc00007f648?})
        c:/go/src/cmd/compile/internal/noder/stmt.go:19 +0xaf
cmd/compile/internal/noder.(*irgen).funcBody(0xc000142240, 0xc00036e6e0, 0x78f320?, 0xc0000577c0, 0xc000057840)
        c:/go/src/cmd/compile/internal/noder/func.go:45 +0x25f
cmd/compile/internal/noder.(*irgen).funcDecl.func1()
        c:/go/src/cmd/compile/internal/noder/decl.go:145 +0xf5
cmd/compile/internal/noder.(*irgen).generate(0xc000142240, {0xc0003d4100, 0x4, 0xc000432000?})
        c:/go/src/cmd/compile/internal/noder/irgen.go:280 +0x227
cmd/compile/internal/noder.check2({0xc0003d4100, 0x4, 0x4})
        c:/go/src/cmd/compile/internal/noder/irgen.go:92 +0x16d
cmd/compile/internal/noder.LoadPackage({0xc00011e120, 0x4, 0x0?})
        c:/go/src/cmd/compile/internal/noder/noder.go:90 +0x335
cmd/compile/internal/gc.Main(0x98f488)
        c:/go/src/cmd/compile/internal/gc/main.go:191 +0xb13
main.main()
        c:/go/src/cmd/compile/main.go:55 +0xdd
FAIL    github.com/akshaybharambe14/go-generics-tests-bug [build failed]
```

Tests run perfectly fine if executed for individual unit:

```sh
$ go1.18beta1 test -v find.go find_test.go
=== RUN   TestFind
=== RUN   TestFind/element_exists_in_slice
=== RUN   TestFind/element_does_not_existsin_slice
=== RUN   TestFind/empty_slice
--- PASS: TestFind (0.00s)
    --- PASS: TestFind/element_exists_in_slice (0.00s)
    --- PASS: TestFind/element_does_not_existsin_slice (0.00s)
    --- PASS: TestFind/empty_slice (0.00s)
PASS
ok      command-line-arguments  0.245s
```

```sh
$ go1.18beta1 test -v delete.go delete_test.go
=== RUN   TestDelete
=== RUN   TestDelete/delete_from_empty_slice
=== RUN   TestDelete/delete_out_of_bound_index
=== RUN   TestDelete/delete_negative_index
=== RUN   TestDelete/delete_valid_index
=== RUN   TestDelete/delete_last_index
=== RUN   TestDelete/delete_first_index
--- PASS: TestDelete (0.00s)
    --- PASS: TestDelete/delete_from_empty_slice (0.00s)
    --- PASS: TestDelete/delete_out_of_bound_index (0.00s)
    --- PASS: TestDelete/delete_negative_index (0.00s)
    --- PASS: TestDelete/delete_valid_index (0.00s)
    --- PASS: TestDelete/delete_last_index (0.00s)
    --- PASS: TestDelete/delete_first_index (0.00s)
PASS
ok      command-line-arguments  0.182s
```

This looks like an issue in compiling test files.

I observed this issue while working on [github.com/akshaybharambe14/gouf](https://github.com/akshaybharambe14/gouf). It is confirmed that this issue is not OS specific as tests failed to run with github actions. You can see that output [here](https://github.com/akshaybharambe14/gouf/runs/4634778561?check_suite_focus=true)
