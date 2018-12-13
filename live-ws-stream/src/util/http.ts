class Http {
    public post(uri: string, parameter?: object): Promise<string> {
        const form: URLSearchParams = new URLSearchParams();
        if (parameter) {
            for (const name in parameter) {
                if (parameter.hasOwnProperty(name)) {
                    form.append(name, parameter[name]);
                }
            }
        }

        return fetch('https://' + location.hostname + ':8193' + uri, {
            method: 'POST',
            mode: 'cors',
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded'
            },
            body: form
        }).then(response => response.ok ? response.text() : '');
    }
}

export default new Http();