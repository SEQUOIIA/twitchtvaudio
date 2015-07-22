<!DOCTYPE html>

<html>
<head>
    <title>Twitch audio-only</title>
    <link rel="stylesheet" href="css/foo.css">
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <meta name="viewport" content="initial-scale = 1.0,maximum-scale = 1.0">
    <script>
        (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
            (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
                m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
        })(window,document,'script','//www.google-analytics.com/analytics.js','ga');

        ga('create', 'UA-38997729-5', 'auto');
        ga('send', 'pageview');

    </script>
</head>
<header>
    <h1><h1><a href="https://letr.it/twitchaudio" class="title">Twitch Audio-only retriever</a></h1></h1>
</header>
<body>
<section>
    <span class="titleerror">Can't find audio-only stream</span>
    <div class="searchbar">
        <form action="get" method="GET">
            <div id="searchform">
                <input type="text" class="input" name="channelname" placeholder="Channel" autofocus>
                <!-- errors will go here -->

                <button type="submit" class="btn btn-success" style="border: none; background-color: transparent;">
                    <i class="fa fa-search fa-lg" id="searchicon"></i>
                </button>

            </div>
        </form>
    </div>
</section>
</body>
<footer>
    <p>Built by <a href="https://github.com/equoia">Sequoia</a> | Version: twitchaudio<a href="https://github.com/equoia/twitchtvaudio/commits/master">{{.Version}}</a></p>
</footer>
</html>