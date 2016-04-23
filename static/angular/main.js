var app = angular.module('heartack', [
    'ngRoute', 'ngAnimate', 'ui.bootstrap',
    'auth', 'home', 'admin', 'nav'
]);

app.config(function($routeProvider, $locationProvider, $httpProvider) {
    $httpProvider.defaults.headers.common["X-Requested-With"] = "XMLHttpRequest";

    $routeProvider.when("/", {
        templateUrl: "angular/home/home.html",
        controller: "PageCtrl"

    }).when("/login", {
        templateUrl: "angular/auth/login.html",
        controller: "AuthCtrl"

    }).when("/admin", {
        templateUrl: "angular/admin/admin.html",
        controller: "AdminCtrl"

    }).otherwise("/", {
        templateUrl: "index.html",
        controller: "PageCtrl"
    });

});

app.config(function($httpProvider) {
    $httpProvider.interceptors.push('TokenAuthInterceptor');
});