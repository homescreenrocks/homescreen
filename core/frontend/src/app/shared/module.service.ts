import { Injectable, Inject } from '@angular/core';
import { Http } from '@angular/http';
import { Observable } from 'rxjs/Observable';
import 'rxjs/add/operator/map';
import 'rxjs/add/operator/retry';
import 'rxjs/add/operator/catch';
import 'rxjs/add/observable/throw';

import { IModule, IModuleSetting } from './module-interface';

@Injectable()
export class ModuleService {

  constructor(
    private http: Http,
    @Inject('API_URL') private api: string
  ) { }

  getSettings(): Observable<Array<IModule>> {
    return this.http
      .get(`/api/v1/modules/`)
      .retry(3)
      .map(res => res.json())
      .map(res => {
        return res.map(m => {
          m.settings.map(s => {
            if (s.type === 'date') {
              s.default = new Date(s.default);
              s.value = new Date(s.value);
            }
            console.log(s);
            return s;
          });
          return m;
        });
      })
      .catch(this.errorHandler);
  }

  private errorHandler(error: Error | any): Observable<any> {
    return Observable.throw(error);
  }
}
