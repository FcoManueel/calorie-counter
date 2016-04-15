$module.controller('SignupController', ['$scope', '$state', 'SignupService', 'TokenService',
    function ($scope, $state, SignupService, TokenService) {
        /*
         * Attrs
         */

        $scope.user = {
            username: '',
            name: '',
            passwd: '',
            goalCalories: ''
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

        $scope.onSignup = function onLogin(user) {
            SignupService.postSignup(user.username, user.name, user.passwd, user.goalCalories)
                .then(function () {
                    $scope.user = {
                        username: '',
                        name: '',
                        passwd: '',
                        goalCalories: ''
                    };
                    $state.go('dashboard');
                });
        };
    }
]);

