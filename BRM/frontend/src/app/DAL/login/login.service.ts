import {HttpClient} from '@angular/common/http';
import {Injectable} from '@angular/core';
import {LoginInputInterface} from './model/login.input.interface';
import {UserInterface} from './model/user.interface';
import {environment} from '../../../environments/environment';
import {Observable} from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class LoginService {
  constructor(private http: HttpClient) {
  }

  login(input: LoginInputInterface): Observable<UserInterface> {
    return this.http.post<UserInterface>(
      `${environment.loginUrl}/login`,
      input
    );
  }
}
