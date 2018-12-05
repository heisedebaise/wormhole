import auth from '../util/auth';

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
            auth: this.unique
        }));
    }

    public receive(event: MessageEvent): void {
        // const video: HTMLVideoElement | null = document.querySelector('#consumer video');
        const source: HTMLSourceElement | null = document.querySelector('#consumer video source');
        if (source === null) {
            return;
        }

        const data = JSON.parse(event.data);
        console.log(data.Content);
        source.src = data.Content;
        // video.play();
    }
}

export default new ProducerWs();