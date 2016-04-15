$module.controller('LogoutController', ['$scope', 'TokenService', 'LogoutService',
    function ($scope, TokenService, LogoutService){
        $scope.onLogout = LogoutService.onLogout;
    }
]);

