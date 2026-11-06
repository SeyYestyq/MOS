# Skills в hotel — единственный источник для агентов

Агенту достаточно **`hotel/AGENTS.md` + один файл из `skills/` по роли**.  
Папку `claude-skills-main` **не открывать** на ночи: там общие enterprise-скиллы (Next, Terraform, load-test) — шум и прыжки по путям.

| Файл | Роль |
|------|------|
| `skills/01-platform-rr.md` | P1 platform / RR |
| `skills/02-backend-activity-bff.md` | P2 BFF |
| `skills/03-frontend-activity-map.md` | P3 UI |
| `skills/04-pitch-demo.md` | P4 питч |
| `skills/05-freeze-review.md` | перед freeze |
| `skills/06-handoff-acid.md` | приёмка стыков |

Промпт агенту:  
`Сначала прочитай ../hotel/AGENTS.md и ../hotel/skills/0N-….md. Другие skill-репозитории не читай.`
