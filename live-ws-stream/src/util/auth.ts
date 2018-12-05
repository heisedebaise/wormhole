import http from './http';

class Auth {
    public consumer(): string {
        const unique: string = 'live-ws-stream:consumer';
        this.post('consumer', 'live-ws-stream', unique);

        return unique;
    }

    public producer(): string {
        const unique: string = 'live-ws-stream:producer';
        this.post('producer', 'live-ws-stream', unique);

        return unique;
    }

    private post(name: string, auth: string, unique: string): Promise<any> {
        return http.post('/auth/' + name, {
            auth: auth,
            unique: unique
        });
    }
}

export default new Auth();