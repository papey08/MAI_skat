import {HttpClient} from '@angular/common/http';
import {Injectable} from '@angular/core';
import {Observable} from 'rxjs';
import {IndustriesResponse} from './model/IndustriesResponse';
import {environment} from '../../../environments/environment';
import {CompanyAndOwnerResponse} from './model/CompanyAndOwnerResponse';
import {AddCompanyAndOwnerRequest} from './model/AddCompanyAndOwnerRequest';

@Injectable({
  providedIn: 'root',
})
export class RegisterService {
  constructor(private _http: HttpClient) {
  }

  getIndustries(): Observable<IndustriesResponse> {
    return this._http.get<IndustriesResponse>(
      `${environment.registerUrl}/companies/industries`
    );
  }

  register(
    input: AddCompanyAndOwnerRequest
  ): Observable<CompanyAndOwnerResponse> {
    return this._http.post<CompanyAndOwnerResponse>(
      `${environment.registerUrl}/register`,
      input
    );
  }
}
