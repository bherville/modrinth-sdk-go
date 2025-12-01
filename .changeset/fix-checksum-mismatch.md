---
"modrinth-sdk-go": patch
---

Fix Go module checksum mismatch by releasing a new version.

The v0.2.0 tag was created before the release commit was pushed, causing the Go module proxy to cache an incorrect checksum. This patch release ensures downstream consumers can download the module without checksum errors.
