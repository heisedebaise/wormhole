import http from './http';

class Auth {
    public producer(): string {
        const unique: string = 'live-ws-stream:producer';
        this.post('producer', 'live-ws-stream', unique);

        return unique;
    }

    public consumer(): string {
        const unique: string = 'live-ws-stream:consumer';
        this.post('consumer', 'live-ws-stream', unique);

        return unique;
    }

    private post(name: string, token: string, ticket: string): Promise<any> {
        return http.post('/whauth/' + name, {
            token: token,
            ticket: ticket
        });
    }
}

export default new Auth();