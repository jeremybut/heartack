var app = angular.module('auth', ['flash', 'ngAnimate']);

app.factory('TokenStorage', function() {
    var storageKey = 'auth_token';
    return {
        store: function(token) {
            return localStorage.setItem(storageKey, token);
        },
        retrieve: function() {
            return localStorage.getItem(storageKey);
        },
        clear: function() {
            return localStorage.removeItem(storageKey);
        }
    };
});

app.service('AuthService', function(TokenStorage) {
    var isAuth = TokenStorage.retrieve() == null || TokenStorage.retrieve() === undefined;
    var auth = {
        logged_in: !isAuth
    };

    this.authenticated = function() {
        return auth;
    };
    this.setAuthentication = function(bool) {
        auth.logged_in = bool;
    };
})

app.factory('TokenAuthInterceptor', function($q, $rootScope, TokenStorage, AuthService, $window) {
    console.log("TokenAuthInterceptor");

    return {
        request: function(config) {
            var authToken = TokenStorage.retrieve();
            if (authToken) {
                // intercept and include our auth token in request
                config.headers['Authorization'] = "bearer " + authToken;
            }
            return config;
        },
        responseError: function(error) {
            if (error.status === 401 || error.status === 403) {
                TokenStorage.clear();
                $window.location.href = "#/login";
                AuthService.setAuthentication(false);
            }
            return $q.reject(error);
        }
    };
});

app.controller('AuthCtrl', ['$scope', '$http', 'TokenStorage', 'AuthService', 'Flash', '$location',
    function($scope, $http, TokenStorage, AuthService, Flash, $location) {

        $scope.login = function() {
            $http.post('/api/login', {
                email: $scope.email,
                password: $scope.password

            }).success(function(result, status, headers) {
                AuthService.setAuthentication(true);
                TokenStorage.store(headers('X-Auth-Token'));
                $location.path('/admin');
                Flash.create('success', 'Logged in successfully !');

            }).error(function() {
                TokenStorage.clear();
            });
        };

    }
]);