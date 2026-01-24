# Application Module Context

Purpose: application services that orchestrate domain use-cases.

Key folders:
- `application/` contains services (e.g., `application/user`) that use domain interfaces.

Notes:
- Services should depend on domain interfaces and be wired to infrastructure implementations in an initialization step.
