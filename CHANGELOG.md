
# CHANGELOG

## [v0.2.0](https://github.com/tab/smartid/releases/tag/v0.2.0)

### Features
- **feat(tls):** Add certificate pinning support in the Smart-ID client using TLS configuration

### Bug Fixes
- **fix(identity):** Fix identity regex
- **fix(requests):** Support custom HTTP status codes: 471, 472, 480, 580

### Tests
- **test(utils):** Enhance certificate extraction tests

### Chore
- **chore(workflow):** Add workflow permissions
- **chore(codecov):** Update ignore patterns in codecov.yaml

## [v0.1.1](https://github.com/tab/smartid/releases/tag/v0.1.1)

### Refactor
- **refactor(models):** Rename AuthenticationSessionRequest to AuthenticationRequest

### Documentation
- **docs(readme):** Update README with sections for installation, client creation, authentication session initiation, session fetching, asynchronous processing, and identity preparation

### CI
- **ci(workflow):** Add Codecov workflow for master branch

## [v0.1.0](https://github.com/tab/smartid/releases/tag/v0.1.0)

### Features
- **feat(client):** Add identity to client using new Identity type for authentication
- **feat(client):** Add Smart-ID client example for authentication
- **feat(staticcheck):** Add staticcheck configuration and Makefile target
- **feat(worker):** Refactor authentication in worker

### Bug Fixes
- **fix(authentication):** Enhance authentication error handling with defined error codes
- **fix(authentication):** Change authentication session naming for consistency
- **fix(workflow):** Update checks workflow with Go version changes

### Refactor
- **refactor(authenticate):** Refactor Authenticate method to return session response
- **refactor(client):** Simplify smartid authentication method
- **refactor(models):** Replace models.Person with a new Person struct
- **refactor(session):** Refactor session handling and CreateSession method
- **refactor(worker):** Refactor Worker configuration and context handling in Start method
- **refactor(worker):** Refactor worker to pass context in Process method

### CI
- **ci(workflow):** Update GitHub Actions workflow and staticcheck configuration
- **ci(workflow):** Disable Go installation in staticcheck action

### Chore
- **chore(module):** Update module path and import statements
- **chore(deps):** Bump codecov/codecov-action from 4 to 5
- **chore(github):** Add CODEOWNERS file
- **chore(ci):** Add Dependabot configuration and CI checks workflow

### Tests
- **test:** Add example tests for session creation and processing multiple identities

### Style
- **style:** Fix ineffectual assignment in linter checks
