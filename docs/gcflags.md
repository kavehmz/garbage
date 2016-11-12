```text
go build -gcflags --help
usage: compile [options] file.go...
  -%	debug non-static initializers
  -+	compiling runtime
  -A	for bootstrapping, allow 'any' type
  -B	disable bounds checking
  -D path
    	set relative path for local imports
  -E	debug symbol export
  -I directory
    	add directory to import search path
  -K	debug missing line numbers
  -M	debug move generation
  -N	disable optimizations
  -P	debug peephole optimizer
  -R	debug register optimizer
  -S	print assembly listing
  -V	print compiler version
  -W	debug parse tree after type checking
  -asmhdr file
    	write assembly header to file
  -buildid id
    	record id as the build id in the export metadata
  -complete
    	compiling complete package (no C or assembly)
  -cpuprofile file
    	write cpu profile to file
  -d list
    	print debug information about items in list
  -dynlink
    	support references to Go symbols defined in other shared libraries
  -e	no limit on number of errors reported
  -f	debug stack frames
  -g	debug code generation
  -h	halt on error
  -i	debug line number stack
  -importmap definition
    	add definition of the form source=actual to import map
  -installsuffix suffix
    	set pkg directory suffix
  -j	debug runtime-initialized variables
  -l	disable inlining
  -largemodel
    	generate code that assumes a large memory model
  -linkobj file
    	write linker-specific object to file
  -live
    	debug liveness analysis
  -m	print optimization decisions
  -memprofile file
    	write memory profile to file
  -memprofilerate rate
    	set runtime.MemProfileRate to rate
  -msan
    	build code compatible with C/C++ memory sanitizer
  -newexport
    	use new export format (default true)
  -nolocalimports
    	reject local (relative) imports
  -o file
    	write output to file
  -p path
    	set expected package import path
  -pack
    	write package file instead of object file
  -r	debug generated wrappers
  -race
    	enable race detector
  -s	warn about composite literals that can be simplified
  -shared
    	generate code that can be linked into a shared library
  -ssa
    	use SSA backend to generate code (default true)
  -trimpath prefix
    	remove prefix from recorded source file paths
  -u	reject unsafe code
  -v	increase debug verbosity
  -w	debug type checking
  -wb
    	enable write barrier (default true)
  -x	debug lexer
```