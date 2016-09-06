(function () {
    'use strict';

    angular.module('homescreenApp').directive('hsNavigation', [hsNavigation]);

    function hsNavigation() {
        return {
            template: `
                <ul class="nav navbar-nav" >
                    <li ui-sref-active="active"><a ui-sref="dashboard">Dashboard</a></li>
                </ul>
                <ul class="nav navbar-nav navbar-right">
                    <li ui-sref-active="active"><a ui-sref="settings"><span class="glyphicon glyphicon-cog" aria-hidden="true"></span><span class="sr-only">Settings</span></a></li>
                </ul>`
        };
    }

})();