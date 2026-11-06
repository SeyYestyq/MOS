# Focused fix (ночь) — из borghei focused-fix

Когда что-то сломалось за час до демо.

## Do

- Чинить **ровно** баг демо: дерево / sync / fixture / port-forward
- Diff минимальный; в commit message — почему фикс работает
- Один smoke после фикса (открыть UI, клик по suspect)

## Don't

- Рефакторить «заодно»
- Менять формат API без синка с другой ролью
- Трогать чужую path-зону
- Апгрейдить зависимости в том же коммите, что fix

## 10-минутный triage RR (P1)

1. `kubectl -n runtime-radar get pods`  
2. port-forward `:9000` жив?  
3. Events не пустые?  
4. Нет → **fixture**, объявить команде; не чинить kind всю ночь  

Детали: `_borghei/focused-fix.md`, `_borghei/incident-commander.md`
