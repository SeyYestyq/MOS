import { useMemo, useState } from "react";
import { ProcessTreeBoard } from "./ProcessTreeBoard";
import {
  NAMESPACE,
  POD,
  events,
  netEdges,
  processes,
  type ProcessNode,
} from "./mockData";

export default function App() {
  const [selectedId, setSelectedId] = useState<string | null>("p5");

  const selected = useMemo(
    () => processes.find((p) => p.id === selectedId) ?? null,
    [selectedId],
  );

  const relatedNet = useMemo(() => {
    if (!selected) return netEdges;
    return netEdges.filter((e) => e.fromPid === selected.pid);
  }, [selected]);

  const relatedEvents = useMemo(() => {
    if (!selected) return events;
    return events.filter((e) => e.pid === selected.pid);
  }, [selected]);

  function onSelect(node: ProcessNode) {
    setSelectedId(node.id);
  }

  return (
    <div className="app-shell">
      <header className="app-header">
        <div>
          <h1>Container Activity Map</h1>
          <p>
            Локальный превью-борд для хакатона: дерево процессов + сеть + таймлайн
            в стиле liquid-metal / topology. Данные — mock сценария reverse-shell.
            В репозиторий OnlyGo не пушится.
          </p>
        </div>
        <div className="badge-row">
          <span className="badge badge--live">preview · local</span>
          <span className="badge">{NAMESPACE}/{POD}</span>
          <span className="badge">Runtime Radar → fixtures</span>
        </div>
      </header>

      <div className="layout">
        <section className="admin-topology-outer">
          <div className="admin-topology-wrap admin-surface liquid-metal-frame liquid-metal-frame--soft">
            <div className="admin-surface-head">
              <h3>Process tree</h3>
              <span>ppid → pid</span>
            </div>
            <ProcessTreeBoard
              processes={processes}
              selectedId={selectedId}
              onSelect={onSelect}
            />
            <p className="hint">
              Клик по узлу синхронизирует сеть и таймлайн справа — как в ТЗ PT.
            </p>
          </div>
        </section>

        <aside className="side-stack">
          <article className="admin-surface liquid-metal-frame liquid-metal-frame--soft">
            <div className="admin-surface-head">
              <h3>Selected process</h3>
              <span>{selected ? `pid ${selected.pid}` : "—"}</span>
            </div>
            {selected ? (
              <ul className="detail-list">
                <li>
                  <strong>binary</strong>
                  <span>{selected.binary}</span>
                </li>
                <li>
                  <strong>cmdline</strong>
                  <span>{selected.cmdline}</span>
                </li>
                <li>
                  <strong>parent</strong>
                  <span>ppid {selected.ppid}</span>
                </li>
                <li>
                  <strong>role</strong>
                  <span>{selected.role}</span>
                </li>
              </ul>
            ) : (
              <p className="hint">Выбери процесс на карте</p>
            )}
          </article>

          <article className="admin-surface liquid-metal-frame liquid-metal-frame--soft">
            <div className="admin-surface-head">
              <h3>Network</h3>
              <span>{relatedNet.length}</span>
            </div>
            <div className="side-stack">
              {relatedNet.length === 0 ? (
                <p className="hint">Нет сетевых связей для этого pid</p>
              ) : (
                relatedNet.map((edge) => (
                  <div key={edge.id} className="net-row">
                    <div>
                      <div style={{ color: "var(--text-soft)", marginBottom: 4 }}>
                        {edge.label} · pid {edge.fromPid}
                      </div>
                      <code>
                        {edge.proto}://{edge.dst}:{edge.port}
                      </code>
                    </div>
                  </div>
                ))
              )}
            </div>
          </article>

          <article className="admin-surface liquid-metal-frame liquid-metal-frame--soft">
            <div className="admin-surface-head">
              <h3>Timeline</h3>
              <span>{relatedEvents.length} events</span>
            </div>
            <div className="timeline">
              {(relatedEvents.length ? relatedEvents : events).map((ev) => {
                const proc = processes.find((p) => p.pid === ev.pid);
                const active = proc?.id === selectedId;
                return (
                  <button
                    key={ev.id}
                    type="button"
                    className={`timeline-item ${active ? "is-active" : ""}`}
                    onClick={() => proc && setSelectedId(proc.id)}
                  >
                    <time>{ev.t}</time>
                    <div>
                      <div className="ev-type">{ev.type}</div>
                      <div className="ev-sum">{ev.summary}</div>
                    </div>
                  </button>
                );
              })}
            </div>
          </article>
        </aside>
      </div>
    </div>
  );
}
