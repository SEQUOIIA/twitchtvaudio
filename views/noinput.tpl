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
        <script>
            // magic.js
            $(document).ready(function() {

                // process the form
                $('form').submit(function(event) {

                    $('.form-group').removeClass('has-error'); // remove the error class
                    $('.help-block').remove(); // remove the error text

                    // get the form data
                    // there are many ways to get this data using jQuery (you can use the class or id also)
                    var formData = {
                        'channelname' 		: $('input[name=channelname]').val()
                    };

                    // process the form
                    $.ajax({
                        type 		: 'POST', // define the type of HTTP verb we want to use (POST for our form)
                        url 		: 'http://192.168.0.133:8089/api/twitch', // the url where we want to POST
                        data 		: formData, // our data object
                        dataType 	: 'json', // what type of data do we expect back from the server
                        encode 		: true,
                        success: function(data)
                        {
                            if (data.Statuscode == 0) {
                                $('#twitchresponse').empty()
                                $('#twitchresponse').append('<p style="color: rgba(255, 0, 0, 0.61)">' + "Can't find audio-only stream.</p>")
                            }

                            if (data.Statuscode == 1) {
                                $('#twitchresponse').empty()
                                $('#twitchresponse').append('<p><a style="text-decoration: none; color: rgba(133, 228, 16, 0.74);" href="' + data.Url + '">Click here</a></p>')
                            }

                        //    $('#twitchresponse').replace('<p>URL: ' + data.Url + '</p>')
                        }
                    })
                        // using the done promise callback
                            .done(function(data) {

                                // log data to the console so we can see
                           //     console.log(data);

                                // here we will handle errors and validation messages
                                if ( ! data.success) {

                                    // handle errors for name ---------------
                                    if (data.errors.name) {
                                        $('#name-group').addClass('has-error'); // add the error class to show red input
                                        $('#name-group').append('<div class="help-block">' + data.errors.name + '</div>'); // add the actual error message under our input
                                    }


                                } else {

                                    // ALL GOOD! just show the success message!
                                    // insert SUCCESS
                                    // usually after form submission, you'll want to redirect
                                    // window.location = '/thank-you'; // redirect a user to another page

                                }
                            })

                        // using the fail promise callback
                            .fail(function(data) {

                                // show any errors
                                // best to remove for production
                                console.log(data);
                            });

                    // stop the form from submitting the normal way and refreshing the page
                    event.preventDefault();
                });

            });

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
				padding-top: 5px;
				padding-left: 5px;
				font-size: 18px;
				font-weight: 200;
				line-height: 30px;
				color: inherit;
			}

			p {
				margin: 0px 0px 10px;
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
			    </p>
			  </div>
			</div>
			</div>
		</header>
    <div id="stuff">
  <!--      <p style="margin-left: auto; margin-right: auto; text-align: center; font-size: 3em; line-height: 1.2em">Oi, you forgot to put a channel name after the URL!</p> -->
        <form action="http://192.168.0.133:8089/api/twitch" method="POST">
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
