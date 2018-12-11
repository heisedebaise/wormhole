import * as React from 'react';
import ws from './ws';
import './index.css';

export default class Consumer extends React.Component {
    constructor(props: object) {
        super(props);

        this.start = this.start.bind(this);
    }

    public render(): JSX.Element {
        return (
            <div id="consumer">
                <video autoPlay={true} />
                <br />
                <button onClick={this.start}>Start</button>
            </div>
        );
    }

    private start(): void {
        ws.auth();
    }
}