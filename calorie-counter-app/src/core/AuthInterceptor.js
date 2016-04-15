$module.factory('AuthInterceptor', ['TokenService',
    function (TokenService) {
        return {
            request: function (config) {
                var accessToken = TokenService.getAccessToken();

                if (accessToken) {
                    config.headers['Authorization'] = 'Bearer ' + accessToken;
                }

                return config;
            }
        };
    }
]);