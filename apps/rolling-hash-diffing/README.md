# Rolling Hash Diffing

## Requirements

- Hashing function gets the data as a parameter. Separate possible filesystem operations.
- Chunk size can be fixed or dynamic, but must be split to at least two chunks on any sufficiently sized data.
- Should be able to recognize changes between chunks. Only the exact differing locations should be added to the delta.
- Well-written unit tests function well in describing the operation, no UI necessary.

## Checklist

- Input/output operations are separated from the calculations
- detects chunk changes and/or additions
- detects chunk removals
- detects additions between chunks with shifted original chunks

## Run

```
go run .
```

## Test

```
go test ./...
```