import React from 'react';
import { Route, Switch} from 'react-router-dom'
import App from './App'

//componentes
import Home from './Componentes/Home.js';
import PageCopy from './Componentes/PageCopy.js';

// http://localhost:3000/google.com
const AppRouter = () =>
    <App>
        <Switch>
            <Route exact path="/" component={Home}/>
            <Route  path="/:urlToClone" component={PageCopy}/>

        </Switch>
    </App>

//lo que exportamos es directamente una variable solo
export default AppRouter;