(function () {
    'use strict';

    angular.module('homescreenApp').directive('hsNavigation', ['ModulesFactory', 'lazyLoader', hsNavigation]);

    function hsNavigation(ModulesFactory, lazyLoader) {
        return {
            template: `
                <ul class="nav navbar-nav" >
                    <li ui-sref-active="active"><a ui-sref="dashboard">Dashboard</a></li>
                    <li ui-sref-active="active" ng-repeat="module in modules"><a ui-sref="{{module.metadata.id}}">{{module.metadata.name}}</a></li>
                </ul>
                <ul class="nav navbar-nav navbar-right">
                    <li ui-sref-active="active"><a ui-sref="settings"><span class="glyphicon glyphicon-cog" aria-hidden="true"></span><span class="sr-only">Settings</span></a></li>
                </ul>`,
            link: (scope, element, attr) => {
                scope.modules = [];
                scope.currentState = scope.currentState;
                ModulesFactory.getModules().then(data => {
                    scope.modules = data;
                    scope.modules.forEach(module => {
                        lazyLoader.addState({
                            name: module.metadata.id,
                            url: '/modules/' + module.metadata.id,
                            views: {
                                "lazyLoadView": {
                                    controller: 'ExampleModuleController',
                                    templateUrl: '/api/v1/modules/' + module.metadata.id + '/proxy/templates/main.html'
                                }
                            },
                            resolve: { // Any property in resolve should return a promise and is executed before the view is loaded
                                loadMyCtrl: ['$ocLazyLoad', function ($ocLazyLoad) {
                                    // you can lazy load files for an existing module
                                    return $ocLazyLoad.load('/api/v1/modules/' + module.metadata.id + '/proxy/js/ExampleModuleController.js');
                                }]
                            }
                        });
                    });
                });
            }
        };
    }

})();