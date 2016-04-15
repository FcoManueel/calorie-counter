$module.factory('TokenService', ['$cookieStore',
    function ($cookieStore) {
        var accessToken = "";

        return {
            getAccessToken: function () {
                return accessToken;
            },

            setAccessToken: function (token) {
                $cookieStore.put('x-access-token', token);

                accessToken = token;

                return this;
            },

            loadFromCookie: function () {
                accessToken = $cookieStore.get('x-access-token');
                return this;
            },

            deleteAccessToken: function() {
                $cookieStore.remove('x-access-token');
                return this;
            },

            checkIsAuthenticated: function () {
                return accessToken !== "";
            },
        };
    }
]);
