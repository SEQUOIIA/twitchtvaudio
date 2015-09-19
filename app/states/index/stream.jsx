import React from 'react';
import Link from 'react-router';

var StreamLoadingComponent = React.createClass({
    render: function() {
        var profileicon;
        if (this.props.profileicon) {
            profileicon = <div className="contentloading"><img className="profileicon" src={this.props.profileicon} /></div>;
        }

        return (
            <div className="headerTitle genericFadeIn">
                <div className="logo">Loading...</div>
                <div className="seperator"></div>
                <br />
                {profileicon}
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
        var profileicon;

        if (this.props.profileicon) {
            profileicon = <div className="contentloading"><img className="profileicon" style={{border: '2px solid #E3101B'}} src={this.props.profileicon} /></div>;
        }

        return (
            <div className="headerTitle genericFadeIn">
                <div className="seperator"></div>
                <br />
                {profileicon}
                <div style={{textAlign: 'center'}}><span style={{fontWeight: 700}}>{this.props.streamer}</span> seems to be offline.</div>
                <br />
                <a href="/">New search</a>
                {/*<Link to="/">Frontpage</Link>*/}
            </div>
        )
    }
});

var streamPage = React.createClass({
    componentDidMount: function() {
        this.getProfile();
        this.getStream();
    },

    getInitialState: function() {
        return {
            streamer: this.props.location.query.Streamer,
            error: false,
            loadingstatus: false,
            apiProfile: {},
            profileloadedStatus: false
        }
    },

    getStream: function() {
        var request = new XMLHttpRequest();
        request.open('GET', 'https://api.twitch.tv/kraken/streams/' + this.state.streamer, true);

        request.onload = function() {
            if (request.status >= 200 && request.status < 400) {
                var data = JSON.parse(request.responseText);
                if (data.stream) {
                    console.log('Stream is live!');
                } else {
                    console.log('Stream is not live.');
                    setTimeout(function() {
                        this.setState({
                            loadingstatus: true,
                            error: true
                        });
                    }.bind(this), 1300);
                }
            } else {
                console.log('API request failed.')
            }
        }.bind(this);

        request.onerror = function() {
            console.log('request error');
        };

        request.send();
    },

    getProfile: function() {
        var request = new XMLHttpRequest();
        request.open('GET', 'https://api.twitch.tv/kraken/users/' + this.state.streamer, true);

        request.onload = function() {
          if (request.status >= 200 && request.status < 400) {
              var data = JSON.parse(request.responseText);
              this.setState({
                  apiProfile: data,
                  profileloadedStatus: true
              });
          } else {
              console.log('API request failed.')
          }
        }.bind(this);

        request.onerror = function() {
            console.log('request error');
        };

        request.send();
    },


    render: function() {
        var content;

        if (this.state.loadingstatus) {
            if (this.state.error) {
                if (this.state.profileloadedStatus) {
                    content = <StreamErrorComponent streamer={this.state.streamer} profileicon={this.state.apiProfile.logo} />;
                } else {
                    content = <StreamErrorComponent streamer={this.state.streamer} />;
                }
            } else {
                // TODO: Loading is successful.
            }
        } else {
            if (this.state.profileloadedStatus) {
                content = <StreamLoadingComponent streamer={this.state.streamer} profileicon={this.state.apiProfile.logo} />
            } else {
                content = <StreamLoadingComponent streamer={this.state.streamer} />;
            }
        }

        return (
            <div>
                {content}
            </div>
        )
    }
});

module.exports.streamPage = streamPage;