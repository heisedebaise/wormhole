import auth from '../util/auth';
import player from './player';

class ProducerWs {
    private unique: string;
    private ws: WebSocket;

    public constructor() {
        this.unique = auth.consumer();
        this.ws = new WebSocket('wss://' + location.hostname + ':8193/whws');
        this.ws.onmessage = this.receive.bind(this);
    }

    public auth(): void {
        this.ws.send(JSON.stringify({
            auth: this.unique,
            operation: 'speech.consumer'
        }));

        player.play();
    }

    public pull(): void {
        this.ws.send(JSON.stringify({
            auth: this.unique,
            operation: 'speech.pull',
            type: 'video'
        }));
    }

    public receive(event: MessageEvent): void {
        player.message(JSON.parse(event.data).content);
    }
}

export default new ProducerWs();