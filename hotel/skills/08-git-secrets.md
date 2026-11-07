# Git + секреты (все роли) — из borghei env-secrets + git standards

## Conventional commits (OnlyGo)

```
feat(web): process tree with react flow
feat(bff): activity endpoint from fixture
fix(web): sync selection across panels
docs: README run instructions
chore(bff): add requirements
```

Scopes ночи: `web` (A) | `bff` (B) | `scripts`/`docs` (C) | `fixtures`

Pacing: scaffold → fixture → tree → timeline → sync → live → polish.  
Не mega-commit на 40 файлов в T+15м.

## Секреты — никогда в OnlyGo

- RR password / encryption keys / `.rr-install/creds.txt` / `keys.txt`
- Токены в коде → только `os.environ` / `.env` (локально, в `.gitignore`)
- Перед push: `git diff --cached` на password|token|secret|BEGIN

`.env.example` без значений — ок. Живые секреты — нет.

## Co-authored-by

Не допускать `Co-authored-by: Cursor` в сообщениях коммитов.

Детали: `_borghei/env-secrets-manager.md`, `_borghei/git-workflow-standards-excerpt.md`
