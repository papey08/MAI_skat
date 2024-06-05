import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {Observable} from 'rxjs';
import {ContactResponse} from './model/ContactResponse';
import {environment} from '../../../environments/environment';
import {CompanyResponse} from './model/CompanyResponse';
import {MainPageResponse} from './model/MainPageResponse';
import {EmployeeResponse} from './model/EmployeeResponse';
import {AdListResponse} from './model/AdListResponse';
import {UpdateContactRequest} from './model/UpdateContactRequest';
import {AdResponse} from './model/AdResponse';
import {AddAdRequest} from './model/AddAdRequest';
import {ResponseResponse} from './model/ResponseResponse';
import {LeadsListResponse} from './model/LeadsListResponse';
import {LeadResponse} from './model/LeadResponse';
import {StatusesResponse} from './model/StatusesResponse';
import {EmployeeListResponse} from './model/EmployeeListResponse';
import {UpdateLeadRequest} from './model/UpdateLeadRequest';
import {UpdateEmployeeRequest} from './model/UpdateEmployeeRequest';
import {AddEmployeeRequest} from './model/AddEmployeeRequest';
import {NotificationListResponse} from './model/NotificationListResponse';
import {NotificationResponse} from './model/NotificationResponse';
import {UpdateCompanyRequest} from "./model/UpdateCompanyRequest";
import {ResponseListResponse} from "./model/ResponseListResponse";

@Injectable({
  providedIn: 'root',
})
export class DalService {
  constructor(private _http: HttpClient) {
  }

  getResponses(
    limit: number,
    offset: number,
  ): Observable<ResponseListResponse> {
    let queryString = `${environment.coreUrl}/responses?limit=${limit}&offset=${offset}`;

    return this._http.get<ResponseListResponse>(queryString);
  }

  getAds(
    limit: number,
    offset: number,
    company_id?: number
  ): Observable<AdListResponse> {
    let queryString = `${environment.coreUrl}/market?limit=${limit}&offset=${offset}`;

    if (company_id) queryString += `&company_id=${company_id}`;

    return this._http.get<AdListResponse>(queryString);
  }

  getAdById(id: number): Observable<AdResponse> {
    return this._http.get<AdResponse>(`${environment.coreUrl}/market/${id}`, {})
  }

  saveAd(addAdRequest: AddAdRequest) {
    return this._http.post<AdResponse>(
      `${environment.coreUrl}/market`,
      addAdRequest
    );
  }

  adResponse(id: number): Observable<ResponseResponse> {
    return this._http.post<ResponseResponse>(
      `${environment.coreUrl}/market/${id}/response`,
      null
    );
  }

  getLeads(limit: number, offset: number): Observable<LeadsListResponse> {
    return this._http.get<LeadsListResponse>(
      `${environment.coreUrl}/leads?limit=${limit}&offset=${offset}`
    );
  }

  getLeadsStatuses(): Observable<StatusesResponse> {
    return this._http.get<StatusesResponse>(
      `${environment.coreUrl}/leads/statuses`
    );
  }

  getLeadById(id: number): Observable<LeadResponse> {
    return this._http.get<LeadResponse>(`${environment.coreUrl}/leads/${id}`);
  }

  editLead(
    id: number,
    updatedLead: UpdateLeadRequest
  ): Observable<LeadResponse> {
    return this._http.put<LeadResponse>(
      `${environment.coreUrl}/leads/${id}`,
      updatedLead
    );
  }

  getContacts(limit: number, offset: number): Observable<ContactResponse> {
    return this._http.get<ContactResponse>(
      `${environment.coreUrl}/contacts?limit=${limit}&offset=${offset}`
    );
  }

  getContactById(id: number): Observable<any> {
    return this._http.get<any>(`${environment.coreUrl}/contacts/${id}`, {})
  }

  addToContacts(employeeId: number): Observable<ContactResponse> {
    return this._http.post<ContactResponse>(
      `${environment.coreUrl}/contacts`, {employee_id: employeeId}
    );
  }

  updateContact(
    id: number,
    updateContactRequest: UpdateContactRequest
  ): Observable<ContactResponse> {
    return this._http.put<ContactResponse>(
      `${environment.coreUrl}/contacts/${id}`,
      updateContactRequest
    );
  }

  getCompanyById(id: number): Observable<CompanyResponse> {
    return this._http.get<CompanyResponse>(
      `${environment.coreUrl}/companies/${id}`
    );
  }

  updateCompanyById(id: number, updateCompanyRequest: UpdateCompanyRequest): Observable<CompanyResponse> {
    return this._http.put<CompanyResponse>(
      `${environment.coreUrl}/companies/${id}`, updateCompanyRequest
    );
  }

  getCompanyMainPage(id: number): Observable<MainPageResponse> {
    return this._http.get<MainPageResponse>(
      `${environment.coreUrl}/companies/${id}/mainpage`
    );
  }

  getEmployees(
    limit: number,
    offset: number
  ): Observable<EmployeeListResponse> {
    return this._http.get<EmployeeListResponse>(
      `${environment.coreUrl}/employees?limit=${limit}&offset=${offset}`
    );
  }

  getEmployeeById(id: number): Observable<EmployeeResponse> {
    return this._http.get<EmployeeResponse>(
      `${environment.coreUrl}/employees/${id}`
    );
  }

  createEmployee(
    addEmployeeRequest: AddEmployeeRequest
  ): Observable<EmployeeResponse> {
    return this._http.post<EmployeeResponse>(
      `${environment.coreUrl}/employees`,
      addEmployeeRequest
    );
  }

  updateEmployee(
    id: number,
    updateEmployeeRequest: UpdateEmployeeRequest
  ): Observable<EmployeeResponse> {
    return this._http.put<EmployeeResponse>(
      `${environment.coreUrl}/employees/${id}`,
      updateEmployeeRequest
    );
  }

  getNotifications(
    limit: number,
    offset: number,
    onlyNotViewed: boolean
  ): Observable<NotificationListResponse> {
    return this._http.get<NotificationListResponse>(
      `${environment.coreUrl}/notifications?limit=${limit}&offset=${offset}&only_not_viewed=${onlyNotViewed}`
    );
  }

  getNotificationById(
    id: number
  ): Observable<NotificationResponse> {
    return this._http.get<NotificationResponse>(
      `${environment.coreUrl}/notifications/${id}`
    );
  }

  markNotificationAsDone(
    id: number,
    submit: boolean
  ): Observable<NotificationResponse> {
    return this._http.post<NotificationResponse>(
      `${environment.coreUrl}/notifications/${id}?submit=${submit}`,
      undefined
    );
  }
}
