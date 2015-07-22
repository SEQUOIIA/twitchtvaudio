<!DOCTYPE html>

<html>
<head id="head">
    <title>Twitch audio-only</title>
    <link rel="stylesheet" href="css/foo.css">
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <meta name="viewport" content="initial-scale = 1.0,maximum-scale = 1.0">
    <script src="//code.jquery.com/jquery-2.1.4.min.js"></script>
    <script>
        (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
            (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
                m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
        })(window,document,'script','//www.google-analytics.com/analytics.js','ga');

        ga('create', 'UA-38997729-5', 'auto');
        ga('send', 'pageview');

    </script>
    <script>
        var iOS = ( navigator.userAgent.match(/iPad|iPhone|iPod/g) ? true : false );
        var Mac = ( navigator.userAgent.match(/Macintosh/g) ? true : false );
        if (iOS || Mac) {
            console.log("Apple device detected, switching to original M3U8 URL due to issues with the modified URL");
            $('head').append('<meta http-equiv="refresh" content="0; url={{.resulturl}}" />');
        } else {
            $('head').append('<meta http-equiv="refresh" content="0; url=https://letr.it/twitchaudio/stream/{{.channel}}.m3u8" />');
        }
    </script>
</head>
<header>
    <h1><h1><a href="https://letr.it/twitchaudio" class="title">Twitch Audio-only retriever</a></h1></h1>
</header>
<body>
<section style="text-align: center">
    <span id="title" style="color: rgba(133, 228, 16, 0.74)">Success!</span>
    <br>
    <span id="title2" style="color: #333">Opening stream right away!</span>
    <br>
    <br>
    <span id="subtext">If it doesn't start to download/stream automatically, you can use <a href="https://letr.it/twitchaudio/stream/{{.channel}}.m3u8">this link</a> instead</span>
</section>
</body>
<footer>
    <p>Built by <a href="https://github.com/equoia">Sequoia</a> | Version: twitchaudio<a href="https://github.com/equoia/twitchtvaudio/commits/master">{{.Version}}</a></p>
</footer>
</html>