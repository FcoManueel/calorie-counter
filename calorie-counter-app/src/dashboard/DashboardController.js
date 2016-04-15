$module.controller('DashboardController', ['$scope', 'IntakesService', 'TokenService', '$state', 'UserService',
    function ($scope, IntakesService, TokenService, $state, UserService) {
        /*
         * Attrs
         */
        $scope.rangeDate = {
            from: "",
            to: ""
        };

        $scope.rangeTime = {
            from: "",
            to: ""
        };

        $scope.moment = moment;

        $scope.intakes = [];

        $scope.newIntake = {
            name: '',
            calories: '',
            time: '',
            date: ''
        };

        $scope.todayCalories = 0;

        $scope.todayIntakes = [];

        $scope.user = {
            id: '',
            email: '',
            role: '',
            name: '',
            goalCalories: 0
        };

        /*
         * Methods
         */

        $scope.init = function init() {
            var accessToken = TokenService.getAccessToken();

            if (!accessToken) {
                $state.go('login');
                return;
            }

            $scope.getUser();
            $scope.getIntakes();
        };

        $scope.getIntakes = function getIntakes() {
            IntakesService.get().then(function (allIntakes) {
                $scope.intakes = allIntakes;

                $scope.todayIntakes = IntakesService.filter(allIntakes, function (intake) {
                    return moment().diff(intake.consumedAt, 'days') === 0;
                });

                $scope.todayCalories = calculateTodayCalories($scope.todayIntakes);
            });
        };

        $scope.getUser = function getUser() {
            UserService.get().then(function (user) {
                $scope.user = user;
            });
        };

        $scope.onAddNewIntake = function onAddNewIntake(newIntake) {

            IntakesService.post(newIntake.name, newIntake.calories, moment()).then(function () {
                $scope.getIntakes();

                $scope.newIntake = {
                    name: '',
                    calories: ''
                };
            });
        };

        function calculateTodayCalories(intakes) {
            var sum = 0;

            intakes.forEach(function (intake) {
                sum += intake.calories;
            });

            return sum;
        }

        $scope.dateFilter = function dateFilter(from, to){
            return function(value, index, array){
                var created = moment(value.createdAt);
                var shouldShow = true;
                if (from !== "" && created.isBefore(from, 'day')) {
                    shouldShow = false
                }
                if (to !== "" && created.isAfter(to, 'day')){
                    shouldShow = false
                }
                return shouldShow;
            }
        };

        $scope.timeFilter = function timeFilter(from, to){
            toHours = function(m){
                return m.minutes() + m.hours()*60
            };

            return function(value, index, array){
                var createdAt = moment(value.createdAt).format("HH:mm");
                var shouldShow = true;
                if (from !== "" && createdAt < from) {
                    console.log("is before");
                    shouldShow = false
                }
                if (to !== "" && createdAt > to){
                    console.log("is after");
                    shouldShow = false
                }
                return shouldShow;
            }
        };
        $scope.formatDate = function(date){
            return moment().diff(date, 'hours') < 24 ? moment(date).startOf('minute').fromNow() : moment(date).format('YYYY-MM-DD [at] hh:mm A')
        }
    }
]);