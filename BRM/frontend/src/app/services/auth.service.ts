import {Injectable, signal} from '@angular/core';
import {TokensDataInterface} from '../DAL/login/model/tokens.data.interface';
import {UserData} from "../DAL/login/model/UserData";

@Injectable({
  providedIn: 'root',
})
export class AuthService {
  currentUserSig = signal<TokensDataInterface | undefined | null>(undefined);
  currentUserDataSig = signal<UserData | undefined | null>(undefined);
}
