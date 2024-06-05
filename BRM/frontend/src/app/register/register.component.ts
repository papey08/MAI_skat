import {HttpClient} from '@angular/common/http';
import {Component, inject, OnInit} from '@angular/core';
import {FormBuilder, FormControl, FormGroup, ReactiveFormsModule, Validators,} from '@angular/forms';
import {MatButtonModule} from '@angular/material/button';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatInputModule} from '@angular/material/input';
import {Router} from '@angular/router';
import {AuthService} from '../services/auth.service';
import {MatSelectModule} from '@angular/material/select';
import {RegisterService} from '../DAL/register/register.service';
import {CommonModule} from '@angular/common';
import {MatSnackBar} from "@angular/material/snack-bar";

@Component({
  selector: 'app-register',
  standalone: true,
  imports: [
    ReactiveFormsModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    MatSelectModule,
    CommonModule,
  ],
  templateUrl: './register.component.html',
  styleUrl: './register.component.scss',
})
export class RegisterComponent implements OnInit {
  industries?: Map<string, number>;
  fb = inject(FormBuilder);
  http = inject(HttpClient);
  register = inject(RegisterService);
  snackBar = inject(MatSnackBar);
  authService = inject(AuthService);
  router = inject(Router);

  form = this.fb.nonNullable.group({
    company: this.fb.nonNullable.group({
      description: this.fb.nonNullable.control('', [Validators.maxLength(1000), Validators.required]),
      industry: this.fb.nonNullable.control('', [Validators.required]),
      name: this.fb.nonNullable.control('', [Validators.maxLength(100), Validators.required]),
    }),
    owner: this.fb.nonNullable.group({
      department: this.fb.nonNullable.control('', [Validators.maxLength(100), Validators.required]),
      email: this.fb.nonNullable.control('', [Validators.email, Validators.required]),
      first_name: this.fb.nonNullable.control('', [Validators.maxLength(100), Validators.required]),
      job_title: this.fb.nonNullable.control('', [Validators.maxLength(100), Validators.required]),
      password: this.fb.nonNullable.control('', [Validators.minLength(8), Validators.maxLength(24), Validators.required]),
      second_name: this.fb.nonNullable.control('', [Validators.maxLength(100), Validators.required]),
    }),
  });

  get email() {
    return (this.form.get('owner') as FormGroup).get('email') as FormControl;
  }

  get first_name() {
    return (this.form.get('owner') as FormGroup).get(
      'first_name'
    ) as FormControl;
  }

  get second_name() {
    return (this.form.get('owner') as FormGroup).get(
      'second_name'
    ) as FormControl;
  }

  get password() {
    return (this.form.get('owner') as FormGroup).get('password') as FormControl;
  }

  get job_title() {
    return (this.form.get('owner') as FormGroup).get(
      'job_title'
    ) as FormControl;
  }

  get department() {
    return (this.form.get('owner') as FormGroup).get(
      'department'
    ) as FormControl;
  }

  get description() {
    return (this.form.get('company') as FormGroup).get(
      'description'
    ) as FormControl;
  }

  get industry() {
    return (this.form.get('company') as FormGroup).get(
      'industry'
    ) as FormControl;
  }

  get name() {
    return (this.form.get('company') as FormGroup).get('name') as FormControl;
  }

  ngOnInit(): void {
    this.register.getIndustries().subscribe({
      next: (success) => {
        this.industries = new Map();
        for (let key in success.data) {
          this.industries.set(key, success.data[key]);
        }
      },
    });
  }

  onSubmit(): void {
    if (this.form.valid)
      this.register.register(this.form.getRawValue()).subscribe(() => {
        this.router.navigateByUrl('/');
      });
    else
      this.snackBar.open('Данные некорректно заполнены', undefined, {
        duration: 5000
      })
  }

  redirectToLogin(): void {
    this.router.navigateByUrl('/login');
  }
}
