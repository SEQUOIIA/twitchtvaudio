import React from 'react';

var StreamLoadingComponent = React.createClass({
    render: function() {
        return (
            <div className="headerTitle">
                <div className="logo">Loading...</div>
                <div className="seperator"></div>
                <br />
                <div style={{textAlign: 'center'}}>Checking if <span style={{fontWeight: 700}}>{this.props.streamer}</span> is streaming. </div>
                <div className="spinner">
                    <div className="double-bounce1"></div>
                    <div className="double-bounce2"></div>
                </div>
                <br/>
                {/*<Link to="/">Frontpage</Link>*/}
            </div>
        )
    }
});

var StreamErrorComponent = React.createClass({
    render: function() {
        return (
            <h2>Dunno, something went wrong.</h2>
        )
    }
});

var streamPage = React.createClass({
    componentDidMount: function() {
    },

    getInitialState: function() {
        return {
            streamer: this.props.location.query.Streamer,
            error: false,
            loadingstatus: false
        }
    },

    render: function() {
        var content;
        console.log(this.state.loadingstatus);
        if (this.state.loadingstatus) {
            content = <StreamErrorComponent />;
        } else {
            content = <StreamLoadingComponent streamer={this.state.streamer} />;
        }

        return (
            <div>
                {content}
            </div>
        )
    }
});

module.exports.streamPage = streamPage;