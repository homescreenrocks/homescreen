(function () {
    'use strict';

    angular.module('homescreenApp', [
        'ngAria',
        'ui.router',
        'ui.bootstrap'
    ]).run(['$rootScope', '$http', function ($rootScope, $http) {

        //$http.defaults.headers.post["Content-Type"] = "application/json";

        $rootScope.$on('$stateNotFound',
            function (event, unfoundState, fromState, fromParams) {
                console.log(unfoundState.to);
                console.log(unfoundState.toParams);
                console.log(unfoundState.options);
            })
    }]);


    /**
     * @name confirmOnExit
     * 
     * @description
     * Prompts user while he navigating away from the current route (or, as long as this directive 
     * is not destroyed) if any unsaved form changes present.
     * 
     * @element Attribute
     * @scope
     * @param confirmOnExit Scope function which will be called on window refresh/close or AngularS $route change to
     *                          decide whether to display the prompt or not.
     * @param confirmMessageWindow Custom message to display before browser refresh or closed.
     * @param confirmMessageRoute Custom message to display before navigating to other route.
     * @param confirmMessage Custom message to display when above specific message is not set.
     * 
     * @example
     * Usage:
     * Example Controller: (using controllerAs syntax in this example)
     * 
     *      angular.module('AppModule', []).controller('pageCtrl', [function () {
     *          this.isDirty = function () {
     *              // do your logic and return 'true' to display the prompt, or 'false' otherwise.
     *              return true;
     *          };
     *      }]);
     * 
     * Template:
     * 
     *      <div confirm-on-exit="pageCtrl.isDirty()" 
     *          confirm-message-window="All your changes will be lost."
     *          confirm-message-route="All your changes will be lost. Are you sure you want to do this?">
     * 
     * @see
     * http://stackoverflow.com/a/28905954/340290
     * 
     * @author Manikanta G
     */
    angular.module('homescreenApp').directive('confirmOnExit', function () {
        return {
            scope: {
                confirmOnExit: '&',
                confirmMessageWindow: '@',
                confirmMessageRoute: '@',
                confirmMessage: '@'
            },
            link: function ($scope, elem, attrs) {
                window.onbeforeunload = function () {
                    if ($scope.confirmOnExit()) {
                        return $scope.confirmMessageWindow || $scope.confirmMessage;
                    }
                }
                var $locationChangeStartUnbind = $scope.$on('$locationChangeStart', function (event, next, current) {
                    if ($scope.confirmOnExit()) {
                        if (!confirm($scope.confirmMessageRoute || $scope.confirmMessage)) {
                            event.preventDefault();
                        }
                    }
                });
                var $stateChangeStartUnbind = $scope.$on('$stateChangeStart', function (event, next, current) {
                    if ($scope.confirmOnExit()) {
                        if (!confirm($scope.confirmMessageRoute || $scope.confirmMessage)) {
                            event.preventDefault();
                        }
                    }
                });

                $scope.$on('$destroy', function () {
                    window.onbeforeunload = null;
                    $locationChangeStartUnbind();
                    $stateChangeStartUnbind();
                });
            }
        };
    });

})();