# antrea_task
This repo is the test task for Antrea project - Automate Dependency updates with Renovate

### What the Go Program Does

This is a simple Go program designed to parse a language tag provided as a command-line argument. The program takes a string from the command line, attempts to interpret it as a language tag using an external library, and prints the result or an error.



### The Vulnerability: (Denial of Service)

The vulnerability resides in the **older versions** of the `golang.org/x/text` module, specifically versions **prior to** `v0.3.7` (the fixed version). Our program explicitly uses an older, vulnerable version like `v0.3.5` for demonstration.

I used `govulncheck` to confirm that my module indeed uses a vulnerable dependency.

<img width="655" alt="Screenshot 2025-05-26 at 3 16 43 PM" src="https://github.com/user-attachments/assets/d21a7aa0-db23-4b46-8e98-f01c2cfd7134" />



## Solution using Renovate

Renovate has access to our repository, it will scan our `go.mod` file. It detects that our `golang.org/x/text@v0.3.5` dependency has a known vulnerability (`GO-2021-0113`) and a fixed version is available (`v0.3.7` or later), it will automatically open a pull request to update that dependency to a safe version.



## References

[https://pkg.go.dev/vuln/GO-2021-0113](https://pkg.go.dev/vuln/GO-2021-0113)

[https://docs.renovatebot.com/configuration-options/#vulnerabilityalerts](https://docs.renovatebot.com/configuration-options/#vulnerabilityalerts)

