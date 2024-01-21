import { ControlMessage } from "./control-msg";

export class WebSocketController {
  private socket: WebSocket | null = null;

  constructor(socket: WebSocket) {
    this.socket = socket;

    this.socket.onmessage = this.onMessage.bind(this);
  }

  get isReady() {
    return this.socket?.readyState === WebSocket.OPEN;
  }

  private onMessage = (event: MessageEvent) => {
    console.log("Received message: ", event.data);
  };

  public close() {
    this.socket?.close();
  }

  public sendControlMessage(message: ControlMessage) {
    if (!this.socket) {
      return;
    }

    this.socket.send(message);
  }
}
