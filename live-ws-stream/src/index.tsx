import * as React from 'react';
import * as ReactDOM from 'react-dom';
import Producer from './producer';
import Consumer from './consumer';
import './index.css';
import registerServiceWorker from './registerServiceWorker';

class Index extends React.Component {
  public render(): JSX.Element {
    return (
      <div>
        <Producer />
        <Consumer />
      </div>
    );
  }
}

ReactDOM.render(
  <Index />,
  document.getElementById('root') as HTMLElement
);
registerServiceWorker();
