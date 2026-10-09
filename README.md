# IRIX GCC

An IRIX 6.5 port of GCC, forked from [`rust-lang/gcc`](https://github.com/rust-lang/gcc) —
the GCC 17 development trunk that also carries the patched `libgccjit` that
`rustc_codegen_gcc` needs. One codebase serves both the IRIX 6.5 rebuild and the
Rust backend for the future IRIX 7 operating system.

Target: **`mips-sgi-irix6.5`** — big-endian MIPS III/IV, ELF32, with the **o32**,
**n32** and **n64** ABIs.

## What this fork provides

- **Cross compiler** — runs on a modern host (x86-64 Linux) and emits IRIX
  binaries. This is the toolchain the IRIX 6.5 rebuild uses. Driven by the
  project's `scripts/build-toolchain.sh --gcc fork`.
- **Native IRIX GCC** *(in progress)* — a GCC that itself runs on IRIX,
  produced by a Canadian cross
  (`--build=x86_64-linux --host=mips-sgi-irix6.5 --target=mips-sgi-irix6.5`) and
  installed on the guest. See
  [issue #15](https://github.com/irix7/project/issues/15).
- **A patched `libgccjit`** — the Rust codegen backend, built from the same tree
  with `--enable-languages=…,jit --enable-host-shared`.

## What's included

| Language | Frontend | Target runtime | Status |
| --- | --- | --- | --- |
| C | `cc1` | `libgcc` | works (o32/n32/n64) |
| C++ | `cc1plus` | `libstdc++` | works (IRIX OS layer ported) |
| Fortran | `f951` | `libgfortran` | works |
| JIT | `libgccjit` | n/a (host) | works — the Rust backend |
| Go | `go1` | `libgo` | **in progress** — [issue #153](https://github.com/irix7/project/issues/153) |
| Objective-C / Obj-C++ | `cc1obj`/`cc1objplus` | `libobjc` | not enabled — [issue #154](https://github.com/irix7/project/issues/154) |
| Modula-2 | `cc1gm2` | `libgm2` | not enabled — [issue #154](https://github.com/irix7/project/issues/154) |
| D | `d21` | `libphobos` | not enabled — [issue #154](https://github.com/irix7/project/issues/154) |
| Ada | `gnat1` | `libgnat` | not enabled — [issue #154](https://github.com/irix7/project/issues/154) |
| COBOL | `cob1` | `libgcobol` | not available (x86-64/AArch64 only) — [issue #154](https://github.com/irix7/project/issues/154) |
| Rust | `gccrs` (experimental) | — | not used — the project uses `rustc` + `rustc_codegen_gcc` |

Also built: `libgcc`, `libstdc++`, `libgfortran`, `libquadmath`, `libgccjit`,
startfiles and the IRIX crt objects.

## Build

Do **not** build in this tree directly. The project drives the build, pins the
sources, applies the series and verifies the result:

```sh
git clone https://github.com/irix7/project
cd project
nix develop --command bash -c \
  'scripts/build-toolchain.sh --gcc fork --languages c,c++,fortran,jit \
     --sysroot /path/to/irix-6.5.7m-sysroot'
scripts/verify-toolchain.sh \
  --prefix .scratch/toolchain-17.0.0/prefix --gcc-version 17.0.0
```

You need GNU Nix and a captured IRIX 6.5.7m sysroot (see the project's
`docs/oracle.md`). The cross is paired with the project's IRIX-patched binutils
series. No SGI material is distributed here or by the build.

## Status — what works, and what does not yet

**Works:** C, C++, Fortran and libgccjit; o32 and n32 (and n64 build/link);
dynamic and static links against the captured sysroot; `verify-toolchain.sh`
passes.

**Not yet:**

- **Go / gccgo** — the frontend builds; the `libgo` IRIX runtime is being
  restored. [issue #153](https://github.com/irix7/project/issues/153)
- **Objective-C, Modula-2, D, Ada, COBOL** — not enabled for IRIX.
  [issue #154](https://github.com/irix7/project/issues/154)
- **n64** — builds and links, but the runtime is not exercised (it cannot run on
  the project's 32-bit guest).
  [issue #155](https://github.com/irix7/project/issues/155)
- **`libatomic`** — disabled for the sysroot build, interim.
  [issue #150](https://github.com/irix7/project/issues/150)
- **Native IRIX GCC** — Canadian cross not built yet.
  [issue #15](https://github.com/irix7/project/issues/15)
- **Rust backend** — `rustc_codegen_gcc` not yet exercised against this fork's
  libgccjit. [issue #156](https://github.com/irix7/project/issues/156)

## Contributing

Pull requests are welcome — IRIX 6.5 build fixes, new frontend/runtime ports
(Go, Objective-C, Modula-2), and GCC 17 forward-port fixes especially.

- Keep the IRIX port as commits on `master` (the series is the source of truth),
  upstream style, one concern per commit, each referencing its tracking issue in
  [irix7/project](https://github.com/irix7/project/issues).
- Match upstream conventions; run the relevant GCC tests where practical.
- The project publishes no SGI-derived material: don't paste IRIX source, headers
  or binaries into patches or the README.

## The IRIX port

The port is a series of commits on `master`, not a branch:

- IRIX target restoration: `config`, multilibs and startfiles
- IRIX MIPS target macros and headers
- IRIX MIPS codegen and bootstrap adjustments; `libgcc`/`libgomp` runtime
- IRIX `configure` changes; `libstdc++` IRIX OS layer
- GCC 17 forward-port fixes: the LRA reload guard
  ([#149](https://github.com/irix7/project/issues/149)), `IRIX_USING_GNU_LD` via
  `ld_flavor` ([#150](https://github.com/irix7/project/issues/150)), n64 sysroot
  paths ([#151](https://github.com/irix7/project/issues/151)), the libstdc++
  namespace/`ctype` files ([#152](https://github.com/irix7/project/issues/152)),
  and the `libgo` IRIX OS layer ([#153](https://github.com/irix7/project/issues/153))

## Upstream, licence and related

- Fork of: <https://github.com/rust-lang/gcc>; upstream GCC:
  <https://gcc.gnu.org/git/gcc.git>
- Licence: GPL-3.0 with the GCC Runtime Library Exception (`COPYING3`,
  `COPYING.RUNTIME`)
- Project hub — patch series, build scripts, docs, decisions and issues:
  <https://github.com/irix7/project>
