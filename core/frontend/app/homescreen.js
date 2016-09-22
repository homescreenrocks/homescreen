(function () {
    'use strict';

    angular.module('homescreenApp', [
        'ui.router',
        'ui.bootstrap'
    ]).run(function ($rootScope) {
        $rootScope.$on('$stateNotFound',
            function (event, unfoundState, fromState, fromParams) {
                console.log(unfoundState.to);
                console.log(unfoundState.toParams);
                console.log(unfoundState.options);
            })
    });

})();