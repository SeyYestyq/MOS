# Git Workflow Standards

## Core Principles

### **1. Conventional Commits**
- All commits follow conventional commits specification
- Enable automated changelog generation
- Clear categorization of all changes
- Explicit marking of breaking changes

### **2. Semantic Versioning**
- Automated versioning based on commit types
- MAJOR version: BREAKING CHANGE
- MINOR version: feat type commits
- PATCH version: fix type commits

### **3. Branch Protection**
- Main branch requires review
- All quality gates must pass
- No force pushes to main
- Signed commits recommended

## Conventional Commits Specification

### **Commit Message Structure**

```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

### **Standard Types**

```yaml
feat:     # New features (agents, skills, tools)
fix:      # Bug fixes
docs:     # Documentation only changes
style:    # Formatting changes (no functional impact)
refactor: # Code restructuring without feature changes
perf:     # Performance improvements
test:     # Adding or correcting tests
chore:    # Maintenance tasks (dependencies, build)
ci:       # CI/CD configuration changes
```

### **Examples**

```bash
# New agent implementation
feat(agents): implement cs-content-creator agent

# Bug fix in Python tool
fix(seo-optimizer): correct keyword density calculation

# Documentation update
docs(README): add agent catalog section

# Breaking change
feat(agents)!: replace agent template structure

BREAKING CHANGE: agent frontmatter now requires 'tools' field
```

### **Scopes for Claude Skills**

```yaml
Common Scopes:
  agents:           # Agent implementation
  skills:           # Skill packages
  standards:        # Standards library
  templates:        # Templates
  docs:             # Documentation
  foundation:       # Core infrastructure
  marketing:        # Marketing domain
  c-level:          # C-level domain
  product:          # Product domain
  engineering:      # Engineering domain
  ra-qm:            # RA/QM domain
```

## Branch Strategy

### **Branch Naming Conventions**

```yaml
# Feature branches (agent/skill development)
feature/agents-{name}:         # New agent implementation
feature/skills-{name}:         # New skill package
feature/{domain}-{component}:  # Domain-specific features

# Documentation branches
docs/{component}:              # Documentation updates
docs/standards-{name}:         # Standards library updates

# Bug fixes
fix/{issue-number}-{description}: # Bug fixes

# Hotfixes (critical production issues)
hotfix/{issue}-{description}:   # Emergency fixes

# Test branches
test/{feature}:                # Testing and validation
```

### **Examples**

```bash
# Create feature branch for new agent
git checkout -b feature/agents-ceo-advisor

# Create documentation branch
git checkout -b docs/installation-guide

# Create fix branch
git checkout -b fix/23-broken-relative-paths

# Create test branch
