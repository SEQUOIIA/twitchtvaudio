var Indexfile = require('./states/index/index.jsx');
var Vodfile = require('./states/index/vod.jsx');


import React from 'react';
import createHashHistory from './node_modules/react-router/node_modules/history/lib/createHashHistory';
import {Router, Route, IndexRoute, DefaultRoute, RouteHandler} from 'react-router';



React.render((
    <Router history={createHashHistory({queryKey: false})}>
        <Route path="/" component={Indexfile.app}>
            <IndexRoute component={Indexfile.indexPage} />
            <Route path="stream" component={Indexfile.streamPage} />
        </Route>
    </Router>
), document.body)