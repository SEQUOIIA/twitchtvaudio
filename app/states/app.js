var TwiVod = angular.module('TwiVod', ['ui.router', 'TwiVod.home', 'TwiVod.vod', 'react'])
    .controller('TwiVodController', [TwiVodController]);

TwiVod.config(function($stateProvider, $urlRouterProvider) {
    $urlRouterProvider.otherwise('/search');

    $stateProvider
        .state('search', {
            url: '/search',
            templateUrl: '/components/home/home.html',
            controller: function($state, $scope) {
                $scope.angscope = {title: "nopelol", scope: $scope};
            }
        })

        .state('vod', {
            url: '/vod',
            templateUrl: '/components/vod/vod.html',
            controller: function() {

            },
            params: {vodurl: null}
        });
});


function TwiVodController () {

}