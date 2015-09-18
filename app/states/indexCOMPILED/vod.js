var VodHeaderComponent = React.createClass({displayName: "VodHeaderComponent",
   render: function() {
       return (
           React.createElement("div", {className: "headerTitle"}, 
               React.createElement("div", {className: "logo"}, 
                    "TWIVOD"
               ), 
               React.createElement("div", {className: "subTitle"}, "Download Twitch.TV VODs"), 
               React.createElement("div", {className: "seperator"})
           )
       );
   }
});

var VodContentComponent = React.createClass({displayName: "VodContentComponent",
   render: function() {
       return (
           React.createElement("div", {className: "content"}
           )
       )
   }
});

var vodRender = function() {
    React.render(
        React.createElement(HomeHeaderComponent, null),
        document.getElementById('header')
    );

    React.render(
        React.createElement(HomeContentComponent, null),
        document.getElementById('content')
    );
};