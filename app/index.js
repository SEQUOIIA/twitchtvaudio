var Indexfile = require('./states/index/index.jsx');
var Streamfile = require('./states/index/stream.jsx');

import React from 'react';
import createHashHistory from './node_modules/react-router/node_modules/history/lib/createHashHistory';
import {Router, Route, IndexRoute, DefaultRoute, RouteHandler} from 'react-router';
import './dist/jquery.js';


React.render((
    <Router history={createHashHistory({queryKey: false})}>
        <Route path="/" component={Indexfile.app}>
            <IndexRoute component={Indexfile.indexPage} />
            <Route path="stream" component={Streamfile.streamPage} />
        </Route>
    </Router>
), document.body)