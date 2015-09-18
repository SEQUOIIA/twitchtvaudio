var Link = ReactRouter.Link

var HomeHeaderComponent = React.createClass({
   render: function() {
       return (
           <div className="headerTitle">
               <div className="logo">
                   {this.props.title}
               </div>
               <div className="subTitle">Listen to Twitch.TV streams instead of watching.</div>
               <div className="seperator"></div>
           </div>
       );
   }
});

var CommentBox = React.createClass({
    render: function() {
        return (
            <div className="commentBox">
                Hello, world! I am a CommentBox.
            </div>
        );
    }
});


var HomeContentComponent = React.createClass({
    onUpdate: function(val) {
        this.setState({
            vodurl: val
        });
    },

    getInitialState: function() {
        return {
            vodurl: 'twitch.tv/username'
        }
    },

   render: function() {
       return (
           <div className="content">
               <h2>Enter URL</h2>
               <InputBar vodurl={this.state.vodurl} onUpdate={this.onUpdate} />
               <button type="submit" form="vodLookup" value="submit" className="downloadButton">LISTEN</button>
               <Link to="/reset">Reset</Link>
           </div>
       );
   }
});

var InputBar = React.createClass({
    handleSubmit: function(e) {
        e.preventDefault();
        var vodUrl = React.findDOMNode(this.refs.vodurl).value.trim();
        if (!vodUrl) {
            return
        }
        console.log(vodUrl);
        React.findDOMNode(this.refs.vodurl).value = '';

        this.props.onUpdate(vodUrl);

        React.render(
            <HomeHeaderComponent title='Nopelol' />,
            document.getElementById('header')
        );
    },

    componentDidMount: function(){
        this.refs.vodurl.getDOMNode().focus();
    },

   render: function() {
       return (
           <form className="inputBar" id="vodLookup" onSubmit={this.handleSubmit} action="lookup">
               <input className="input" type="text" placeholder={this.props.vodurl} ref="vodurl" />
           </form>
       )
   }
});

var indexPage = React.createClass({
    render() {
        return (
            <div id="app">
                <HomeHeaderComponent title='Twitch audio-only retriever' />
                <HomeContentComponent />
            </div>
        )
    }
});


var Route = ReactRouter.Route;
var Router = ReactRouter.Router;
var DefaultRoute = ReactRouter.DefaultRoute;
var RouteHandler = ReactRouter.RouteHandler;



React.render((
    <Router>
        <Route path="/" component={indexPage}>
            <Route path="reset" component={HomeContentComponent} />
        </Route>
    </Router>
), document.body)

/*
React.render(
    <HomeHeaderComponent title='Twitch audio-only retriever' />,
    document.getElementById('header')
);

React.render(
    <HomeContentComponent />,
    document.getElementById('content')
);
*/