import {Component, inject} from '@angular/core';
import {MatCardModule} from "@angular/material/card";
import {DalService} from '../DAL/core/dal.service';
import {AuthService} from "../services/auth.service";
import {CompanyResponse} from "../DAL/core/model/CompanyResponse";
import {MatFormFieldModule} from "@angular/material/form-field";
import {MatIconModule} from "@angular/material/icon";
import {MatInput} from "@angular/material/input";
import {FormBuilder, ReactiveFormsModule, Validators} from "@angular/forms";
import {MatButton} from "@angular/material/button";
import {KeyValuePipe} from "@angular/common";
import {MatOption} from "@angular/material/autocomplete";
import {MatSelect} from "@angular/material/select";
import {RegisterService} from "../DAL/register/register.service";
import {MatSnackBar} from "@angular/material/snack-bar";

@Component({
  selector: 'app-company',
  standalone: true,
  imports: [MatCardModule, MatFormFieldModule, MatIconModule, MatInput, ReactiveFormsModule, MatButton, KeyValuePipe, MatOption, MatSelect],
  templateUrl: './company.component.html',
  styleUrl: './company.component.scss'
})
export class CompanyComponent {
  dal = inject(DalService);
  register = inject(RegisterService);
  authService = inject(AuthService);
  fb = inject(FormBuilder)
  snackBar = inject(MatSnackBar)
  company!: CompanyResponse
  industries?: Map<string, number>;

  formGroupChanged: boolean = false;

  companyFormGroup = this.fb.group({
    description: this.fb.nonNullable.control('', [Validators.maxLength(1000), Validators.required]),
    industry: this.fb.nonNullable.control('', [Validators.required]),
    name: this.fb.nonNullable.control('', [Validators.maxLength(100), Validators.required]),
    rating: this.fb.nonNullable.control(''),
  })

  constructor() {
    this.dal.getCompanyById(+this.authService.currentUserDataSig()?.
      ["company-id"]!).subscribe((company) => {
      this.company = company;
      this.companyFormGroup.setValue({
        description: company.data.description!,
        industry: company.data.industry!,
        name: company.data.name!,
        rating: `${company.data.rating!}`
      })
    })

    this.companyFormGroup.valueChanges.subscribe(value => {
      this.formGroupChanged = value.description != this.company.data.description
        || value.industry != this.company.data.industry || value.name != this.company.data.name
    })

    this.register.getIndustries().subscribe({
      next: (success) => {
        this.industries = new Map();
        for (let key in success.data) {
          this.industries.set(key, success.data[key]);
        }
      },
    });

  }

  updateCompany() {
    if (this.companyFormGroup.valid)
      this.dal.updateCompanyById(this.company.data.id!,
        {
          description: this.companyFormGroup.getRawValue().description,
          name: this.companyFormGroup.getRawValue().name,
          industry: this.companyFormGroup.getRawValue().industry,
          owner_id: this.company.data.owner_id
        }).subscribe({
        next: value => {
          this.company = value;
          this.companyFormGroup.updateValueAndValidity()
        }
      })
    else
      this.snackBar.open('Введены некорректные данные')
  }
}
