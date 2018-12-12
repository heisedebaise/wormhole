import auth from '../util/auth';

class ProducerWs {
    private unique: string;
    private ws: WebSocket;

    public constructor() {
        this.unique = auth.producer();
        this.ws = new WebSocket('wss://' + location.hostname + ':8193/whws');
        this.ws.onmessage = this.receive.bind(this);
    }

    public send(blob: Blob): void {
        const reader: FileReader = new FileReader();
        reader.onload = () => {
            this.ws.send(JSON.stringify({
                auth: this.unique,
                operation: 'speech.produce',
                unique: '' + new Date().getTime(),
                type: 'video',
                content: reader.result
            }));
        };
        reader.readAsDataURL(blob);
    }

    public receive(event: MessageEvent): void {
        console.log(event.data);
    }

    public close(): void {
        this.ws.close();
    }
}

export default new ProducerWs();