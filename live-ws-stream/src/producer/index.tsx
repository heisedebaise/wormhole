import * as React from 'react';
import ws from './ws';
import './index.css';

export default class Producer extends React.Component {
    constructor(props: object) {
        super(props);

        this.start = this.start.bind(this);
    }

    public render(): JSX.Element {
        return (
            <div id="producer">
                <video autoPlay={true} playsInline={true} />
                <br />
                <button onClick={this.start}>Start</button>
            </div>
        );
    }

    private start(): void {
        navigator.mediaDevices.getUserMedia({
            audio: true,
            video: true
        }).then((stream: MediaStream) => {
            const video: HTMLVideoElement | null = document.querySelector('#producer video');
            if (video === null) {
                return;
            }

            video.srcObject = stream;
            const mediaRecorder: MediaRecorder = new MediaRecorder(stream, {
                mimeType: 'video/webm',
                bitsPerSecond: 50 * 1024
            });
            mediaRecorder.ondataavailable = (event: BlobEvent) => {
                ws.send(event.data);
            };
            mediaRecorder.start(1000);
        });
    }
}
