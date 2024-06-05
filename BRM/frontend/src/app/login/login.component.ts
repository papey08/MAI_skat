import {HttpClient} from '@angular/common/http';
import {Component, inject} from '@angular/core';
import {FormBuilder, ReactiveFormsModule, Validators} from '@angular/forms';
import {Router} from '@angular/router';
import {AuthService} from '../services/auth.service';
import {LoginService} from '../DAL/login/login.service';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatInputModule} from '@angular/material/input';
import {MatButtonModule} from '@angular/material/button';
import {jwtDecode} from "jwt-decode";

@Component({
  selector: 'app-login',
  standalone: true,
  imports: [
    ReactiveFormsModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
  ],
  templateUrl: './login.component.html',
  styleUrl: './login.component.scss',
})
export class LoginComponent {
  fb = inject(FormBuilder);
  http = inject(HttpClient);
  authService = inject(AuthService);
  router = inject(Router);
  loginService = inject(LoginService);

  form = this.fb.nonNullable.group({
    email: ['', Validators.required],
    password: ['', Validators.required],
  });

  onSubmit(): void {
    this.loginService.login(this.form.getRawValue()).subscribe((response) => {
      localStorage.setItem('token', response.data.access);
      this.authService.currentUserSig.set(response.data);
      this.authService.currentUserDataSig.set(jwtDecode(response.data.access));
      this.router.navigateByUrl('/');
    });
  }

  signUp(): void {
    this.router.navigateByUrl('/sign-up');
  }
}
