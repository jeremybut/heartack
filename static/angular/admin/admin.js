var app = angular.module('admin', ['auth', 'flash', 'ngAnimate', 'ngTable']);

app.controller('AdminCtrl', ['$scope', '$http', 'Flash', '$filter', 'ngTableParams',
    function($scope, $http, Flash, $filter, ngTableParams) {

        $http.get('/api/users').success(function(result, status, headers) {
            $scope.patientsTable = result;
        }).error(function(data, status) {
            Flash.create('danger', data);
        });

    }
]);