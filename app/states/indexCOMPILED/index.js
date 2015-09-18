var HomeHeaderComponent = React.createClass({displayName: "HomeHeaderComponent",
   render: function() {
       return (
           React.createElement("div", {className: "headerTitle"}, 
               React.createElement("div", {className: "logo"}, 
                    "Twitch audio-only retriever"
               ), 
               React.createElement("div", {className: "subTitle"}, "Download Twitch.TV VODs"), 
               React.createElement("div", {className: "seperator"})
           )
       );
   }
});

var CommentBox = React.createClass({displayName: "CommentBox",
    render: function() {
        return (
            React.createElement("div", {className: "commentBox"}, 
                "Hello, world! I am a CommentBox."
            )
        );
    }
});


var HomeContentComponent = React.createClass({displayName: "HomeContentComponent",
    onUpdate: function(val) {
        this.setState({
            vodurl: val
        });
    },

    getInitialState: function() {
        return {
            vodurl: 'twitch.tv/username/v/01010101'
        }
    },

   render: function() {
       return (
           React.createElement("div", {className: "content"}, 
               React.createElement("h2", null, "Enter URL"), 
               React.createElement(InputBar, {vodurl: this.state.vodurl, onUpdate: this.onUpdate}), 
               React.createElement("button", {type: "submit", form: "vodLookup", value: "submit", className: "downloadButton"}, "DOWNLOAD")
           )
       );
   }
});

var InputBar = React.createClass({displayName: "InputBar",
    handleSubmit: function(e) {
        e.preventDefault();
        var vodUrl = React.findDOMNode(this.refs.vodurl).value.trim();
        if (!vodUrl) {
            return
        }
        console.log(vodUrl);
        React.findDOMNode(this.refs.vodurl).value = '';

        this.props.onUpdate(vodUrl);
    },

    componentDidMount: function(){
        this.refs.vodurl.getDOMNode().focus();
    },

   render: function() {
       return (
           React.createElement("form", {className: "inputBar", id: "vodLookup", onSubmit: this.handleSubmit, action: "lookup"}, 
               React.createElement("input", {className: "input", type: "text", placeholder: this.props.vodurl, ref: "vodurl"})
           )
       )
   }
});

React.render(
    React.createElement(HomeHeaderComponent, null),
    document.getElementById('header')
);

React.render(
    React.createElement(HomeContentComponent, null),
    document.getElementById('content')
);