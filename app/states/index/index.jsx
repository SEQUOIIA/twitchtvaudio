import React from 'react';
import {Router, Link} from 'react-router';

var HomeHeaderComponent = React.createClass({
   render: function() {
       return (
           <div className="headerTitle">
               <div className="logo">
                   {this.props.title}
               </div>
               <div className="subTitle">If you'd rather listen to Twitch.TV streams instead of watching.</div>
               <div className="seperator"></div>
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
            vodurl: 'e.g. twitch.tv/riotgames = "riotgames"'
        }
    },

   render: function() {
       return (
           <div className="content">
               <h2>Enter streamer username</h2>
               <InputBar vodurl={this.state.vodurl} onUpdate={this.onUpdate} />
               <button type="submit" form="vodLookup" value="submit" className="downloadButton">LISTEN</button>
               {/* <Link query={{Streamer: "twitchchannel"}} to={`/stream`}>Reset</Link> */}
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

        window.location = "#/stream?Streamer=" + vodUrl;
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

var app = React.createClass({
    render() {
        return (
            <div id="app">
                {this.props.children}
            </div>
        )
    }
});

var indexPage = React.createClass({
   render: function() {
       return (
           <div>
               <HomeHeaderComponent title='Twitch audio-only retriever' />
               <HomeContentComponent />
           </div>
       )
   }
});

module.exports.app = app;
module.exports.indexPage = indexPage;