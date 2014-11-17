<!DOCTYPE html>

<html>
  	<head>
    	<title>Twitch audio-only</title>
    	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
        <script>
            (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
                (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
                    m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
            })(window,document,'script','//www.google-analytics.com/analytics.js','ga');

            ga('create', 'UA-38997729-5', 'auto');
            ga('send', 'pageview');

        </script>
        <script src="https://code.jquery.com/jquery-2.1.1.min.js"></script>

		<style type="text/css">
			body {
				margin: 0px;
				font-family: "Helvetica Neue",Helvetica,Arial,sans-serif;
				font-size: 14px;
				line-height: 20px;
				color: rgb(51, 51, 51);
				background-color: rgb(255, 255, 255);
			}

			.hero-unit {
				padding: 30px;
				margin-bottom: 30px;
			}

			.container {
				width: 100%;
				margin-right: auto;
				margin-left: auto;
			}

			.row {
				margin-left: -20px;
			}

			h1 {
				margin: 10px 0px;
				font-family: inherit;
				font-weight: bold;
				text-rendering: optimizelegibility;
			}

			.hero-unit h1 {
				margin-bottom: 0px;
				font-size: 60px;
				line-height: 1;
				letter-spacing: -1px;
				color: inherit;
			}

			.description {
				padding-top: 0px;
				padding-left: 5px;
				font-size: 18px;
				font-weight: 200;
				line-height: 25px;
				color: inherit;
			}

			p {
				margin: 0px 0px 0px;
			}
		</style>
	</head>

  	<body>
  		<header class="hero-unit" style="background-color:rgba(166, 142, 210, 0.76)">
			<div class="container">
			<div class="row">
			  <div class="hero-text" style="margin-left: 3em">
			    <h1>Twitch audio-only retriever</h1>
			    <p class="description">
                    Developer: <a href="https://github.com/equoia">Sequoia</a>
                    <br>
                    Version: <a href="https://github.com/equoia/twitchtvaudio/commits/master">{{.Version}}</a>
			    </p>
			  </div>
			</div>
			</div>
		</header>
    <div id="stuff">
  <!--      <p style="margin-left: auto; margin-right: auto; text-align: center; font-size: 3em; line-height: 1.2em">Oi, you forgot to put a channel name after the URL!</p> -->
        <form action="http://192.168.0.133:8089/get" method="GET">
            <div id="name-group" class="form-group" style="margin-left: auto; margin-right: auto; width: 13em">
                <label for="channelname">Channel</label>
                <input type="text" class="form-control" name="channelname" placeholder="totalbiscuit">
                <!-- errors will go here -->
            </div>
            <div style="margin-left: auto; margin-right: auto; width: 1em">
            <button type="submit" class="btn btn-success">Submit <span class="fa fa-arrow-right"></span></button>
            </div>
            <div id="twitchresponse" style="text-align: center; font-size: 3em; padding-top: 0.6em;"></div>
        </form>

    </div>
	</body>
</html>
