class Http {
    public post(uri: string, parameter?: object): Promise<any> {
        return fetch('https://' + location.hostname + ':8193' + uri, {
            method: 'POST',
            mode: 'cors',
            body: JSON.stringify(parameter)
        }).then(response => response.ok ? response.json() : null);
    }
}

export default new Http();