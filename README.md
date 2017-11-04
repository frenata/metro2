# Metro 2 File Parser

* to parse either Header Records or Base Segments, call `parseFixed` with the text of the file.
* to run the tests: `go test` 
* Header and Base segments are printed by default in JSON format
* to verify Header data integrity, `.metro()` can be called to convert back to the Metro2 format.
