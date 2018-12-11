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
            auth: this.unique,
            operation: 'speech.consumer'
        }));
    }

    public receive(event: MessageEvent): void {
        const video: HTMLVideoElement | null = document.querySelector('#consumer video');
        if (video === null) {
            return;
        }

        video.src=JSON.parse(event.data).content;

        // const source: HTMLSourceElement = document.createElement('source');
        // source.type = 'video/webm';
        // source.src = JSON.parse(event.data).content;
        // video.appendChild(source);
        video.load();
        video.play();
    }
}

export default new ProducerWs();