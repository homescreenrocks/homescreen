(function () {
    'use strict';

    angular.module('homescreenApp').config(['$stateProvider', '$urlRouterProvider', homescreenAppRouter]);

    function homescreenAppRouter($stateProvider, $urlRouterProvider) {
        $urlRouterProvider.when('', '/');
        $urlRouterProvider.otherwise('/');
        $stateProvider.state({
            name: 'dashboard',
            url: '/',
            template: '<h3>hello world!</h3>'
        });

        $stateProvider.state({
            name: 'about',
            url: '/core/about',
            template: '<h3>Its the UI-Router hello world app!</h3>'
        });

        $stateProvider
            .state({
                name: 'settings',
                url: '/core/settings/:tab',
                controller: 'SettingsCtrl',
                controllerAs: 'vm',
                templateUrl: 'www/app/components/main/Settings.html'
            });
    }

})();