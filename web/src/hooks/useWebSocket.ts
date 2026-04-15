import { useEffect, useRef, useCallback } from 'react';

export function useWebSocket(projectSlug: string | null, onReload: () => void) {
  const wsRef = useRef<WebSocket | null>(null);
  const retryRef = useRef(1000);

  const connect = useCallback(() => {
    if (!projectSlug) return;
    const proto = location.protocol === 'https:' ? 'wss:' : 'ws:';
    const ws = new WebSocket(`${proto}//${location.host}/ws/${projectSlug}`);

    ws.onmessage = () => {
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
  }, [projectSlug, onReload]);

  useEffect(() => {
    connect();
    return () => { wsRef.current?.close(); };
  }, [connect]);
}
