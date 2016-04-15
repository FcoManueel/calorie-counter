$module.controller('LoginController', ['$scope', '$state', 'LoginService', 'TokenService',
    function ($scope, $state, LoginService, TokenService) {
        /*
         * Attrs
         */

        $scope.user = {
            username: '',
            passwd: ''
        };

        /*
         * Methods
         */

        $scope.init = function init() {
            var accessToken = TokenService.loadFromCookie().getAccessToken();
            if (accessToken) {
                $state.go('dashboard');
            }
        };

        $scope.onLogin = function onLogin(user) {
            LoginService.postLogin(user.username, user.passwd)
                .then(function () {
                    $state.go('dashboard');
                    $scope.user = {
                        name: '',
                        passwd: '',
                    };
                });
        };
    }
]);

