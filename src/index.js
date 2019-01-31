import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import * as serviceWorker from './serviceWorker';
import AppRouter from './Router';
import { BrowserRouter } from 'react-router-dom';

ReactDOM.render(<BrowserRouter>
    <AppRouter/>
</BrowserRouter>, document.getElementById('root'));
serviceWorker.unregister();
