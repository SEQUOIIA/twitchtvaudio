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
                <a href="#/" style={{textDecoration: 'none'}}><button type="button" form="vodLookup" value="submit" className="genericButton">NEW SEARCH</button></a>
                {/*<Link to="/">Frontpage</Link>*/}
            </div>
        )
    }
});

var StreamSuccessComponent = React.createClass({
   componentDidMount: function() {
       setTimeout(function() {
           window.open(this.props.hls, '_blank');
       }.bind(this), 2000);
   },

   render: function() {
       var profileicon;
       if (this.props.profileicon) {
           profileicon = <div className="contentloading"><img className="profileicon" style={{border: '2px solid #92D62B'}} src={this.props.profileicon} /></div>;
       }
       return (
           <div className="headerTitle genericFadeIn">
               {/*<meta httpEquiv="refresh" content={"2; url=" + this.props.hls} /> */}
               <div className="seperator"></div>
               <br />
               {profileicon}
               <div style={{textAlign: 'center'}}><span style={{fontWeight: 700}}>{this.props.streamer}</span> is live!</div>
               <br />
               <div style={{textAlign: 'center'}}>Opening the audio stream in a new tab</div>
               <div style={{textAlign: 'center', fontSize: 'smaller'}}>Enable pop-ups for this site, if you haven't already.</div>
               <br />
               <a href="#/" style={{textDecoration: 'none'}}><button type="button" form="vodLookup" value="submit" className="genericButton">NEW SEARCH</button></a>
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
            profileloadedStatus: false,
            hlsStream: ''
        }
    },

    getStream: function() {
        var request = new XMLHttpRequest();
        request.open('GET', 'https://api.twitch.tv/kraken/streams/' + this.state.streamer, true);

        request.onload = function() {
            if (request.status >= 200 && request.status < 400) {
                var data = JSON.parse(request.responseText);
                if (data.stream) {
                    //console.log('Stream is live!');

                    var m3u8request = new XMLHttpRequest();
                    m3u8request.open('GET', 'http://twitchaudio.seq.tf:8089/api/stream/' + this.state.streamer, true);
                    m3u8request.onload = function() {
                        if (m3u8request.status >= 200 && m3u8request.status < 400) {
                            var data = JSON.parse(m3u8request.responseText);

                            setTimeout(function() {
                                this.setState({
                                    loadingstatus: true,
                                    error: false,
                                    hlsStream: data.url
                                });
                            }.bind(this), 1300);
                        } else {
                            //console.log('Stream is not live.');
                            setTimeout(function() {
                                this.setState({
                                    loadingstatus: true,
                                    error: true
                                });
                            }.bind(this), 1300);
                        }
                    }.bind(this);

                    m3u8request.onerror = function() {
                        console.log('request error');
                    };

                    m3u8request.send();
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
                if (this.state.profileloadedStatus) {
                    content = <StreamSuccessComponent streamer={this.state.streamer} profileicon={this.state.apiProfile.logo} hls={this.state.hlsStream} />;
                } else {
                    content = <StreamSuccessComponent streamer={this.state.streamer} hls={this.state.hlsStream} />;
                }
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