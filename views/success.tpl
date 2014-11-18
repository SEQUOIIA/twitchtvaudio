<!DOCTYPE html>

<html>
  	<head>
    	<title>Twitch audio-only</title>
    	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
        <meta http-equiv="refresh" content="0; url={{.resulturl}}" />
        <script>
            (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
                (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
                    m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
            })(window,document,'script','//www.google-analytics.com/analytics.js','ga');

            ga('create', 'UA-38997729-5', 'auto');
            ga('send', 'pageview');

        </script>
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

            #title {
                color: #333;
                text-decoration: none;
            }

            #title:hover {
                text-decoration: none;
                color: rgb(84, 83, 83);
            }
		</style>
	</head>

  	<body>
  		<header class="hero-unit" style="background-color:rgba(166, 142, 210, 0.76)">
			<div class="container">
			<div class="row">
			  <div class="hero-text" style="margin-left: 3em">
                  <a id="title" href="https://letr.it/twitchaudio"><h1>Twitch audio-only retriever</h1></a>
			    <p class="description">
                    Developer: <a href="https://github.com/equoia">Sequoia</a>
                    <br>
                    Version: <a href="https://github.com/equoia/twitchtvaudio/commits/master">{{.Version}}</a>
			    </p>
			  </div>
			</div>
			</div>
		</header>
    <div style="padding-top: 0.2em; margin-left: auto; margin-right: auto; text-align: center; line-height: 1.5em;">
        <h1 style="font-size: 5em; color: rgba(133, 228, 16, 0.74)">Success!</h1>
        <p style="font-size: 5em; padding-top: 0.2em; line-height: 1em;">
        Opening stream right away!</p>
        <p style="font-size: 0.9em; padding-top: 0.5em;">
        <span>If it doesn't start to download/stream automatically, you can use <a href="{{.resulturl}}">this link</a> instead</span>
        </p>
    </div>
	</body>
</html>
