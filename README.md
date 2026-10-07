# GCC — IRIX fork

This repository is the IRIX GCC fork. It is based on
[`rust-lang/gcc`](https://github.com/rust-lang/gcc) — the GCC 17 development
trunk carrying the libgccjit changes that `rustc_codegen_gcc` requires —
rather than on upstream GCC (ADR-0019).

## The IRIX port

The IRIX port lives on `master` as a series of commits, not a branch:

- `IRIX target restoration: config, multilibs and startfiles`
- `IRIX MIPS target macros and headers`
- `IRIX MIPS codegen and bootstrap adjustments`
- `libgcc/libgomp IRIX runtime adjustments`
- `IRIX configure changes for the gcc directory`
- `libstdc++ IRIX OS layer`
- `Guard GCC's stdint types against IRIX <inttypes.h>`
- `Build fixes for native IRIX hosts and legacy IRIX headers`

These restore the IRIX 6.5 target — `mips-sgi-irix6.5`, ELF32 big-endian
MIPS III/IV, with o32, n32 and n64 ABIs.

## Why rust-lang/gcc

The fork carries the patched libgccjit that `rustc_codegen_gcc` uses as the
Rust codegen backend (ADR-0018), so one codebase serves both the stock rebuild
of IRIX 6.5.7m and the Rust backend for IRIX 7.

## Where the series lives

The patch series is the source of truth. The published series and its
documentation live in [`irix7/project`](https://github.com/irix7/project):

- `patches/` — the GCC and binutils series
- `docs/toolchain.md` — the series and the build

## Build

Do not build from this repository directly. Use the project build script,
which pins the sources, applies the series and verifies the result:

```sh
nix develop --command bash -c \
  'scripts/build-toolchain.sh --languages c'
```

See https://github.com/irix7/project/blob/main/docs/toolchain.md.

## Upstream

- Fork of: https://github.com/rust-lang/gcc
- Upstream GCC: https://gcc.gnu.org/git/gcc.git
- Licence: GPL-3.0 with the GCC Runtime Library Exception
  (`COPYING3`, `COPYING.RUNTIME`)
