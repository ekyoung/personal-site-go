var slideShowApp = angular.module('slideShowApp', [
    'ngRoute',
    'slideShowControllers'
]);

slideShowApp.config(['$routeProvider',
    function($routeProvider) {
        $routeProvider.
            when('/slides/:slideId', {
                templateUrl: '/client-side/trips/slide-show/SlideTemplate.html',
                controller: 'slideController'
            }).
            otherwise({
                templateUrl: '/client-side/trips/slide-show/LoadingTemplate.html',
                controller: 'loadingController'
            });
    }]);

slideShowApp.service('slideData', function() {
    var slides = [];
    var firstSlideId = "";
    return {
        slides: slides,
        firstSlideId: firstSlideId
    };
});