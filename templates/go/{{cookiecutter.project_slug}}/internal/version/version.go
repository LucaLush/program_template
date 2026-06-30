package version

var (
	// Version is the version of the application (injected at build time)
	Version = "dev"
	// Commit is the git commit hash (injected at build time)
	Commit = "none"
	// BuildTime is the time when the binary was built (injected at build time)
	BuildTime = "unknown"
)
