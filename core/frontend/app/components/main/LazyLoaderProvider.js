(function () {
    'use strict';

    angular.module('homescreenApp').provider('lazyLoader', ['$controllerProvider', '$provide', '$stateProvider', function ($controllerProvider, $provide, $stateProvider) {
        this.loaders = {
            factory: $provide.factory,
            service: $provide.service,
            controller: $controllerProvider.register,
            state: $stateProvider
        };
        this.$get = () => {
            var self = this;
            return {
                factoryLoader: self.loaders.factory,
                serviceLoader: self.loaders.service,
                controllerLoader: self.loaders.controller,
                addState: (state) => {
                    self.loaders.state.state(state);
                }
            }
        };
    }]);

})();