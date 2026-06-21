package main

import (
	"errors"
	"fmt"
)

/*
	19. Error Chains — Sequential Pipeline with Explicit Propagation:

	In TS with async/await + try/catch, errors propagate automatically:
	  try {
	    const config = await loadConfig()      // throws => jumps to catch
	    const db     = await connectDB(config) // throws => jumps to catch
	    const user   = await fetchUser(db)     // throws => jumps to catch
	  } catch (err) {
	    if (err instanceof ConfigError) { ... }
	  }

	In Go, there is NO automatic propagation. Every function call that can fail
	must be explicitly checked, and the error must be explicitly forwarded.
	This IS verbose. That verbosity is intentional — every failure path is visible
	at the call site, not hidden until a catch block somewhere far above.

	The pipeline below is the direct Go equivalent of that TS pattern.
	Feel the repetition of `if err != nil { return ..., err }` — that is Go.
	Do not try to compress it into a helper that mimics try/catch.
	The repetition is the documentation.
*/

// Sentinel errors for this pipeline — same pattern as 14-errors/sentinel-errors.go
var ErrConfigMissing = errors.New("config file not found")
var ErrDBConnection = errors.New("database connection refused")
var ErrUserNotFound = errors.New("user not found")

// Step 1: loadConfig simulates reading app configuration
// In a real app: reads from a file, env vars, or a config service
func loadConfig(env string) (string, error) {
	configs := map[string]string{
		"production":  "prod-config-v2",
		"development": "dev-config-v1",
	}

	cfg, ok := configs[env]
	if !ok {
		// wrap with context: what were we doing + which sentinel is in the chain
		return "", fmt.Errorf("loadConfig (env=%s): %w", env, ErrConfigMissing)
	}

	return cfg, nil
}

// Step 2: connectDB uses the config to connect to the database
// In a real app: establishes a *sql.DB connection pool
func connectDB(config string) (string, error) {
	// simulating: prod config works, anything else fails
	if config != "prod-config-v2" {
		return "", fmt.Errorf("connectDB (config=%s): %w", config, ErrDBConnection)
	}

	return "db-connection-pool", nil
}

// Step 3: fetchUser queries a user from the database connection
// In a real app: runs SELECT query, scans into struct
func fetchUser(db string, userID int) (string, error) {
	if db == "" {
		// guard against bad input — explicit check, not a panic
		return "", fmt.Errorf("fetchUser: received empty db connection")
	}

	users := map[int]string{1: "akshat", 2: "alan"}
	user, ok := users[userID]
	if !ok {
		return "", fmt.Errorf("fetchUser (userID=%d): %w", userID, ErrUserNotFound)
	}

	return user, nil
}

// startupPipeline chains all three steps — this is the Go equivalent of
// the TS try { await loadConfig(); await connectDB(); await fetchUser() } catch {}
// Notice: every error is explicitly forwarded with if err != nil { return ..., err }
// There is no implicit propagation. Every level adds its own context via %w.
func startupPipeline(env string, userID int) (string, error) {

	// Step 1
	config, err := loadConfig(env)
	if err != nil {
		// forward upward — wrapping adds this layer's context to the chain
		return "", fmt.Errorf("startupPipeline: %w", err)
	}

	// Step 2
	db, err := connectDB(config)
	if err != nil {
		return "", fmt.Errorf("startupPipeline: %w", err)
	}

	// Step 3
	user, err := fetchUser(db, userID)
	if err != nil {
		return "", fmt.Errorf("startupPipeline: %w", err)
	}

	return user, nil
}

func ErrorChainsDemo() {
	println("\n\n01. Error Chains — Sequential Pipeline:")

	// Case 1: valid env, valid user — full happy path
	println("\nCase 1: production env, userID=1 (should succeed):")
	user, err := startupPipeline("production", 1)
	if err != nil {
		fmt.Println("Pipeline failed:", err)
	} else {
		fmt.Println("Pipeline succeeded, user:", user)
	}

	// Case 2: invalid env — fails at Step 1, never reaches Steps 2 or 3
	println("\nCase 2: unknown env (fails at loadConfig):")
	_, err = startupPipeline("staging", 1)
	if err != nil {
		fmt.Println("Pipeline failed:", err)
		// errors.Is() unwraps through all the fmt.Errorf %w layers to find the sentinel
		if errors.Is(err, ErrConfigMissing) {
			fmt.Println("Root cause confirmed via errors.Is(): ErrConfigMissing")
		}
	}

	// Case 3: valid env, user not found — fails at Step 3
	println("\nCase 3: production env, userID=99 (fails at fetchUser):")
	_, err = startupPipeline("production", 99)
	if err != nil {
		fmt.Println("Pipeline failed:", err)
		if errors.Is(err, ErrUserNotFound) {
			fmt.Println("Root cause confirmed via errors.Is(): ErrUserNotFound")
		}
	}
}

/*
	The Broken %v Version:
	This demonstrates what happens when you forget to use %w (wrapping) and use %v instead.
	%v formats the error as a string but DISCARDS the original error value.
	errors.Is() and errors.As() work by unwrapping the chain — if %v was used,
	the chain is broken silently. No compile error. Just wrong runtime behaviour.

	This is the specific failure mode your TS custom-error instanceof habit won't warn you about.
	In TS, throw propagates the actual object through the chain.
	In Go, fmt.Errorf with %v creates a brand new string error with no chain link.
*/

func loadConfigBroken(env string) (string, error) {
	configs := map[string]string{"production": "prod-config-v2"}
	cfg, ok := configs[env]
	if !ok {
		// WRONG: %v converts the sentinel to a string and discards it
		// errors.Is(err, ErrConfigMissing) will return false after this
		return "", fmt.Errorf("loadConfig (env=%s): %v", env, ErrConfigMissing)
	}
	return cfg, nil
}

func startupPipelineBroken(env string) error {
	_, err := loadConfigBroken(env)
	if err != nil {
		return fmt.Errorf("startupPipeline: %w", err) // %w here doesn't help — chain already broken above
	}
	return nil
}

func ErrorChainBrokenDemo() {
	println("\n\n02. Broken Error Chain — %v Instead of %w:")

	err := startupPipelineBroken("staging")
	if err != nil {
		fmt.Println("Error string (still readable):", err)

		// This returns FALSE — even though ErrConfigMissing was the original error
		// %v broke the chain: the sentinel was converted to a string, not wrapped
		if errors.Is(err, ErrConfigMissing) {
			fmt.Println("errors.Is() found ErrConfigMissing") // never prints
		} else {
			fmt.Println("errors.Is() returned false — chain broken by %v, sentinel lost")
		}
	}

	fmt.Println("\nLesson: always use %w (not %v or %s) when wrapping errors you want to inspect later.")
}
