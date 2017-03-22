export interface IModule {
  'metadata': IModuleMetadata;
  'module-url': string;
  'settings': Array<IModuleSetting>;
}

export interface IModuleSetting {
  'name': string;
  'default': any;
  'type': string;
  'mandatory': boolean;
  'description': string;
  'value': any;
  'possible-values'?: Array<string>;
}

export interface IModuleMetadata {
  'id': string;
  'name': string;
  'version': string;
  'description': string;
}
