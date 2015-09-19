var browserify = require('browserify');
var gulp = require('gulp');
var source = require("vinyl-source-stream");
var reactify = require('reactify');
var watchify = require('watchify');
var babelify = require('babelify');
var concat = require('gulp-concat');
var sass = require('gulp-sass');

gulp.task('browserify', function(){
    var bundler = browserify({
        entries: ['./index.js'], // Only need initial file, browserify finds the deps
        transform: [babelify], // We want to convert JSX to normal javascript
        debug: true, // Gives us sourcemapping
        cache: {}, packageCache: {}, fullPaths: true // Requirement of watchify
    });
    var watcher  = watchify(bundler);

    return watcher
        .on('update', function () { // When any files update
            var updateStart = Date.now();
            console.log('Updating!');
            watcher.bundle() // Create new bundle that uses the cache for high performance
                .pipe(source('index.js'))
                // This is where you add uglifying etc.
                .pipe(gulp.dest('./dist/'));
            console.log('Updated!', (Date.now() - updateStart) + 'ms');
        })
        .bundle() // Create the initial bundle when starting the task
        .pipe(source('index.js'))
        .pipe(gulp.dest('./dist/'));
});

gulp.task('sass', function() {
    gulp.src('./css/sass/main.scss')
        .pipe(sass({outputStyle: 'compressed'}).on('error', sass.logError))
        .pipe(gulp.dest('./css/sass'));
});

gulp.task('watchsass', function() {
    return gulp
        // Watch the input folder for change,
        // and run `sass` task when something happens
        .watch('./css/sass/main.scss', ['sass'])
        // When there is a change,
        // log a message in the console
        .on('change', function(event) {
            console.log('File ' + event.path + ' was ' + event.type + ', running tasks...');
        });
});


gulp.task('default', ['browserify', 'watchsass']);