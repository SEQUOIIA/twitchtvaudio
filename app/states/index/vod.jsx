var VodHeaderComponent = React.createClass({
   render: function() {
       return (
           <div className="headerTitle">
               <div className="logo">
                    TWIVOD
               </div>
               <div className="subTitle">Download Twitch.TV VODs</div>
               <div className="seperator"></div>
           </div>
       );
   }
});

var VodContentComponent = React.createClass({
   render: function() {
       return (
           <div className="content">
           </div>
       )
   }
});

var vodRender = function() {
    React.render(
        <HomeHeaderComponent />,
        document.getElementById('header')
    );

    React.render(
        <HomeContentComponent />,
        document.getElementById('content')
    );
};