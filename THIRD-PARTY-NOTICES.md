# THIRD-PARTY SOFTWARE NOTICES

bunchotext incorporates the following third-party open source libraries, all of which are compatible with the MIT License under which bunchotext is distributed.

---
## Command-Line Framework

- github.com/spf13/cobra v1.10.2
- License: Apache License 2.0
- LICENSE File: [Link](https://github.com/spf13/cobra/blob/main/LICENSE.txt)

> Cobra is a library for creating powerful modern CLI applications. It provides a simple interface to create commands, flags, and help text.

---
## Gitignore Parsing

- github.com/sabhiram/go-gitignore v0.0.0-20210923224102-525f6e181f06
- License: MIT
- LICENSE File: [Link](https://github.com/sabhiram/go-gitignore/blob/master/LICENSE)

> A Go library for parsing .gitignore files and matching file paths against ignore patterns.

---
# License Compliance

All dependencies are used in unmodified form via Go modules. Their respective license texts are preserved in the Go module cache and remain accessible to users who build bunchotext from source.

Full license texts for each dependency can be viewed at the URLs above or by examining the LICENSE file within each module's repository.

To view licenses locally after building:
```bash
go list -m -f '{{.Dir}}' github.com/spf13/cobra
go list -m -f '{{.Dir}}' github.com/sabhiram/go-gitignore
```

# bunchotext itself is licensed under MIT:

```
MIT License

Copyright (c) 2026 Tsupko Nikita "quadroc0rp" Romanovich

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```