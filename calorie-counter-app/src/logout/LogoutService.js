$module.factory('LogoutService', ['$http', 'API_ROUTE', 'TokenService', '$state', '$rootScope',
    function ($http, API_ROUTE, TokenService, $state, $rootScope) {
        return {
            onLogout: function () {
                TokenService.deleteAccessToken();
                $state.go('index');
                if(!$scope.$$phase) {
                    $rootScope.$digest();
                }
            }
        };
    }
]);