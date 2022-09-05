import React from 'react';
import ReactDOM from 'react-dom';
import './index.scss';
import App from './App';
import {store} from './app/store';
import {Provider} from 'react-redux';
import * as serviceWorker from './serviceWorker';
import { BrowserRouter } from "react-router-dom";
import {
    RelayEnvironmentProvider,
} from 'react-relay/hooks';

import "./i18n";
import RelayEnvironment from "./RelayEnvironment";
import {LoaderSuspense} from "./components/common/Loader";


ReactDOM.render(
    <React.StrictMode>
        <Provider store={store}>
            <RelayEnvironmentProvider environment={RelayEnvironment({store})}>
                <LoaderSuspense>
                    <BrowserRouter>
                        <App />
                    </BrowserRouter>
                </LoaderSuspense>
            </RelayEnvironmentProvider>
        </Provider>
    </React.StrictMode>,
    document.getElementById('root')
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
