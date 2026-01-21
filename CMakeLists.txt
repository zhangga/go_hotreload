cmake_minimum_required(VERSION 3.20.0)

project(go_hotreload NONE)

# --- Config ---
find_program(GO_EXECUTABLE go)
if(NOT GO_EXECUTABLE)
  message(FATAL_ERROR "go not found in PATH")
endif()

# Current directory / output directory (similar to Makefile variables)
set(CUR_DIR "${CMAKE_CURRENT_LIST_DIR}")
set(OUT_DIR "${CUR_DIR}/bin" CACHE PATH "Output directory")

# Branch name (best-effort, like Makefile CUR_BRANCH)
execute_process(
  COMMAND git branch --show-current
  WORKING_DIRECTORY "${CUR_DIR}"
  OUTPUT_VARIABLE CUR_BRANCH
  OUTPUT_STRIP_TRAILING_WHITESPACE
  ERROR_QUIET
)
if(CUR_BRANCH STREQUAL "")
  execute_process(
    COMMAND git rev-parse --abbrev-ref HEAD
    WORKING_DIRECTORY "${CUR_DIR}"
    OUTPUT_VARIABLE CUR_BRANCH
    OUTPUT_STRIP_TRAILING_WHITESPACE
    ERROR_QUIET
  )
endif()
if(CUR_BRANCH STREQUAL "")
  set(CUR_BRANCH "unknown")
endif()

set(EXAMPLE_DIR "${CUR_DIR}/example")

# --- Targets ---

# target: example  (go run)
add_custom_target(example
  COMMAND "${GO_EXECUTABLE}" run ./main.go
  WORKING_DIRECTORY "${EXAMPLE_DIR}"
  USES_TERMINAL
  COMMENT "Running example (branch=${CUR_BRANCH})"
)

# target: patch  (go build plugin)
add_custom_target(patch
  COMMAND "${GO_EXECUTABLE}" build -buildmode=plugin -o=patch_v1.so ./patch_v1/main.go
  WORKING_DIRECTORY "${EXAMPLE_DIR}"
  USES_TERMINAL
  COMMENT "Building patch plugin: example/patch_v1.so (branch=${CUR_BRANCH})"
)

set_property(DIRECTORY PROPERTY VS_STARTUP_PROJECT help)
