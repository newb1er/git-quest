import { useEffect, useState } from "react";
import { WebSocketController } from "../websocket/controller";
import { ControlMessage } from "../websocket/control-msg";

const sendControlMessage = (
  socketController: WebSocketController | null,
  message: ControlMessage
) => {
  if (!socketController) {
    return;
  }

  socketController.sendControlMessage(message);
};

export const useWebsocket = (url: string) => {
  const [socketController, setSocketController] =
    useState<WebSocketController | null>(null);

  useEffect(() => {
    const ws = new WebSocket(url);
    setSocketController(new WebSocketController(ws));
    return () => {
      socketController?.close();
    };
  }, [url]);

  return {
    isReady: socketController?.isReady ?? false,
    sendControlMessage: sendControlMessage.bind(null, socketController),
  };
};
