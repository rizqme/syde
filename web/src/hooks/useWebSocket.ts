import { useEffect, useRef, useCallback } from 'react';

type DashboardMessage =
  | { type: 'reload'; file?: string }
  | { type: 'navigate'; path: string };

export function useWebSocket(
  projectSlug: string | null,
  onReload: () => void,
  onNavigate?: (path: string) => void,
) {
  const wsRef = useRef<WebSocket | null>(null);
  const retryRef = useRef(1000);

  const connect = useCallback(() => {
    if (!projectSlug) return;
    const proto = location.protocol === 'https:' ? 'wss:' : 'ws:';
    const ws = new WebSocket(`${proto}//${location.host}/ws/${projectSlug}`);

    ws.onmessage = (event) => {
      let msg: DashboardMessage | null = null;
      try {
        msg = JSON.parse(event.data);
      } catch {
        onReload();
        return;
      }
      if (msg?.type === 'navigate' && msg.path) {
        onNavigate?.(msg.path);
        return;
      }
      onReload();
    };

    ws.onclose = () => {
      setTimeout(() => {
        retryRef.current = Math.min(retryRef.current * 2, 30000);
        connect();
      }, retryRef.current);
    };

    ws.onopen = () => {
      retryRef.current = 1000;
    };

    wsRef.current = ws;
  }, [projectSlug, onReload, onNavigate]);

  useEffect(() => {
    connect();
    return () => { wsRef.current?.close(); };
  }, [connect]);
}
